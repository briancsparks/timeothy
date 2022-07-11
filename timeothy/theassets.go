package timeothy

import (
  "bytes"
  _ "embed"
  "encoding/xml"
  "fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "image"
  _ "image/png"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

//go:embed assets/kenney/pixel-platformer/tilemap/tiles.png
var pixelPlatformerSsTilemapBytes []byte
//var pixelPlatformerSsTilemapImage *ebiten.Image
//var platformTilemap *TileMap

//go:embed assets/kenney/pixel-platformer/tilemap/characters.png
var pixelPlatformerSsCharactersBytes []byte
//var pixelPlatformerSsCharactersImage *ebiten.Image
//var characterTilemap *TileMap

//go:embed assets/kenney/roguelike-characters-pack/spritesheet/roguelikechar_transparent.png
var roguelikeSsCharactersBytes []byte
//var roguelikeSsCharactersImage *ebiten.Image
//var roguelikeCharacterTilemap *TileMap

//go:embed assets/colorblend.png
var colorBlendBytes []byte

//go:embed assets/timothy-zero-tilemap.png
var timothyZeroBytes []byte


// --------------------------------------------------------------------------------------------------------------------

var roguelikecityAsset *TilemapAsset

//go:embed assets/kenney/roguelike-city-pack/spritesheet/roguelikecity_magenta.png
var roguelikecitySSBytes []byte

//go:embed assets/roguelikecity_magenta.tsx
var roguelikecityTsxBytes []byte

// --------------------------------------------------------------------------------------------------------------------

//go:embed assets/roguelike.tmx
var roguelikecityTmxBytes []byte

var platformTilemapAsset              *TilemapAsset
var characterTilemapAsset             *TilemapAsset
var roguelikeCharacterTilemapAsset    *TilemapAsset

var colorBlendTilemapAsset            *TilemapAsset
var timothyZeroTilemapAsset            *TilemapAsset

func init() {
  roguelikecityAsset = MakeTilemapFromBytesAndXml(roguelikecitySSBytes, roguelikecityTsxBytes, "assets/kenney/roguelike-city-pack/spritesheet/roguelikecity_magenta.png")

  //pixelPlatformerSsTilemapImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(pixelPlatformerSsTilemapBytes))
  //pixelPlatformerSsTilemapImage = ebiten.NewImageFromImage(pixelPlatformerSsTilemapImageIm)
  //platformTilemap = NewTileMap(pixelPlatformerSsTilemapImage, 20, 9, 1, 1, 19, 19, "")

  platformTilemapAsset = MakeTilemapFromBytes(pixelPlatformerSsTilemapBytes, "", 20, 9, 1, 1, 19, 19, "")

  //pixelPlatformerSsCharactersImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(pixelPlatformerSsCharactersBytes))
  //pixelPlatformerSsCharactersImage = ebiten.NewImageFromImage(pixelPlatformerSsCharactersImageIm)
  //characterTilemap = NewTileMap(pixelPlatformerSsCharactersImage, 9, 3, 3, 3, 23, 23, "")

  characterTilemapAsset = MakeTilemapFromBytes(pixelPlatformerSsCharactersBytes, "", 9, 3, 3, 3, 23, 23, "")
  colorBlendTilemapAsset = MakeTilemapFromBytes(colorBlendBytes, "", 10, 5, 0, 0, 40, 40, "")
  timothyZeroTilemapAsset = MakeTilemapFromBytes(timothyZeroBytes, "", -1, -1, 0, 0, 32, 128, "")

  //roguelikeSsCharactersImageIm, _, _ /*err*/ := image.Decode(bytes.NewReader(roguelikeSsCharactersBytes))
  //roguelikeSsCharactersImage = ebiten.NewImageFromImage(roguelikeSsCharactersImageIm)
  //roguelikeCharacterTilemap = NewTileMap(roguelikeSsCharactersImage, -1, -1, 1, 1, 16, 16, "")

  roguelikeCharacterTilemapAsset = MakeTilemapFromBytes(roguelikeSsCharactersBytes, "", -1, -1, 1, 1, 16, 16, "")
}
func (tt *TiledTileset) Rows() int {
  return tt.Tilecount / tt.Columns
}

type TilemapAsset struct {
  Tilemap       *TileMap
  Image         *ebiten.Image

  Name           string
  Data          []byte
  ImImage      *image.Image
  DataFormat    string
}

// --------------------------------------------------------------------------------------------------------------------

func MakeTilemapFromBytes(data []byte, filename string,
    sheetSpriteWidth int,
    sheetSpriteHeight int,
    sheetPixelXSpace int,
    sheetPixelYSpace int,
    spritePixelWidth int,
    spritePixelHeight int,
    dataFormat string,

) *TilemapAsset {

  imageIm, _, _ := image.Decode(bytes.NewReader(data))
  myImage := ebiten.NewImageFromImage(imageIm)

  tilemap := NewTileMap(
    myImage,
    sheetSpriteWidth,
    sheetSpriteHeight,
    sheetPixelXSpace,
    sheetPixelYSpace,
    spritePixelWidth,
    spritePixelHeight,
    filename,
  )
  sheetSpriteWidth = tilemap.sheetSpriteWidth
  sheetSpriteHeight = tilemap.sheetSpriteHeight

  if len(dataFormat) <= 0 {
    dataFormat = fmt.Sprintf("Tiles: (%dpx x %dpx), Spacing: (%d, %d), #Sprites: %d (%d x %d)",
      spritePixelWidth, spritePixelHeight, sheetPixelXSpace, sheetPixelYSpace, sheetSpriteWidth * sheetSpriteHeight, sheetSpriteWidth, sheetSpriteHeight)
  }

  ebitenTilemap := TilemapAsset{
    Image:      myImage,
    Tilemap:    tilemap,
    Name:       filename,
    Data:       data,
    ImImage:    &imageIm,
    DataFormat: dataFormat,
  }

  return &ebitenTilemap
}

// --------------------------------------------------------------------------------------------------------------------

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
    DataFormat: string(xmlBytes),
  }

  return &ebitenTilemap
}

