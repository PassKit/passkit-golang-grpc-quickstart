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

// ValidateTicket takes a ticketId to validate an event ticket.
func ValidateTicket(ticketId string) {
	fmt.Println("Start validating an event ticket")

	// Generate PassKit Client object for Event Tickets protocol.
	pkEventsClient := event_tickets.NewEventTicketsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	validateDate := time.Now()
	tvalidateDate := timestamppb.New(validateDate)

	eventTicket := &event_tickets.TicketId{
		Id: &event_tickets.TicketId_TicketId{TicketId: ticketId},
	}

	validateDetail := event_tickets.ValidateDetails{
		ValidateDate: tvalidateDate,
	}

	// Create the ticket to validate.
	ticketToValidate := event_tickets.ValidateTicketRequest{
		MaxNumberOfValidations: 1,
		Ticket:                 eventTicket,
		ValidateDetails:        &validateDetail,
	}

	// Send gRPC request to validate a ticket.
	_, err := pkEventsClient.ValidateTicket(ctx, &ticketToValidate)
	if err != nil {
		log.Fatalf("Validate Ticket err: %v", err)
	}

	log.Printf("Validate Ticket Success: You have successfully validated an event ticket.\n")
}
