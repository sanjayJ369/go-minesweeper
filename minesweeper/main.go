package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	winWidth  = 450
	winHeight = 800
	tileSize  = 30
)

func main() {
	// inital game state
	game := &state{
		cols:  9,
		rows:  9,
		mines: 10,
	}
	game.reset()

	rl.InitWindow(int32(winWidth), int32(winHeight), "minesweeper")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if game.gameOver || game.gameWon {
			// center window - default
			centerWindow(winWidth, winHeight)
		} else {
			// center window based on game state dimensions
			centerWindow(game.getWidth(), game.getHeight())
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if game.gameWon {
			game.drawCongrats()
		} else if game.menu {
			game.drawMenu()
		} else {
			game.drawTiles()
			game.drawField()
		}

		rl.EndDrawing()
	}
}
