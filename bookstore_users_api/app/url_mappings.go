package app

import (
	"../controllers/ping"
	"../controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("users/:user_id", users.GetUser)
	// router.GET("users/search", controllers.SearchUser)
	router.POST("users", users.CreateUser)

	router.POST("/login", users.LoginUser)
}