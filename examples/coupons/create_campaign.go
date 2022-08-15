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

// CreateCampaign takes a new campaign name and creates a new campaign. The method returns the campaign id.
// A campaign needs to be created because campaign functions as a class object for offer and coupons.
func CreateCampaign() string {
	fmt.Println("Start creating a coupon campaign...")

	// Generate PassKit Client object for Coupon protocol.
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your coupon campaign object.
	campaign := &single_use_coupons.CouponCampaign{
		Name: "Coupon Campaign",
		Status: []io.ProjectStatus{
			io.ProjectStatus_PROJECT_DRAFT,
			io.ProjectStatus_PROJECT_ACTIVE_FOR_OBJECT_CREATION,
		},
	}

	// Send gRPC request to create a coupon campaign record.
	campaignId, err := pkCouponsClient.CreateCouponCampaign(ctx, campaign)
	if err != nil || campaignId == nil {
		log.Fatalf("Create campaign err: %v", err)
	}

	// You need this campaign id to create Offer and Coupon objects in order to issue coupons.
	log.Printf("Create Campaign Success: You have successfully created your coupon campaign. Your campaign id is %s.\n", campaignId.Id)

	return campaignId.Id
}
