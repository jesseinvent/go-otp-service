package api

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/otp", app.SendSMS())
	app.Router.POST("/verify-otp", app.VerifySMS())
}
