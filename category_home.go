package main

import (
	"fmt"

	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

type CategoryHomeScreen struct {
	r           *Regenesis
	Category    Category
	Results     []libgen.Book
	SearchQuery string
}

func NewCategoryHomeScreen(r *Regenesis, category Category) simple.Scene {
	return &CategoryHomeScreen{r, category, []libgen.Book{}, ""}
}

func (c *CategoryHomeScreen) Render() (simple.Widget, error) {
	return simple.WidgetList{
		Header(),
		c.searchWidget(),
		c.bookListWidget(c.Results, 15),
		BackButton(c.r),
	}, nil
}

func (c *CategoryHomeScreen) searchInputHandler(a *simple.App, t *simple.TextInput, newValue string) error {
	var err error
	c.SearchQuery = newValue
	c.Results, err = c.Category.Provider.Find(newValue)
	if err != nil {
		// TODO: display search err
		panic(err)
	}

	return nil
}

func (c *CategoryHomeScreen) searchWidget() simple.Widget {
	return simple.WidgetList{
		simple.FontSize(64),
		simple.NewLabel(
			simple.Pos(simple.Abs(100), simple.Abs(300), simple.Percent(100), simple.Abs(100)),
			c.Category.Name,
		),
		simple.FontSize(48),
		simple.NewLabel(
			simple.Pos(simple.Abs(150), simple.Step, simple.Percent(100), simple.Abs(100)),
			"Search",
		),
		simple.NewTextInput(
			"search",
			simple.Pos(simple.Step, simple.Same, simple.Percent(60), simple.Abs(50)),
			c.SearchQuery,
			c.searchInputHandler,
		),
	}
}

func (c *CategoryHomeScreen) bookListWidget(books []libgen.Book, maxResults int) simple.Widget {
	widgets := simple.WidgetList{
		simple.FontSize(32),
	}

	end := maxResults
	if len(books) < end {
		end = len(books)
	}

	for idx, book := range books[:end] {
		widget := simple.NewButton(
			fmt.Sprintf("book_%d", idx),
			simple.Pos(simple.Abs(150), simple.Step, simple.Percent(80), simple.Abs(50)),
			fmt.Sprintf("%s - %s", book.Author(), book.Title()),
			nil,
		)
		widgets = append(widgets, widget)
	}

	return widgets
}
