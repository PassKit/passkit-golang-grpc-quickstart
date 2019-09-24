package examples

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-sdk/io"
	"github.com/PassKit/passkit-golang-sdk/io/members"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// EnrolMember takes programId, tierId and memberDetails, creates a new member record, and sends a welcome email to deliver membership card url.
// The method returns the member id. Member id is a part of card url.
func EnrolMember(programId, tierId, emailAddress string) string {
	fmt.Println("Start enrolling a member...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	newMember := &members.Member{
		TierId:    tierId,
		ProgramId: programId,
		MemberDetails: &io.Person{
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

	// TODO: Change the pass url prefix to the production one
	fmt.Printf("Enrol Member Success: You have successfully enrolled a member. Your member id is %s.\n", memberId.Id)
	fmt.Printf("To check this member's membership card, please visit https://dev.pskt.io/%s\n", memberId.Id)

	return memberId.Id
}
