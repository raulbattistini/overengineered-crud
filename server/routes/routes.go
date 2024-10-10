package routes

import (
	_ "fmt"

	echo "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/users/all", ListPosts)
	e.GET("/users/user/:id", GetPostById)
	e.POST("/users/create", CreatePost)
	e.DELETE("/users/delete/:id", RemovePost)
	e.PATCH("/users/atualize-fields/:id", UpdatePostById)
}
