package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {
	healtcheckSvc := &services.Healthcheck{
	}
	healtcheckAPI := &api.Healthcheck{
		HealthcheckServices: healtcheckSvc,
	}

	r := gin.New()
	r.GET("/health", healtcheckAPI.HealthCheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT","8083"))
	if err != nil {
		log.Fatal(err.Error())
	}
}