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

// UpdateMember takes memberId and memberDetails, and updates existing member record.
// UpdateMember_EmailAddress shows how to update member's registered email address.
func UpdateMember_EmailAddress(memberId, tierId, programId string) {
	fmt.Println("Updating member's email address...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	updateTo := &members.Member{
		Id:        memberId,
		TierId:    tierId,
		ProgramId: programId,
		Person: &io.Person{
			Surname:      "Smith",
			Forename:     "Bailey",
			DisplayName:  "Bailey",
			EmailAddress: "replace_with_new_email_address@gmail.com",
		},
	}

	response, err := pkMembersClient.UpdateMember(ctx, updateTo)
	if err != nil {
		log.Fatalf("Update member email err: %v", err)
	}

	fmt.Printf("You have successfully updated email address of [%v].\n", response.Id)
}
