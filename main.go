package main

import (
	"embed"
	_ "embed"
	"fmt"
	"image/color"
	"io/fs"
	"strconv"
	"strings"

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

//go:embed *.ttf
var localFonts embed.FS

func main() {

	fonts := make([]text.FontFace, 0)

	fs.WalkDir(localFonts, ".", func(path string, d fs.DirEntry, err error) error {
		split := strings.Split(strings.Replace(strings.TrimRight(d.Name(), ".ttf"), "-", "_", -1), "_")
		if len(split) != 2 {
			return nil
		}

		name := split[0]
		weight, _ := strconv.ParseInt(split[1], 10, 32)
		fmt.Println(name)
		fmt.Println(weight)

		file, err := fs.ReadFile(localFonts, path)
		if err != nil {
			return err
		}

		face, err := opentype.Parse(file)
		if err != nil {
			return err
		}

		fonts = append(fonts, text.FontFace{
			Font: font.Font{
				Typeface: font.Typeface(name),
				Weight:   font.Weight(weight),
			},
			Face: face,
		})

		return nil
	})

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
