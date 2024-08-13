package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) TwilioSendOtp(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}

	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(envSERVICESSID(), params)

	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func (app *Config) TwilioVerifyOtp(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESSID(), params)

	if err != nil {
		return err
	}

	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}
