package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game settings
const (
	screenWidth  = 500
	screenHeight = 500

	cellsWide = 5
	cellsHigh = 5

	cellWidth  = screenWidth / cellsWide
	cellHeight = screenHeight / cellsHigh
)

type Vec2i struct {
	X int
	Y int
}

func NewVec2i(x, y int) Vec2i {
	return Vec2i{X: x, Y: y}
}

func (v Vec2i) ToVector2() rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}

type Player struct {
	pos Vec2i
}

type Game struct {
	width  int
	height int
	player Player
}

func (g *Game) Init() {

}

func (g *Game) Update() {

}

func (g *Game) Draw() {

	rl.DrawRectangleV(g.player.pos.ToVector2(), rl.NewVector2(cellWidth, cellHeight), rl.Black)

	// We need to fix this
	rl.DrawTextEx(rl.GetFontDefault(), "Abcde", g.player.pos.ToVector2(), cellHeight, cellWidth-60, rl.Red)

}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Vimcleaning")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var game Game

	game.Init()

	for !rl.WindowShouldClose() {

		game.Update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		game.Draw()

		rl.EndDrawing()
	}
}
