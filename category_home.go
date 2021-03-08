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

func (c *CategoryHomeScreen) BookWidget(id string, b libgen.Book) simple.Widget {
	return simple.NewButton(
		id,
		simple.Pos(simple.Abs(150), simple.Step, simple.Percent(80), simple.Abs(50)),
		fmt.Sprintf("%s - %s", b.Author(), b.Title()),
		nil,
	)
}

func (c *CategoryHomeScreen) searchInputHandler(a *simple.App, t *simple.TextInput, newValue string) error {
	var err error
	c.SearchQuery = newValue
	c.Results, err = c.Category.Provider.Find(newValue)
	if err != nil {
		// TODO: display search err
		panic(err)
		return nil
	}

	c.r.Replace(c)
	return nil
}

func (c *CategoryHomeScreen) Render() (simple.Widget, error) {
	println("rendering home screen")
	books := simple.WidgetList{}
	end := 15
	if len(c.Results) < end {
		end = len(c.Results)
	}
	println("halko")
	for idx, book := range c.Results[:end] {
		println(book.Title())
		books = append(books, c.BookWidget(fmt.Sprintf("book_%d", idx), book))
	}
	print("\n[[[")
	co, _ := books.Render()
	print(co)
	println("]]]")

	return simple.WidgetList{
		Header(),
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
		simple.FontSize(32),
		books,
		simple.NewLabel(
			simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(50)),
			" ",
		),
		simple.NewButton(
			"back",
			simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(50)),
			"(back)",
			func(a *simple.App, b *simple.Button) error { c.r.Pop(); return nil },
		),
	}, nil
}
