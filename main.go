package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/coupons"
	event_tickets "github.com/PassKit/passkit-golang-grpc-quickstart/examples/event-tickets"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/flights"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/membership"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-quickstart/internal/config"
	"github.com/PassKit/passkit-golang-grpc-quickstart/internal/passkit"
	"google.golang.org/grpc"
)

var emailAddressToReceiveSamplePassUrl string
var cleanupDelay time.Duration

// These variables will be used by EngageWithMembers methods.
var programId string
var tierId string
var memberId string
var campaignId string
var offerId string
var couponId string
var boardingPassId string
var templateId string
var productionId string
var venueId string
var ticketTypeId string
var ticketId string
var eventId string

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" || os.Args[1] == "--help" {
		fmt.Println("Usage: go run . <membership|coupons|event-tickets|flights>")
		fmt.Println("Copy .env.example to .env and add your PassKit credentials first.")
		return
	}
	workflow := strings.ToLower(os.Args[1])
	if workflow == "loyalty" {
		workflow = "membership"
	}
	if workflow == "events" {
		workflow = "event-tickets"
	}
	if workflow != "membership" && workflow != "coupons" && workflow != "event-tickets" && workflow != "flights" {
		log.Fatalf("unknown workflow %q; use membership, coupons, event-tickets, or flights", os.Args[1])
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := cfg.Validate(workflow); err != nil {
		log.Fatal(err)
	}
	if err := shared.SetGRPCPoolSize(cfg.PoolSize); err != nil {
		log.Fatal(err)
	}
	shared.InitializeGRPCPool(cfg.Certificate, cfg.Key, cfg.CAChain, cfg.Host, cfg.Port)
	defer shared.ShutdownGRPCPool()
	if _, err := passkit.New(func() grpc.ClientConnInterface { return shared.GetGRPCConnection() }, false); err != nil {
		log.Fatal(err)
	}
	cleanupDelay = time.Duration(cfg.CleanupDelaySeconds) * time.Second

	switch workflow {
	case "membership":
		emailAddressToReceiveSamplePassUrl = cfg.Email
		IssueMembershipCard()
		pauseBeforeFollowUp()
		EngageWithMembers()
	case "coupons":
		emailAddressToReceiveSamplePassUrl = cfg.Email
		IssueCoupon()
		pauseBeforeFollowUp()
		EngageWithCoupons()
	case "event-tickets":
		emailAddressToReceiveSamplePassUrl = cfg.Email
		IssueEventTicket()
		pauseBeforeFollowUp()
		EngageWithEventTicket()
	case "flights":
		emailAddressToReceiveSamplePassUrl = cfg.Email
		flights.AppleCertificateID = cfg.AppleCertificateID
		IssueBoardingPass()
		pauseBeforeFollowUp()
		EngageWithBoardingPass()
	}
}

func pauseBeforeFollowUp() {
	if cleanupDelay <= 0 {
		return
	}
	fmt.Printf("Waiting %s before follow-up and cleanup operations...\n", cleanupDelay)
	time.Sleep(cleanupDelay)
}

// Each method has the minimum information needed to execute the method, if you
// would like to add more details please refer to
// https://docs.passkit.io/protocols/member/
// for fields that can be added.

// IssueMembershipCard shows the methods needed to issue a membership card
// In order to create a membership card for your member, you need to take following process:
// 1. Create a program
// 2. Create a tier
// 3. Enrol a member (enrolling a member will automatically issue a membership card to your member).
func IssueMembershipCard() {
	newProgramId := membership.CreateProgram()

	newTierId := membership.CreateTier(newProgramId)

	newMemberId := membership.EnrolMember(newProgramId, newTierId, emailAddressToReceiveSamplePassUrl)

	// These variables will be used by EngageWithMembers().
	programId = newProgramId
	tierId = newTierId
	memberId = newMemberId
}

// EngageWithMembers show methods you can use to engage with loyalty members.
// When you execute EngageWithMembers method by itself, please establish connection with the server first by using
// ConnectWithPasskitServer
func EngageWithMembers() {
	membership.GetSingleMember(memberId)

	membership.ListMembers(programId)

	membership.CountMembers(programId, tierId)

	membership.SendWelcomeEmail(memberId)

	membership.UpdateMember_EmailAddress(memberId, tierId, programId)

	membership.AddPoints(programId, memberId)

	membership.UsePoints(programId, memberId)

	membership.DeleteMember(programId, tierId, memberId)

}

// Each method has the minimum information needed to execute the method, if
// you would like to add more details please refer to
// https://docs.passkit.io/protocols/coupon/
//for fields that can be added.

// IssueCoupon shows the methods needed to issue a coupon
// In order to create a coupon, you need to take following process:
// 1. Create a campaign
// 2. Create an offer
// 3. Enrol someone on a couppn (enrolling someone will automatically issue a coupon card to your customer).
func IssueCoupon() {
	newCampaignId := coupons.CreateCampaign()
	newOfferId := coupons.CreateOffer(newCampaignId)
	newCouponId := coupons.CreateCoupon(newCampaignId, newOfferId, emailAddressToReceiveSamplePassUrl)

	// These variables will be used by EngageWithCoupons().
	campaignId = newCampaignId
	offerId = newOfferId
	couponId = newCouponId
}

// EngageWithCoupons show methods you can use to engage with coupons.
// When you execute EngageWithCoupons method by itself, please establish connection with the server first by using
// ConnectWithPasskitServer
func EngageWithCoupons() {
	coupons.GetSingleCoupon(couponId)

	coupons.ListCoupons(campaignId)

	coupons.CountCoupons(campaignId, offerId)

	coupons.UpdateCoupon(campaignId, couponId)

	coupons.RedeemCoupon(campaignId, couponId)

	coupons.VoidCoupon(campaignId, offerId, couponId)

	coupons.DeleteCouponOffer(campaignId)

}

// Each method has the minimum information needed to execute the method, if you
// would like to add more details please refer to
// https://docs.passkit.io/protocols/boarding/
// for fields that can be added.

// IssueBoardingPass shows the methods needed to issue a boarding pass
// In order to create a boarding pass, you need to take following process:
// 1. Create a carrier
// 2. Create departure and arrival airport
// 3. Create boarding pass template
// 4. Create a flight
// 5. Create flight designator
// 6. Enrol someone on a boarding pass (enrolling someone will automatically issue a boarding pass to your customer).
func IssueBoardingPass() {
	newTemplateId := flights.CreateTemplate()

	flights.CreateCarrier()

	flights.CreateAirports()

	flights.WaitForFlightResources()

	flights.CreateFlight(newTemplateId)

	flights.CreateFlightDesignator(newTemplateId)

	newBoardingPass := flights.CreateBoardingPass(newTemplateId, emailAddressToReceiveSamplePassUrl)

	boardingPassId = newBoardingPass
	templateId = newTemplateId

}

// EngageWithBoardingPass show methods you can use to engage with boarding passes.
// When you execute EngageWithBoardingPass method by itself, please establish connection with the server first by using
// ConnectWithPasskitServer
func EngageWithBoardingPass() {
	flights.DeleteFlight()

	flights.DeleteFlightDesignator()

	flights.DeleteAirport(flights.Origin)
	flights.DeleteAirport(flights.Destination)

	flights.DeleteCarrier()
}

// Each method has the minimum information needed to execute the method, if you
// would like to add more details please refer to
// https://docs.passkit.io/protocols/event-tickets/
// for fields that can be added.

// IssueEventTicket shows the methods needed to issue an event ticket
// In order to create an event ticket, you need to take following process:
// 1. Create a venue
// 2. Create a production
// 3. Create event ticket template
// 4. Create a event ticket type
// 5. Issue an event ticket
func IssueEventTicket() {

	newTemplateId := event_tickets.CreateTemplate()

	newProductionId := event_tickets.CreateProduction()

	newVenueId := event_tickets.CreateVenue()

	newTicketTypeId := event_tickets.CreateTicketType(newProductionId, newTemplateId)

	newEventId := event_tickets.CreateEvent(newVenueId, newProductionId)

	newTicketId := event_tickets.IssueEventTicket(newTicketTypeId, newProductionId, newEventId, emailAddressToReceiveSamplePassUrl)

	venueId = newVenueId

	productionId = newProductionId

	ticketTypeId = newTicketTypeId

	ticketId = newTicketId

}

// EngageWithEventTicket show methods you can use to engage with event tickets.
// When you execute EngageWithEventTicket method by itself, please establish connection with the server first by using
// ConnectWithPasskitServer
func EngageWithEventTicket() {
	event_tickets.ValidateTicket(ticketId)
	event_tickets.RedeemTicket(ticketId)
}
