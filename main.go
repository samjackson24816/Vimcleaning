package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)


// Game settings
const (
	screenWidth = 500 
	screenHeight = 500 

	cellsWide = 5
	cellsHigh = 5

	cellWidth = screenWidth / cellsWide
	cellHeight = screenHeight /  cellsHigh
)


type Player struct {
	pos rl.Vector2
}


type Game struct {
	width int
	height int
	player Player
}

func (g *Game) Init() {
	

}


func (g *Game) Update() {

}


func (g *Game) Draw() {

	rl.DrawRectangleV(g.player.pos, rl.NewVector2(cellWidth, cellHeight), rl.Black)


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
