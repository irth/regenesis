package main

import "github.com/irth/regenesis/pkg/simple"

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

func BackButton(r *Regenesis) simple.Widget {
	return simple.WidgetList{
		// padding
		simple.NewLabel(
			simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(50)),
			" ",
		),
		// button
		simple.NewButton(
			"back",
			simple.Pos(simple.Abs(100), simple.Step, simple.Percent(100), simple.Abs(50)),
			"(back)",
			func(a *simple.App, b *simple.Button) error { r.Pop(); return nil },
		),
	}
}
