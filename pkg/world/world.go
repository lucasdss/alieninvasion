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
				// I am assming the city direction might not exist.
				c.AddReference(worldCity)
			}

		}
	}
}

func (w *World) City() *city.City {

	size := len(w.worldMap)
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(size)

	var city *city.City
	for _, c := range w.worldMap {
		if i == 0 {
			city = c
		}
		i--
	}

	return city
}

func (w *World) Travel(id int64, c *city.City) *city.City {
	// If there is no more direction
	// the alien is trapped.
	cityName := c.Next(id)
	if cityName == "" {
		return nil
	}

	// In this case there is a direction
	// but the city was never created.
	// I decided to keep the alien at the
	// same city until its interactions is over.
	city, ok := w.worldMap[cityName]
	if !ok {
		return c
	}

	return city
}
