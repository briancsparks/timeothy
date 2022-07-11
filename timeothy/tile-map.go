package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/hajimehoshi/ebiten/v2"
  "image"
)

// -------------------------------------------------------------------------------------------------------------------

type TileMap struct {
  sheetSpriteWidth  int
  sheetSpriteHeight int
  sheetPixelXSpace  int
  sheetPixelYSpace int
  spritePixelWidth  int
  spritePixelHeight int

  tilemap              *ebiten.Image
  tiles                 []*Tile
  //subRect               image.Rectangle

  filename              string
}

// -------------------------------------------------------------------------------------------------------------------

func NewTileMap(tilemap *ebiten.Image, sheetSpriteWidth int, sheetSpriteHeight int, sheetPixelXSpace int, sheetPixelYSpace int,
      spritePixelWidth int, spritePixelHeight int, filename string) *TileMap {

  if sheetSpriteWidth == -1 {
    sheetPixelWidth, sheetPixelHeight := tilemap.Size()
    sheetSpriteWidth = (sheetPixelWidth + 1) / (spritePixelWidth + sheetPixelXSpace)
    sheetSpriteHeight = (sheetPixelHeight + 1) / (spritePixelHeight + sheetPixelYSpace)
  }

  tm := &TileMap{
    sheetSpriteWidth:  sheetSpriteWidth,
    sheetSpriteHeight: sheetSpriteHeight,
    sheetPixelXSpace:  sheetPixelXSpace,
    sheetPixelYSpace:  sheetPixelYSpace,
    spritePixelWidth:  spritePixelWidth,
    spritePixelHeight: spritePixelHeight,
    tilemap:           tilemap,
    filename:          filename,
  }

  index := 0
  for j := 0; j < tm.sheetSpriteHeight; j++ {
    for i := 0; i < tm.sheetSpriteWidth; i++ {
      sub := tm.subSpriteRect(i, j)
      t := NewTile(tm, sub, i, j, index)
      tm.tiles = append(tm.tiles, t)

      index += 1
    }
  }

  return tm
}

// -------------------------------------------------------------------------------------------------------------------

func (tm *TileMap) getTiles(indexes []int) []*Tile {
  var result []*Tile
  for _, index := range indexes {
    result = append(result, tm.tiles[index])
  }
  return result
}

// -------------------------------------------------------------------------------------------------------------------

func (tm *TileMap) subSpriteRect( /*pxWidth, pxHeight, pxXSpace, pxYSpace,*/ x, y int) image.Rectangle {
  left := x * (tm.spritePixelWidth + tm.sheetPixelXSpace)
  right := left + tm.spritePixelWidth

  top := y * (tm.spritePixelHeight + tm.sheetPixelYSpace)
  bottom := top + tm.spritePixelHeight

  return image.Rect(left, top, right, bottom)
}

// -------------------------------------------------------------------------------------------------------------------

func (tm *TileMap) tileWidth() int {
  return tm.sheetSpriteWidth
}

// -------------------------------------------------------------------------------------------------------------------

func (tm *TileMap) tileHeight() int {
  return tm.sheetSpriteHeight
}
