package world

import (
	"math/rand"
	"strings"
	"time"
)

type City struct {
	name       string
	directions []string

	referencedBy []*City

	invaders []int64
}

func (c *City) Name() string {
	return c.name
}

func (c *City) MoveIn(id int64) bool {
	for i := range c.invaders {
		if c.invaders[i] == id {
			return false
		}
	}
	c.invaders = append(c.invaders, id)

	return true
}

func (c *City) Invaders() []int64 {
	return c.invaders
}

func (c *City) Destroy() {
	for i := range c.referencedBy {
		c.referencedBy[i].deleteDirection(c.name)
	}
}

func (c *City) moveOut(id int64) {
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

func (c *City) next() string {

	size := len(c.directions)
	if size == 0 {
		return ""
	}

	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(c.directions))

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
