package world

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/lucasdss/alieninvasion/pkg/world/city"
)

func loadMap(fd io.Reader) (cities []*city.City, err error) {

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()

		cityData := strings.Split(strings.Trim(line, "\n"), " ")
		if len(cityData) < 1 {
			return nil, fmt.Errorf("could not load city from line; %s", line)
		}

		name := strings.Trim(cityData[0], "\t\n")
		directions := cityData[1:]

		cities = append(cities, city.New(name, directions))
	}

	err = scanner.Err()
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return cities, nil
}

func (w *World) PrintMap() {

	fmt.Printf("\n\n########## REMAINING WORLD ##########\n\n")
	for _, c := range w.worldMap {
		if !c.Destroyed() {
			fmt.Printf("%s\n", c)
		}
	}
}
