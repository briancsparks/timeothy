package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "sort"
  "strconv"
  "strings"
)

// -------------------------------------------------------------------------------------------------------------------

type Gamepad struct {
  // Taken from ebiten gamepad example
  gamepadIDsBuf  []ebiten.GamepadID
  gamepadIDs     map[ebiten.GamepadID]struct{}
  axes           map[ebiten.GamepadID][]string
  pressedButtons map[ebiten.GamepadID][]string

}

// -------------------------------------------------------------------------------------------------------------------

func NewGamepad() *Gamepad {
  gp := &Gamepad{}

  gp.gamepadIDs = map[ebiten.GamepadID]struct{}{}

  return gp
}

// -------------------------------------------------------------------------------------------------------------------

func (gp *Gamepad) Update() error {

  // Log the gamepad connection events.
  gp.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(gp.gamepadIDsBuf[:0])
  for _, id := range gp.gamepadIDsBuf {
    //log.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
    gp.gamepadIDs[id] = struct{}{}
  }

  for id := range gp.gamepadIDs {
    if inpututil.IsGamepadJustDisconnected(id) {
      //log.Printf("gamepad disconnected: id: %d", id)
      delete(gp.gamepadIDs, id)
    }
  }

  gp.axes = map[ebiten.GamepadID][]string{}
  gp.pressedButtons = map[ebiten.GamepadID][]string{}
  for id := range gp.gamepadIDs {
    maxAxis := ebiten.GamepadAxisNum(id)
    for a := 0; a < maxAxis; a++ {
      v := ebiten.GamepadAxisValue(id, a)
      gp.axes[id] = append(gp.axes[id], fmt.Sprintf("%d:%+0.2f", a, v))
    }

    maxButton := ebiten.GamepadButton(ebiten.GamepadButtonNum(id))
    for b := ebiten.GamepadButton(id); b < maxButton; b++ {
      if ebiten.IsGamepadButtonPressed(id, b) {
        gp.pressedButtons[id] = append(gp.pressedButtons[id], strconv.Itoa(int(b)))
      }

      // Log button events.
      if inpututil.IsGamepadButtonJustPressed(id, b) {
        //log.Printf("button pressed: id: %d, button: %d", id, b)
      }
      if inpututil.IsGamepadButtonJustReleased(id, b) {
        //log.Printf("button released: id: %d, button: %d", id, b)
      }
    }

    if ebiten.IsStandardGamepadLayoutAvailable(id) {
      for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
        // Log button events.
        if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
          //log.Printf("standard button pressed: id: %d, button: %d", id, b)
        }
        if inpututil.IsStandardGamepadButtonJustReleased(id, b) {
          //log.Printf("standard button released: id: %d, button: %d", id, b)
        }
      }
    }
  }

  return nil
}

// -------------------------------------------------------------------------------------------------------------------

func (gp *Gamepad) Draw(screen *ebiten.Image) {
  // Draw the current gamepad status.
  str := ""
  if len(gp.gamepadIDs) > 0 {
    ids := make([]ebiten.GamepadID, 0, len(gp.gamepadIDs))
    for id := range gp.gamepadIDs {
      ids = append(ids, id)
    }
    sort.Slice(ids, func(a, b int) bool {
      return ids[a] < ids[b]
    })
    for _, id := range ids {
      var standard string
      if ebiten.IsStandardGamepadLayoutAvailable(id) {
        standard = " (Standard Layout)"
      }
      str += fmt.Sprintf("Gamepad (ID: %d, SDL ID: %s)%s:\n", id, ebiten.GamepadSDLID(id), standard)
      str += fmt.Sprintf("  Name:    %s\n", ebiten.GamepadName(id))
      str += fmt.Sprintf("  Axes:    %s\n", strings.Join(gp.axes[id], ", "))
      str += fmt.Sprintf("  Buttons: %s\n", strings.Join(gp.pressedButtons[id], ", "))
      if ebiten.IsStandardGamepadLayoutAvailable(id) {
        str += "\n"
        str += standardMap(id) + "\n"
      }
      str += "\n"
    }
  } else {
    str = "Please connect your gamepad."
  }
  ebitenutil.DebugPrint(screen, str)
}

// -------------------------------------------------------------------------------------------------------------------

var standardButtonToString = map[ebiten.StandardGamepadButton]string{
  ebiten.StandardGamepadButtonRightBottom:      "RB",
  ebiten.StandardGamepadButtonRightRight:       "RR",
  ebiten.StandardGamepadButtonRightLeft:        "RL",
  ebiten.StandardGamepadButtonRightTop:         "RT",
  ebiten.StandardGamepadButtonFrontTopLeft:     "FTL",
  ebiten.StandardGamepadButtonFrontTopRight:    "FTR",
  ebiten.StandardGamepadButtonFrontBottomLeft:  "FBL",
  ebiten.StandardGamepadButtonFrontBottomRight: "FBR",
  ebiten.StandardGamepadButtonCenterLeft:       "CL",
  ebiten.StandardGamepadButtonCenterRight:      "CR",
  ebiten.StandardGamepadButtonLeftStick:        "LS",
  ebiten.StandardGamepadButtonRightStick:       "RS",
  ebiten.StandardGamepadButtonLeftBottom:       "LB",
  ebiten.StandardGamepadButtonLeftRight:        "LR",
  ebiten.StandardGamepadButtonLeftLeft:         "LL",
  ebiten.StandardGamepadButtonLeftTop:          "LT",
  ebiten.StandardGamepadButtonCenterCenter:     "CC",
}

// -------------------------------------------------------------------------------------------------------------------

func standardMap(id ebiten.GamepadID) string {
  m := `       [FBL ]                    [FBR ]
       [FTL ]                    [FTR ]

       [LT  ]       [CC  ]       [RT  ]
    [LL  ][LR  ] [CL  ][CR  ] [RL  ][RR  ]
       [LB  ]                    [RB  ]
             [LS  ]       [RS  ]
`

  for b, str := range standardButtonToString {
    placeholder := "[" + str + strings.Repeat(" ", 4-len(str)) + "]"
    v := ebiten.StandardGamepadButtonValue(id, b)
    if ebiten.IsStandardGamepadButtonPressed(id, b) {
      m = strings.Replace(m, placeholder, fmt.Sprintf("[%0.2f]", v), 1)
    } else {
      m = strings.Replace(m, placeholder, fmt.Sprintf(" %0.2f ", v), 1)
    }
  }

  m += fmt.Sprintf("    Left Stick:  X: %+0.2f, Y: %+0.2f\n    Right Stick: X: %+0.2f, Y: %+0.2f",
    ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal),
    ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical),
    ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickHorizontal),
    ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickVertical))
  return m
}



