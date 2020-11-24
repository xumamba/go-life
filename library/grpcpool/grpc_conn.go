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

	pool *ConnectcionPool

	addr       string
	entity     *grpc.ClientConn
	status     connectivity.State
	expires    time.Time
	retryTimes int
}

func (c *gRPCConn) activateConn(ctx context.Context) error {
	c.locker.Lock()
	defer c.locker.Unlock()

	if c.entity != nil{
		c.entity.Close()
	}
	clientConn, err := c.pool.dial(c.addr)
	if err != nil{
		return err
	}
	c.entity = clientConn

	readyCtx, cancelFunc := context.WithTimeout(ctx, c.pool.checkReadyTimeout)
	defer cancelFunc()
	entityStatus := c.pool.stateCheck(readyCtx, c.entity)
	if entityStatus != connectivity.Ready{
		return err
	}

}
