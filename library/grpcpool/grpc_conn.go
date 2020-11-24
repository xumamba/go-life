package grpcpool

/**
 * @DateTime   : 2020/11/24
 * @Author     : xumamba
 * @Description: gRPC single connection
 **/

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type gRPCConn struct {
	locker sync.RWMutex

	pool *ConnectionPool

	addr       string
	entity     *grpc.ClientConn
	state      connectivity.State
	expires    time.Time
	retryTimes int
	cancelFunc context.CancelFunc
}

func (c *gRPCConn) activateConn(ctx context.Context, force bool) error {
	c.locker.Lock()
	defer c.locker.Unlock()

	if !force && c.entity != nil {
		if c.state == connectivity.Ready {
			return nil
		}
		if c.state == connectivity.Idle {
			return errNoReady
		}
	}

	if c.entity != nil {
		c.entity.Close()
	}
	clientConn, err := c.pool.dial(c.addr)
	if err != nil {
		return err
	}
	c.entity = clientConn

	readyCtx, cancelFunc := context.WithTimeout(ctx, c.pool.checkReadyTimeout)
	defer cancelFunc()
	entityStatus := c.pool.stateCheck(readyCtx, c.entity)
	if entityStatus != connectivity.Ready {
		return errNoReady
	}

	heartbeatCtx, hbCancelFunc := context.WithCancel(ctx)
	c.cancelFunc = hbCancelFunc
	go c.heartbeat(heartbeatCtx)

	c.ready()
	return nil
}

func (c *gRPCConn) ready() {
	c.state = connectivity.Ready
	c.expires = time.Now().Add(c.pool.timeout)
	c.retryTimes = 0
	c.pool.connReady(c)
}

func (c *gRPCConn) idle() {
	c.state = connectivity.Idle
	c.retryTimes++
	c.pool.connUnavailable(c.addr)
}

func (c *gRPCConn) shutdown() {
	c.state = connectivity.Shutdown
	c.entity.Close()
	c.cancelFunc()
	c.pool.connUnavailable(c.addr)
}

func (c *gRPCConn) isExpired() bool {
	return c.expires.Before(time.Now())
}

func (c *gRPCConn) getState() connectivity.State {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.state
}

func (c *gRPCConn) heartbeat(ctx context.Context) {
	ticker := time.NewTicker(c.pool.heartbeatInterval)
	for c.getState() != connectivity.Shutdown {
		select {
		case <-ctx.Done():
			c.shutdown()
			break
		case <-ticker.C:
			c.healthCheck(ctx)
		}
	}
}

func (c *gRPCConn) healthCheck(ctx context.Context) {
	ctx, cancelFunc := context.WithTimeout(ctx, c.pool.checkReadyTimeout)
	defer cancelFunc()

	switch c.pool.stateCheck(ctx, c.entity) {
	case connectivity.Ready:
		c.ready()
	case connectivity.Shutdown:
		c.shutdown()
	case connectivity.Idle:
		if c.isExpired() {
			c.shutdown()
		} else {
			c.idle()
		}
	}
}
