package routes

import (
	"moviedle/src/apps/api/dicontainer"
	"moviedle/src/apps/api/handlers"
	"github.com/labstack/echo/v4"
)

type movieRouter struct {
	handler handlers.MovieHander
}

func NewMovieRouter() Router {
	service := dicontainer.MovieService()
	handler := handlers.NewMovieHandler(service)
	return &movieRouter{handler}
}

func (r *movieRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/movies")
	router.GET("/:id", r.handler.Get)
}