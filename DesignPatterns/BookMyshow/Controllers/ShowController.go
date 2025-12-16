package controllers

import "github.com/satyamdash/LLD-Golang/DesignPatterns/BookMyshow/model"

type ShowController struct {
	Shows map[string]*model.Show
}

func NewShowController() *ShowController {
	return &ShowController{
		Shows: make(map[string]*model.Show),
	}
}

func (sc *ShowController) AddShow(show *model.Show) {
	sc.Shows[show.ID] = show
}

func (sc *ShowController) GetShow(showID string) *model.Show {
	return sc.Shows[showID]
}
