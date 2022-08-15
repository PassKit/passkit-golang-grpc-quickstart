package coupons

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// GetSingleCoupon takes couponId and returns the record of that coupon.
func GetSingleCoupon(couponId string)  {
	fmt.Println("Getting a coupon record...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Generates id to find coupon with
	id := &io.Id{
		Id: couponId,
	}

	couponRecord, err := pkCouponsClient.GetCouponById(ctx, id)
	if err != nil {
		log.Fatalf("Get coupon err: %v", err)
	}

	fmt.Printf("You have successfully retrieved a coupon record for [%s] %v\n", couponId, *couponRecord)
}

