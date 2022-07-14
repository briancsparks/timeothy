package timeothy

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  hook "github.com/robotn/gohook"
  "strconv"
  "sync"
)

func f() {
	fmt.Printf("\n")
}

type DynamicConfig struct {
  Data      map[string]string

  actions   map[string]func()
}
var dynFig *DynamicConfig
var evChan chan hook.Event

func init() {
  dynFig = &DynamicConfig{}
  dynFig.actions = map[string]func(){}
  dynFig.Data = map[string]string{}

  evChan = hook.Start()

  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    wg.Done()

    for {
      select {
      case ev := <- evChan:
        asciiCode := int(ev.Keychar)
        s := string(rune(asciiCode))
        fn, present := dynFig.actions[s]
        if !present {
          //fmt.Printf("Dynamic config: unknown action: %v\n", ev)
          continue
        }
        fn()

      }
    }

  }()
  wg.Wait()
}

func (dy *DynamicConfig) MakeFloat64IncDecActions(key, incName, decName string, delta float64) {
  incFn := dy.MakeFloat64IncFn(key, delta)
  decFn := dy.MakeFloat64IncFn(key, -delta)

  dy.actions[incName] = incFn
  dy.actions[decName] = decFn
}

func (dy *DynamicConfig) MakeFloat64IncFn(key string, delta float64) func() {
  return func() {
    x, _ := dy.GetFloat64(key)
    dy.SetFloat64(key, x+delta)
    fmt.Printf("%s = %f\n", key, x+delta)
  }
}

func (dy *DynamicConfig) Get(key string) string {
  _, present := dy.Data[key]
  if !present {
    dy.Data[key] = ""
  }
  return dy.Data[key]
}

func (dy *DynamicConfig) Set(key, value string) {
  dy.Data[key] = value
}

func (dy *DynamicConfig) GetFloat64(key string) (float64,error) {
  data := dy.Get(key)
  if data == "" {
    return 0.0, nil
  }
  return strconv.ParseFloat(data, 64)
}

func (dy *DynamicConfig) SetFloat64(key string, value float64) {
  s := fmt.Sprintf("%f", value)
  dy.Set(key, s)
}
