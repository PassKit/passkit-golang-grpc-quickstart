package coupons

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateCoupon takes campaignId, offerId and couponDetails, creates a new coupon record, and sends a welcome email to deliver coupon card url.
// The method returns the coupon id. Coupon id is a part of card url.
func CreateCoupon(campaignId, offerId, emailAddress string) string {
	fmt.Println("Start creating a coupon...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Generates coupon with mandatory fields, more fields can be added, refer to docs.passkit.io and select Coupons for the full list
	newCoupon := &single_use_coupons.Coupon{
		OfferId:    offerId,
		CampaignId: campaignId,
		Person: &io.Person{
			Surname:      "Smith",
			Forename:     "Bailey",
			DisplayName:  "Bailey",
			EmailAddress: emailAddress,
		},
	}

	couponId, err := pkCouponsClient.CreateCoupon(ctx, newCoupon)
	if err != nil {
		log.Fatalf("Create coupon err: %v", err)
	}

	fmt.Printf("Create Coupon Success: You have successfully created a coupon. Your coupon id is %s.\n", couponId.Id)
	fmt.Printf("To check this user's coupon card, please visit https://pub1.pskt.io/%s\n", couponId.Id)

	return couponId.Id
}

