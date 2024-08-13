package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matiaseiglesias/storiChallenge/controllers"
	"github.com/matiaseiglesias/storiChallenge/services"
)

type AppServer struct {
	Router      *gin.Engine
	Services    *services.Services
	Controllers *controllers.Controllers
	Host        string
}

func SetupApp() (*AppServer, error) {
	r := gin.Default()
	config_ := cors.DefaultConfig()
	config_.AllowAllOrigins = true
	r.Use(cors.New(config_))

	var services = services.CreateServices()
	var controllers = controllers.CreateControllers(services, r)

	var host = "localhost:8080"
	return &AppServer{
		Router:      r,
		Services:    services,
		Controllers: controllers,
		Host:        host,
	}, nil

}

func main() {

	app, _ := SetupApp()

	srv := &http.Server{
		Addr:    app.Host,
		Handler: app.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 10)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()

	log.Println("Server exiting")
}
