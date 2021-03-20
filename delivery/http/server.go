package http

import (
	"jar-service/app"
	"jar-service/domain"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseJarCode struct {
	JarCode string `json:"jarCode"`
}

type server struct {
	*app.App
}

func NewHandler(e *echo.Echo, app *app.App) {
	handler := &server{
		app,
	}
	e.POST("/jar", handler.AddJar)
	e.GET("/jar/:code", handler.GetJar)
}

func (s *server) AddJar(c echo.Context) error {

	return c.JSON(http.StatusOK, ResponseJarCode{JarCode: "123"})
}

func (s *server) GetJar(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound, domain.ErrRecordNotFound, gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
