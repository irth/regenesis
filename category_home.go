package main

import (
	"fmt"

	libgen "github.com/irth/golibgen"
	ui "github.com/irth/regenesis/pkg/simple"
)

type CategoryHomeScreen struct {
	r           *Regenesis
	Category    Category
	Results     []libgen.Book
	SearchQuery string
}

func NewCategoryHomeScreen(r *Regenesis, category Category) ui.Scene {
	return &CategoryHomeScreen{r, category, []libgen.Book{}, ""}
}

func (c *CategoryHomeScreen) Render() (ui.Widget, error) {
	return ui.WidgetList{
		Header(),
		c.searchWidget(),
		c.bookListWidget(c.Results, 10),
		BackButton(c.r),
	}, nil
}

func (c *CategoryHomeScreen) searchInputHandler(a *ui.App, t *ui.TextInputWidget, newValue string) error {
	var err error
	c.SearchQuery = newValue
	c.Results, err = c.Category.Provider.Find(newValue)
	if err != nil {
		// TODO: display search err
		panic(err)
	}

	return nil
}

func (c *CategoryHomeScreen) searchWidget() ui.Widget {
	return ui.WidgetList{
		ui.FontSize(64),
		ui.Label(
			ui.Pos(ui.Abs(100), ui.Abs(300), ui.Percent(100), ui.Abs(100)),
			c.Category.Name,
		),
		ui.FontSize(48),
		ui.Label(
			ui.Pos(ui.Abs(150), ui.Step, ui.Percent(100), ui.Abs(100)),
			"Search",
		),
		ui.TextInput(
			"search",
			ui.Pos(ui.Step, ui.Same, ui.Percent(60), ui.Abs(50)),
			c.SearchQuery,
			c.searchInputHandler,
		),
	}
}

func (c *CategoryHomeScreen) bookListWidget(books []libgen.Book, maxResults int) ui.Widget {
	widgets := ui.WidgetList{
		ui.FontSize(32),
	}

	end := maxResults
	if len(books) < end {
		end = len(books)
	}

	for idx, book := range books[:end] {
		widgets = append(
			widgets,
			c.bookWidget(fmt.Sprintf("book_%d", idx), book),
		)
	}

	return widgets
}

func (c *CategoryHomeScreen) bookWidget(id string, book libgen.Book) ui.Widget {
	pos := ui.Pos(ui.Abs(150), ui.Step, ui.Percent(80), ui.Abs(50))
	return ui.WidgetList{
		ui.FontSize(32),
		ui.Button(
			id,
			pos,
			book.Title(),
			func(a *ui.App, b *ui.ButtonWidget) error {
				c.r.Push(NewBookScreen(c.r, book))
				return nil
			},
		),
		ui.FontSize(28),
		ui.Label(
			ui.Pos(ui.Abs(150), ui.Step, ui.Percent(80), ui.Abs(10)),
			fmt.Sprintf("%s · %s · %s · %s", book.Size(), book.Format(), book.Language(), book.Author()),
		),
	}
}
