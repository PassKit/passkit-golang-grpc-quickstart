package event_tickets

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/event_tickets"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateProdutions takes a new production name and creates a new production.
// If the production name  already exists it cannot be created.
func CreateProduction() string {
	fmt.Println("Start creating a production")

	// Generate PassKit Client object for Event Tickets protocol.
	pkEventsClient := event_tickets.NewEventTicketsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create the production.
	production := &event_tickets.Production{
		Name:                              "Quickstart Production",
		FinePrint:                         "Quickstart Fine print",
		AutoInvalidateTicketsUponEventEnd: io.Toggle_ON,
	}

	// Send gRPC request to create a production.
	productionId, err := pkEventsClient.CreateProduction(ctx, production)
	if err != nil || productionId == nil {
		log.Fatalf("Create production err: %v", err)
	}

	log.Printf("Create Production Success: You have successfully created a production.\n")
	return productionId.Id
}
