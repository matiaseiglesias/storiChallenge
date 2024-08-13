package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func CreateTestController(router *gin.Engine) *TestController {
	var c = &TestController{}
	c.Register(router)
	return c
}

// swagger:route GET /test/ Test hello_world
//
// Testing endpoint.
//
// Testing endpoint
//
//	Produces:
//	- application/json
//
//	Responses:
//	- 200: "Hello world"

func (c *TestController) getAll(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "Hello World")
}

func (results *TestController) Register(router *gin.Engine) {
	userRoute := router.Group("/test")
	userRoute.GET("/helloWorld", results.getAll)
}
