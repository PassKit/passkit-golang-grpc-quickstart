package shared

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// newCertAuthTLSGRPCClient replaces the broken router helper in SDK v1.1.163.
func newCertAuthTLSGRPCClient(address, certFile, keyFile, caFile string) (*grpc.ClientConn, error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("load client key pair: %w", err)
	}
	ca, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("read CA certificate: %w", err)
	}
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("CA chain contains no valid certificates")
	}
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, fmt.Errorf("parse gRPC address: %w", err)
	}
	creds := credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS12, ServerName: host,
		Certificates: []tls.Certificate{certificate}, RootCAs: roots,
	})
	return grpc.NewClient(address, grpc.WithTransportCredentials(creds))
}
