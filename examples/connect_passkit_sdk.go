package examples

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/PassKit/passkit-golang-sdk/helpers/router"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

// ConnectPasskitSdk takes your credentials and establish connection with PassKit SDK.
func ConnectPasskitSdk(clientCertFile, clientKeyFile, clientCAFile, gRPCHost, gRPCPort string) {
	var err error

	cert, err := ioutil.ReadFile(clientCertFile)

	if err != nil {
		log.Fatalf("could not load certificate file: %v", err)
	}

	key, err := ioutil.ReadFile(clientKeyFile)

	if err != nil {
		log.Fatalf("could not load key file: %v", err)
	}

	ca, err := ioutil.ReadFile(clientCAFile)

	if err != nil {
		log.Fatalf("could not load ca file: %v", err)
	}

	// Generate context object to connect to the server.
	if conn, err = router.NewCertAuthTLSGRPCClient(fmt.Sprintf("%s:%s", gRPCHost, gRPCPort), string(cert), string(key), string(ca)); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connect SDK Success: Established connection to the server successfully.")
}
