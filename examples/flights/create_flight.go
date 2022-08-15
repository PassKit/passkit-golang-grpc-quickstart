package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
)

// CreateFlight takes templateId to use as base template and uses a carrier code and creates a new flight.
func CreateFlight(templateId string){
	fmt.Println("Start creating a Flight...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	localDate := &io.LocalDateTime{
		DateTime: "2022-03-28",
	}
	// Create your flight.
	flight := &flights.Flight{
		CarrierCode: "YY",
		FlightNumber: "12345",
		BoardingPoint: "YY4",
		DeplaningPoint: "ADP",
		DepartureDate: &io.Date{Year: 2022, Month: 03, Day: 28},
		ScheduledDepartureTime: localDate,
		PassTemplateId: templateId,
	}

	// Send gRPC request to create a flight.
	_, err := pkFlightsClient.CreateFlight(ctx, flight)
	if err != nil {
		log.Fatalf("Create flight err: %v", err)
	}

	log.Printf("Create Flight Success: You have successfully created a flight.\n")

}
