package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"

	"github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Swagger Elevenia API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "elevenia service authentication is ready")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	h := &handler{}

	e.POST("/login", h.login)

	e.GET("/private", h.private, isLoggedIn)

	e.GET("/admin", h.private, isLoggedIn, isAdmin)

	e.POST("/token", h.token)

	e.Logger.Fatal(e.Start(":1323"))
}
