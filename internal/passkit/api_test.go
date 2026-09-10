package passkit

import (
	"reflect"
	"strings"
	"testing"

	"github.com/PassKit/passkit-golang-grpc-sdk/io/analytics"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/event_tickets"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/raw"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"google.golang.org/grpc"
)

func TestCoreClientMethodParity(t *testing.T) {
	groups := []struct {
		name    string
		client  any
		methods string
	}{
		{"loyalty", (*members.MembersClient)(nil), "CreateProgram GetProgram UpdateProgram DeleteProgram ListPrograms CreateTier GetTier UpdateTier DeleteTier ListTiers EnrolMember GetMemberRecordById GetMemberRecordByExternalId UpdateMember PatchPerson DeleteMember CountMembers ChangeMemberTier EarnPoints BurnPoints SetPoints CheckInMember CheckOutMember GetProgramEnrolment"},
		{"coupons", (*single_use_coupons.SingleUseCouponsClient)(nil), "CreateCouponCampaign GetCouponCampaign UpdateCouponCampaign DeleteCouponCampaign ListCouponCampaigns CreateCouponOffer GetCouponOffer UpdateCouponOffer DeleteCouponOffer CreateCoupon GetCouponById GetCouponByExternalId UpdateCoupon RedeemCoupon VoidCoupon CountCouponsByCouponCampaign GetAnalytics"},
		{"event tickets", (*event_tickets.EventTicketsClient)(nil), "CreateProduction GetProduction UpdateProduction DeleteProduction CreateVenue GetVenueById UpdateVenue DeleteVenue CreateEvent GetEventById UpdateEvent DeleteEvent CreateTicketType GetTicketTypeById UpdateTicketType DeleteTicketType IssueTicket GetTicketById UpdateTicket ValidateTicket RedeemTicket DeleteTicket CountTickets GetAnalytics"},
		{"flights", (*flights.FlightsClient)(nil), "CreateCarrier GetCarrier UpdateCarrier DeleteCarrier CreatePort GetPort UpdatePort DeletePort CreateFlightDesignator GetFlightDesignator UpdateFlightDesignator DeleteFlightDesignator CreateFlight GetFlight UpdateFlight DeleteFlight CreateBoardingPass GetBoardingPass GetBoardingPassRecord UpdateBoardingPass DeleteBoardingPass"},
		{"analytics", (*analytics.AnalyticsClient)(nil), "GetAnalytics"},
		{"raw", (*raw.RawClient)(nil), "StreamPassUpdates ListPassesByPassProject ListPassesByPassTemplate CreatePassProject GetPassProject UpdatePassProject CopyPassProject DeletePassProject CreatePass GetPassById GetPassByExternalId UpdatePass DeletePass"},
	}
	for _, group := range groups {
		typeOf := reflect.TypeOf(group.client).Elem()
		for _, method := range strings.Fields(group.methods) {
			if _, ok := typeOf.MethodByName(method); !ok {
				t.Errorf("%s is missing %s", group.name, method)
			}
		}
	}
}

func TestAdvancedOperationsRequireOptIn(t *testing.T) {
	api, err := New(func() grpc.ClientConnInterface { return nil }, false)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := api.Advanced(); err == nil {
		t.Fatal("advanced operations should be disabled by default")
	}
}
