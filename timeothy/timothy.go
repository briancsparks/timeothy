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
  anim    []*Sprite

  joystickMovementMult float64        /* Joystick movement multiplier */

  g       *Game
}

func NewTimothy(g *Game, tileId int) *Timothy {

  t := &Timothy{sprite: g.allTiles[tileId], g: g}

  t.anim = append(t.anim, g.allTiles[tileId])
  t.anim = append(t.anim, g.allTiles[tileId + 1])
  t.anim = append(t.anim, g.allTiles[tileId + 2])
  t.anim = append(t.anim, g.allTiles[tileId + 3])
  t.anim = append(t.anim, g.allTiles[tileId + 4])
  t.anim = append(t.anim, g.allTiles[tileId + 5])

  for _, sprite := range t.anim {
    sprite.CloneProps(t.sprite)
  }

  t.joystickMovementMult = 1.0

  dynFig.MakeFloat64IncDecActions("joystickMoveMult", "b", "n", 0.1)

  return t
}

func (t *Timothy) Update(g *Game) {
  if g.updateIterNum % 10 == 0 {
    t.joystickMovementMult, _ = dynFig.GetFloat64("joystickMoveMult")
  }

  dx, dy := 0.0, 0.0

  var id ebiten.GamepadID
  for gamepadID, _ := range t.g.gamepad.gamepadIDs {
    id = gamepadID
    break
  }

  dx += ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal) * t.joystickMovementMult
  dy += ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical) * t.joystickMovementMult
  //dx += ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickHorizontal)
  //dy += ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickVertical)

  t.sprite.dx = dx
  t.sprite.dy = dy
  t.sprite.fx += t.sprite.dx
  t.sprite.fy += t.sprite.dy
  t.sprite.x   = int(t.sprite.fx)
  t.sprite.y   = int(t.sprite.fy)
  //fmt.Printf("x: %v, y: %v\n", t.sprite.fx, t.sprite.fy)

  for _, sprite := range t.anim {
    sprite.CloneProps(t.sprite)
  }

}

func (t *Timothy) Draw(screen *ebiten.Image, g *Game, op *ebiten.DrawImageOptions) {
  animCellId := g.frameTimings[6]

  t.anim[animCellId].Draw(screen, g, op)
}




