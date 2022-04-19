package coupons

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateOffer takes a campaignId of an existing campaign, creates a new template (based of default template), creates an offer, and links this offer to the campaign.
// The method returns the offer id.
func CreateOffer(campaignId string) string {
	fmt.Println("Start creating an offer...")

	// Generate a template module client
	pkTemplatesClient := io.NewTemplatesClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// In order to create an offer, we need a pass template id which holds pass design data. Let's use the default pass template for now.
	defaultTemplateRequest := &io.DefaultTemplateRequest{
		Protocol: io.PassProtocol_SINGLE_USE_COUPON,
		Revision: uint32(1),
	}

	defaultPassTemplate, err := pkTemplatesClient.GetDefaultTemplate(ctx, defaultTemplateRequest)
	if err != nil {
		log.Fatalf("Create offer template err: %v", err)
	}

	// If you use the default template, you need to set name, description and timezone because these fields are mandatory.
	defaultPassTemplate.Name = "quick_start"
	defaultPassTemplate.Description = "quick start sample template"
	defaultPassTemplate.Timezone = "America/New_York"

	templateId, err := pkTemplatesClient.CreateTemplate(ctx, defaultPassTemplate)
	if templateId == nil || err != nil {
		log.Fatalf("Could not create offer template: %v", err)
	}

	// We now have a campaign id which we created in create_campaign.go example and default pass template.
	// Let's create an offer. First, let's create a coupons module client.
	pkCouponsClient := single_use_coupons.NewSingleUseCouponsClient(shared.Conn)

	// Generates offer with required fields, more fields can be added, refer to docs.passkit.io and select Coupons for the full list
	offer := &single_use_coupons.CouponOffer{
		CampaignId: campaignId,
		BeforeRedeemPassTemplateId: templateId.Id,
		Id:             "base",
		OfferTitle: "Base Offer",
		OfferShortTitle: "Base Offer",
		OfferDetails: "Base offer",
		IssueStartDate: timestamppb.Now(),
		IssueEndDate: timestamppb.New(time.Now().AddDate(0,0,1)),
	}

	offerId, err := pkCouponsClient.CreateCouponOffer(ctx, offer)
	if err != nil {
		log.Fatalf("Create offer err: %v", err)
	}

	fmt.Printf("Create Offer Success: You have successfully created your offer. Your offer id is %s.\n", offerId.Id)

	return offerId.Id
}