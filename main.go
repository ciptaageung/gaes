package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"

	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.conf")

	if viper.ReadInConfig() != nil {
		e.Logger.Fatal(viper.ReadInConfig())
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "elevenia service authentication is ready"})
	})

	h := &handler{}

	e.POST("/auth/authenticate", h.authenticate)

	e.GET("/private", h.private, isAuthenticate)

	e.GET("/admin", h.private, isAuthenticate, isAdmin)

	e.POST("/token", h.token)

	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))

}
