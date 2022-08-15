package flights


import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateAirports takes a new airport code and creates a new airport.
// If the airport already exists it cannot be created.
// Please make sure an arrival and departure airport exist.
func CreateAirports(){
	fmt.Println("Start creating an airport...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create the departure airport.
	departureAirport := &flights.Port{
		IataAirportCode: "YY4",
		IcaoAirportCode: "YYYY",
		CityName: "Insert City Name",
		AirportName: "Insert Airport Name",
		CountryCode: "IE",
		Timezone: "Europe/London",
	}

	// Send gRPC request to create an airport.
	_, err := pkFlightsClient.CreatePort(ctx, departureAirport)
	if err != nil {
		log.Fatalf("Create departure airport err: %v", err)
	}

	log.Printf("Create Airport Success: You have successfully created the departure airport.\n")

	// Create the arrival airport.
	arrivalAirport := &flights.Port{
		IataAirportCode: "ADP",
		IcaoAirportCode: "ADPY",
		CityName: "Insert City Name",
		AirportName: "Insert Airport Name",
		CountryCode: "IE",
		Timezone: "Europe/London",
	}

	// Send gRPC request to create an airport.
	_, err = pkFlightsClient.CreatePort(ctx, arrivalAirport)
	if err != nil {
		log.Fatalf("Create arrival airport err: %v", err)
	}

	log.Printf("Create Airport Success: You have successfully created the arrival airport.\n")

}
