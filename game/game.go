package game

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var pidleu time.Time

type Game struct {
	Pawn   Pawn
	PFrame []*ebiten.Image
}

func (g *Game) Update() error {
	if g.Pawn.Idle && g.Pawn.Building {
		g.PFrame = g.Pawn.BuildSprite
	} else if g.Pawn.Idle && g.Pawn.Cutting {
		g.PFrame = g.Pawn.CuttingSprtie
	} else if g.Pawn.Idle {
		g.PFrame = g.Pawn.IdleSprite
	} else {
		g.PFrame = g.Pawn.MoveSprite
	}
	if time.Since(pidleu) >= 100*time.Millisecond {
		g.Pawn.CurrentFrame = (g.Pawn.CurrentFrame + 1) % len(g.Pawn.IdleSprite)
		pidleu = time.Now()
	}
	g.movePawn()
	return nil
}

func (g *Game) movePawn() {
	g.Pawn.Idle = true
	g.Pawn.Building = false
	g.Pawn.Cutting = false
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.Pawn.Building = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.Pawn.Cutting = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Pawn.Idle = false
		g.Pawn.Y -= g.Pawn.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Pawn.Idle = false
		g.Pawn.X -= g.Pawn.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Pawn.Idle = false
		g.Pawn.Y += g.Pawn.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Pawn.Idle = false
		g.Pawn.X += g.Pawn.Speed
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawWater(screen)
	p := g.PFrame[g.Pawn.CurrentFrame]
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1, 1)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		opts.GeoM.Scale(-1, 1)
		x := p.Bounds().Dx()
		opts.GeoM.Translate(float64(x), float64(0))
	}
	opts.GeoM.Translate(float64(g.Pawn.X), float64(g.Pawn.Y))
	screen.DrawImage(p, opts)
}

func (g *Game) drawWater(screen *ebiten.Image) {
	waterImage, _, err := ebitenutil.NewImageFromFile("assets/Terrain/Water/Water.png")
	if err != nil {
		log.Fatal(err)
	}
	waterWidth, waterHeight := waterImage.Bounds().Dx(), waterImage.Bounds().Dy()
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	for x := 0; x < screenWidth; x += waterWidth {
		for y := 0; y < screenHeight; y += waterHeight {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(waterImage, opts)
		}
	}
}

func (g *Game) Layout(int, int) (int, int) { return 1280, 720 }
