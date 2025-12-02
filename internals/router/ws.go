package router

import (
	"github.com/labstack/echo/v4"
	"github.com/suhas-developer07/GuessVibe-Server/internals/ws"
)


func LoadWSRoutes(e *echo.Echo,hub *ws.Hub) {

	e.GET("/ws",ws.WebsocketHandler(hub))
}