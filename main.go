package main

import (
	libgen "github.com/irth/golibgen"
	"github.com/irth/regenesis/pkg/simple"
)

func Header() simple.Widget {
	return simple.WidgetList{
		simple.Left,
		simple.FontSize(100),
		simple.NewLabel(
			simple.Pos(simple.Abs(100), simple.Abs(100), simple.Percent(100), simple.Abs(50)),
			"LibGen",
		),
		simple.FontSize(32),
	}
}

func main() {
	categorySelectScreen := NewCategorySelectScreen([]Category{
		{
			ID:       "libgen",
			Name:     "Sci-Tech",
			Provider: &libgen.LibgenSearchProvider{},
		},
		{
			ID:       "fiction",
			Name:     "Fiction",
			Provider: &libgen.FictionSearchProvider{},
		},
	})

	app := simple.NewApp(categorySelectScreen.Scene())

	err := app.RunForever()
	if err != nil {
		panic(err)
	}
}
