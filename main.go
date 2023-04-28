package main

import (
	_ "embed"
	"image/color"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
)

//go:embed JetBrainsMono-500.ttf
var j500 []byte

//go:embed Montserrat-300.ttf
var m300 []byte

func main() {

	fonts := make([]text.FontFace, 0)

	{
		face, err := opentype.Parse(j500)
		if err != nil {
			panic(err)
		}

		fonts = append(fonts, text.FontFace{
			Font: font.Font{Typeface: "JetBrainsMono", Weight: 500},
			Face: face,
		})
	}

	{
		face, err := opentype.Parse(m300)
		if err != nil {
			panic(err)
		}
		fonts = append(fonts, text.FontFace{
			Font: font.Font{Typeface: "Montserrat", Weight: 300},
			Face: face,
		})
	}

	shaper := text.NewShaper(fonts)

	w := app.NewWindow()
	ops := new(op.Ops)
	go func() {
		for evt := range w.Events() {
			switch e := evt.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(ops, e)

				r := op.Record(gtx.Ops)
				paint.ColorOp{Color: color.NRGBA{A: 255}}.Add(gtx.Ops)
				material := r.Stop()

				widget.Label{}.Layout(gtx, shaper, font.Font{Typeface: "Montserrat", Weight: 300}, 16, "Hello world!", material)

				e.Frame(ops)
			}
		}
	}()

	app.Main()
}
