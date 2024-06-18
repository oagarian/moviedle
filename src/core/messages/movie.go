package messages

import()

const (
    Movie       = "Movie"
    MovieID     = "ID of a movie"
    MovieTitle  = "Title of a movie"
	MovieYear   = "Year of a movie"
	MovieSlogan = "Slogan of a movie"
	MovieResume = "Resume of a movie"

	MovieInvalidIDErrorMessage      = "The provided ID is invalid"
	MovieIDCannotBeEmptyMessage     = "The movie ID cannot be empty"
	MovieTitleCannotBeEmptyMessage  = "The movie title cannot be empty"
	MovieYearCannotBeEmptyMessage   = "The movie year cannot be empty"
	MovieSloganCannotBeEmptyMessage = "The movie slogan cannot be empty"
	MovieResumeCannotBeEmptyMessage = "The movie resume cannot be empty"
	MovieInvalidYearErrorMessage 	= "The provided year is invalid"
	MovieInvalidYearRangeMessage 	= "The provided year is out of range"
)
