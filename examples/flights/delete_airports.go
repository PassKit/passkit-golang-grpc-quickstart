package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// DeleteAirport takes an existing airport code and deletes the airport associated with it.
//If the airport doesn't exist it cannot be deleted.
func DeleteAirport(airport string){
	fmt.Println("Start deleting airport...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your  existing airport.
	airportCode := &flights.AirportCode{
		AirportCode: airport,
	}

	// Send gRPC request to delete an airport.
	_, err := pkFlightsClient.DeletePort(ctx, airportCode)
	if err != nil {
		log.Fatalf("Delete airport err: %v", err)
	}

	log.Printf("Delete Airport Success: You have successfully deleted an airport.\n")

}
