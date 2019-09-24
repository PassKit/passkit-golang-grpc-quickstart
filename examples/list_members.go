package examples

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-sdk/io"
	"github.com/PassKit/passkit-golang-sdk/io/members"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// ListMembers takes search conditions as pagination object and returns list of member records which match with the conditions.
func ListMembers(programId string) {
	fmt.Println("Getting a member record...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	listRequest := &members.ListRequest{
		ProgramId: programId,
		Pagination: &io.Pagination{
			FilterField:    []string{"programId"},
			FilterValue:    []string{programId},
			FilterOperator: []string{"eq"},
		},
	}

	res, err := pkMembersClient.ListMembers(ctx, listRequest)
	if err != nil {
		log.Fatalf("List member err: %v", err)
	}

	membersList, err := res.Recv()

	fmt.Printf("You have successfully retrieved member list.\n%v\n", *membersList)
}
