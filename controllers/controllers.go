package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matiaseiglesias/storiChallenge/services"
)

type Controllers struct {
	services *services.Services
	Test     *TestController
}

func CreateControllers(services *services.Services, router *gin.Engine) *Controllers {
	return &Controllers{
		services: services,
		Test:     CreateTestController(router),
	}
}
