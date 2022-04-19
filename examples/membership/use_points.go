package membership

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// UsePoints takes a programId of an existing program and memberId of existing member to use points from a chosen member.
func UsePoints(programId,memberId string)  {
	fmt.Println("Start using points from a member...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	//The points to use to should be from whatever point scheme is used on your card e.g. Points, TierPoints or SecondaryPoints
	usePoints := &members.EarnBurnPointsRequest{
		Id:        memberId,
		ProgramId: programId,
		TierPoints: 12,
	}

	response, err := pkMembersClient.BurnPoints(ctx, usePoints)
	if err != nil {
		log.Fatalf("Use members points err: %v", err)
	}

	fmt.Printf("You have successfully used points from [%v].\n", response.Id)
}
