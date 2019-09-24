package main

import (
	"flag"
	"log"
	"os"

	"github.com/lucasdss/alieninvasion/internal/invasion"
	"github.com/lucasdss/alieninvasion/pkg/world"
)

func main() {

	var (
		mapfile   string
		numAliens int64
	)

	flag.StringVar(&mapfile, "map", "assets/world_map.txt", "path to world map file.")
	flag.Int64Var(&numAliens, "aliens", 5, "number os aliens invading the planet.")
	flag.Parse()

	fd, err := os.Open(mapfile)
	if err != nil {
		log.Fatal(err)
	}

	w, err := world.New("X", fd)
	if err != nil {
		log.Fatal(err)
	}

	inv := invasion.New(numAliens, w)

	inv.Start()

	// Print the world map again

}
