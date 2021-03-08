package main

import (
	"fmt"
	"os/user"
	"path/filepath"

	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

type Regenesis struct {
	app *simple.App
	simple.SceneStack
	BookLocation string
}

func getDefaultBookLocation() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("while getting home directory: %w", err)
	}
	return filepath.Join(user.HomeDir, "Books"), nil
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
	var err error
	r.BookLocation, err = getDefaultBookLocation()
	if err != nil {
		panic(err)
	}

	r.app = simple.NewApp(r)

	err = r.app.RunForever()
	if err != nil {
		panic(err)
	}
}
