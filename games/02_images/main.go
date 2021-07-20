package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var coin *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(coinImg))
	if err != nil {
		log.Fatal(err)
	}
	coin = ebiten.NewImageFromImage(img)

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	w, h := coin.Size()
	x, y := screen.Size()
	tx := x/2 - w/2
	ty := y/2 - h/2

	op.GeoM.Translate(float64(tx), float64(ty))
	screen.DrawImage(coin, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Draw image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
