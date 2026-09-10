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

// CreateAirports takes a new airport code and creates a new airport.
// If an airport already exists, creation is skipped and the existing airport is reused.
// Please make sure an arrival and departure airport exist.
func CreateAirports() {
	fmt.Println("Start creating an airport...")

	// Generate PassKit Client object for Flights protocol.
	pkFlightsClient := flights.NewFlightsClient(shared.GetGRPCConnection())

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	Origin = createOrReuseAirport(pkFlightsClient, ctx, "departure", Origin, runSeed+101, "YYYY", "GB", "Europe/London")
	Destination = createOrReuseAirport(pkFlightsClient, ctx, "arrival", Destination, runSeed+503, "VHHH", "HK", "Asia/Hong_Kong")

}

func createOrReuseAirport(client flights.FlightsClient, ctx context.Context, label, preferred string, seed int64, preferredICAO, country, timezone string) string {
	for attempt := int64(0); attempt < 20; attempt++ {
		code := preferred
		port := &flights.Port{
			IataAirportCode: code,
			IcaoAirportCode: preferredICAO,
			CityName:        "Quickstart City", AirportName: "Quickstart Airport",
			CountryCode: country, Timezone: timezone,
		}
		if attempt > 0 {
			code = alphaCode(seed+attempt*7919, 4)
			port.IataAirportCode = ""
			port.IcaoAirportCode = code
		}

		_, err := client.CreatePort(ctx, port)
		if err == nil {
			CreatedAirports[code] = true
			log.Printf("Created %s airport %s.\n", label, code)
			return code
		}
		if status.Code(err) != codes.AlreadyExists {
			log.Fatalf("Create %s airport err: %v", label, err)
		}
		_, lookupErr := client.GetPort(ctx, &flights.AirportCode{AirportCode: code})
		if lookupErr == nil {
			CreatedAirports[code] = false
			log.Printf("%s airport %s already exists in this account; skipping creation.\n", label, code)
			return code
		}
		log.Printf("Airport code %s is unavailable to this account; trying another ICAO code.\n", code)
	}
	log.Fatalf("Could not find an available ICAO code for the %s airport after 20 attempts", label)
	return ""
}
