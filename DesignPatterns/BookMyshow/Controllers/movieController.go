package controllers

import "github.com/satyamdash/LLD-Golang/DesignPatterns/BookMyshow/model"

type MovieController struct {
	Movies map[string]*model.Movie
}

func NewMovieController() *MovieController {
	return &MovieController{
		Movies: make(map[string]*model.Movie),
	}
}

func (mc *MovieController) AddMovie(movie *model.Movie) {
	mc.Movies[movie.ID] = movie
}

func (mc *MovieController) ListAllMovies() []*model.Movie {
	result := []*model.Movie{}
	for _, movie := range mc.Movies {
		result = append(result, movie)
	}
	return result
}
