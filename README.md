# passkit-golang-members-grpc-quickstart

The PassKit Golang SDK makes it quick and easy to create and install your branded membership passes for Apple Wallet and Google Pay.

This repository has following structure with each purpose.
- `certs` folder is a place to store your credential files.
- `examples/membership` folder contains SDK methods you can use to create membership cards and engage with members.
- `examples/coupons` folder contains SDK methods you can use to create coupons and engage with coupons.
- `examples/flights` folder contains SDK methods you can use to create boarding passes and engage with boarding passes.

## Table of Content
* [Prerequisites](#prerequisites)
* [Quickstart](#quickstart)
* [Examples](#examples)
* [GUI Tool](#gui-tool)
* [Documentation](#documentation)
* [Getting Help](#getting-help)
* [License](#license)

## Prerequisites
1. Create a PassKit account. Sign up for free [HERE](https://app.passkit.com/).

2. Generate & Download your SDK credentials by clicking the 'GENERATE NEW SDK CREDENTIALS' button from the Developer Tools page in the [portal](https://app.passkit.com/app/account/developer-tools).

3. Your Apple Wallet certificate id (for boarding passes only)
   
## Quickstart
By completing this Quickstart, you will be able to up and running with the PassKit SDK as quickly as possible.

1. Ensure your followed the steps in [prerequisites](#prerequisites).

2. Install PassKit Golang SDK with:
   ``` go
   go get -u github.com/PassKit/passkit-golang-sdk
   ```

3. Place your SDK credential files (`certificate.pem`, `key.pem` and `ca-chain.pem`) in the certs folder in this repoo. The SDK uses these .pem files to authenticate against the PassKit server.

4. Now we need to decrypt your `key.pem`. At your project root directory, run `cd ./certs`  `openssl ec -in key.pem -out key.pem`. Your `key.pem` file should look like below.
   ![ScreenShot](https://raw.githubusercontent.com/PassKit/passkit-golang-members-quickstart/master/images/decrypted_key_pem.png)
   If you do not see `Proc-Type: 4,ENCEYPTED` on line 2, you have successfully decrypted `key.pem`. 

5. Replace `YOUR_EMAIL_ADDRESS@EMAIL.COM` in `main.go` with your email address in order to receive the welcome email with card url which your member will also receive.

6. Go back to root directory with `cd ../..`. Then run `go mod tidy` , then `go run main.go` to create a sample membership card, coupon card and boarding pass (with default templates & tiers/offers) and issue them.

## Examples
###  Membership Cards
#### Issue A Membership Card.
Follow the steps of the [Quickstart](#quickstart) to create a sample membership card and experience it in your Mobile Wallet.

#### Engage With Your Members
`EngageWithMembers()` contains multiple methods you can use to engage with your members. 
For example, you can update contents of digital membership card or send a push notification.

###  Coupons
#### Issue A Coupon.
Follow the steps of the [Quickstart](#quickstart) to create a sample coupon card and experience it in your Mobile Wallet.

#### Engage With Coupons.
`EngageWithCoupons()` contains multiple methods you can use to engage with coupons.
For example, you can redeem a coupon or list all coupons under an offer.

### Boarding Passes
#### Issue A Boarding Pass.
Follow the steps of the [Quickstart](#quickstart) to create a sample boarding pass and experience it in your Mobile Wallet.

#### Engage With Boarding Passes.
`EngageWithBoardingPass()` contains multiple methods you can use to engage with boarding passes.

## GUI Tool
GUI tool can be accessed from [your PassKit account](https://app.passkit.com/login).

## Documentation
* [PassKit Membership Official Documentation](https://docs.passkit.io/protocols/member)
* [PassKit Coupons Official Documentation](https://docs.passkit.io/protocols/coupon)
* [PassKit Boarding Passes Official Documentation](https://docs.passkit.io/protocols/boarding)


## Getting Help
* Email [support@passkit.com](email:support@passkit.com)
* [Online chat support](https://passkit.com/)

## License
Distributed under MIT License. Details available on [license file](#).
