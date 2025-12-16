package controllers

import "github.com/satyamdash/LLD-Golang/DesignPatterns/BookMyshow/model"

type TheatreController struct {
	CityTheatreMap map[string][]*model.Theatre
}

func NewTheatreController() *TheatreController {
	return &TheatreController{
		CityTheatreMap: make(map[string][]*model.Theatre),
	}
}

func (tc *TheatreController) AddTheatre(city string, theatre *model.Theatre) {
	tc.CityTheatreMap[city] = append(tc.CityTheatreMap[city], theatre)
}

func (tc *TheatreController) GetTheatresByCity(city string) []*model.Theatre {
	return tc.CityTheatreMap[city]
}
