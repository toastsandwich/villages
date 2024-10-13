package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Grass struct {
	T *ebiten.Image //tile
	H *ebiten.Image //horizontal
	V *ebiten.Image //vertical
	G *ebiten.Image //grid
	A *ebiten.Image //addon
}

/*

// layout for grass ans index in the slice
+-----+-----+-----+
|--0--|--3--|--6--|
+-----+-----+-----+
|--1--|--4--|--7--|
+-----+-----+-----+
|--2--|--5--|--8--|
+-----+-----+-----+

0 -> (0,0) and (64, 64)
1 -> (0,64) and (64, 128)
2 -> (0, 128) and (64, 128)
6 -> (128, 0) and (192, 64)
..
8 -> (128, 128) and (192, 192)
*/

// 9 will be for grass tile
// 10 will be water
func (g *Grass) SliceG() []*ebiten.Image {
	cx, cy := 0, 0
	slc := make([]*ebiten.Image, 0, 9)
	for x := 0; x < 3; x++ {
		cx = (x * 64)
		for y := 0; y < 3; y++ {
			cy = (y * 64)
			sub := g.G.SubImage(image.Rect(cx, cy, cx+64, cy+64)).(*ebiten.Image)
			slc = append(slc, sub)
		}
	}
	return slc
}
