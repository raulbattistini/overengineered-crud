package routes

import (
	"fmt"
	"io"
	"server/api"
	"server/enums"
	"server/hepers"

	echo "github.com/labstack/echo/v4"
)

func UpdatePostById(c echo.Context) error {
	id := c.QueryParam("id")
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		hepers.Log("error reading request body", &err, enums.Error)
		return err
	}

	response, err := api.AtualizePostById(body, id)

	responseCode := response.ResponseCode()
	responseStatus := response.ResponseStatus()

	switch responseCode {
	case enums.NonAuthoritativeUpdated:
		hepers.Log(fmt.Sprintf("atualized pot with id %s to content %s", id, string(body)), nil, enums.Info)
		return c.JSON(responseStatus, response)
	default:
		return c.JSON(responseStatus, response)
	}
}
