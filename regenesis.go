package main

import (
	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

type Screen interface {
	Scene() *simple.Scene
}

type Regenesis struct {
	CategorySelectScreen Screen
}

func NewRegenesis() *Regenesis {
	r := &Regenesis{}
	r.CategorySelectScreen = NewCategorySelectScreen(r, []Category{
		{
			Name:     "Sci-Tech",
			Provider: &libgen.LibgenSearchProvider{},
		},
		{
			Name:     "Fiction",
			Provider: &libgen.FictionSearchProvider{},
		},
	})
	return r
}

func (r *Regenesis) Run() {
	app := simple.NewApp(r.CategorySelectScreen.Scene())

	err := app.RunForever()
	if err != nil {
		panic(err)
	}
}
