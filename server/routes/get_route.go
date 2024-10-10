package routes

import (
	"fmt"

	"server/api"
	"server/hepers"
	"server/types"

	echo "github.com/labstack/echo/v4"
)

func GetPostById(c echo.Context) error {
	id := c.Param("id")

	response, err := api.GetPostIdHandler(id)

	responseCode := response.ResponseCode()
	responseStatus := response.ResponseStatus()

	switch responseCode {
	case types.NoIdProvided:
		hepers.Log("no id provided to find post", &err, types.Error)
		return c.JSON(responseStatus, response)
	case types.NotFound:
		hepers.Log(fmt.Sprintf("the id %s was not found on thisserver", id), &err, types.Error)
		return c.JSON(responseStatus, response)
	case types.Success:
		hepers.Log(fmt.Sprintf("debugging clause for get post by id:\n %v", id), nil, types.Info)
		return c.JSON(responseStatus, response)
	case types.InternalServerError:
		hepers.Log(fmt.Sprintf(`error when founding post with id %s`, id), &err, types.Error)
		return c.JSON(responseStatus, response)
	default:
		hepers.Log(fmt.Sprintf("interal server error when founding post with id %s", id), &err, types.Error)
		return c.JSON(responseStatus, response)
	}
}
