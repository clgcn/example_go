package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	audioContext *audio.Context
	click        *audio.Player
)

const debouncer = 100 * time.Millisecond

type Game struct {
	lastClickAt time.Time
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) && time.Since(g.lastClickAt) > debouncer {
		log.Printf("A pressed")
		click.Rewind()
		click.Play()
		g.lastClickAt = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Click A to play a sound")
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func main() {
	oggF, _ := ebitenutil.OpenFile("ragtime.ogg")
	wavF, _ := ebitenutil.OpenFile("click.wav")
	audioContext = audio.NewContext(44100)

	oggS, _ := vorbis.DecodeWithSampleRate(audioContext.SampleRate(), oggF)
	sournd, _ := wav.DecodeWithSampleRate(audioContext.SampleRate(), wavF)

	s := audio.NewInfiniteLoop(oggS, oggS.Length())
	background, _ := audio.NewPlayer(audioContext, s)
	click, _ = audio.NewPlayer(audioContext, sournd)

	background.Play()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("sounds")
	ebiten.RunGame(&Game{})
}
