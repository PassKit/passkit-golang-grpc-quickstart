package event_tickets

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/event_tickets"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateVenue creates a new venue for an event.
// If the venue name  already exists it cannot be created.
func CreateVenue() string {
	fmt.Println("Start creating a venue")

	// Generate PassKit Client object for Event Tickets protocol.
	pkEventsClient := event_tickets.NewEventTicketsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create the venue.
	venue := &event_tickets.Venue{
		Name:     "Quickstart Venue",
		Address:  "123 Abc Street",
		Timezone: "Europe/London",
	}

	// Send gRPC request to create a venue.
	venueId, err := pkEventsClient.CreateVenue(ctx, venue)
	if err != nil || venueId == nil {
		log.Fatalf("Create venue err: %v", err)
	}

	log.Printf("Create Venue Success: You have successfully created a venue.\n")
	return venueId.Id
}
