package shared

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/helpers/router"
	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn

// ConnectPasskitSdk takes your credentials and establish connection with PassKit SDK.
func ConnectPasskitSdk(clientCertFile, clientKeyFile, clientCAFile, gRPCHost, gRPCPort string) {
	var err error

	// Generate context object to connect to the server.
	if Conn, err = router.NewCertAuthTLSGRPCClient(fmt.Sprintf("%s:%s", gRPCHost, gRPCPort), clientCertFile, clientKeyFile, clientCAFile); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connect SDK Success: Established connection to the server successfully.")
}
