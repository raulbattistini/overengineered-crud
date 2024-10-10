package routes

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func RemovePost(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Removed post with id %s", id))
}
