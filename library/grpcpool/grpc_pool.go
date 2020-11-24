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

// DialFunc
type DialFunc func(addr string) (*grpc.ClientConn, error)

// StateCheckFunc
type StateCheckFunc func(ctx context.Context, conn *grpc.ClientConn) connectivity.State

type ConnectionPool struct {
	locker sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc

	dial        DialFunc
	stateCheck StateCheckFunc
	connections map[string]*gRPCConn
	viableConn  map[string]*gRPCConn

	timeout           time.Duration
	checkReadyTimeout time.Duration
	heartbeatInterval time.Duration
}

func (cp *ConnectionPool) getConn(addr string) (*grpc.ClientConn, error) {
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

	err := conn.activateConn(cp.ctx)
	if err != nil {
		return nil, err
	}
	return conn.entity, nil
}

// connReady place the conn into the viable connection pool.
func (cp *ConnectionPool) connReady(conn *gRPCConn)  {
	cp.locker.Lock()
	defer cp.locker.Unlock()
	cp.viableConn[conn.addr] = conn
}

// connUnavailable remove addr from the viable connection pool.
func (cp *ConnectionPool) connUnavailable(addr string)  {
	cp.locker.Lock()
	defer cp.locker.Unlock()
	delete(cp.viableConn, addr)
}

// getViableConn get the currently valid connections.
func (cp *ConnectionPool) getViableConn() []string {
	cp.locker.RLock()
	defer cp.locker.RUnlock()
	var viableConnections []string
	for addr := range cp.viableConn{
		viableConnections = append(viableConnections, addr)
	}
	return viableConnections
}


// InitOption functional options
type InitOption func(pool *ConnectionPool)


// New initialize grpc connection pool
func New(dial DialFunc, opts ...InitOption) *ConnectionPool {
	ctx, cancelFunc := context.WithCancel(context.Background())
	cp := &ConnectionPool{
		ctx:    ctx,
		cancel: cancelFunc,

		dial:        dial,c
		stateCheck: defaultStateCheck,
		connections: make(map[string]*gRPCConn),
		viableConn:  make(map[string]*gRPCConn),

		timeout:           100 * time.Second,
		checkReadyTimeout: 5 * time.Second,
		heartbeatInterval: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(cp)
	}

	return cp
}
