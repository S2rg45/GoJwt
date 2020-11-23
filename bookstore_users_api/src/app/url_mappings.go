package app

import (
	"../controllers/ping"
	"../controllers/users"
	"../middleware"
)

func mapUrlsToken() {
	router.POST("/login", users.LoginUser)
	mapUrls()
}

func mapUrls(){
	apiRoutes := router.Group("/API",  middleware.AuthorizeJWT())
	{
		apiRoutes.GET("/ping", ping.Ping)
		apiRoutes.GET("users/:user_id", users.GetUser)
		// apiRoutes.GET("users/search", controllers.SearchUser)
		apiRoutes.POST("users", users.CreateUser)
	}
}