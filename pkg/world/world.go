package world

import (
	"errors"
	"io"
	"math/rand"
	"time"
)

var (
	ErrNoDirection = errors.New("trapped")
)

type World struct {
	name string

	worldMap map[string]*City

	invaders map[int64]*City
}

func New(name string, worldMap io.Reader) (*World, error) {
	cities, err := loadMap(worldMap)
	if err != nil {
		return nil, err
	}

	world := World{
		name:     name,
		worldMap: make(map[string]*City),
		invaders: make(map[int64]*City),
	}

	for i := range cities {
		c := cities[i]
		world.worldMap[c.name] = &cities[i]
	}

	world.buildReferences()

	return &world, nil
}

func (w *World) buildReferences() {

	for _, city := range w.worldMap {
		for _, d := range city.directions {
			n := getCityName(d)
			c, ok := w.worldMap[n]
			if ok {
				c.referencedBy = append(c.referencedBy, city)
			}
		}
	}
}

func (w *World) City() *City {

	size := len(w.worldMap)
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(size)

	var city *City
	for _, c := range w.worldMap {
		if i == 0 {
			city = c
		}
		i--
	}

	return city
}

func (w *World) Travel(c *City) (*City, error) {
	// If there is no more direction
	// the alien is trapped.
	cityName := c.next()
	if cityName == "" {
		return nil, ErrNoDirection
	}

	// In this case there is a direction
	// but the city was never created.
	// I decided to keep the alien at the
	// same city until its interactions is over.
	city, ok := w.worldMap[cityName]
	if !ok {
		return c, nil
	}

	return city, nil
}
