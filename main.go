package main

import (
	"log"

	"github.com/willroberts/tactics/unit"
)

func main() {
	u := unit.NewRandomizedUnit()
	log.Println("unit:", u)
}
