package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/toastsandwich/villages/game"
)

func ConfigurePawn() (game.Pawn, error) {
	esprite, _, err := ebitenutil.NewImageFromFile("assets/Factions/Knights/Troops/Pawn/Blue/Pawn_Blue.png")
	if err != nil {
		return game.Pawn{}, err
	}
	var minx, miny int = 0, 0
	var maxx, maxy int = 192, 192
	idleSprite := make([]*ebiten.Image, 6)
	for i := 0; i < 6; i++ {
		x0, y0 := minx+(i*192), miny
		x1, y1 := maxx+(i*192), maxy
		frame := esprite.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
		idleSprite[i] = frame
	}
	moveSprite := make([]*ebiten.Image, 6)
	minx, miny = 0, 192
	maxx, maxy = 192, 384
	for i := 0; i < 6; i++ {
		x0, y0 := minx+(i*192), miny
		x1, y1 := maxx+(i*192), maxy
		frame := esprite.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
		moveSprite[i] = frame
	}
	buildSprite := make([]*ebiten.Image, 6)
	minx, miny = 0, 384
	maxx, maxy = 192, 576
	for i := 0; i < 6; i++ {
		x0, y0 := minx+(i*192), miny
		x1, y1 := maxx+(i*192), maxy
		frame := esprite.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
		buildSprite[i] = frame
	}
	cuttingSprite := make([]*ebiten.Image, 6)
	minx, miny = 0, 576
	maxx, maxy = 192, 768
	for i := 0; i < 6; i++ {
		x0, y0 := minx+(i*192), miny
		x1, y1 := maxx+(i*192), maxy
		frame := esprite.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
		cuttingSprite[i] = frame
	}
	return game.NewPawn(idleSprite, moveSprite, buildSprite, cuttingSprite), nil
}

func ConfigureGrass() (game.Grass, error) {
	eimg, _, err := ebitenutil.NewImageFromFile("assets/Terrain/Ground/Tilemap_Flat.png")
	if err != nil {
		return game.Grass{}, err
	}

	var minx, miny int = 0, 0
	var maxx, maxy int = 192, 192
	g := eimg.SubImage(image.Rect(minx, miny, maxx, maxy)).(*ebiten.Image)

	minx, miny = 0, 192
	maxx, maxy = 192, 256
	h := eimg.SubImage(image.Rect(minx, miny, maxx, maxy)).(*ebiten.Image)

	minx, miny = 192, 0
	maxx, maxy = 256, 192
	v := eimg.SubImage(image.Rect(minx, miny, maxx, maxy)).(*ebiten.Image)

	minx, miny = 192, 192
	maxx, maxy = 256, 256
	t := eimg.SubImage(image.Rect(minx, miny, maxx, maxy)).(*ebiten.Image)
	return game.Grass{
		G: g,
		H: h,
		V: v,
		T: t,
	}, nil
}

func Foam() ([]*ebiten.Image, error) {
	eimg, _, err := ebitenutil.NewImageFromFile("assets/Terrain/Water/Foam/Foam.png")
	if err != nil {
		return nil, err
	}
	foam := make([]*ebiten.Image, 8)

	for i := 0; i < 8; i++ {
		minx := i * 192
		maxx := 192 + minx

		f := eimg.SubImage(image.Rect(minx, 0, maxx, 192)).(*ebiten.Image)
		foam[i] = f
	}
	return foam, nil
}

func main() {
	ebiten.SetWindowTitle("game")
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(1280, 720)

	pawn, err := ConfigurePawn()
	if err != nil {
		log.Fatal(err)
	}
	grass, err := ConfigureGrass()
	if err != nil {
		log.Fatal(err)
	}
	foam, err := Foam()
	if err != nil {
		log.Fatal(err)
	}
	game := &game.Game{
		Pawn:  pawn,
		Grass: grass,
		Foam:  foam,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
