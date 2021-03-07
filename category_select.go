package main

import (
	"fmt"

	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

type Category struct {
	Provider libgen.SearchProvider
	Name     string
	ID       string
}

type CategorySelectScreen struct {
	r          *Regenesis
	Categories []Category
}

func NewCategorySelectScreen(r *Regenesis, categories []Category) Screen {
	return &CategorySelectScreen{r, categories}
}

func (c *CategorySelectScreen) clickHandler(a *simple.App, cat Category) error {
	categoryHomeScreen := NewCategoryHomeScreen(c.r, cat)
	a.NextScene(categoryHomeScreen.Scene())
	return nil
}

func (c *CategorySelectScreen) categoryWidget(id string, category Category) simple.Widget {
	return simple.NewButton(
		id,
		simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(100)),
		category.Name,
		func(a *simple.App, b *simple.Button) error {
			return c.clickHandler(a, category)
		},
	)
}

func (c *CategorySelectScreen) Scene() *simple.Scene {
	categories := simple.WidgetList{}
	for idx, cat := range c.Categories {
		categories = append(categories, c.categoryWidget(fmt.Sprintf("category_%d", idx), cat))
	}

	return &simple.Scene{
		Widgets: []simple.Widget{
			Header(),
			simple.NewLabel(
				simple.Pos(simple.Abs(100), simple.Abs(300), simple.Percent(100), simple.Abs(100)),
				"Choose a category",
			),
			simple.FontSize(64),
			categories,
		},
	}
}
