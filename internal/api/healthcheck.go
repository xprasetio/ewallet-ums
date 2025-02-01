package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	HealthcheckServices interfaces.IHealthcheckServices
}

func (api *Healthcheck) HealthcheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	helpers.SendResponseHttp(c, http.StatusOK, msg, nil)
}