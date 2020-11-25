package grpcpool

/**
 * @DateTime   : 2020/11/23
 * @Author     : xumamba
 * @Description: gRPC connections pool
 **/

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// DialFunc creates a client connection to the given address.
type DialFunc func(addr string) (*grpc.ClientConn, error)

// StateCheckFunc check the connectivity.State of ClientConn.
type StateCheckFunc func(ctx context.Context, conn *grpc.ClientConn) connectivity.State

// Pool connections pool
type Pool interface {
	// GetConn get a connection
	GetConn(addr string) (*grpc.ClientConn, error)
	// Dial force a new connection
	Dial(addr string) (*grpc.ClientConn, error)
	// GetViableConn get current viable connections
	GetViableConn() []string
}

type ConnectionPool struct {
	locker sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc

	dialFunc       DialFunc
	stateCheckFunc StateCheckFunc
	connections    map[string]*gRPCConn
	viableConn     map[string]*gRPCConn

	expiresTime       time.Duration
	checkReadyTimeout time.Duration
	heartbeatInterval time.Duration
}

func (cp *ConnectionPool) getConn(addr string, force bool) (*grpc.ClientConn, error) {
	cp.locker.Lock()
	conn, ok := cp.connections[addr]
	if !ok {
		conn = &gRPCConn{
			pool: cp,
			addr: addr,
		}
		cp.connections[addr] = conn
	}
	cp.locker.Unlock()

	err := conn.activateConn(cp.ctx, force)
	if err != nil {
		return nil, err
	}
	return conn.entity, nil
}

// connReady place the conn into the viable connection pool.
func (cp *ConnectionPool) connReady(conn *gRPCConn) {
	cp.locker.Lock()
	defer cp.locker.Unlock()
	cp.viableConn[conn.addr] = conn
}

// connUnavailable remove addr from the viable connection pool.
func (cp *ConnectionPool) connUnavailable(addr string) {
	cp.locker.Lock()
	defer cp.locker.Unlock()
	delete(cp.viableConn, addr)
}

// InitOption functional options
type InitOption func(pool *ConnectionPool)

func SetExpiresTime(expiresTime time.Duration) InitOption {
	return func(pool *ConnectionPool) {
		pool.expiresTime = expiresTime
	}
}

func SetCheckReadyTimeout(timeout time.Duration) InitOption {
	return func(pool *ConnectionPool) {
		pool.checkReadyTimeout = timeout
	}
}

func SetHeartbeatInterval(interval time.Duration) InitOption {
	return func(pool *ConnectionPool) {
		pool.heartbeatInterval = interval
	}
}

func SetDialFunc(dial DialFunc) InitOption  {
	return func(pool *ConnectionPool) {
		pool.dialFunc = dial
	}
}

func SetStateCheckFunc(stateCheck StateCheckFunc) InitOption {
	return func(pool *ConnectionPool) {
		pool.stateCheckFunc = stateCheck
	}
}

func (cp *ConnectionPool) GetConn(addr string) (*grpc.ClientConn, error) {
	return cp.getConn(addr, false)
}

func (cp *ConnectionPool) Dial(addr string) (*grpc.ClientConn, error) {
	return cp.getConn(addr, true)
}

// GetViableConn get the currently valid connections.
func (cp *ConnectionPool) GetViableConn() []string {
	cp.locker.RLock()
	defer cp.locker.RUnlock()
	var viableConnections []string
	for addr := range cp.viableConn {
		viableConnections = append(viableConnections, addr)
	}
	return viableConnections
}

// New initialize grpc connection pool
func New(opts ...InitOption) *ConnectionPool {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cp := &ConnectionPool{
		ctx:    ctx,
		cancel: cancelFunc,

		dialFunc:       defaultDial,
		stateCheckFunc: defaultStateCheck,
		connections:    make(map[string]*gRPCConn),
		viableConn:     make(map[string]*gRPCConn),

		expiresTime:       defaultTimeout,
		checkReadyTimeout: defaultCheckReadyTimeout,
		heartbeatInterval: defaultHeartbeatInterval,
	}

	for _, opt := range opts {
		opt(cp)
	}

	return cp
}

var (
	connectionPool *ConnectionPool
	once           sync.Once
)

func pool() *ConnectionPool {
	once.Do(func() {
		connectionPool = New()
	})
	return connectionPool
}

func GetConn(addr string) (*grpc.ClientConn, error) {
	return pool().GetConn(addr)
}

func Dial(addr string) (*grpc.ClientConn, error) {
	return pool().Dial(addr)
}

func CurrentViableConn() []string {
	return pool().GetViableConn()
}
