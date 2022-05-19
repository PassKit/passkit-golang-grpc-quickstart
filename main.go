package main

import (
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/coupons"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/flights"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/membership"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
)

/*
	Prerequisites to run this file:
		- store your certificate.pem in passkit-golang-members-quickstart/certs/ directory
		- store your key.pem in passkit-golang-members-quickstart/certs/ directory
		- store your ca-chain.pem in passkit-golang-members-quickstart/certs/ directory
		- decrypt your key.pem with `cd ./certs openssl ec -in key.pem -out key.pem` from project root
*/

const (
	// The address & port of the PassKit gRPC service.
	gRPCHost = "grpc.pub1.passkit.io"
	gRPCPort = "443"

	// The location of your client certificates.
	clientCertFile = "certs/certificate.pem" // [Required] Please store your certificate.pem at ./certs directory. Your client certificate you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login).
	clientKeyFile  = "certs/key.pem"         // [Required] Please store your key.pem at ./certs directory. Your private key you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login). You need to decrypt the key before use. Check README.md for details.
	clientCAFile   = "certs/ca-chain.pem"    // [Required] Please store your ca-chain.pem at ./certs directory. The certificate chain you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login).

	emailAddressToReceiveSamplePassUrl = "YOUR_EMAIL_ADDRESS@EMAIL.COM" // [Required] Please set your email address to receive digital card url by email.
)

// These variables will be used by EngageWithMembers methods.
var programId string
var tierId string
var memberId string
var campaignId string
var offerId string
var couponId string
var boardingPassId string
var templateId string

func main() {
	if emailAddressToReceiveSamplePassUrl == "" {
		log.Fatal("Please set emailAddressToReceiveSamplePassUrl with your email address in coupons.go so that you can receive sample welcome email with digital card url.")
	}

	ConnectWithPasskitServer()
	// Membership functions
	//IssueMembershipCard()
	//EngageWithMembers()

	// Coupon functions
	//IssueCoupon()
	//EngageWithCoupons()

	// Flight functions
	IssueBoardingPass()
	EngageWithBoardingPass()
}

// In order to use PassKit SDK, you need to establish the connection to the PassKit server first.
func ConnectWithPasskitServer() {
	shared.ConnectPasskitSdk(clientCertFile, clientKeyFile, clientCAFile, gRPCHost, gRPCPort)
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

	// Delete arrival airport
	flights.DeleteAirport("YY4")
	// Delete departure airport
	flights.DeleteAirport("ADP")

	flights.DeleteCarrier()
}
