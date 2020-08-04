package examples

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"

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
		},
	}

	countResponse, err := pkMembersClient.CountMembers(ctx, listRequest)
	if err != nil {
		log.Fatalf("Count member err: %v\n", err)
	}

	fmt.Printf("Count result was %v.\n", countResponse.Total)
}
