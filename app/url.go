package app

import "github.com/letstalkdata/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
