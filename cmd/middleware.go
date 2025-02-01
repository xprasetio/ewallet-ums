package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}
	_, err := d.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("ailed to get user session on DB: ", err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}
	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("Unauthorized ", err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}
	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("Jwt token expired ", claim.ExpiresAt)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}
	ctx.Set("token", claim)
	ctx.Next()
}
