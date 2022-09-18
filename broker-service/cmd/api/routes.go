package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() *gin.Engine {
	mux := gin.New()
	mux.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		AllowWildcard:    true,
	}))

	mux.GET("/ping", heartBeat)
	mux.HEAD("/ping", heartBeat)
	mux.POST("/", app.Broker)

	return mux
}

func heartBeat(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
