package event_tickets

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-grpc-quickstart/examples/shared"
	"github.com/PassKit/passkit-golang-grpc-sdk/io"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// CreateTemplate creates the pass template for event tickets.
func CreateTemplate() string {
	fmt.Println("Start creating pass template...")
	// Generate PassKit Client object for Templates protocol.
	pkTemplatesClient := io.NewTemplatesClient(shared.Conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	// In order to create an event ticket, we need a pass template id which holds pass design data. Let's use the default pass template for now.
	defaultTemplateRequest := &io.DefaultTemplateRequest{
		Protocol: io.PassProtocol_EVENT_TICKETING,
		Revision: uint32(1),
	}

	defaultPassTemplate, err := pkTemplatesClient.GetDefaultTemplate(ctx, defaultTemplateRequest)
	if err != nil {
		log.Fatalf("Create event ticket template err: %v", err)
	}

	// If you use the default template, you need to set name, description and timezone because these fields are mandatory.
	defaultPassTemplate.Name = "quick_start"
	defaultPassTemplate.Description = "quick start sample template"
	defaultPassTemplate.Timezone = "Europe/London"

	templateId, err := pkTemplatesClient.CreateTemplate(ctx, defaultPassTemplate)
	if templateId == nil || err != nil {
		log.Fatalf("Could not create event ticket template: %v", err)
	}
	return templateId.Id
}
