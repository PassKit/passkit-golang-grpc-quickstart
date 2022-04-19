package membership

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// AddPoints takes a programId of an existing program and memberId of existing member to add points to chosen member.
func AddPoints(programId,memberId string)  {
	fmt.Println("Start adding points to a member...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	//The points to add to should be whatever point scheme is used on your card e.g. Points, TierPoints or SecondaryPoints
	addPointsTo := &members.EarnBurnPointsRequest{
		Id:        memberId,
		ProgramId: programId,
		TierPoints: 24,
	}

	response, err := pkMembersClient.EarnPoints(ctx, addPointsTo)
	if err != nil {
		log.Fatalf("Add members points err: %v", err)
	}

	fmt.Printf("You have successfully added points to [%v].\n", response.Id)
}
