package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// DeleteFlightDesignator takes an existing flight designation and deletes the flight designator associated with it.
//If the flight designator doesn't exist it cannot be deleted.
func DeleteFlightDesignator(){
	fmt.Println("Start deleting a flight designation...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your existing flight designator.
	flightDesignator := &flights.FlightDesignatorRequest{
		CarrierCode: "YY",
		FlightNumber: "YY123",
		Revision: 0,
	}

	// Send gRPC request to delete flight designator.
	_, err := pkFlightsClient.DeleteFlightDesignator(ctx, flightDesignator)
	if err != nil {
		log.Fatalf("Delete carrier err: %v", err)
	}

	log.Printf("Delete Flight Designator Success: You have successfully deleted the flight designator.\n")

}
