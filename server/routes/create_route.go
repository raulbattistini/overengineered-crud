package routes

import (
	"fmt"
	"io"
	"server/api"
	"server/enums"
	"server/hepers"

	echo "github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		hepers.Log("error reading request body", &err, enums.Error)
		return err
	}
	response, err := api.CreateAPost(body)
	responseCode := response.ResponseCode()
	responseStatus := response.ResponseStatus()

	switch responseCode {
	case enums.BadRequestInvalidBody:
		hepers.Log("error reading request body", &err, enums.Error)
		return c.JSON(responseStatus, response)
	case enums.Created:
		hepers.Log(fmt.Sprintf("created post with body %s", string(body)), nil, enums.Info)
		return c.JSON(responseStatus, response)
	case enums.AlreadyExists:
		hepers.Log(fmt.Sprintf("post with body %s had a duplicat contente", string(body)), &err, enums.Error)
		return c.JSON(responseStatus, response)
	case enums.InternalServerError:
		hepers.Log("internal server error when creating a new post", nil, enums.Error)
		return c.JSON(responseStatus, response)
	default:
		hepers.Log(fmt.Sprintf("created post with body %s", string(body)), nil, enums.Info)
		return c.JSON(responseStatus, response)
	}
}
