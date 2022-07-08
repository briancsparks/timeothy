package timeothy

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */
/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "image"
)

func asdf() {
	fmt.Printf("\n")
}

type TileMap struct {
  sheetTileWidth        int
  sheetTileHeight       int
  sheetPixelXSpace      int
  sheetPixelYSpace      int
  tilePixelWidth        int
  tilePixelHeight       int

  tilemap              *ebiten.Image
  tiles                 []*Tile
  //subRect               image.Rectangle

  filename              string
}

func NewTileMap(tilemap *ebiten.Image, sheetSpriteWidth int, sheetSpriteHeight int, sheetPixelXSpace int, sheetPixelYSpace int,
      spritePixelWidth int, spritePixelHeight int, filename string) *TileMap {

  tm := &TileMap{
    sheetTileWidth:   sheetSpriteWidth,
    sheetTileHeight:  sheetSpriteHeight,
    sheetPixelXSpace: sheetPixelXSpace,
    sheetPixelYSpace: sheetPixelYSpace,
    tilePixelWidth:   spritePixelWidth,
    tilePixelHeight:  spritePixelHeight,
    tilemap:          tilemap,
    filename:         filename,
  }

  index := 0
  for i := 0; i < tm.sheetTileWidth; i++ {
    for j := 0; j < tm.sheetTileHeight; j++ {
      sub := tm.subSpriteRect(i, j)
      t := NewTile(tm, sub, i, j, index)
      tm.tiles = append(tm.tiles, t)

      index += 1
    }
  }

  return tm
}

func (tm *TileMap) getTiles(indexes []int) []*Tile {
  var result []*Tile
  for _, index := range indexes {
    result = append(result, tm.tiles[index])
  }
  return result
}


func (tm *TileMap) subSpriteRect( /*pxWidth, pxHeight, pxXSpace, pxYSpace,*/ x, y int) image.Rectangle {
  left := x * (tm.tilePixelWidth + tm.sheetPixelXSpace)
  right := left + tm.tilePixelWidth

  top := y * (tm.tilePixelHeight + tm.sheetPixelYSpace)
  bottom := top + tm.tilePixelHeight

  return image.Rect(left, top, right, bottom)
}


func (tm *TileMap) tileWidth() int {
  return tm.sheetTileWidth
}

func (tm *TileMap) tileHeight() int {
  return tm.sheetTileHeight
}
