package main

import (
	"github.com/labstack/echo/v4"
	handler "github.com/shotu/craft/handler"
)

func main() {

	e := echo.New()
	// api := e.Group("/api/v1", serverHeader)
	e.POST("/sl-game", handler.CreateBoard)
	e.GET("/sl-game/:id", handler.GetBoard)

	e.PUT("/sl-game/:board_id/players/:id", handler.RollDice)

	e.Logger.Fatal(e.Start(":8080"))

}
