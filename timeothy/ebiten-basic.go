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

var (
  ebitenImage *ebiten.Image
)

func foobar() {
	fmt.Printf("\n")
}

type Game struct {
  allTiles    []*Sprite
  xscl, yscl  float64
  touchIDs    []ebiten.TouchID
  op          ebiten.DrawImageOptions
  inited      bool
}

func (g *Game) init() {
  defer func() {
    g.inited = true
  }()

  g.xscl, g.yscl = 2.0, 2.0

  for _, tile := range characterTilemap.tiles {
   tilemap := tile.parentMap
   g.allTiles = append(g.allTiles, NewSprite2(
     tile,
     tile.x * tilemap.tilePixelWidth,
     tile.y * tilemap.tilePixelHeight,
     1,
     1,
     math.Pi / 4.0,
     rand.Float64() * (math.Pi / 6.0),
   ))
  }

  for _, tile := range platformTilemap.tiles {
   tilemap := tile.parentMap
   g.allTiles = append(g.allTiles, NewSprite2(
     tile,
     tile.x * tilemap.tilePixelWidth,
     tile.y * tilemap.tilePixelHeight,
     1,
     1,
     math.Pi / 4.0,
     rand.Float64() * (math.Pi / 6.0),
   ))
  }
}

//func subSpriteRect(pxWidth, pxHeight, pxXSpace, pxYSpace, x, y int) image.Rectangle {
//  left := x * (pxWidth + pxXSpace)
//  right := left + pxWidth
//
//  top := y * (pxHeight + pxYSpace)
//  bottom := top + pxHeight
//
//  return image.Rect(left, top, right, bottom)
//}

func (g *Game) leftTouched() bool {

  return false
}

func (g *Game) rightTouched() bool {

  return false
}

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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return outsideWidth, outsideHeight
}

var startTime, lastTime time.Time

func init() {
  startTime = time.Now()
  //lastTime = time.Now()
}

func (g *Game) Draw(screen *ebiten.Image) {
  drawTime := time.Now()
  /*elapsed*/_ = drawTime.Sub(lastTime)
  //fmt.Printf("delta: %v\n", elapsed)

  //w, h := ebitenImage.Size()

  op := &ebiten.DrawImageOptions{}
  op.GeoM.Scale(2.0, 2.0)

  for _, sprite := range g.allTiles {
    //screen.DrawImage(sprite.subImage, op)
    sprite.Draw(screen, g, op)
  }

  lastTime = drawTime
}

func TimEbitenMain() {
  ebiten.SetWindowSize(1200, 900)
  ebiten.SetWindowTitle("Timeothy")
  ebiten.SetWindowResizable(true)
  if err := ebiten.RunGame(&Game{}); err != nil {
    log.Fatal(err)
  }
}

