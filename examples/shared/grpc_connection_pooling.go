package shared

import (
	"fmt"
	"log"
	"sync"

	"github.com/PassKit/passkit-golang-grpc-sdk/helpers/router"
	"google.golang.org/grpc"
)

var (
	connPool     []*grpc.ClientConn
	poolSize     = 5
	currentIndex = 0
	mu           sync.Mutex
)

// InitializeGRPCPool initializes a pool of gRPC connections.
func InitializeGRPCPool(clientCertFile, clientKeyFile, clientCAFile, gRPCHost, gRPCPort string) {
	if len(connPool) > 0 {
		fmt.Println("GRPC Connection Pool already initialized.")
		return
	}

	fmt.Println("Initializing GRPC Connection Pool...")

	for i := 0; i < poolSize; i++ {
		conn, err := router.NewCertAuthTLSGRPCClient(fmt.Sprintf("%s:%s", gRPCHost, gRPCPort), clientCertFile, clientKeyFile, clientCAFile)
		if err != nil {
			log.Fatalf("Failed to create gRPC connection: %v", err)
		}
		connPool = append(connPool, conn)
	}

	if len(connPool) == 0 {
		log.Fatal("No gRPC connections were created. Exiting.")
	}

	fmt.Println("GRPC Connection Pool Initialized with", len(connPool), "connections.")
}

// GetGRPCConnection returns a connection from the pool using round-robin selection.
func GetGRPCConnection() *grpc.ClientConn {
	mu.Lock()
	defer mu.Unlock()

	if len(connPool) == 0 {
		log.Fatal("GRPC Connection Pool is empty. Ensure InitializeGRPCPool() was called before using GetGRPCConnection().")
	}

	conn := connPool[currentIndex]
	if conn == nil {
		log.Fatal("Retrieved a nil gRPC connection from the pool. Exiting.")
	}

	currentIndex = (currentIndex + 1) % len(connPool)
	return conn
}

// ShutdownGRPCPool closes all gRPC connections.
func ShutdownGRPCPool() {
	mu.Lock()
	defer mu.Unlock()

	for _, conn := range connPool {
		if conn != nil {
			conn.Close()
		}
	}
	connPool = nil
	fmt.Println("GRPC Connection Pool Shut Down.")
}
