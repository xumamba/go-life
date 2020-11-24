package grpcpool

/**
 * @DateTime   : 2020/11/24
 * @Author     : xumamba
 * @Description:
 **/

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

const (
	defaultTimeout           = 100 * time.Second
	defaultCheckReadyTimeout = 5 * time.Second
	defaultHeartbeatInterval = 10 * time.Second
)

var errNoReady = errors.New("connection state is not ready")

func defaultStateCheck(ctx context.Context, conn *grpc.ClientConn) connectivity.State {
	for {
		state := conn.GetState()
		if state == connectivity.Ready || state == connectivity.Shutdown {
			return state
		}
		if !conn.WaitForStateChange(ctx, state) {
			return connectivity.Idle
		}
	}
}

func defaultDialFunc(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure())
}
