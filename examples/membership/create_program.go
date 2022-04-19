package membership

import (
	"fmt"
	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"log"

	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateProgram takes a new program name and creates a new program. The method returns the program id.
// A program needs to be created because program functions as a class object for tier and members.
func CreateProgram() string {
	fmt.Println("Start creating a membership program...")

	// Generate PassKit Client object for Membership protocol.
	pkMembersClient := members.NewMembersClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// Create your membership program object.
	program := &members.Program{
		Name: "Membership Program",
		Status: []io.ProjectStatus{
			io.ProjectStatus_PROJECT_DRAFT,
			io.ProjectStatus_PROJECT_ACTIVE_FOR_OBJECT_CREATION,
		},
	}

	// Send gRPC request to create a membership program record.
	programId, err := pkMembersClient.CreateProgram(ctx, program)
	if err != nil || programId == nil {
		log.Fatalf("Create program err: %v", err)
	}

	// You need this program id to create Tier and Member objects in order to issue membership card.
	log.Printf("Create Program Success: You have successfully created your membership program. Your program id is %s.\n", programId.Id)

	return programId.Id
}
