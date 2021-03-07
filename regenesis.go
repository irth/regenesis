package main

import (
	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

type Screen interface {
	Scene() *simple.Scene
}

type Regenesis struct {
	app         *simple.App
	ScreenStack []Screen
}

func NewRegenesis() *Regenesis {
	r := &Regenesis{}
	r.ScreenStack = []Screen{
		NewCategorySelectScreen(r, []Category{
			{
				Name:     "Sci-Tech",
				Provider: &libgen.LibgenSearchProvider{},
			},
			{
				Name:     "Fiction",
				Provider: &libgen.FictionSearchProvider{},
			},
		}),
	}
	return r
}

func (r *Regenesis) Run() {
	r.app = simple.NewApp(r.ActiveScreen().Scene())

	err := r.app.RunForever()
	if err != nil {
		panic(err)
	}
}

func (r *Regenesis) ActiveScreen() Screen {
	return r.ScreenStack[len(r.ScreenStack)-1]
}

func (r *Regenesis) Push(screen Screen) {
	r.ScreenStack = append(r.ScreenStack, screen)
	if r.app != nil {
		r.app.NextScene(r.ActiveScreen().Scene())
	}
}

func (r *Regenesis) Pop() {
	if len(r.ScreenStack) == 1 {
		return
	}
	r.ScreenStack = r.ScreenStack[:len(r.ScreenStack)-1]
	if r.app != nil {
		r.app.NextScene(r.ActiveScreen().Scene())
	}
}
