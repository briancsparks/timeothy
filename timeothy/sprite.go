package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
  //subImage     *ebiten.Image
  tiles         []*Tile

  x, y          int
  dx,dy         int
  angle         float64
  dtheta        float64
}

func NewSprite2(mainTile *Tile, x int, y int, dx int, dy int, angle, dtheta float64) *Sprite {
  s := &Sprite{
    x: x,
    y: y,
    dx: dx,
    dy: dy,
    angle: angle,
    dtheta: dtheta,
  }
  s.tiles = append(s.tiles, mainTile)

  return s
}

func (s *Sprite) getMainTile() *Tile {
  return s.tiles[0]
}

func (s *Sprite) Update(g *Game) error {
  s.angle += s.dtheta
  return nil
}

func (s *Sprite) Draw(screen *ebiten.Image, g *Game, op *ebiten.DrawImageOptions) {
  //nw, nh := s.getMainTile().subImage.Size()
  //w, h := float64(nw), float64(nh)
  //xr, yr := w / 2.0, h / 2.0              /* x and y radius */

  op.GeoM.Reset()
  op.GeoM.Scale(g.xscl, g.yscl)
  //op.GeoM.Translate(-xr, -yr)
  //op.GeoM.Rotate(s.angle)
  //op.GeoM.Translate(xr, yr)
  op.GeoM.Translate(g.xscl * float64(s.x), g.yscl * float64(s.y))

  screen.DrawImage(s.getMainTile().subImage, op)
}


//func NewSprite(image *ebiten.Image, x int, y int, dx int, dy int, angle, dtheta float64) *Sprite {
// return &Sprite{
//   subImage: image,
//   x: x,
//   y: y,
//   dx: dx,
//   dy: dy,
//   angle: angle,
//   dtheta: dtheta,
// }
//}