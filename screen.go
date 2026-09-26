package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Screen struct {
	width, height int
}

func InitScreen(width, height int) *Screen {
	return &Screen{
		width:  width,
		height: height,
	}
}

func (s *Screen) Update() error {
	return nil
}

func (s *Screen) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
}
