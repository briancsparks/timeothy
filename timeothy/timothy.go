package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/hajimehoshi/ebiten/v2"
)

func t() {
	fmt.Printf("\n")
}

type Timothy struct {
  sprite  *Sprite
}

func (t *Timothy) Update() {

}

func (t *Timothy) Draw(screen *ebiten.Image, g *Game, op *ebiten.DrawImageOptions) {

}




