package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matiaseiglesias/storiChallenge/internal/dto"
	"github.com/matiaseiglesias/storiChallenge/internal/services"
)

type TransactionsController struct {
	emailSender  *services.EmailSenderService
	transactions *services.TransactionsService
}

func CreateTransactionsController(services *services.Services, router *gin.Engine) *TransactionsController {
	var c = &TransactionsController{
		emailSender:  services.EmailSender,
		transactions: services.Transaction,
	}
	c.Register(router)
	return c
}

// swagger:route POST /transactions/ Transactions makeSummary
//
// Process the transaction file and send summary information to a user in the form of an email.
//
// Process the transaction file and send summary information to a user in the form of an email.
//
// Produces:
//   - application/json
//
// Responses:
//   - 200: SummaryResponse
func (c *TransactionsController) makeSummary(ctx *gin.Context) {
	// c.emailSender.CreateSummaryTemplate()
	c.emailSender.Send("matiaseiglesias@yahoo.com")
	name := "/home/matiasei/Documentos/storiChallenge/test.csv"
	test := c.transactions.CalculateSummary(name)
	log.Println(test)
	result := &dto.SummaryResponse{
		Status: "ok",
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *TransactionsController) Register(router *gin.Engine) {
	userRoute := router.Group("/transactions")
	userRoute.POST("/summaries", c.makeSummary)
}
