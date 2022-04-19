package coupons

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// VoidCoupon takes the couponId, offerId and campaignId to void an existing coupon.
func VoidCoupon( campaignId, offerId, couponId string) {
	fmt.Println("Start voiding a coupon...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Generates coupon to delete
	coupon := &single_use_coupons.Coupon{
		Id: couponId,
		OfferId: offerId,
		CampaignId: campaignId,
	}

	_, err := pkCouponsClient.VoidCoupon(ctx, coupon)
	if err != nil {
		log.Fatalf("Delete coupon err: %v", err)
	}

	fmt.Printf("Delete Coupon Success: You have successfully deleted a coupon.")

}
