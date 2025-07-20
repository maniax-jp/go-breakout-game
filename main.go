package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 600
	screenHeight = 1000
	paddleWidth  = 100
	paddleHeight = 20
	ballSize     = 10
	blockWidth   = 60
	blockHeight  = 20
	blockRows    = 8
	blockCols    = 9
)

type Paddle struct {
	X, Y          float64
	Width, Height float64
}

type Ball struct {
	X, Y   float64
	VX, VY float64
	Size   float64
}

type Block struct {
	X, Y          float64
	Width, Height float64
	Active        bool
}

type Game struct {
	paddle       Paddle
	ball         Ball
	blocks       []Block
	gameStarted  bool
	gameOver     bool
	gameWin      bool
}

func NewGame() *Game {
	g := &Game{
		paddle: Paddle{
			X:      screenWidth/2 - paddleWidth/2,
			Y:      screenHeight - 50,
			Width:  paddleWidth,
			Height: paddleHeight,
		},
		ball: Ball{
			X:    screenWidth / 2,
			Y:    screenHeight - 70,
			VX:   3,
			VY:   -3,
			Size: ballSize,
		},
		blocks: make([]Block, blockRows*blockCols),
	}

	// ブロックの初期化
	blockAreaWidth := float64(blockCols)*(blockWidth+5) - 5
	startX := (screenWidth - blockAreaWidth) / 2
	for row := 0; row < blockRows; row++ {
		for col := 0; col < blockCols; col++ {
			index := row*blockCols + col
			g.blocks[index] = Block{
				X:      startX + float64(col)*(blockWidth+5),
				Y:      100 + float64(row)*(blockHeight+5),
				Width:  blockWidth,
				Height: blockHeight,
				Active: true,
			}
		}
	}

	return g
}

func (g *Game) Update() error {
	if g.gameOver || g.gameWin {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			*g = *NewGame()
		}
		return nil
	}

	if !g.gameStarted {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.gameStarted = true
		}
		return nil
	}

	// パドルの操作
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.paddle.X > 0 {
		g.paddle.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.paddle.X < screenWidth-g.paddle.Width {
		g.paddle.X += 5
	}

	// ボールの移動
	g.ball.X += g.ball.VX
	g.ball.Y += g.ball.VY

	// 壁との衝突判定
	if g.ball.X <= 0 || g.ball.X >= screenWidth-g.ball.Size {
		g.ball.VX = -g.ball.VX
	}
	if g.ball.Y <= 0 {
		g.ball.VY = -g.ball.VY
	}

	// パドルとの衝突判定
	if g.ball.Y+g.ball.Size >= g.paddle.Y &&
		g.ball.X+g.ball.Size >= g.paddle.X &&
		g.ball.X <= g.paddle.X+g.paddle.Width {
		g.ball.VY = -math.Abs(g.ball.VY)
		// パドルの位置によってボールの方向を変える
		hitPos := (g.ball.X + g.ball.Size/2 - g.paddle.X) / g.paddle.Width
		g.ball.VX = (hitPos - 0.5) * 6
	}

	// ブロックとの衝突判定
	for i := range g.blocks {
		if !g.blocks[i].Active {
			continue
		}

		if g.ball.X+g.ball.Size >= g.blocks[i].X &&
			g.ball.X <= g.blocks[i].X+g.blocks[i].Width &&
			g.ball.Y+g.ball.Size >= g.blocks[i].Y &&
			g.ball.Y <= g.blocks[i].Y+g.blocks[i].Height {
			g.blocks[i].Active = false
			g.ball.VY = -g.ball.VY
			break
		}
	}

	// ゲームオーバー判定
	if g.ball.Y > screenHeight {
		g.gameOver = true
	}

	// ゲームクリア判定
	activeBlocks := 0
	for _, block := range g.blocks {
		if block.Active {
			activeBlocks++
		}
	}
	if activeBlocks == 0 {
		g.gameWin = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	if !g.gameStarted {
		ebitenutil.DebugPrint(screen, "BREAKOUT GAME\n\nPress SPACE to start\nUse Arrow Keys to move paddle")
		return
	}

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "GAME OVER\n\nPress SPACE to restart")
		return
	}

	if g.gameWin {
		ebitenutil.DebugPrint(screen, "YOU WIN!\n\nPress SPACE to restart")
		return
	}

	// パドルの描画
	ebitenutil.DrawRect(screen, g.paddle.X, g.paddle.Y, g.paddle.Width, g.paddle.Height, color.RGBA{255, 255, 255, 255})

	// ボールの描画
	ebitenutil.DrawRect(screen, g.ball.X, g.ball.Y, g.ball.Size, g.ball.Size, color.RGBA{255, 255, 0, 255})

	// ブロックの描画
	for _, block := range g.blocks {
		if block.Active {
			ebitenutil.DrawRect(screen, block.X, block.Y, block.Width, block.Height, color.RGBA{255, 100, 100, 255})
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Breakout Game")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}