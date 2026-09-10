package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateFlightDesignator creates flight designator using flight code.
func CreateFlightDesignator(templateId string) {
	fmt.Println("Start creating a flight designation...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.GetGRPCConnection())

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your flight designator.
	flightDesignator := &flights.FlightDesignator{
		CarrierCode:  CarrierCode,
		FlightNumber: FlightNumber,
		Revision:     DesignatorRevision,
		Schedule: &flights.FlightSchedule{
			Monday: &flights.FlightTimes{
				ScheduledDepartureTime: &io.Time{Hour: 13},
				GateClosingTime:        &io.Time{Hour: 13, Minute: 30},
				ScheduledArrivalTime:   &io.Time{Hour: 14},
				BoardingTime:           &io.Time{Hour: 13},
			},
			Tuesday:   nil,
			Wednesday: nil,
			Thursday:  nil,
			Friday:    nil,
			Saturday:  nil,
			Sunday:    nil,
		},
		Origin:         Origin,
		Destination:    Destination,
		PassTemplateId: templateId,
	}

	// Send gRPC request to create a carrier.
	_, err := pkFlightsClient.CreateFlightDesignator(ctx, flightDesignator)
	if err != nil {
		log.Fatalf("Create flight designator err: %v", err)
	}

	log.Printf("Create Flight Designator Success: You have successfully created flight designator.\n")

}
