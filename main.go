package main

import (
	"fmt"
	"strings"

	"github.com/irth/regenesis/pkg/simple"
)

func main() {
	butt := &simple.Button{
		Position: simple.Position{
			X:      simple.Absolute(10),
			Y:      simple.Absolute(10),
			Width:  simple.Percent(30),
			Height: simple.Absolute(40),
		},
		ID:   "mybutton",
		Name: "myname1",
		OnClick: func(a *simple.App, b *simple.Button) error {
			fmt.Println("button clicked!")
			return nil
		},
	}

	butt2 := &simple.Button{
		Position: simple.Position{
			X:      simple.Same,
			Y:      simple.Step,
			Width:  simple.Percent(30),
			Height: simple.Absolute(40),
		},
		ID:   "mybutton2",
		Name: "myname2",
		OnClick: func(a *simple.App, b *simple.Button) error {
			fmt.Println("second button clicked!")
			return nil
		},
	}

	input := &simple.TextInput{
		Position: simple.Position{
			X:      simple.Same,
			Y:      simple.Step,
			Width:  simple.Percent(30),
			Height: simple.Absolute(40),
		},
		ID: "textinput",
		OnUpdate: func(a *simple.App, t *simple.TextInput) error {
			t.Value = strings.ToTitle(t.Value)
			return nil
		},
	}

	scene := simple.Scene{
		Widgets: []simple.Widget{butt, butt2, input},
	}

	app := simple.NewApp(&scene)

	err := app.RunForever()
	if err != nil {
		panic(err)
	}
}
