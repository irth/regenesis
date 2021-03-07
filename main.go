package main

import (
	"fmt"

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
	var selectedProvider libgen.SearchProvider = nil

	clickHandler := func(a *simple.App, b *simple.Button) error {
		switch b.ID {
		case "category_libgen":
			selectedProvider = &libgen.LibgenSearchProvider{}
		case "category_fiction":
			selectedProvider = &libgen.FictionSearchProvider{}
		}
		if selectedProvider != nil {
			a.NextScene(&simple.Scene{
				[]simple.Widget{
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

	categorySelection := simple.Scene{
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
				clickHandler,
			),
			simple.NewButton(
				"category_fiction",
				simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(100)),
				"Fiction",
				clickHandler,
			),
		},
	}

	app := simple.NewApp(&categorySelection)

	err := app.RunForever()
	if err != nil {
		panic(err)
	}
}
