package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"log"


	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// DeleteCarrier takes an existing carrier code and deletes the carrier associated with it.
//If the carrier doesn't exist it cannot be deleted.
func DeleteCarrier(){
	fmt.Println("Start deleting a carrier...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your existing carrier.
	carrierCode := &flights.CarrierCode{
		CarrierCode: "YY",
	}

	// Send gRPC request to create an airport.
	_, err := pkFlightsClient.DeleteCarrier(ctx, carrierCode)
	if err != nil {
		log.Fatalf("Delete carrier err: %v", err)
	}

	log.Printf("Delete Carrier Success: You have successfully deleted a carrier.\n")

}
