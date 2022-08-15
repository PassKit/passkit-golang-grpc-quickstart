package coupons

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// RedeemCoupon takes a campaignId of an existing campaign and couponId of existing coupon to redeem that coupon.
func RedeemCoupon(campaignId,couponId string)  {
	fmt.Println("Start redeeming a coupon...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	//The points to add to should be whatever point scheme is used on your card e.g. Points, TierPoints or SecondaryPoints
	redeemCoupon := &single_use_coupons.Coupon{
		Id:        couponId,
		CampaignId: campaignId,
		Status: single_use_coupons.CouponStatus_REDEEMED,
	}

	response, err := pkCouponsClient.RedeemCoupon(ctx, redeemCoupon)
	if err != nil {
		log.Fatalf("Redeem coupon err: %v", err)
	}

	fmt.Printf("You have successfully redeemed coupon [%v].\n", response.Id)
}

