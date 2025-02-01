package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {
	dependency := dependencyInject()

	r := gin.Default()
	r.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)
	userV1.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.Logout)

	err := r.Run(":" + helpers.GetEnv("PORT","8083"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Dependency struct { 
	UserRepository interfaces.IUserRepository
	HealthcheckAPI interfaces.IHealthcheckHandler
	RegisterAPI interfaces.IRegisterHandler
	LoginAPI interfaces.ILoginHandler
	LogoutAPI interfaces.ILogoutHandler
}

func dependencyInject() Dependency {
	healtcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
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

	logoutSvc := &services.LogoutService{
		UserRepo: userRepo,
	}
	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	return Dependency{
		UserRepository: userRepo,
		HealthcheckAPI: healthcheckAPI,
		RegisterAPI: registerAPI,
		LoginAPI: loginAPI,
		LogoutAPI: logoutAPI,
	}
}

