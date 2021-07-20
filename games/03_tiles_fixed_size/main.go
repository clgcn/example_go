package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	imgSize   = 16
	numFrames = 8
)

type Game struct {
	tick  float64
	speed float64
}

var coins *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(coinImg))
	if err != nil {
		log.Fatal(err)
	}
	coins = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	g.tick++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	frameNum := int(g.tick/g.speed) % numFrames
	frameX := frameNum * imgSize
	subImg := coins.SubImage(image.Rect(frameX, 0, frameX+imgSize, imgSize)).(*ebiten.Image)
	screen.DrawImage(subImg, op)

}

func (g *Game) Layout(w, h int) (screenWidth, screenHeight int) {
	return w / 2, h / 2
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Draw tiles")

	g := &Game{
		speed: 60 / 6,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
