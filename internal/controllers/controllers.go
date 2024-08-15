package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matiaseiglesias/storiChallenge/internal/services"
)

type Controllers struct {
	services     *services.Services
	Transactions *TransactionsController
}

func CreateControllers(services *services.Services, router *gin.Engine) *Controllers {
	return &Controllers{
		services:     services,
		Transactions: CreateTransactionsController(services, router),
	}
}
