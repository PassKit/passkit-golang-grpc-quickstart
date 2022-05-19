package coupons

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// DeleteCouponOffer takes the offerId to delete an existing offer.
func DeleteCouponOffer(campaignId string) {
	fmt.Println("Start deleting a coupon offer...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Generates offer to delete
	campaign := &io.Id{
		Id: campaignId,
	}

	_, err := pkCouponsClient.DeleteCouponCampaign(ctx, campaign)
	if err != nil {
		log.Fatalf("Delete coupon offer err: %v", err)
	}

	fmt.Printf("Delete Coupon Offer Success: You have successfully deleted a coupon offer.")

}
