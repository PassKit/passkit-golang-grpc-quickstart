package flights

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/flights"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
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
	pkFlightsClient := flights.NewFlightsClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	boardingPass := &flights.BoardingPassRecord{
		CarrierCode:         "YY",
		BoardingPoint:       "YYY",
		DeplaningPoint:      "LHR",
		OperatingCarrierPNR: "",
		FlightNumber:        "12345",
		DepartureDate:       &io.Date{Year: 2022, Month: 03, Day: 28},
		SequenceNumber:      2,
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