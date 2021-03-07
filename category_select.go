package main

import (
	"fmt"

	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

var selectedProvider libgen.SearchProvider = nil

type Category struct {
	Provider libgen.SearchProvider
	Name     string
	ID       string
}

type CategorySelectScreen struct {
	Categories []Category
}

func NewCategorySelectScreen(categories []Category) *CategorySelectScreen {
	return &CategorySelectScreen{categories}
}

func (c *CategorySelectScreen) clickHandler(a *simple.App, b *simple.Button) error {
	switch b.ID {
	case "category_libgen":
		selectedProvider = &libgen.LibgenSearchProvider{}
	case "category_fiction":
		selectedProvider = &libgen.FictionSearchProvider{}
	}
	if selectedProvider != nil {
		a.NextScene(&simple.Scene{
			Widgets: simple.WidgetList{
				Header(),
				simple.NewLabel(
					simple.Pos(simple.Abs(100), simple.Abs(300), simple.Percent(100), simple.Abs(100)),
					fmt.Sprintf("selected: %s", b.Name),
				),
			},
		})
	}
	return nil
}

func (c *CategorySelectScreen) Scene() *simple.Scene {
	return &simple.Scene{
		Widgets: []simple.Widget{
			Header(),
			simple.NewLabel(
				simple.Pos(simple.Abs(100), simple.Abs(300), simple.Percent(100), simple.Abs(100)),
				"Choose a category",
			),
			simple.FontSize(64),
			simple.NewButton(
				"category_libgen",
				simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(100)),
				"Sci-Tech",
				c.clickHandler,
			),
			simple.NewButton(
				"category_fiction",
				simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(100)),
				"Fiction",
				c.clickHandler,
			),
		},
	}
}
