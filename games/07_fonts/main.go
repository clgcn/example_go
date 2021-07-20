package main

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	sampleText = `Hello, world!`
	dpi        = 72
	fontSize   = 36
)

type Game struct{}

var mplusNormalFont font.Face

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bounds := text.BoundString(mplusNormalFont, sampleText)
	text.Draw(screen, sampleText, mplusNormalFont, 10, bounds.Dy(), color.White)
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func init() {
	tt, err := truetype.Parse(Font)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("inputs")
	ebiten.RunGame(&Game{})
}
