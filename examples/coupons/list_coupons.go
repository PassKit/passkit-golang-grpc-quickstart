package coupons

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	goio "io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// ListCoupons takes search conditions as pagination object and returns list of coupon records which match with the conditions.
func ListCoupons(campaignId string) {
	fmt.Println("Listing coupons")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Generates list request
	listRequest := &single_use_coupons.ListRequest{
		CouponCampaignId: campaignId,

	}

	res, err := pkCouponsClient.ListCouponsByCouponCampaign(ctx, listRequest)
	if err != nil {
		log.Fatalf("List coupons err: %v", err)
	}

	couponCount := 0
	for {
		m, err := res.Recv()
		if err == goio.EOF {
			break
		}

		if err != nil {
			fmt.Print("Error listing coupons: %s", err.Error())
			res.CloseSend()
			break
		}

		fmt.Printf("Listing coupon: %s\n", m.Id)
		couponCount++
	}

	fmt.Printf("Listed a total of %d coupons\n", couponCount)
}
