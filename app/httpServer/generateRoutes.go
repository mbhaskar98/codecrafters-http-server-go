package httpServer

import (
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/handlers"
)

func GenerateRoutes() map[string]handlers.IHandler {
	routes := make(map[string]handlers.IHandler)
	routes["/echo/"] = handlers.NewEchoRequestHandler()
	routes["/"] = handlers.NewDefaultRequestHomeHandler()
	routes["/user-agent"] = handlers.NewUserAgentRequestHandler()
	return routes
}
