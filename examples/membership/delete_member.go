package membership

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// DeleteMember takes programId, tierId, memberId and memberDetails, deletes an existing member record.
func DeleteMember(programId, tierId, memberId string) {
	fmt.Println("Start deleting a member...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	member := &members.Member{
		Id: memberId,
		TierId:    tierId,
		ProgramId: programId,
	}

	_, err := pkMembersClient.DeleteMember(ctx, member)
	if err != nil {
		log.Fatalf("Delete member err: %v", err)
	}

	fmt.Printf("Delete Member Success: You have successfully deleted a member.")

}
