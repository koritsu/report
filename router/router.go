package router

import (
	"github.com/labstack/echo/v4"
	"restapi-go/api"
)

func CreateRouting(e *echo.Echo) {
	// example
	// e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	// })

	// test
	e.GET("/", api.Hello)

	e.POST("/v1/login", api.Login)
	e.GET("/v1/testData", api.GetTestData)
	e.GET("/v1/user/", api.GetUser)
}
