package examples

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-sdk/io"
	"github.com/PassKit/passkit-golang-sdk/io/members"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// GetSingleMember takes memberId and returns the record of that member.
func GetSingleMember(memberId string) {
	fmt.Println("Getting a member record...")

	// Generate a members module client
	pkMembersClient := members.NewMembersClient(conn)

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
