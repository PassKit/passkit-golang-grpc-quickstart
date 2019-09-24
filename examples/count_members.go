package examples

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-sdk/io"
	"github.com/PassKit/passkit-golang-sdk/io/members"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CountMembers takes search conditions as pagination object and returns the number of members who match with the condition.
func CountMembers(programId string) {
	fmt.Println("Counting member records match with conditions...")

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

	countResponse, err := pkMembersClient.CountMembers(ctx, listRequest)
	if err != nil {
		log.Fatalf("Count member err: %v\n", err)
	}

	fmt.Printf("Count result was %v.\n", countResponse.Total)
}
