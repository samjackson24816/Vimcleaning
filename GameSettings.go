package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameSettings struct {
	screenWidth  int32
	screenHeight int32

	cellsWide int
	cellsHigh int

	font rl.Font

	wordFilePath string

	textGapRatio float32
}

func (gs GameSettings) CellWidth() float32 {
	return float32(gs.screenWidth) / float32(gs.cellsWide)
}

func (gs GameSettings) CellHeight() float32 {
	return float32(gs.screenHeight) / float32(gs.cellsHigh)
}

func (gs GameSettings) WorldToScreen(v Vec2i) rl.Vector2 {
	return rl.NewVector2((float32(v.X))*gs.CellWidth(), (float32(v.Y))*gs.CellHeight())
}
