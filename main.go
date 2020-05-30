package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lytics/logrus"
	"net/http"
)


type config struct {
	PORT string `envconfig:"PORT"`
}

func (c *config) init() error {
	err := envconfig.Process("", c)
	return err
}

func main() {
	e := echo.New()

	logrus.SetLevel(logrus.DebugLevel)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://www.bellesdespins.com", "http://www.bellesdespins.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "ok")
	})

	e.GET("/health/", func(context echo.Context) error {
		return context.JSON(200, nil)
	})

	e.Logger.Fatal(e.Start(":1323"))
}