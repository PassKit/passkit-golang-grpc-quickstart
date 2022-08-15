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

// UpdateCoupon takes a campaignId of an existing campaign and couponId of existing coupon to update that coupon.
func UpdateCoupon(campaignId,couponId string)  {
	fmt.Println("Start updating a coupon...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Generates coupon with updated fields, more fields can be updated, refer to docs.passkit.io and select Coupons for the full list
	updateCoupon := &single_use_coupons.Coupon{
		Id:        couponId,
		CampaignId: campaignId,
		Person: &io.Person{
			Surname:      "Jones",
		},
	}

	response, err := pkCouponsClient.UpdateCoupon(ctx, updateCoupon)
	if err != nil {
		log.Fatalf("Update coupon err: %v", err)
	}

	fmt.Printf("You have successfully updated coupon [%v].\n", response.Id)
}
