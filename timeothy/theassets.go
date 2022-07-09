package timeothy

import (
  "bytes"
  _ "embed"
  "encoding/xml"
  "github.com/hajimehoshi/ebiten/v2"
  "image"
  _ "image/png"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

//go:embed assets/kenney/pixel-platformer/tilemap/tiles.png
var pixelPlatformerSsTilemapBytes []byte
var pixelPlatformerSsTilemapImage *ebiten.Image
var platformTilemap *TileMap

//go:embed assets/kenney/pixel-platformer/tilemap/characters.png
var pixelPlatformerSsCharactersBytes []byte
var pixelPlatformerSsCharactersImage *ebiten.Image
var characterTilemap *TileMap

//go:embed assets/kenney/roguelike-characters-pack/spritesheet/roguelikechar_transparent.png
var roguelikeSsCharactersBytes []byte
var roguelikeSsCharactersImage *ebiten.Image
var roguelikeCharacterTilemap *TileMap

//go:embed assets/kenney/roguelike-city-pack/spritesheet/roguelikecity_magenta.png
var roguelikecitySSBytes []byte
var roguelikecitySSImage *ebiten.Image
var roguelikecityTilemap *TileMap

//go:embed assets/roguelike.tmx
var roguelikecityTmxBytes []byte

//go:embed assets/roguelikecity_magenta.tsx
var roguelikecityTsxBytes []byte


var roguelikecityAsset *TilemapAsset

func init() {
  roguelikecityAsset = MakeTilemapFromBytesAndXml(roguelikecitySSBytes, roguelikecityTsxBytes, "assets/kenney/roguelike-city-pack/spritesheet/roguelikecity_magenta.png")

  var tiledMap TiledMap
  _ = xml.Unmarshal(roguelikecityTmxBytes, &tiledMap)

  pixelPlatformerSsTilemapImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(pixelPlatformerSsTilemapBytes))
  pixelPlatformerSsTilemapImage = ebiten.NewImageFromImage(pixelPlatformerSsTilemapImageIm)
  platformTilemap = NewTileMap(pixelPlatformerSsTilemapImage, 20, 9, 1, 1, 19, 19, "")

  pixelPlatformerSsCharactersImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(pixelPlatformerSsCharactersBytes))
  pixelPlatformerSsCharactersImage = ebiten.NewImageFromImage(pixelPlatformerSsCharactersImageIm)
  characterTilemap = NewTileMap(pixelPlatformerSsCharactersImage, 9, 3, 3, 3, 23, 23, "")

  roguelikeSsCharactersImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(roguelikeSsCharactersBytes))
  roguelikeSsCharactersImage = ebiten.NewImageFromImage(roguelikeSsCharactersImageIm)
  roguelikeCharacterTilemap = NewTileMap(roguelikeSsCharactersImage, -1, -1, 1, 1, 16, 16, "")
}
func (tt *TiledTileset) Rows() int {
  return tt.Tilecount / tt.Columns
}

type TilemapAsset struct {
  Tilemap       *TileMap
  Image         *ebiten.Image

  Name           string
  Data          []byte
  ImImage       *image.Image
  Xml            string
}

func MakeTilemapFromBytesAndXml(data []byte, xmlBytes []byte, filename string) *TilemapAsset {
  var tiledTileset TiledTileset
  _ = xml.Unmarshal(xmlBytes, &tiledTileset)

  imageIm, _, _ := image.Decode(bytes.NewReader(data))
  myImage := ebiten.NewImageFromImage(imageIm)
  tilemap := NewTileMap(
    myImage,
    tiledTileset.Columns,
    tiledTileset.Rows(),
    tiledTileset.Spacing,
    tiledTileset.Spacing,
    tiledTileset.Tilewidth,
    tiledTileset.Tileheight,
    filename,
  )

  ebitenTilemap := TilemapAsset{
    Image:      myImage,
    Tilemap:    tilemap,
    Name:       filename,
    Data:       data,
    ImImage:    &imageIm,
    Xml:        string(xmlBytes),
  }

  return &ebitenTilemap
}

