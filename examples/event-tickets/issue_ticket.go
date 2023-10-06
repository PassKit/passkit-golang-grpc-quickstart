package event_tickets

import (
	"fmt"
	"log"
	"time"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/event_tickets"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// IssueEventTicket takes a ticketTypeId, productionId and eventId to issue an event ticket.
func IssueEventTicket(ticketTypeId string, productionId string, eventId string) string {
	fmt.Println("Start issuing an event ticket")

	// Generate PassKit Client object for Event Tickets protocol.
	pkEventsClient := event_tickets.NewEventTicketsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	endDate := time.Date(2023, 12, 13, 13, 0, 0, 0, time.UTC)
	tEndDate := timestamppb.New(endDate)

	ticketEvent := event_tickets.IssueTicketRequest_EventId{
		EventId: eventId,
	}

	// Create the ticket to issue.
	ticket := &event_tickets.IssueTicketRequest{
		TicketTypeId: ticketTypeId,
		EventInfo:    &ticketEvent,
		ExpiryDate:   tEndDate,
		OrderNumber:  "1",
		Person: &io.Person{
			Surname:      "Loopy",
			Forename:     "Larry",
			DisplayName:  "Larry",
			EmailAddress: "",
		},
		FaceValue: &event_tickets.FaceValue{
			Amount:   100,
			Currency: "GBR",
		},
		SeatInfo: &event_tickets.Seat{
			Gate:    "12",
			Section: "A",
			Row:     "123",
			Seat:    "D",
		},
		TicketNumber: "123",
	}

	// Send gRPC request to issue a ticket.
	ticketId, err := pkEventsClient.IssueTicket(ctx, ticket)
	if err != nil || ticketId == nil {
		log.Fatalf("Issue Ticket err: %v", err)
	}

	log.Printf("Issue Ticket Success: You have successfully issued an event ticket.\n")
	return ticketId.Id
}
