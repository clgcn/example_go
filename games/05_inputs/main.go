package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	lastClickAT  time.Time
	currentColor int
	x, y         int
}

var colors = []color.Color{
	color.RGBA{0, 0, 0, 0xff},
	color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0, 0xff},
	color.RGBA{0xff, 0, 0xff, 0xff},
	color.RGBA{0, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

const debouncer = 100 * time.Millisecond

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && time.Since(g.lastClickAT) > debouncer {
		log.Printf("A pressed")
		g.lastClickAT = time.Now()
		g.currentColor = (g.currentColor + 1) % len(colors)
	}
	g.x, g.y = ebiten.CursorPosition()
	return nil
}

func (g *Game) Layout(x, y int) (int, int) {
	return x, y
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("press A to switch background color. Mouse position: %s", g.mousePosition())
	screen.Fill(colors[g.currentColor])
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) mousePosition() string {
	return fmt.Sprintf("(%d, %d)", g.x, g.y)
}

func main() {
	log.Println(time.Time{})
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("inputs")
	ebiten.RunGame(&Game{})
}
