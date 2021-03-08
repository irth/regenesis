package main

import (
	"github.com/irth/regenesis/pkg/simple"
	ui "github.com/irth/regenesis/pkg/simple"
)

func Header() simple.Widget {
	return simple.WidgetList{
		simple.Left,
		simple.FontSize(100),
		simple.Label(
			simple.Pos(simple.Abs(100), simple.Abs(100), simple.Percent(100), simple.Abs(50)),
			"LibGen",
		),
		simple.FontSize(32),
	}
}

func BackButton(r *Regenesis) simple.Widget {
	return simple.WidgetList{
		ui.FontSize(28),
		// padding
		simple.Label(
			simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(50)),
			" ",
		),
		// button
		simple.Button(
			"back",
			simple.Pos(simple.Abs(100), simple.Step, simple.Abs(100), simple.Abs(50)),
			"(back)",
			func(a *simple.App, b *simple.ButtonWidget) error { r.Pop(); return nil },
		),
	}
}
