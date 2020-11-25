# gRPC 连接池

## 设计目的：

&emsp;&emsp;连接复用，节省创建连接和销毁连接的开销；

## 设计原则：

- 连接超时处理
- 连接心跳保活
- 自定义连接建立方式、状态检测

## 代码设计：

```go
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
```

## 连接池使用：

```go
import (
    "context"
    "fmt"
    
    "google.golang.org/grpc"
    "google.golang.org/grpc/connectivity"
    
    pb "go-life/example/grpc/base"
)

func TestGetConn(t *testing.T) {
	healthStateCheck := func(ctx context.Context, conn *grpc.ClientConn) connectivity.State {
		client := pb.NewGreeterClient(conn)
		_, err := client.HealthCheck(ctx, &pb.Request{Ping: "health check"})
		if err != nil {
			return connectivity.Idle
		}
		return connectivity.Ready
	}
	dial := func(addr string) (*grpc.ClientConn, error) {
		return grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	}
	cp := New(
		SetDialFunc(dial),
		SetStateCheckFunc(healthStateCheck),
		SetExpiresTime(5*time.Second),
		SetCheckReadyTimeout(1*time.Second),
		SetHeartbeatInterval(2*time.Second),
	)

	conn, err := cp.GetConn(serverAddr)
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "hello gRPC connection pool"})
	fmt.Println(reply, err)
}
```