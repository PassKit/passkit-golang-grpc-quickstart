package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// DeleteFlight takes an existing flight number as well as other details and deletes the flight associated with it.
// If the flight doesn't exist it cannot be deleted.
func DeleteFlight() {
	fmt.Println("Start deleting flight...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.GetGRPCConnection())

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your existing flight.
	flight := &flights.FlightRequest{
		CarrierCode:    CarrierCode,
		FlightNumber:   FlightNumber,
		BoardingPoint:  Origin,
		DeplaningPoint: Destination,
		DepartureDate:  DepartureDate,
	}

	// Send gRPC request to delete a flight.
	_, err := pkFlightsClient.DeleteFlight(ctx, flight)
	if err != nil {
		log.Fatalf("Delete flight err: %v", err)
	}

	log.Printf("Delete Flight Success: You have successfully deleted a flight.\n")

}
