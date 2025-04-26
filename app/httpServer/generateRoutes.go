package httpServer

import (
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/handlers"
)

func GenerateRoutes(config Config) map[string]handlers.IHandler {
	routes := make(map[string]handlers.IHandler)
	routes["/echo/"] = handlers.NewEchoRequestHandler()
	routes["/user-agent"] = handlers.NewUserAgentRequestHandler()
	routes["/files/"] = handlers.NewFileRequestHandler(config.Directory)
	return routes
}
