package flights

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AppleCertificateID is configured by the root quickstart from PASSKIT_APPLE_CERTIFICATE_ID.
var AppleCertificateID string

// CreateCarrier takes a new carrier code and creates a new carrier.
// If the carrier already exists, creation is skipped and the existing carrier is reused.
func CreateCarrier() {
	fmt.Println("Start creating a carrier...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.GetGRPCConnection())

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	carrier := &flights.Carrier{
		IataCarrierCode:    CarrierCode,
		AirlineName:        "PassKit Quickstart Air",
		PassTypeIdentifier: AppleCertificateID,
	}

	_, err := pkFlightsClient.CreateCarrier(ctx, carrier)
	if status.Code(err) == codes.AlreadyExists {
		CarrierCreated = false
		log.Printf("Carrier %s already exists; skipping creation.\n", CarrierCode)
		return
	}
	if err != nil {
		log.Fatalf("Create carrier err: %v", err)
	}
	CarrierCreated = true
	log.Printf("Create Carrier Success: created IATA carrier %s.\n", CarrierCode)

}
