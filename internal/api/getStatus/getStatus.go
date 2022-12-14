package GetStatusAPI

import (
	"log"
	"net/http"

	"github.com/kimulaco/trophy-comp-server/pkg/httputil"
	"github.com/kimulaco/trophy-comp-server/pkg/steamworks"
	"github.com/labstack/echo/v4"
)

type GetUserResponse struct {
	StatusCode int `json:"statusCode"`
}

func GetStatus(c echo.Context) error {
	s := steamworks.NewSteamworks()
	if s.InvalidEnv() {
		log.Print("STATUS_ENVIROMENT_ERROR: undefined steam API key")
		return c.JSON(httputil.NewError500("STATUS_ENVIROMENT_ERROR", ""))
	}

	return c.JSON(http.StatusOK, GetUserResponse{
		StatusCode: http.StatusOK,
	})
}
