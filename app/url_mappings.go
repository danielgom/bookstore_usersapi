package app

import (
	"github.com/danielgom/bookstore_usersapi/controllers/ping"
	"github.com/danielgom/bookstore_usersapi/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)
	router.GET("/users/:userId", users.GetUser)
	router.POST("/users", users.CreateUser)
}
