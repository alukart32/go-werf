package main

import (
	"echo-reverse-service/internal/app/reverse"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/:wordToReverse", func(c echo.Context) error {
		return c.String(http.StatusOK, reverse.ReverseRunes(c.Param("wordToReverse")))
	})

	e.Logger.Fatal(e.Start(":9090"))
}
