package examples

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	goio "io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// ListMembers takes search conditions as pagination object and returns list of member records which match with the conditions.
func ListMembers(programId string) {
	fmt.Println("Listing members")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	listRequest := &members.ListRequest{
		ProgramId: programId,
		Filters: &io.Filters{
			FilterGroups: []*io.FilterGroup{
				{
					Condition: io.Operator_AND,
					FieldFilters: []*io.FieldFilter{
						{
							FilterField: "programId",
							FilterValue: programId,
							FilterOperator: "eq",
						},
					},
				},
			},
			Limit: 5,
		},
	}

	res, err := pkMembersClient.ListMembers(ctx, listRequest)
	if err != nil {
		log.Fatalf("List member err: %v", err)
	}

	memberCount := 0
	for {
		m, err := res.Recv()
		if err == goio.EOF {
			break
		}

		fmt.Printf("Listing member: %s\n", m.Id)
		memberCount++
	}

	fmt.Printf("Listed a total of %d members\n", memberCount)
}
