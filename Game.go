package main

import (
	"bufio"
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	pos           Vec2i
	heldCharacter string
}

func NewPlayer(startPos Vec2i) Player {
	return Player{
		pos:           startPos,
		heldCharacter: " ",
	}
}

type Game struct {
	gs         GameSettings
	cellVals   [][]string
	word       string
	wordLayout []string
	player     Player
}

func (g *Game) Init(gs GameSettings) {
	g.gs = gs

	centerPos := NewVec2i(gs.cellsWide/2, gs.cellsHigh/2)
	g.player = NewPlayer(centerPos)

	g.loadWord()

	// We are making an array the width of the top row that has the word in the middle
	g.wordLayout = make([]string, g.gs.cellsWide)

	centerRoundedDown := int(math.Floor(float64(gs.cellsWide) / 2.0))

	wordHalfLengthRoundedUp := int(math.Ceil(float64(len(g.word)) / 2.0))

	offset := centerRoundedDown - wordHalfLengthRoundedUp

	for i := 0; i < len(g.wordLayout); i++ {
		g.wordLayout[i] = " "
	}

	for i := 0; i < len(g.word); i++ {
		g.wordLayout[offset+i] = string(g.word[i])
	}

	// Make array
	g.cellVals = make([][]string, g.gs.cellsWide)
	for i := range g.cellVals {
		g.cellVals[i] = make([]string, g.gs.cellsHigh)
	}

	charsToPlace := make([]string, 0, len(g.word))
	println("Word to place: ", g.word)

	for _, char := range g.word {
		charsToPlace = append(charsToPlace, string(char))
	}

	println("charsToPlace: ", charsToPlace)

	// Fill array
	for i := 0; i < g.gs.cellsWide; i++ {
		for j := 0; j < g.gs.cellsHigh; j++ {

			g.cellVals[i][j] = " "

		}
	}

	// Place characters in random positions
	for len(charsToPlace) > 0 {

		randX := rl.GetRandomValue(0, int32(g.gs.cellsWide-1))
		randY := rl.GetRandomValue(1, int32(g.gs.cellsHigh-1)) // Dont spawn on the top row

		if g.cellVals[randX][randY] == " " {
			g.cellVals[randX][randY] = charsToPlace[0]
			charsToPlace = charsToPlace[1:]
		}
	}

}

func (g *Game) loadWord() {
	f, err := os.Open(g.gs.wordFilePath)
	if err != nil {
		println("Error opening file:", err)
		return
	}
	print("Loading words from file: ", g.gs.wordFilePath)
	// remember to close the file at the end of the program
	defer f.Close()

	wordArray := make([]string, 0, 1000) // allocate 1000 elements as a guess

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		wordArray = append(wordArray, scanner.Text())
	}

	scanner.Err()

	wordCount := len(wordArray)

	randIdx := rl.GetRandomValue(0, int32(wordCount-1))

	g.word = wordArray[randIdx]

	println("Loaded word: ", g.word)
}

func (g *Game) Update() {
	g.updatePlayerMovement()
	g.updateWin()
}

func (g *Game) updatePlayerMovement() {

	// HJKL movement of player

	if rl.IsKeyPressed(rl.KeyH) {
		g.player.pos.X--
		if g.player.pos.X < 0 {
			g.player.pos.X = 0
		}
	} else if rl.IsKeyPressed(rl.KeyL) {
		g.player.pos.X++
		if g.player.pos.X >= g.gs.cellsWide {
			g.player.pos.X = g.gs.cellsWide - 1
		}
	}

	if rl.IsKeyPressed(rl.KeyK) {
		g.player.pos.Y--
		if g.player.pos.Y < 0 {
			g.player.pos.Y = 0
		}
	} else if rl.IsKeyPressed(rl.KeyJ) {
		g.player.pos.Y++
		if g.player.pos.Y >= g.gs.cellsHigh {
			g.player.pos.Y = g.gs.cellsHigh - 1

		}
	}

	// Pick up and put down character with x an p

	if (rl.IsKeyPressed(rl.KeyX) || rl.IsKeyPressed(rl.KeyV)) && g.player.heldCharacter == " " {

		g.player.heldCharacter = g.cellVals[g.player.pos.X][g.player.pos.Y]
		g.cellVals[g.player.pos.X][g.player.pos.Y] = " "

	} else if (rl.IsKeyPressed(rl.KeyP) || rl.IsKeyPressed(rl.KeyV)) && g.player.heldCharacter != " " && g.cellVals[g.player.pos.X][g.player.pos.Y] == " " {
		g.cellVals[g.player.pos.X][g.player.pos.Y] = g.player.heldCharacter
		g.player.heldCharacter = " "
	}

}

func (g *Game) updateWin() {
	if g.checkWin() {

		if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace) {
			if g.checkWin() {
				println("Player won! Restarting game...")
				g.Init(g.gs)
			}
		}

	}

}

func (g *Game) checkWin() bool {

	for i := 0; i < g.gs.cellsWide; i++ {
		if g.wordLayout[i] != " " && g.cellVals[i][0] != g.wordLayout[i] {
			return false
		}
	}

	return true

}

func (g *Game) Draw() {

	var gs = g.gs

	playerRect := rl.NewRectangle(gs.WorldToScreen(g.player.pos).X, gs.WorldToScreen(g.player.pos).Y, gs.CellWidth(), gs.CellHeight())

	rl.DrawRectangleRounded(playerRect, 0.1, 0, rl.LightGray)

	for i := 0; i < gs.cellsWide; i++ {

		for j := 0; j < gs.cellsHigh; j++ {

			var character string
			var textColor rl.Color
			if g.player.pos.X == i && g.player.pos.Y == j && g.player.heldCharacter != " " {

				// if we are the held character, we are orange
				character = g.player.heldCharacter
				textColor = rl.Orange
			} else if j == 0 && g.wordLayout[i] != " " {
				// Inside the place we put letters

				if g.cellVals[i][j] == " " {
					// If the cell is empty, make it a greyed out version of what it should be
					character = g.wordLayout[i]
					textColor = rl.LightGray
				} else if g.cellVals[i][j] == string(g.wordLayout[i]) {
					// If the cell is not empty, and the character is right, we make it green
					character = g.cellVals[i][j]
					textColor = rl.Green
				} else {
					// If the cell is not empty, and the character is wrong, we make it red
					character = g.cellVals[i][j]
					textColor = rl.Red
				}
			} else {
				character = g.cellVals[i][j]
				textColor = rl.Black
			}

			pos := NewVec2i(i, j)
			screenPos := gs.WorldToScreen(pos)

			var textSize = float32(gs.CellHeight()) * (1.0 - gs.textGapRatio)

			fontSize := rl.MeasureTextEx(gs.font, character, textSize, 0)

			textX := math.Round(float64(screenPos.X + (gs.CellWidth()-fontSize.X)/2))
			textY := math.Round(float64(screenPos.Y + (gs.CellHeight()-textSize)/2))

			rl.DrawTextEx(gs.font, character, rl.NewVector2(float32(textX), float32(textY)), textSize, 0, textColor)
		}
	}

}
