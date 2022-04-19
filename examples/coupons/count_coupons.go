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

// CountCoupons takes search conditions as pagination object and returns the number of coupons who match with the condition.
func CountCoupons(campaignId, offerId string) {
	fmt.Println("Counting coupon records that match with conditions...")

	// Generate a coupons module client
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	//Filters coupons based on chosen filter/s e.g. based on offers, more than one field can be chosen
	listRequest := &single_use_coupons.ListRequest{
		CouponCampaignId: campaignId,
		Filters: &io.Filters{
			FilterGroups: []*io.FilterGroup{
				{
					Condition: io.Operator_AND,
					FieldFilters: []*io.FieldFilter{
						{
							FilterField: "offerId",
							FilterValue: offerId,
							FilterOperator: "eq",
						},
					},
				},
			},
		},
	}

	countResponse, err := pkCouponsClient.CountCouponsByCouponCampaign(ctx, listRequest)
	if err != nil {
		log.Fatalf("Count coupons err: %v\n", err)
	}

	fmt.Printf("Count result was %v.\n", countResponse.Total)
}
