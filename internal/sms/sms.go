package sms

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

// This function sends an OTP to your phone number
func SendOtp(to string) {
	accountSid := "ACf0a6aa5e36ea840310faf30145b84249"
	authToken := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
	verifyServiceSid := "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(verifyServiceSid, params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
}

// This function waits for you to input the OTP sent to your phone,
// and validates that the code is approved
func CheckOtp(to string, client twilio.RestClient, verifyServiceSid string) {
	var code string
	fmt.Println("Please check your phone and enter the code:")
	fmt.Scanln(&code)

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(verifyServiceSid, params)
	if err != nil {
		fmt.Println(err.Error())
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}
}

/* func main() {
   to := "<your phone number here>"

   sendOtp(to)
   checkOtp(to)
} */
