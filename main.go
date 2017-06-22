package main

import (
	"log"

	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/game/scenes"
)

const (
	winW int = 900
	winH int = 500
)

func main() {
	eng, err := engine.NewSDLEngine("tactics", winW, winH)
	if err != nil {
		log.Fatalln("failed to initalize sdl engine:", err)
	}
	defer eng.Window().Destroy()

	// SCENE 1: MAIN MENU
	scene := scenes.NewMainMenuScene(eng)
	if err = scene.Setup(); err != nil {
		log.Fatalln("error during menu scene setup:", err)
	}

	for {
		err = scene.Main()
		if err == scenes.ErrQuitting {
			_ = scene.Teardown()
			log.Println("quitting")
			return
		}
		if err != nil {
			log.Println("error:", err)
			break
		}
	}

	if err = scene.Teardown(); err != nil {
		log.Println("error tearing down menu scene:", err)
	}

	// SCENE 2: ISOMETRIC
	scene = scenes.NewIsometricScene(eng)
	if err = scene.Setup(); err != nil {
		log.Fatalln("error during scene setup:", err)
	}

	for {
		err = scene.Main()
		if err != nil {
			log.Println("error:", err)
			break
		}
	}

	if err = scene.Teardown(); err != nil {
		log.Println("error during scene teardown:", err)
	}
}
