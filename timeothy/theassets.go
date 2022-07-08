package timeothy

import (
  "bytes"
  _ "embed"
  "github.com/hajimehoshi/ebiten/v2"
  "image"
  _ "image/png"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

//go:embed assets/kenney/pixel-platformer/tilemap/tiles.png
var pixelPlatformerSsTilemapBytes []byte
//var pixelPlatformerSsTilemapImageIm *image.Image
var pixelPlatformerSsTilemapImage *ebiten.Image
var platformTilemap *TileMap

//go:embed assets/kenney/pixel-platformer/tilemap/characters.png
var pixelPlatformerSsCharactersBytes []byte
//var pixelPlatformerSsCharactersImageIm *image.Image
var pixelPlatformerSsCharactersImage *ebiten.Image
var characterTilemap *TileMap

func init() {
  pixelPlatformerSsTilemapImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(pixelPlatformerSsTilemapBytes))
  pixelPlatformerSsTilemapImage = ebiten.NewImageFromImage(pixelPlatformerSsTilemapImageIm)
  platformTilemap = NewTileMap(pixelPlatformerSsTilemapImage, 20, 9, 1, 1, 18, 18, "")

  pixelPlatformerSsCharactersImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(pixelPlatformerSsCharactersBytes))
  pixelPlatformerSsCharactersImage = ebiten.NewImageFromImage(pixelPlatformerSsCharactersImageIm)
  characterTilemap = NewTileMap(pixelPlatformerSsCharactersImage, 9, 3, 1, 1, 18, 18, "")
}

