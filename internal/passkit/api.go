package passkit

import (
	"errors"

	sdk "github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/analytics"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/event_tickets"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/raw"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/single_use_coupons"
	"google.golang.org/grpc"
)

type ConnectionProvider func() grpc.ClientConnInterface

// API exposes the same product-oriented client groups as the other quickstarts.
type API struct {
	next             ConnectionProvider
	allowDestructive bool
}

func New(next ConnectionProvider, allowDestructive bool) (*API, error) {
	if next == nil {
		return nil, errors.New("connection provider is required")
	}
	return &API{next: next, allowDestructive: allowDestructive}, nil
}

func (a *API) Loyalty() members.MembersClient { return members.NewMembersClient(a.next()) }
func (a *API) Coupons() single_use_coupons.SingleUseCouponsClient {
	return single_use_coupons.NewSingleUseCouponsClient(a.next())
}
func (a *API) EventTickets() event_tickets.EventTicketsClient {
	return event_tickets.NewEventTicketsClient(a.next())
}
func (a *API) Flights() flights.FlightsClient       { return flights.NewFlightsClient(a.next()) }
func (a *API) Templates() sdk.TemplatesClient       { return sdk.NewTemplatesClient(a.next()) }
func (a *API) Images() sdk.ImagesClient             { return sdk.NewImagesClient(a.next()) }
func (a *API) Analytics() analytics.AnalyticsClient { return analytics.NewAnalyticsClient(a.next()) }
func (a *API) Distribution() sdk.DistributionClient { return sdk.NewDistributionClient(a.next()) }
func (a *API) Integrations() sdk.IntegrationsClient { return sdk.NewIntegrationsClient(a.next()) }
func (a *API) Scanners() sdk.UsersClient            { return sdk.NewUsersClient(a.next()) }
func (a *API) Certificates() sdk.CertificatesClient { return sdk.NewCertificatesClient(a.next()) }
func (a *API) Messages() sdk.MessagesClient         { return sdk.NewMessagesClient(a.next()) }
func (a *API) Raw() raw.RawClient                   { return raw.NewRawClient(a.next()) }

type Advanced struct {
	Loyalty      members.MembersClient
	Coupons      single_use_coupons.SingleUseCouponsClient
	EventTickets event_tickets.EventTicketsClient
}

func (a *API) Advanced() (*Advanced, error) {
	if !a.allowDestructive {
		return nil, errors.New("advanced bulk operations are disabled; call passkit.New(provider, true) to enable them")
	}
	return &Advanced{Loyalty: a.Loyalty(), Coupons: a.Coupons(), EventTickets: a.EventTickets()}, nil
}
