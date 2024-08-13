package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-otp-service/data"
)

const appTimeout = time.Second * 10

func (app *Config) SendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)

		defer cancel()

		var payload *data.OTPData

		err := app.ValidateBody(c, &payload)

		if err != nil {
			app.ErrorJSON(c, err)
		}

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err = app.TwilioSendOtp(newData.PhoneNumber)

		if err != nil {
			app.ErrorJSON(c, err)
			return
		}

		app.WriteJSON(c, http.StatusAccepted, "OTP send successful")
	}
}

func (app *Config) VerifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)

		var payload *data.VerifyData

		defer cancel()

		app.ValidateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.TwilioVerifyOtp(newData.User.PhoneNumber, newData.Code)

		fmt.Println("err: ", err)

		if err != nil {
			app.ErrorJSON(c, err)
			return
		}

		app.WriteJSON(c, http.StatusAccepted, "OTP successfully verified")
	}
}
