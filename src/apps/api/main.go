package main

import (
	"fmt"
	"moviedle/src/apps/api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	api := NewAPI("0.0.0.0", 8080)
	api.Serve()
}

type API interface {
	Serve()
}
type api struct {
	host   string
	port   int
	server *echo.Echo
}

func (a *api) Serve() {
	a.loadRoutes()
	a.start()
}

func NewAPI(host string, port int) API {
	server := echo.New()
	return &api{host, port, server}
}

func (a *api) Start() error {
    return a.server.Start(fmt.Sprintf("%s:%d", a.host, a.port))
}

func (a *api) rootEndpoint() *echo.Group {
	return a.server.Group("/api")
}

func (a *api) loadRoutes() {
	manager := routes.New()
	manager.Load(a.rootEndpoint())
}

func (a *api) start() {
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	if err := a.server.Start(address); err != nil {
		a.server.Logger.Fatal(err)
	}
}