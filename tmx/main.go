package main

import (
	"log"

	"github.com/mewspring/tmx"
)

func main() {
	m, err := tmx.Open("grass.tmx")
	if err != nil {
		log.Fatal("failed to open grass.tmx: ", err)
	}
	log.Println("m: ", m)
	log.Println("first layer: ", m.Layers[0])
	log.Println("tiles: ", m.Layers[0].Data.Tiles)
	tiles := m.Layers[0].Data.Tiles
	for _, t := range tiles {
		log.Println("tile: ", t.GID)
	}
}
