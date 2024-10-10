package routes

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func ListPosts(c echo.Context) error {
	return c.String(http.StatusOK, "List al potss")
}
