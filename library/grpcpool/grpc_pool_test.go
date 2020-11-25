package grpcpool

/**
 * @DateTime   : 2020/11/25
 * @Author     : xumamba
 * @Description:
 **/
import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	pb "go-life/example/grpc/base"
)

// start test server: go-life/example/grpc/base/greeter_server
var (
	serverAddr = "localhost:50001"
	ctx        = context.Background()
	content    = "test gRPC pool"
)

func TestGRPCPool(t *testing.T) {
	t.Run("CurrentViableConn", func(t *testing.T) {
		viableConn := CurrentViableConn()
		assert.Equal(t, viableConn, []string(nil))
	})

	t.Run("GetConn", func(t *testing.T) {
		conn, err := GetConn(serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		sayHello(t, conn)
	})

	t.Run("Dial", func(t *testing.T) {
		conn, err := Dial(serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		sayHello(t, conn)
		viableConn := CurrentViableConn()
		assert.Equal(t, viableConn, []string{serverAddr})
	})

	t.Run("GetConn_pool", func(t *testing.T) {
		cp := New()
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				if _, err := cp.GetConn(serverAddr); err != nil {
					t.Fatal(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	})

	t.Run("connection idle", func(t *testing.T) {
		cp := New()
		_, err := cp.GetConn(serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		conn := cp.connections[serverAddr]
		conn.idle()
		_, err = cp.GetConn(serverAddr)
		assert.Equal(t, err, errNoReady)
	})

	t.Run("connection expired", func(t *testing.T) {
		flag := true
		mockCheckState := func(ctx context.Context, conn *grpc.ClientConn) connectivity.State {
			if flag == true {
				flag = false
				return connectivity.Ready
			} else {
				return connectivity.Idle
			}
		}
		cp := New(
			SetDialFunc(func(addr string) (*grpc.ClientConn, error) {
				return grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
			}),
			SetExpiresTime(3*time.Second),
			SetCheckReadyTimeout(1*time.Second),
			SetHeartbeatInterval(2*time.Second),
			SetStateCheckFunc(mockCheckState))
		conn, err := cp.GetConn(serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		sayHello(t, conn)

		viableConn := cp.GetViableConn()
		assert.Equal(t, viableConn, []string{serverAddr})

		time.Sleep(4 * time.Second)
		viableConn = cp.GetViableConn()
		assert.Equal(t, viableConn, []string(nil))
	})

	t.Run("SetHealthCheck", func(t *testing.T) {
		healthStateCheck := func(ctx context.Context, conn *grpc.ClientConn) connectivity.State {
			client := pb.NewGreeterClient(conn)
			_, err := client.SayHello(ctx, &pb.HelloRequest{Name: "health check"})
			if err != nil {
				return connectivity.Idle
			}
			return connectivity.Ready
		}
		cp := New(
			SetStateCheckFunc(healthStateCheck),
			SetExpiresTime(5*time.Second),
			SetCheckReadyTimeout(1*time.Second),
			SetHeartbeatInterval(2*time.Second),
		)

		conn, err := cp.GetConn(serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		sayHello(t, conn)

		viableConn := cp.GetViableConn()
		assert.Equal(t, viableConn, []string{serverAddr})

		time.Sleep(6 * time.Second)
		viableConn = cp.GetViableConn()
		assert.Equal(t, viableConn, []string{serverAddr})
	})
}

func sayHello(t *testing.T, conn *grpc.ClientConn) {
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: content})
	if err != nil {
		t.Fatalf("failed to say hello: %v", err)
	}
	t.Logf("Greetings from success: %v", reply.Message)
}
