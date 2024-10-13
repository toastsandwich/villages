package game

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	pidleu     time.Time
	foamUpdate time.Time
	foamFrame  = 0
)

type Game struct {
	Pawn   Pawn
	Grass  Grass
	Foam   []*ebiten.Image
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

	if time.Since(foamUpdate) >= 50*time.Millisecond {
		foamFrame = (foamFrame + 1) % len(g.Foam)
		foamUpdate = time.Now()
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
	g.drawMap(screen)
	g.drawGrid(screen)
	// g.drawGreen(screen)
	g.drawPawn(screen)
}

func (g *Game) drawMap(screen *ebiten.Image) {
	sx, sy := 192, 128
	f := getFoamPos()
	for x, row := range f {
		for y, val := range row {
			opts := &ebiten.DrawImageOptions{}
			if val == 1 {
				opts.GeoM.Translate(float64(sx+(x*64)-64), float64(sy+(y*64)-64))
				screen.DrawImage(g.Foam[foamFrame], opts)
			}
		}
	}
	slc := g.Grass.SliceG()
	for x, row := range GRASSMAP {
		for y, val := range row {
			opts := &ebiten.DrawImageOptions{}
			if val != 10 {
				opts.GeoM.Translate(float64(sx+(x*64)), float64(sy+(y*64)))
				screen.DrawImage(slc[val], opts)
			}
		}
	}

}

func (g *Game) drawGrid(screen *ebiten.Image) {
	// Define grid size
	gridSize := 64                              // Change this value to increase or decrease the size of each grid cell
	lineColor := color.RGBA{255, 255, 255, 128} // White color with some transparency

	// Draw vertical lines
	for x := 0; x < screen.Bounds().Dx(); x += gridSize {
		ebitenutil.DrawLine(screen, float64(x), 0, float64(x), float64(screen.Bounds().Dy()), lineColor)
	}
	for y := 0; y < screen.Bounds().Dy(); y += gridSize {
		ebitenutil.DrawLine(screen, 0, float64(y), float64(screen.Bounds().Dx()), float64(y), lineColor)
	}
}

func (g *Game) drawPawn(screen *ebiten.Image) {
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

func (g *Game) Layout(int, int) (int, int) { return 1920, 1080 }
