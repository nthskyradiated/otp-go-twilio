package main

import (
	"github.com/nthskyradiated/otp-go-twilio/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := api.Config{Router: router}

	api.Routes()

	router.Run(":8000")
}