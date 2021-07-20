package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	w, h := screen.Size()
	i1 := ebiten.NewImage(w/2, h/2)
	i1.Fill(color.RGBA{0, 0xff, 0, 0xff})
	screen.DrawImage(i1, nil)
	i1w, i1h := i1.Size()

	i2 := ebiten.NewImage(w/3, h/3)
	i2.Fill(color.RGBA{0, 0, 0xff, 0x88})
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(i1w), float64(i1h))
	opts.GeoM.Rotate(0.5)
	opts.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(i2, &opts)
}

func (g Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 280)
	ebiten.SetWindowTitle("Colors and ImageOptions")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
