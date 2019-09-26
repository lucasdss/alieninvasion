package world

import (
	"errors"
	"io"
	"math/rand"
	"time"

	"github.com/lucasdss/alieninvasion/pkg/world/city"
)

var (
	ErrNoDirection = errors.New("trapped")
)

type World struct {
	name string

	worldMap map[string]*city.City
}

func New(name string, worldMap io.Reader) (*World, error) {
	cities, err := loadMap(worldMap)
	if err != nil {
		return nil, err
	}

	world := World{
		name:     name,
		worldMap: make(map[string]*city.City),
	}

	for _, c := range cities {
		world.worldMap[c.Name()] = c
	}

	world.buildReferences()

	return &world, nil
}

func (w *World) buildReferences() {

	for _, worldCity := range w.worldMap {
		for _, n := range worldCity.Directions() {

			c, ok := w.worldMap[n]
			if ok {
				// I am assuming the city direction might not exist.
				c.AddReference(worldCity)
			}

		}
	}
}

// RandomCity is used to pick up any city inside the world.
func (w *World) RandomCity() *city.City {

	size := len(w.worldMap)
	ri := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(size)

	var i int
	for _, c := range w.worldMap {
		if i >= ri {
			if !c.Destroyed() {
				return c
			}
		}
		i++
	}

	return nil
}

// City returns a pointer to city.City if the
// city name is found. Otherwise the default
// value is returned, it means a nil value.
// In the case, there is a direction but the city
// was never created, I decided to keep the alien at the
// same city until its interaction is over.
func (w *World) City(name string) *city.City {

	return w.worldMap[name]
}
