package api

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct { 
	RegisterService interfaces.IRegisterService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	resp, err := api.RegisterService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("failed to register: ", err)
		helpers.SendResponseHttp(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, resp)
	return
}