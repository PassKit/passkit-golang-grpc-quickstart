# passkit-golang-members-quickstart

__Important Disclaimer:__ _This repository is currently under development. Below instructions are meant as a reference for beta partners, and could be subject to changes. Please note that the environment mentioned in below Quickstart is our Development Environment, which should not be used in a production setting; this is purely for exploratory testing. Data created in our Development Environment can be purged at anytime at our discretion without notice._

The PassKit IO SDK makes it quick and easy to create and install your branded membership passes for Apple Wallet and Google Pay.


This repository has following structure with each purpose.
- `certs` folder is a place to store your credential files.
- `examples` folder contains SDK methods you can use to create membership cards and engage with members.

## Table of Content
* [Installation](#installation)
* [Prerequisites](#prerequisites)
* [Quickstart](#quickstart)
* [Examples](#examples)
* [GUI Tool](#gui-tool)
* [Documentation](#documentation)
* [Check Other Passes](#check-other-passes)
* [Getting Help](#getting-help)
* [License](#license)

## Installation
Install passkit-io-go with:
```go
go get -u github.com/PassKit/passkit-golang-sdk
```
Then, import SDK with:
```go
import(
    "github.com/PassKit/passkit-golang-sdk/io/members"
    "github.com/PassKit/passkit-golang-sdk/io"
)
```
## Prerequisites
1. Create a PassKit account. Sign up for free [HERE](https://dev-app.passkit.io/).

2. Download three `.pem` files you received by email after sign up. 
   
   To re-generate credentials, visit Settings (click gear icon in top right of the PassKit.IO web app window) > Developer Credential page and click 'Generate' (Login [HERE](https://dev-app.passkit.io/)).
   
## Quickstart
By completing this Quickstart, you will be able to issue a membership card for a new member.

1. Ensure your followed the steps in [prerequisites](#prerequisites).

2. Install passkit-io-go with:
   ```go
   go get -u github.com/PassKit/passkit-golang-sdk
   ```

3. When you created an account (Prerequisites #1), you should have received 3 files: `certificate.pem`, `key.pem` and `ca-chain.pem` by email. Please save those 3 files under `passkit-golang-members-quickstart/certs` directory. These .pem files are required to authenticate your accesss to PassKit.IO server.

4. Now we need to decrypt your `key.pem`. At your project root directory, run `cd ./certs openssl ec -in key.pem -out key.pem`. Your `key.pem` file should look like below.
   ![ScreenShot](https://raw.githubusercontent.com/PassKit/passkit-golang-members-quickstart/master/images/decrypted_key_pem.png)
   If you do not see `Proc-Type: 4,ENCEYPTED` on line 2, you have successfully decrypted `key.pem`. 

5. Replace `YOUR_EMAIL_ADDRESS@EMAIL.COM` in `main.go` with your email address in order to receive the welcome email with card url which your member will also receive.

6. Go back to root directory with `cd ../..`. Then run `go run main.go` to create a sample program (with default template & tier) and issue a membership card against that.

## Examples
#### Issue A Membership Card.
Follow the steps of the [Quickstart](#quickstart) to create a sample membership card and experience it in your Mobile Wallet.

#### Engage With Your Members
`EngageWithMembers()` contains multiple methods you can use to engage with your members. 
For example, you can update contents of digital membership card or send a push notification.

## GUI Tool
GUI tool can be accessed from [your PassKit.IO account](https://dev-app.passkit.io/login).

## Documentation
* [PassKit.IO Membership Official Documentation](https://docs.passkit.io/protocols/member)

## Check Other Passes
* Coupons (coming soon)
* Flight Ticket (coming soon)

## Getting Help
* Email [support@passkit.com](email:support@passkit.com)
* [Online chat support](https://passkit.com/)

## License
Distributed under MIT License. Details available on [license file](#).