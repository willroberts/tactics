package main

import (
	"log"

	"github.com/willroberts/tactics/engine"
	"github.com/willroberts/tactics/game/scenes"
)

func main() {
	eng, err := engine.NewSDLEngine("tactics")
	if err != nil {
		log.Fatalln("failed to initalize sdl engine:", err)
	}
	defer eng.Window().Destroy()

	sceneList := []engine.Scene{
		scenes.NewMainMenuScene(eng),
		scenes.NewIsometricScene(eng),
	}

	for _, scene := range sceneList {
		if err = scene.Setup(); err != nil {
			log.Fatalln("error during menu scene setup:", err)
		}
		for {
			if err := scene.Main(); err != nil {
				if err == scenes.ErrQuitGame {
					_ = scene.Teardown()
					return
				} else if err == scenes.ErrEndScene {
					break
				} else {
					log.Println("error:", err)
					break
				}
			}
		}
		if err = scene.Teardown(); err != nil {
			log.Println("error tearing down menu scene:", err)
		}
	}
}
