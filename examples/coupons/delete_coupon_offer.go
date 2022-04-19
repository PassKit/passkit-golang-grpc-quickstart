package coupons

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// DeleteCouponOffer takes the offerId to delete an existing offer.
func DeleteCouponOffer( offerId string) {
	fmt.Println("Start deleting a coupon offer...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Generates offer to delete
	offer := &io.Id{
		Id: offerId,
	}

	_, err := pkCouponsClient.DeleteCouponOffer(ctx, offer)
	if err != nil {
		log.Fatalf("Delete coupon offer err: %v", err)
	}

	fmt.Printf("Delete Coupon Offer Success: You have successfully deleted a coupon offer.")

}
