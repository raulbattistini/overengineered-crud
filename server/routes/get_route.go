package routes

import (
	"fmt"

	"server/api"
	"server/enums"
	"server/hepers"

	echo "github.com/labstack/echo/v4"
)

func GetPostById(c echo.Context) error {
	id := c.Param("id")

	response, err := api.GetPostIdHandler(id)

	responseCode := response.ResponseCode()
	responseStatus := response.ResponseStatus()

	switch responseCode {
	case enums.NoIdProvided:
		hepers.Log("no id provided to find post", &err, enums.Error)
		return c.JSON(responseStatus, response)
	case enums.NotFound:
		hepers.Log(fmt.Sprintf("the id %s was not found on thisserver", id), &err, enums.Error)
		return c.JSON(responseStatus, response)
	case enums.Success:
		hepers.Log(fmt.Sprintf("debugging clause for get post by id:\n %v", id), nil, enums.Info)
		return c.JSON(responseStatus, response)
	case enums.InternalServerError:
		hepers.Log(fmt.Sprintf(`error when founding post with id %s`, id), &err, enums.Error)
		return c.JSON(responseStatus, response)
	default:
		hepers.Log(fmt.Sprintf("interal server error when founding post with id %s", id), &err, enums.Error)
		return c.JSON(responseStatus, response)
	}
}
