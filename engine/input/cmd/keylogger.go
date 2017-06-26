package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/tactics/engine"
)

func main() {
	eng, _ := engine.NewSDLEngine("keylogger")
	defer eng.Window().Destroy()

	cont := true
	for cont {
		_ = eng.ClearScreen()
		for _, e := range eng.Events() {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				cont = false
			case *sdl.KeyDownEvent:
				log.Println("key pressed:", t.Keysym.Sym)
			case *sdl.KeyUpEvent:
				//
			}
		}
		_ = eng.UpdateSurface()
		eng.PauseRendering(1000)
	}
}
