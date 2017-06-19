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

	scene := scenes.NewNineByFiveScene(eng)
	if err = scene.Setup(); err != nil {
		log.Fatalln("error during scene setup:", err)
	}

	for {
		scene.Main()
	}

	if err = scene.Teardown(); err != nil {
		log.Println("error during scene teardown:", err)
	}
}
