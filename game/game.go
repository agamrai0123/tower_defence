package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const gridSize = 200

const (
	ScreenWidth  = 1200
	ScreenHeight = 800
)

type touch struct {
	id  ebiten.TouchID
	pos pos
}

type pos struct {
	x, y int
}

// Game represents a game state.
type Game struct {
	cursor  pos
	touches []touch
	// count   int

	canvasImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() *Game {
	g := &Game{
		canvasImage: ebiten.NewImage(ScreenWidth, ScreenHeight),
	}
	g.canvasImage.Fill(color.Transparent)

	return g
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.cursor.x, g.cursor.y = ebiten.CursorPosition()
	g.touches = g.touches[:0]
	var touchIDs []ebiten.TouchID
	touchIDs = ebiten.AppendTouchIDs(touchIDs[:0]) // Get active touch IDs

	for _, id := range touchIDs {
		tx, ty := ebiten.TouchPosition(id)
		g.touches = append(g.touches, touch{id: id, pos: pos{x: tx, y: ty}})
	}

	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Fill screen with background color
	screen.Fill(backgroundColor)

	// Draw grid lines
	drawGrid(screen)

	// Draw canvas image on top
	screen.DrawImage(g.canvasImage, nil)

	// Display cursor and touch info
	msg := fmt.Sprintf("(%d, %d)", g.cursor.x, g.cursor.y)
	for _, t := range g.touches {
		msg += fmt.Sprintf("\n(%d, %d) touch %d", t.pos.x, t.pos.y, t.id)
	}
	ebitenutil.DebugPrint(screen, msg)
}

func drawGrid(screen *ebiten.Image) {
	gridColor := color.RGBA{150, 150, 150, 255} // Light gray grid
	// textColor := color.White

	// Draw vertical lines
	for x := 0; x <= ScreenWidth; x += gridSize {
		vector.StrokeLine(screen, float32(x), 0, float32(x), float32(ScreenHeight), 1, gridColor, false)
	}

	// Draw horizontal lines
	for y := 0; y <= ScreenHeight; y += gridSize {
		vector.StrokeLine(screen, 0, float32(y), float32(ScreenWidth), float32(y), 1, gridColor, false)
	}

	cellIndex := 1
	for y := 0; y < ScreenHeight; y += gridSize {
		for x := 0; x < ScreenWidth; x += gridSize {
			// Draw the number at the top-left of the cell
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d, %d", x, y), x+5, y+5)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", cellIndex), x+100, y+100)
			cellIndex++
		}
	}
}
