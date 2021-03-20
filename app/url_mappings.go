package app

import (
	"github.com/danielgom/bookstore_usersapi/controllers/ping"
	"github.com/danielgom/bookstore_usersapi/controllers/users"
)

func mapUrls() {

	router.GET("/health", ping.Health)

	router.POST("/users", users.Create)
	router.GET("/users/:userId", users.GetById)
	router.PUT("/users/:userId", users.Update)
	router.PATCH("/users/:userId", users.UpdatePartial)
	router.DELETE("/users/:userId", users.Delete)
	router.POST("/users/login", users.Login)

	// These endpoints are going to be internal
	internal := router.Group("/internal")
	internal.GET("/users/search", users.FindByStatus)
}
