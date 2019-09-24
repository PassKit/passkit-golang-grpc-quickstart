package examples

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-sdk/io"
	"github.com/PassKit/passkit-golang-sdk/io/members"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateTier takes a programId of an existing program, creates a new template (based of default template), creates a tier, and links this tier to the program.
// The method returns the tier id.
func CreateTier(programId string) string {
	fmt.Println("Start creating a membership tier...")

	// Generate a template module client
	pkTemplatesClient := io.NewTemplatesClient(conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// In order to create a tier, we need a pass template id which holds pass design data. Let's use the default pass template for now.
	defaultTemplateRequest := &io.DefaultTemplateRequest{
		Protocol: io.PassProtocol_MEMBERSHIP,
		Revision: uint32(1),
	}

	defaultPassTemplate, err := pkTemplatesClient.GetDefaultTemplate(ctx, defaultTemplateRequest)
	if err != nil {
		log.Fatalf("Create tier err: %v", err)
	}

	// If you use the default template, you need to set name, description and timezone because these fields are mandatory.
	defaultPassTemplate.Name = "quick_start"
	defaultPassTemplate.Description = "quick start sample template"
	defaultPassTemplate.Timezone = "America/New_York"

	templateId, err := pkTemplatesClient.CreateTemplate(ctx, defaultPassTemplate)
	if templateId == nil || err != nil {
		log.Fatalf("Could not create tier: %v", err)
	}

	// We now have a program id which we created in create_program.go example and default pass template.
	// Let's create a tier. First, let's create a members module client.
	pkMembersClient := members.NewMembersClient(conn)

	tier := &members.Tier{
		ProgramId:      programId,
		PassTemplateId: templateId.Id,
		Id:             "base",
		TierIndex:      uint32(1),
		Name:           "Base Tier",
	}

	tierId, err := pkMembersClient.CreateTier(ctx, tier)
	if err != nil {
		log.Fatalf("Create tier err: %v", err)
	}

	fmt.Printf("Create Tier Success: You have successfully created your tier. Your tier id is %s.\n", tierId.Id)

	return tierId.Id
}
