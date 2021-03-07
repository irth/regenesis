package main

import (
	"github.com/irth/regenesis/pkg/simple"
)

type CategoryHomeScreen struct {
	r        *Regenesis
	Category Category
}

func NewCategoryHomeScreen(r *Regenesis, category Category) Screen {
	return &CategoryHomeScreen{r, category}
}

func (c *CategoryHomeScreen) clickHandler(a *simple.App, b *simple.Button) error {
	return nil
}

func (c *CategoryHomeScreen) Scene() *simple.Scene {
	return &simple.Scene{
		Widgets: []simple.Widget{
			Header(),
			simple.FontSize(64),
			simple.NewLabel(
				simple.Pos(simple.Abs(100), simple.Abs(300), simple.Percent(100), simple.Abs(100)),
				c.Category.Name,
			),
		},
	}
}
