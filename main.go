package main

import (
	"log"

	"passkit-golang-members-quickstart/examples"
)

/*
	Prerequisites to run this file:
		- store your certificate.pem in passkit-golang-members-quickstart/certs/ directory
		- store your key.pem in passkit-golang-members-quickstart/certs/ directory
		- store your ca-chain.pem in passkit-golang-members-quickstart/certs/ directory
		- decrypt your key.pem with `cd ./certs openssl ec -in key.pem -out key.pem` from project root
*/

const (
	// The address & port of the PassKit gRPC service.
	gRPCHost = "grpc-dev.passkit.io"
	gRPCPort = "443"

	// The location of your client certificates.
	clientCertFile = "certs/certificate.pem" // [Required] Please store your certificate.pem at ./certs directory. Your client certificate you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login).
	clientKeyFile  = "certs/key.pem"         // [Required] Please store your key.pem at ./certs directory. Your private key you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login). You need to decrypt the key before use. Check README.md for details.
	clientCAFile   = "certs/ca-chain.pem"    // [Required] Please store your ca-chain.pem at ./certs directory. The certificate chain you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login).

	emailAddressToReceiveSamplePassUrl = "" // [Required] Please set your email address to receive digital card url by email.
)

// These variables will be used by EngageWithMembers methods.
var programId string
var tierId string
var memberId string

func main() {
	if emailAddressToReceiveSamplePassUrl == "" {
		log.Fatal("Please set emailAddressToReceiveSamplePassUrl with your email address in main.go so that you can receive sample welcome email with digital card url.")
	}

	ConnectWithPasskitServer()

	IssueMembershipCard()

	EngageWithMembers()
}

// In order to use PassKit SDK, you need to establish the connection to the PassKit server first.
func ConnectWithPasskitServer() {
	examples.ConnectPasskitSdk(clientCertFile, clientKeyFile, clientCAFile, gRPCHost, gRPCPort)
}

// In order to create a membership card for your member, you need to take following process:
// 1. Create a program
// 2. Create a tier
// 3. Enrol a member (enrolling a member will automatically issue a membership card to your member).
func IssueMembershipCard() {
	newProgramId := examples.CreateProgram()

	newTierId := examples.CreateTier(newProgramId)

	newMemberId := examples.EnrolMember(newProgramId, newTierId, emailAddressToReceiveSamplePassUrl)

	// These variables will be used by EngageWithMembers().
	programId = newProgramId
	tierId = newTierId
	memberId = newMemberId
}

// EngageWithMembers show methods you can use to engage with loyalty members.
// When you execute EngageWithMembers method by itself, please establish connection with the server first by
func EngageWithMembers() {
	examples.GetSingleMember(memberId)

	examples.ListMembers(programId)

	examples.CountMembers(programId)

	// examples.SendWelcomeEmail(memberId) ::: under maintenance :::
	examples.UpdateMember_EmailAddress(memberId, tierId, programId)

	// earn points (coming soon)
	// use points (coming soon)
	// delete member (coming soon)
}
