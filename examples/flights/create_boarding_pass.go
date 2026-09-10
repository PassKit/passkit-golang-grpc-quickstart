package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateBoardingPass takes templateId and customer details creates a new boarding pass, and sends a welcome email to deliver boarding pass url.
// The method returns the boarding pass id. Boarding Pass id is a part of card url.
func CreateBoardingPass(templateId, emailAddress string) string {
	fmt.Println("Start creating a boarding pass...")

	// Generate a flights module client
	pkFlightsClient := flights.NewFlightsClient(shared.GetGRPCConnection())

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	boardingPass := &flights.BoardingPassRecord{
		CarrierCode:         CarrierCode,
		BoardingPoint:       Origin,
		DeplaningPoint:      Destination,
		OperatingCarrierPNR: OperatingCarrierPNR,
		FlightNumber:        FlightNumber,
		DepartureDate:       DepartureDate,
		SequenceNumber:      2,
		SeatNumber:          "12A",
		Class:               "Economy",
		Passenger: &flights.Passenger{
			PassengerDetails: &io.Person{
				Surname:      "Smith",
				Forename:     "Bailey",
				DisplayName:  "Bailey",
				EmailAddress: emailAddress,
			},
		},
	}

	boardingPassId, err := pkFlightsClient.CreateBoardingPass(ctx, boardingPass)
	if err != nil {
		log.Fatalf("Create boarding pass err: %v", err)
	}

	fmt.Printf("Create Boarding Pass Success: You have successfully created a boarding pass. Your boarding pass id is %s.\n", boardingPassId.BoardingPasses)
	fmt.Printf("To check this user's boarding pass, please visit https://pub1.pskt.io/%s\n", boardingPassId.BoardingPasses)

	return boardingPassId.String()
}
