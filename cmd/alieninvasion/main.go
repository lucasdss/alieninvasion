package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lucasdss/alieninvasion/internal/goinvasion"
	"github.com/lucasdss/alieninvasion/internal/invasion"
	"github.com/lucasdss/alieninvasion/pkg/world"
)

func main() {

	var (
		mapfile    string
		numAliens  int64
		goInvasion bool
	)

	flag.StringVar(&mapfile, "map", "assets/world_map.txt", "path to world map file.")
	flag.Int64Var(&numAliens, "n", 5, "number os aliens invading the planet.")

	flag.BoolVar(&goInvasion, "c", true, "version with goroutines and channels.")
	flag.Parse()

	fd, err := os.Open(mapfile)
	if err != nil {
		log.Fatal(err)
	}

	w, err := world.New("X", fd)
	if err != nil {
		log.Fatal(err)
	}

	if goInvasion {
		fmt.Println("USING GOROUTINE VERSION")
		inv := goinvasion.New(numAliens, w)
		inv.Start()
	} else {
		fmt.Println("USING SIMPLE VERSION")
		inv := invasion.New(numAliens, w)

		inv.Start()
	}

	w.PrintMap()
}
