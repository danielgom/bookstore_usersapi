package app

import (
	"github.com/danielgom/bookstore_usersapi/logger"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {

	logger.Info("About to start the application")

	mapUrls()

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
