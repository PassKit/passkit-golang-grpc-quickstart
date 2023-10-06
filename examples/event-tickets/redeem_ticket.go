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

// RedeemTicket takes a ticketId to redeems an event ticket.
func RedeemTicket(ticketId string) {
	fmt.Println("Start redeeming an event ticket")

	// Generate PassKit Client object for Event Tickets protocol.
	pkEventsClient := event_tickets.NewEventTicketsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	redemptionDate := time.Now()
	tRedemptionDate := timestamppb.New(redemptionDate)

	eventTicket := &event_tickets.TicketId{
		Id: &event_tickets.TicketId_TicketId{TicketId: ticketId},
	}

	redemptionDetail := event_tickets.RedemptionDetails{
		RedemptionDate: tRedemptionDate,
	}

	// Create the ticket to redeem.
	ticketToRedeem := event_tickets.RedeemTicketRequest{
		Ticket:            eventTicket,
		RedemptionDetails: &redemptionDetail,
	}

	// Send gRPC request to redeem a ticket.
	_, err := pkEventsClient.RedeemTicket(ctx, &ticketToRedeem)
	if err != nil {
		log.Fatalf("Redeem Ticket err: %v", err)
	}

	log.Printf("Redem Ticket Success: You have successfully redeemed an event ticket.\n")
}
