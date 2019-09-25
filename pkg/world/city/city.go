package city

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type City struct {
	destroyed bool

	name       string
	directions []string

	referencedBy []*City

	invaders []int64
}

func New(name string, directions []string) *City {

	return &City{name: name, directions: directions}
}

func (c *City) Name() string {
	return c.name
}

func (c *City) MoveIn(id int64) {
	c.invaders = append(c.invaders, id)
}

func (c *City) MoveOut(id int64) {
	size := len(c.invaders)
	for i, n := range c.invaders {
		if n == id {
			if size > 0 {
				c.invaders[i] = c.invaders[size-1]
				c.invaders = c.invaders[:size-1]
			} else {
				c.invaders = []int64{}
			}
		}
	}
}

func (c *City) Invaders() []int64 {
	return c.invaders
}

func (c *City) Destroy() {
	c.destroyed = true
	for i := range c.referencedBy {
		c.referencedBy[i].deleteDirection(c.name)
	}
}

func (c *City) Destroyed() bool {
	return c.destroyed
}

func (c *City) Directions() (cityNames []string) {

	for _, d := range c.directions {
		cityNames = append(cityNames, getCityName(d))
	}

	return cityNames
}

func (c *City) AddReference(city *City) {
	c.referencedBy = append(c.referencedBy, city)
}

func (c *City) String() string {
	out := fmt.Sprintf("%s", c.name)
	for _, d := range c.directions {
		out = fmt.Sprintf("%s %s", out, d)
	}

	return out
}

func (c *City) Next() string {

	size := len(c.directions)
	if size == 0 {
		return ""
	}

	i := rand.New(
		rand.NewSource(time.Now().UnixNano()),
	).Intn(len(c.directions))

	return getCityName(c.directions[i])
}

func (c *City) deleteDirection(cityName string) {
	size := len(c.directions)

	for i, d := range c.directions {
		name := getCityName(d)
		if cityName == name {
			if size > 1 {
				c.directions[i] = c.directions[size-1]
				c.directions = c.directions[:size-1]
			} else {
				c.directions = []string{}
			}
		}
	}
}

func getCityName(direction string) (cityName string) {
	v := strings.Split(direction, "=")
	if len(v) == 2 {
		return v[1]
	}

	return ""
}
