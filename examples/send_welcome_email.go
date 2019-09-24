package examples

import (
	"fmt"
	"log"

	"github.com/PassKit/passkit-golang-sdk/io"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// SendWelcomeEmail takes memberId and sends a welcome email (contains membership card url) to the member.
// Welcome email settings can be changed with updateProgram method.
func SendWelcomeEmail(memberId string) {
	fmt.Println("Sending a welcome email...")

	// Generate a distribution module client
	pkDistributionClient := io.NewDistributionClient(conn)

	// Generate context object to connect to the server.
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, nil)

	sendEmailRequest := &io.EmailDistributionRequest{
		Id:       memberId,
		Protocol: io.PassProtocol_MEMBERSHIP,
	}

	_, err := pkDistributionClient.SendWelcomeEmail(ctx, sendEmailRequest)
	if err != nil {
		log.Fatalf("Send welcome email err: %v", err)
	}

	fmt.Println("Successfully sent a welcome email to the member.")
}
