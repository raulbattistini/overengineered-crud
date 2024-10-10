package routes

import (
	"fmt"
	"io"
	"server/api"
	"server/hepers"
	"server/types"

	echo "github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		hepers.Log("error reading request body", &err, types.Error)
		return err
	}
	response, err := api.CreatePost(body)
	responseCode := response.ResponseCode()
	responseStatus := response.ResponseStatus()

	switch responseCode {
	case types.BadRequestInvalidBody:
		hepers.Log("error reading request body", &err, types.Error)
		return c.JSON(responseStatus, response)
	case types.Created:
		hepers.Log(fmt.Sprintf("created post with body %s", string(body)), nil, types.Info)
		return c.JSON(responseStatus, response)
	case types.AlreadyExists:
		hepers.Log(fmt.Sprintf("post with body %s had a duplicat contente", string(body)), &err, types.Error)
		return c.JSON(responseStatus, response)
	case types.InternalServerError:
		hepers.Log("internal server error when creating a new post", nil, types.Error)
		return c.JSON(responseStatus, response)
	default:
		hepers.Log(fmt.Sprintf("created post with body %s", string(body)), nil, types.Info)
		return c.JSON(responseStatus, response)
	}
}
