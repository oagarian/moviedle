package handlers

import (
	"moviedle/src/apps/api/handlers/dto/response"
	"moviedle/src/core/domain/errors"
	"moviedle/src/core/messages"
	"moviedle/src/core/ports/primary"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MovieHander interface {
	Get(echo.Context) error
	All(echo.Context) error
}

type movieHandler struct {
	service primary.MoviePort
}

func NewMovieHandler(service primary.MoviePort) MovieHander {
    return &movieHandler{service}
}

// Get
// @ID Movies.Get
// @Summary View movie details.
// @Description Allow a user to see details of a movie.
// @Security bearerAuth
// @Accept json
// @Param id path string true "Movie ID." default(a92adbf5-930c-41ce-bff3-b11798505f1c)
// @Tags Movie
// @Produce json
// @Success 200 {object} response.Movie "Success."
// @Failure 400 {object} response.ErrorMessage "Bad Request."
// @Failure 500 {object} response.ErrorMessage "Internal error."
// @Failure 503 {object} response.ErrorMessage "Database out of function."
// @Router /movies/{id} [get]
func (h *movieHandler) Get(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		badReqErr := errors.NewFromString(messages.MovieInvalidIDErrorMessage)
		return response.ErrorBuilder().NewFromDomain(ctx, badReqErr)
	} else if result, err := h.service.Get(&id); err != nil {
		return response.ErrorBuilder().NewFromDomain(ctx, err)
	} else {
		return ctx.JSON(http.StatusOK, response.NewMovieBuilder().BuildFromDomain(result))
	}
}


// All
// @ID Movies.All
// @Summary List all movies.
// @Description Allow a user list all movies in database.
// @Security bearerAuth
// @Accept json
// @Tags Movie
// @Produce json
// @Success 200 {object} response.Movie "Success."
// @Failure 400 {object} response.ErrorMessage "Bad Request."
// @Failure 500 {object} response.ErrorMessage "Internal error."
// @Failure 503 {object} response.ErrorMessage "Database out of function."
// @Router /movies [get]
func (h *movieHandler) All(ctx echo.Context) error {
	if result, err := h.service.All(); err!= nil {
        return response.ErrorBuilder().NewFromDomain(ctx, err)
    } else {
        return ctx.JSON(http.StatusOK, response.NewMovieBuilder().BuildFromDomainList(result))
    }
}