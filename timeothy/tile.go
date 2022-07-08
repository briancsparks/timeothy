package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "image"
)

type Tile struct {
  parentMap    *TileMap
  subImage     *ebiten.Image

  subRect       image.Rectangle
  x, y, n       int
}

func NewTile(parentMap *TileMap, subRect image.Rectangle, x, y, n int) *Tile {
  return &Tile{
    parentMap:  parentMap,
    subImage:   parentMap.tilemap.SubImage(subRect).(*ebiten.Image),
    subRect:    subRect,
    x:          x,
    y:          y,
    n:          n,
  }
}

func tile() {
	fmt.Printf("\n")
}
