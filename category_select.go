package main

import (
	"fmt"

	libgen "github.com/irth/golibgen"
	ui "github.com/irth/regenesis/pkg/simple"
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

func NewCategorySelectScreen(r *Regenesis, categories []Category) ui.Scene {
	return &CategorySelectScreen{r, categories}
}

func (c *CategorySelectScreen) clickHandler(a *ui.App, cat Category) error {
	categoryHomeScreen := NewCategoryHomeScreen(c.r, cat)
	c.r.Push(categoryHomeScreen)
	return nil
}

func (c *CategorySelectScreen) categoryWidget(id string, category Category) ui.Widget {
	return ui.Button(
		id,
		ui.Pos(ui.Abs(100), ui.Step, ui.Percent(100), ui.Abs(100)),
		category.Name,
		func(a *ui.App, b *ui.ButtonWidget) error {
			return c.clickHandler(a, category)
		},
	)
}

func (c *CategorySelectScreen) Render() (ui.Widget, error) {
	categories := ui.WidgetList{}
	for idx, cat := range c.Categories {
		categories = append(categories, c.categoryWidget(fmt.Sprintf("category_%d", idx), cat))
	}

	return ui.WidgetList{
		Header(),
		ui.Label(
			ui.Pos(ui.Abs(100), ui.Abs(300), ui.Percent(100), ui.Abs(100)),
			"Choose a category",
		),
		ui.FontSize(64),
		categories,
	}, nil
}
