package main

import (
	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

type Regenesis struct {
	app *simple.App
	simple.SceneStack
}

func NewRegenesis() *Regenesis {
	r := &Regenesis{}
	r.Push(NewCategorySelectScreen(r, []Category{
		{
			Name:     "Sci-Tech",
			Provider: &libgen.LibgenSearchProvider{},
		},
		{
			Name:     "Fiction",
			Provider: &libgen.FictionSearchProvider{},
		},
	}))
	return r
}

func (r *Regenesis) Run() {
	r.app = simple.NewApp(r)

	err := r.app.RunForever()
	if err != nil {
		panic(err)
	}
}
