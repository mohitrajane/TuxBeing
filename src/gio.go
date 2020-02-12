package main

import (
	"image/color"
	"log"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func ui(w *app.Window) error {
	gofont.Register()
	var list = &layout.List{
		Axis: layout.Vertical,
	}
	th := material.NewTheme()
	gtx := layout.NewContext(w.Queue())
	var button = new(widget.Button)
	var clicked = false
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			widgets := []func(){
				func() {
					var test1 = th.H3("Hello")
					test1.Alignment = text.Middle
					if clicked {
						test1.Color = color.RGBA{127, 255, 255, 0}
					} else {
						test1.Color = color.RGBA{0, 127, 0, 255}
					}
					test1.Layout(gtx)
				},
				func() {
					in := layout.UniformInset(unit.Dp(8))
					layout.Flex{Alignment: layout.Middle}.Layout(gtx, //what is use of Alignment:layout.Middle?????
						layout.Rigid(func() {
							in.Layout(gtx, func() {
								for button.Clicked(gtx) {
									clicked = !clicked
								}
								th.Button("Hide Text").Layout(gtx, button)
							})
						}),
					)
				},
				func() {
					th.Button("Button!")
				},
			}
			// l := th.H6("Hello")
			// btn := th.Button("button")
			// maroon := color.RGBA{127, 0, 0, 255}
			// l.Color = maroon
			// l.Alignment = text.Middle
			// btn.Layout(gtx, new(widget.Button))
			// l.Layout(gtx)
			list.Layout(gtx, len(widgets), func(i int) {
				layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
			})
			e.Frame(gtx.Ops)
		}
	}
}
func main() {

	go func() {
		w := app.NewWindow()
		if err := ui(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}
