package event_tickets

import (
	"fmt"
	"log"
	"time"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/event_tickets"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateEvent takes venueId and productionId to create a new event for an event ticket.
func CreateEvent(venueId string, productionId string) string {
	fmt.Println("Start creating an event")

	// Generate PassKit Client object for Event Tickets protocol.
	pkEventsClient := event_tickets.NewEventTicketsClient(shared.GetGRPCConnection())

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	startDate := time.Now().UTC().AddDate(0, 0, 14).Truncate(time.Hour).Add(13 * time.Hour)
	tStartDate := timestamppb.New(startDate)

	doorsOpen := startDate.Add(-time.Hour)
	tDoorsOpen := timestamppb.New(doorsOpen)

	endDate := startDate.Add(3 * time.Hour)
	tEndDate := timestamppb.New(endDate)

	// Create the event.
	event := &event_tickets.Event{
		Production: &event_tickets.Production{
			Id: productionId,
		},
		Venue: &event_tickets.Venue{
			Id: venueId,
		},
		ScheduledStartDate: tStartDate,
		DoorsOpen:          tDoorsOpen,
		EndDate:            tEndDate,
		RelevantDate:       tStartDate,
	}

	// Send gRPC request to create a venue.
	eventId, err := pkEventsClient.CreateEvent(ctx, event)
	if err != nil || eventId == nil {
		log.Fatalf("Create Event err: %v", err)
	}

	log.Printf("Create Event Success: You have successfully created an event.\n")
	return eventId.Id
}
