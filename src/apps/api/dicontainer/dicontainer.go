package dicontainer

import (
	"moviedle/src/core/ports/primary"
	"moviedle/src/core/services"
	"moviedle/src/infra/repository/postgres"
)

func MovieService() primary.MoviePort {
	repo := postgres.NewMovieRepository()
	return services.NewMovieService(repo)
}