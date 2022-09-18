package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Config) Broker(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted,gin.H{
		"Error":   false,
		"Message": "Hitted broker service",
	})
}
