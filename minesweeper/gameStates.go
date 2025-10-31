package main

import (
	"fmt"
	"math/rand"
	"time"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// game state
type state struct {
	menu       bool
	gameWon    bool
	gameOver   bool
	rows       int
	cols       int
	mines      int
	tiles      [][]tile
	startedAt  time.Time
	finishedAt time.Time
}

type tile struct {
	isMine      bool
	minesAround int
	marked      bool
	open        bool
}

func getTextColor(neighbors int) rl.Color {
	switch neighbors {
	case 1:
		return rl.Blue
	case 2:
		return rl.Green
	case 3:
		return rl.Red
	default:
		return rl.Black
	}
}

func (s *state) revealTile(x, y int) {
	if s.tiles[x][y].open {
		return
	}

	s.tiles[x][y].open = true

	tile := s.tiles[x][y]

	if tile.isMine {
		s.gameOver = true
		s.finishedAt = time.Now()
		return
	}

	s.gameWon = s.isGameWon()

	if tile.minesAround == 0 {
		s.doForNeighbours(x, y, func(nx, ny int) {
			s.revealTile(nx, ny)
		})
	}

}

func (s *state) drawCongrats() {
	w := winWidth
	var lineHeight int32 = 50

	if s.gameWon {
		rl.DrawText("WELL DONE !", 0, lineHeight, int32(tileSize), rl.White)
	}

	clicked := gui.Button(rl.NewRectangle(0, float32(2*lineHeight), float32(w), float32(tileSize)), "PLAY AGAIN")
	if clicked {
		s.reset()
	}
}

func (s *state) drawTiles() {
	for x := 0; x < s.rows; x++ {
		for y := 0; y < s.cols; y++ {
			if s.gameOver {
				// reveal the tile
				var (
					text  string
					color rl.Color
				)

				tile := s.tiles[x][y]
				if tile.isMine {
					text = "*"
					color = rl.Red
				} else if tile.minesAround > 0 {
					text = fmt.Sprintf("%d", tile.minesAround)
					color = getTextColor(tile.minesAround)
				}

				rl.DrawText(text, int32(5+(x*tileSize)), int32(5+(y*tileSize)), int32(20), color)
				continue
			}

			// mark
			rect := rl.NewRectangle(float32(x*tileSize), float32(y*tileSize), float32(tileSize), float32(tileSize))

			if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
				if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
					if !s.tiles[x][y].open {
						s.tiles[x][y].marked = !s.tiles[x][y].marked
					}
				}
			}

			if s.tiles[x][y].marked {
				rl.DrawText("M", int32(5+(x*tileSize)), int32(5+(y*tileSize)), int32(20), rl.Violet)
			} else if s.tiles[x][y].open {
				text := ""
				if s.tiles[x][y].minesAround > 0 {
					text = fmt.Sprint(s.tiles[x][y].minesAround)
				}

				rl.DrawText(text,
					int32(5+(x*tileSize)),
					int32(5+(y*tileSize)),
					int32(20),
					getTextColor(s.tiles[x][y].minesAround))
			} else {
				if gui.Button(rect, "") {
					s.revealTile(x, y)
				}
			}
		}
	}
}

func (s *state) drawMenu() {
	baseX, baseY := float32(0), float32(50)
	spacing := float32(2 * tileSize)
	// difficualty levels
	// beginner
	if gui.Button(rl.NewRectangle(baseX, baseY, float32(winWidth), float32(tileSize)), "BEGINNER") {
		s.rows = 9
		s.cols = 9
		s.mines = 10
	}
	baseY += spacing

	// intermeidatte
	if gui.Button(rl.NewRectangle(baseX, baseY, float32(winWidth), float32(tileSize)), "INTERMEDIATE") {
		s.rows = 16
		s.cols = 16
		s.mines = 40
	}
	baseY += spacing

	// expert
	if gui.Button(rl.NewRectangle(baseX, baseY, float32(winWidth), float32(tileSize)), "EXPERT") {
		s.rows = 30
		s.cols = 30
		s.mines = 99
	}
	baseY += spacing

	// start game
	if gui.Button(rl.NewRectangle(baseX, baseY, float32(winWidth), float32(tileSize)), "START GAME") {
		s.start()
	}
}

func (s *state) start() {
	s.menu = false
	s.startedAt = time.Now()
	// build grid
	s.tiles = make([][]tile, s.rows)
	for i := 0; i < s.rows; i++ {
		s.tiles[i] = make([]tile, s.cols)
	}

	// plant the times
	count := s.mines
	for count > 0 {
		x, y := rand.Intn(s.rows), rand.Intn(s.cols)
		if s.tiles[x][y].isMine {
			continue
		}

		s.tiles[x][y].isMine = true
		s.doForNeighbours(x, y, func(x, y int) {
			s.tiles[x][y].minesAround += 1
		})
		count--
	}
}

func (s *state) getStatus() string {
	fps := rl.GetFPS()
	var duration time.Duration

	// If the game hasn't started yet, show zero time.
	if s.startedAt.IsZero() {
		duration = 0
	} else if s.gameOver || s.gameWon {
		// when finished, compute finished - started
		duration = s.finishedAt.Sub(s.startedAt)
	} else {
		// ongoing game: elapsed since started
		duration = time.Since(s.startedAt)
	}

	return fmt.Sprintf("FPS: %d, time: %.2f", fps, duration.Seconds())
}

func (s *state) getWidth() int {
	return (s.cols + 3) * tileSize
}

func (s *state) getHeight() int {
	return (s.rows + 3) * tileSize
}

func (s *state) drawField() {
	height := float32(s.getHeight())
	width := float32(s.getWidth())

	gui.StatusBar(rl.NewRectangle(0,
		float32(height)-float32(tileSize),
		float32(width),
		float32(tileSize)), s.getStatus())

	if gui.Button(rl.NewRectangle(0,
		float32(height)-float32(tileSize*2),
		float32(width),
		float32(tileSize)), "RESTART") {
		s.reset()
		return
	}
}

func (s *state) isGameWon() bool {
	opened := 0
	for x := 0; x < s.rows; x++ {
		for y := 0; y < s.cols; y++ {
			if s.tiles[x][y].open {
				opened += 1
			}
		}
	}

	return opened == (s.rows*s.cols)-s.mines
}

func (s *state) doForNeighbours(x, y int, fn func(x int, y int)) {
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := 0; i < len(dx); i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		// check range
		if 0 <= nx && nx < s.rows && 0 <= ny && ny < s.cols {
			fn(x+dx[i], y+dy[i])
		}
	}
}
func (s *state) reset() {
	s.gameOver = false
	s.gameWon = false
	s.menu = true
}
