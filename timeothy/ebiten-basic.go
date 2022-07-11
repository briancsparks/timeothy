package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "log"
  "math"
  "math/rand"
  "time"
)

// -------------------------------------------------------------------------------------------------------------------

var (
  ebitenImage *ebiten.Image
)

// -------------------------------------------------------------------------------------------------------------------

var startTime, lastTime time.Time

func init() {
  startTime = time.Now()
  //lastTime = time.Now()
}

// -------------------------------------------------------------------------------------------------------------------

type Game struct {
  allTiles    []*Sprite
  lucky       *Sprite
  xscl, yscl  float64
  touchIDs    []ebiten.TouchID
  op          ebiten.DrawImageOptions
  inited      bool
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) addTiles(tilemap *TileMap) {

  for _, tile := range tilemap.tiles {
    tilemap := tile.parentMap
    g.allTiles = append(g.allTiles, NewSprite2(
      tile,
      tile.x * tilemap.spritePixelWidth,
      tile.y * tilemap.spritePixelHeight,
      1,
      1,
      math.Pi / 4.0,
      rand.Float64() * (math.Pi / 6.0),
    ))
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) init() {
  defer func() {
    g.inited = true
  }()

  g.xscl, g.yscl = 2.0, 2.0

  //g.addTiles(roguelikecityAsset.Tilemap)
  //g.addTiles(platformTilemapAsset.Tilemap)
  //g.addTiles(characterTilemapAsset.Tilemap)
  //g.addTiles(roguelikeCharacterTilemapAsset.Tilemap)
  //g.addTiles(colorBlendTilemapAsset.Tilemap)
  g.addTiles(timothyZeroTilemapAsset.Tilemap)

  seed := time.Now().UnixNano()
  fmt.Printf("Using random seed: %v\n", seed)
  rand.Seed(seed)
  //luckyIndex := int(float64(len(g.allTiles)) * rand.Float64())
  luckyIndex := 30
  g.lucky = g.allTiles[luckyIndex]

}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) leftTouched() bool {

  return false
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) rightTouched() bool {

  return false
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) Update() error {
  if !g.inited {
    g.init()
  }

  for _, sprite := range g.allTiles {
    //screen.DrawImage(sprite.subImage, op)
    sprite.Update(g)
  }

  return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return outsideWidth, outsideHeight
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
  drawTime := time.Now()
  /*elapsed*/_ = drawTime.Sub(lastTime)
  //fmt.Printf("delta: %v\n", elapsed)

  //w, h := ebitenImage.Size()

  op := &ebiten.DrawImageOptions{}
  op.GeoM.Scale(2.0, 2.0)

  //for _, sprite := range g.allTiles {
  // sprite.Draw(screen, g, op)
  //}

  g.lucky.Draw(screen, g, op)

  lastTime = drawTime
}

// -------------------------------------------------------------------------------------------------------------------

func TimEbitenMain() {
  ebiten.SetWindowSize(1200, 900)
  ebiten.SetWindowTitle("Timeothy")
  ebiten.SetWindowResizable(true)
  if err := ebiten.RunGame(&Game{}); err != nil {
    log.Fatal(err)
  }
}

