package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/hajimehoshi/ebiten/v2"
  "log"
)

// -------------------------------------------------------------------------------------------------------------------

func init() {
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) Update() error {
  err := g.PreUpdate()
  if err != nil {
    return err
  }

  for _, sprite := range g.allTiles {
    //screen.DrawImage(sprite.subImage, op)
    sprite.Update(g)
  }

  return g.PostUpdate()
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return outsideWidth, outsideHeight
}

// -------------------------------------------------------------------------------------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
  g.PreDraw(screen)

  g.PostDraw(screen)
}

// -------------------------------------------------------------------------------------------------------------------

func TimEbitenMain() {
  ebiten.SetWindowSize(1200, 900)
  ebiten.SetWindowTitle("Timeothy")
  ebiten.SetWindowResizable(true)

  game := NewGame(2.0, 2.0)
  if err := ebiten.RunGame(game); err != nil {
    log.Fatal(err)
  }
}

