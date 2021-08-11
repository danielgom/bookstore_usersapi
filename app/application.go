package app

import (
	"context"
	"fmt"
	"github.com/danielgom/bookstore_usersapi/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const (
	PORT = ":8081"
)

var (
	router = gin.Default()
)

func StartApplication() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Info(fmt.Sprintln("Listening and serving HTTP on ", PORT))
	mapUrls()

	srv := &http.Server{
		Addr:         PORT,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  2 * time.Second,
	}


	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("listen: ", err.Error())
		}
	}()

	<-ctx.Done()

	stop()

	log.Println("shutting down gracefully")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	logger.Info("Server exiting")
}
