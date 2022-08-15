package flights

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateCarrier takes a new carrier code and creates a new carrier.
//If the carrier already exists it cannot be created.
func CreateCarrier() {
	fmt.Println("Start creating a carrier...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your carrier.
	carrier := &flights.Carrier{
		IataCarrierCode:    "YY",
		AirlineName:        "Insert Airline Name",
		PassTypeIdentifier: "", //Change to your apple certificate
	}

	// Send gRPC request to create a carrier.
	_, err := pkFlightsClient.CreateCarrier(ctx, carrier)
	if err != nil {
		log.Fatalf("Create carrier err: %v", err)
	}

	log.Printf("Create Carrier Success: You have successfully created a carrier.\n")

}
