package flights

import (
	"fmt"
	"log"
	"time"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	pkflights "github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"golang.org/x/net/context"
)

// WaitForFlightResources avoids a read-after-create race across pooled gRPC
// connections by waiting until the carrier and both airports are resolvable.
func WaitForFlightResources() {
	client := pkflights.NewFlightsClient(shared.GetGRPCConnection())
	ctx := context.Background()
	for attempt := 1; attempt <= 20; attempt++ {
		_, carrierErr := client.GetCarrier(ctx, &pkflights.CarrierCode{CarrierCode: CarrierCode})
		_, originErr := client.GetPort(ctx, &pkflights.AirportCode{AirportCode: Origin})
		_, destinationErr := client.GetPort(ctx, &pkflights.AirportCode{AirportCode: Destination})
		if carrierErr == nil && originErr == nil && destinationErr == nil {
			log.Printf("Carrier %s and airports %s/%s are ready.\n", CarrierCode, Origin, Destination)
			return
		}
		if attempt == 20 {
			log.Fatalf("Flight resources did not become ready: carrier=%v origin=%v destination=%v", carrierErr, originErr, destinationErr)
		}
		fmt.Println("Waiting for carrier and airport records to become available...")
		time.Sleep(500 * time.Millisecond)
	}
}
