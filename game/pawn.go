package game

import "github.com/hajimehoshi/ebiten/v2"

type Pawn struct {
	IdleSprite    []*ebiten.Image
	MoveSprite    []*ebiten.Image
	BuildSprite   []*ebiten.Image
	CuttingSprtie []*ebiten.Image
	CurrentFrame  int
	X             float64
	Y             float64
	Speed         float64
	Idle          bool
	Building      bool
	Cutting       bool
}

func NewPawn(idleSprite, moveSprite, buildSprite, cuttingSprite []*ebiten.Image) Pawn {
	return Pawn{
		IdleSprite:    idleSprite,
		MoveSprite:    moveSprite,
		BuildSprite:   buildSprite,
		CuttingSprtie: cuttingSprite,
		CurrentFrame:  0,
		Speed:         3,
		Idle:          true,
		Building:      false,
		X:             0,
		Y:             0,
	}
}
