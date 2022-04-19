package membership

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// EnrolMember takes programId, tierId and memberDetails, creates a new member record, and sends a welcome email to deliver membership card url.
// The method returns the member id. Member id is a part of card url.
func EnrolMember(programId, tierId, emailAddress string) string {
	fmt.Println("Start enrolling a member...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	newMember := &members.Member{
		TierId:    tierId,
		ProgramId: programId,
		Person: &io.Person{
			Surname:      "Smith",
			Forename:     "Bailey",
			DisplayName:  "Bailey",
			EmailAddress: emailAddress,
		},
	}

	memberId, err := pkMembersClient.EnrolMember(ctx, newMember)
	if err != nil {
		log.Fatalf("Create tier err: %v", err)
	}

	fmt.Printf("Enrol Member Success: You have successfully enrolled a member. Your member id is %s.\n", memberId.Id)
	fmt.Printf("To check this member's membership card, please visit https://pub1.pskt.io/%s\n", memberId.Id)

	return memberId.Id
}
