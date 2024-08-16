package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	customerrors "github.com/matiaseiglesias/storiChallenge/internal/custom_errors"
	"github.com/matiaseiglesias/storiChallenge/internal/dto"
	"github.com/matiaseiglesias/storiChallenge/internal/services"
)

type TransactionsController struct {
	emailSender  services.EmailSenderService
	transactions services.TransactionService
}

func CreateTransactionsController(services *services.Services, router *gin.Engine) *TransactionsController {
	var c = &TransactionsController{
		emailSender:  services.EmailSender,
		transactions: services.Transaction,
	}
	c.Register(router)
	return c
}

// swagger:route POST /transactions/summaries Transactions makeSummary
//
// Process the transaction file and send summary information to a user in the form of an email.
//
// This endpoint processes the transaction file associated with a given account and sends a summary
// of the transactions to the specified email address.
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// Responses:
//   200: Response
//   400: Response
//   500: Response

func (c *TransactionsController) makeSummary(ctx *gin.Context) {
	var summaryRequest dto.SummaryRequest
	if err := ctx.ShouldBind(&summaryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Status:  "Error",
			Message: "Error while parsing request info",
		})
		return
	}

	account := summaryRequest.Account
	email := summaryRequest.Email

	err := c.transactions.NotifyTransactionSummary(account, email)
	if err != nil {
		if _, ok := err.(*customerrors.EmptyFieldError); ok {
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		log.Println("Internal error:", err)
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Status:  "error",
			Message: "An internal error occurred while processing the request.",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  "success",
		Message: "Email sent successfully",
	})
}

func (c *TransactionsController) Register(router *gin.Engine) {
	userRoute := router.Group("/transactions")
	userRoute.POST("/summaries", c.makeSummary)
}
