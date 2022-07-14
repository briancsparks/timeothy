package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "math"
  "math/rand"
  "time"
)


// -------------------------------------------------------------------------------------------------------------------

type Game struct {
  allTiles    []*Sprite
  timothy     *Timothy
  lucky       *Sprite
  xscl, yscl  float64
  touchIDs    []ebiten.TouchID
  op          ebiten.DrawImageOptions
  inited      bool

  startTime       time.Time
  iterTime        time.Time
  dt              time.Duration
  drawTime        time.Time
  updateIterNum   int
  drawIterNum     int

  frameTimings    map[int]int

  gamepad         *Gamepad
}

// -------------------------------------------------------------------------------------------------------------------

func NewGame(xscl float64, yscl float64) *Game {
  g:= &Game{xscl: xscl, yscl: yscl}
  g.startTime = time.Now()
  g.updateIterNum = -1
  g.drawIterNum   = -1

  g.frameTimings = map[int]int{}

  g.frameTimings[2] = 0
  g.frameTimings[3] = 0
  g.frameTimings[4] = 0
  g.frameTimings[5] = 0
  g.frameTimings[6] = 0
  g.frameTimings[10] = 0
  g.frameTimings[15] = 0
  g.frameTimings[30] = 0
  g.frameTimings[60] = 0

  g.gamepad = NewGamepad()

  return g
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) init() {
  defer func() {
    g.inited = true
  }()

  seed := time.Now().UnixNano()
  fmt.Printf("Using random seed: %v\n", seed)
  rand.Seed(seed)

  // Must be first!
  g.addTiles(timothyZeroTilemapAsset.Tilemap)
  g.timothy = NewTimothy(g, 30)

  //g.addTiles(roguelikecityAsset.Tilemap)
  //g.addTiles(platformTilemapAsset.Tilemap)
  //g.addTiles(characterTilemapAsset.Tilemap)
  //g.addTiles(roguelikeCharacterTilemapAsset.Tilemap)
  //g.addTiles(colorBlendTilemapAsset.Tilemap)

  //luckyIndex := int(float64(len(g.allTiles)) * rand.Float64())
  //luckyIndex := 30
  //g.lucky = g.allTiles[luckyIndex]

}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) PreUpdate() error {
  if !g.inited {
    g.init()
  }

  now := time.Now()
  g.dt = now.Sub(g.iterTime)
  g.iterTime = now
  g.updateIterNum += 1

  for i, count := range g.frameTimings {
    newCount := count + 1
    g.frameTimings[i] = newCount
    if newCount >= i {
      g.frameTimings[i] = 0
    }
  }

  g.gamepad.Update()

  if g.lucky != nil {
    g.lucky.Update(g)
  }

  if g.timothy != nil {
    g.timothy.Update(g)
  }

  return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) PostUpdate() error {
  return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) PreDraw(screen *ebiten.Image) {
  now := time.Now()
  /*elapsed*/_ = now.Sub(g.drawTime)
  drawTime := time.Now()
  _= drawTime
  //fmt.Printf("delta: %v\n", elapsed)
  g.drawIterNum += 1

  g.gamepad.Draw(screen)

  // -----------------------------------------------------
  op := &ebiten.DrawImageOptions{}
  op.GeoM.Scale(2.0, 2.0)

  //for _, sprite := range g.allTiles {
  // sprite.Draw(screen, g, op)
  //}

  if g.lucky != nil {
    g.lucky.Draw(screen, g, op)
  }

  if g.timothy != nil {
    g.timothy.Draw(screen, g, op)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) PostDraw(screen *ebiten.Image) {
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
      1,
      1,
      math.Pi / 4.0,
      rand.Float64() * (math.Pi / 6.0),
    ))
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) leftTouched() bool {

  return false
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) rightTouched() bool {

  return false
}


