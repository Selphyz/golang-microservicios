package app

import (
	"bookstore/controllers/ping"
	"bookstore/controllers/users"
)
func mapUrls()  {
	router.GET("/ping", ping.Ping)
	router.GET("/users", users.CreateUser)
}