package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {
	dependency := depencyInjection()

	r := gin.Default()
	r.GET("/health", dependency.HealtCheckAPI.HealthCheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)

	err := r.Run(":" + helpers.GetEnv("PORT","8083"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Dependency struct { 
	HealtCheckAPI *api.Healthcheck
	RegisterAPI *api.RegisterHandler
	LoginAPI *api.LoginHandler
}

func depencyInjection() Dependency {
	healtcheckSvc := &services.Healthcheck{
	}
	healtcheckAPI := &api.Healthcheck{
		HealthcheckServices: healtcheckSvc,
	}

	userRepo := &repository.UserRepository {
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		UserRepo: userRepo,
	}	
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}
	loginSvc := &services.LoginService{ 
		UserRepo: userRepo,
	}
	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	return Dependency{
		HealtCheckAPI: healtcheckAPI,
		RegisterAPI: registerAPI,
		LoginAPI: loginAPI,
	}
}
