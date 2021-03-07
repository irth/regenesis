package main

import (
	"github.com/irth/regenesis/pkg/simple"
)

func main() {
	butt := simple.Button{
		Position: simple.Position{
			X:      simple.Absolute(10),
			Y:      simple.Absolute(10),
			Width:  simple.Percent(30),
			Height: simple.Absolute(30),
		},
		ID:   "mybutton",
		Name: "myname",
	}

	scene := simple.Scene{
		Widgets: []simple.Widget{butt},
	}
	app := simple.NewApp(&scene)
	err := app.Render()
	if err != nil {
		panic(err)
	}
}
