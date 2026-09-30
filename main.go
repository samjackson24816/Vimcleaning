package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	gs := GameSettings{
		screenWidth:  500,
		screenHeight: 500,
		cellsWide:    10,
		cellsHigh:    10,
		font:         rl.GetFontDefault(),
		textGapRatio: 0.02,
		wordFilePath: "words.txt",
	}

	rl.InitWindow(int32(gs.screenWidth), int32(gs.screenHeight), "Vimcleaning")

	defer rl.UnloadFont(gs.font)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var font = rl.LoadFontEx("JetBrainsMono-Bold.ttf", int32(math.Round(float64(gs.CellHeight()))), nil, 0)
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)

	gs.font = font

	var game Game

	game.Init(gs)

	game.player.pos = NewVec2i(0, 0)

	for !rl.WindowShouldClose() {

		game.Update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		game.Draw()

		rl.EndDrawing()
	}
}
