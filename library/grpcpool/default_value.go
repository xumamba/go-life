package grpcpool

/**
 * @DateTime   : 2020/11/24
 * @Author     : xumamba
 * @Description:
 **/

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)


const (
	defaultTimeout = 100 * time.Second
	defaultCheckReadyTimeout = 5 * time.Second
	defaultHeartbeatInterval = 10 * time.Second
)

func defaultStateCheck(ctx context.Context, conn *grpc.ClientConn) connectivity.State {
	conn.GetState()
}