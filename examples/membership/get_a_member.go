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

// GetSingleMember takes memberId and returns the record of that member.
func GetSingleMember(memberId string) {
	fmt.Println("Getting a member record...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	id := &io.Id{
		Id: memberId,
	}

	memberRecord, err := pkMembersClient.GetMemberRecordById(ctx, id)
	if err != nil {
		log.Fatalf("Get member err: %v", err)
	}

	fmt.Printf("You have successfully retrieved a member record for [%s] %v\n", memberId, *memberRecord)
}
