package event_tickets

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/event_tickets"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateTicketType takes the productionId and the templateId creating a new ticket type for an event.
func CreateTicketType(productionId string, templateId string) string {
	fmt.Println("Start creating a ticket type")

	// Generate PassKit Client object for Event Tickets protocol.
	pkEventsClient := event_tickets.NewEventTicketsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create the ticket type.
	ticketType := &event_tickets.TicketType{
		Name:                       "Quickstart Ticket Type",
		ProductionId:               productionId,
		BeforeRedeemPassTemplateId: templateId,
		Uid:                        "",
	}

	// Send gRPC request to create a ticket type.
	ticketTypeId, err := pkEventsClient.CreateTicketType(ctx, ticketType)
	if err != nil || ticketTypeId == nil {
		log.Fatalf("Create ticket type err: %v", err)
	}

	log.Printf("Create Ticket Type Success: You have successfully created a ticket type.\n")
	return ticketTypeId.Id
}
