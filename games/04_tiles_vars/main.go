package main

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type framesSpec struct {
	Frames []frameSpec `json:"frames"`
}

type frameSpec struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type Game struct {
	tick      float64
	speed     float64
	frames    []frameSpec
	numFrames int
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

func (g *Game) Layout(w, h int) (int, int) {
	return w / 2, h / 2
}

func (g *Game) Draw(screen *ebiten.Image) {
	frameNum := int(g.tick/g.speed) % g.numFrames
	frame := g.frames[frameNum]
	subImg := coins.SubImage(image.Rect(frame.X, frame.Y, frame.X+frame.W, frame.H)).(*ebiten.Image)
	x, y := screen.Size()
	tx := x/2 - frame.W/2
	ty := y/2 - frame.H/2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tx), float64(ty))
	screen.DrawImage(subImg, op)

}

func (g *Game) buildFrames(path string) error {
	j, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fSpec := &framesSpec{}
	json.Unmarshal(j, fSpec)
	g.frames = fSpec.Frames
	g.numFrames = len(g.frames)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("excepting json spec path as command line argument")
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Draw tiles with differente size")

	g := &Game{
		speed: 60 / 12,
	}

	if err := g.buildFrames(os.Args[1]); err != nil {
		log.Fatal(err)
	}
	ebiten.RunGame(g)
}
