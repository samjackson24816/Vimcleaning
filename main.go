package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game settings

type GameSettings struct {
	screenWidth  int32
	screenHeight int32

	cellsWide int16
	cellsHigh int16

	font rl.Font

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

func (g *Game) Draw(gs GameSettings) {

	rl.DrawRectangleV(gs.WorldToScreen(g.player.pos), rl.NewVector2(gs.CellWidth(), gs.CellHeight()), rl.Black)

	var allChars = []string{"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S"}

	for i := 0; i < 10; i++ {

		for j := 0; j < 10; j++ {
			pos := NewVec2i(i, j)
			screenPos := gs.WorldToScreen(pos)
			var character string = allChars[(i+j)%len(allChars)]

			var textSize = float32(gs.CellHeight()) * (1.0 - gs.textGapRatio)
			println("textSize: ", textSize)

			fontSize := rl.MeasureTextEx(gs.font, character, textSize, 0)
			println("fontSize: ", fontSize.X)

			textX := math.Round(float64(screenPos.X + (gs.CellWidth()-fontSize.X)/2))
			textY := math.Round(float64(screenPos.Y + (gs.CellHeight()-textSize)/2))

			rl.DrawRectangleV(screenPos, rl.NewVector2(gs.CellWidth(), gs.CellHeight()), rl.LightGray)

			rl.DrawTextEx(gs.font, character, rl.NewVector2(float32(textX), float32(textY)), textSize, 0, rl.Black)
		}
	}
}

func main() {

	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	gs := GameSettings{
		screenWidth:  500,
		screenHeight: 500,
		cellsWide:    5,
		cellsHigh:    5,
		font:         rl.GetFontDefault(),
		textGapRatio: 0.02,
	}

	rl.InitWindow(int32(gs.screenWidth), int32(gs.screenHeight), "Vimcleaning")

	defer rl.UnloadFont(gs.font)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var font = rl.LoadFontEx("JetBrainsMono-Bold.ttf", int32(math.Round(float64(gs.CellHeight()))), nil, 0)
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)

	gs.font = font

	var game Game

	game.Init()

	game.player.pos = NewVec2i(0, 0)

	for !rl.WindowShouldClose() {

		game.Update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		game.Draw(gs)

		rl.EndDrawing()
	}
}
