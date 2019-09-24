package invasion

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lucasdss/alieninvasion/pkg/alien"
	"github.com/lucasdss/alieninvasion/pkg/world"
)

type Invasion struct {
	alienCity map[int64]*world.City
	aliens    []*alien.Alien
	world     *world.World
}

func New(numAliens int64, w *world.World) *Invasion {

	inv := Invasion{
		world:     w,
		alienCity: make(map[int64]*world.City),
	}

	for i := int64(0); i < numAliens; i++ {
		inv.aliens = append(inv.aliens, alien.New(i))
	}

	return &inv
}

func (inv *Invasion) Start() {

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		size := int64(len(inv.aliens))

		i := rd.Int63n(size)

		alien := inv.aliens[i]

		id := alien.ID()

		c, ok := inv.alienCity[id]
		if !ok {
			c = inv.world.City()

			inv.alienCity[id] = c
			invaders, destroyed := alien.Attack(c)
			if destroyed {
				fmt.Printf("%s has been destroyed by alien %d and alien %d.\n", c.Name(), invaders[0], invaders[1])
			}
			continue
		}

		city, err := inv.world.Travel(c)
		if err != nil {
			inv.stop(id)
			fmt.Printf("%d in city %s got %s\n", id, c.Name(), err)
			continue
		}

		inv.alienCity[id] = city
		invaders, destroyed := alien.Attack(city)
		if destroyed {
			fmt.Printf("%s has been destroyed by alien %d and alien %d.\n", city.Name(), invaders[0], invaders[1])
		}

		if !inv.aliens[i].Continue() {
			inv.stop(i)
			fmt.Printf("%d leaved the planed\n", id)
		}

		if len(inv.aliens) < 1 {
			fmt.Printf("No more aliens to move; last %d", id)
			break
		}

	}
}

func (inv *Invasion) stop(i int64) {
	size := int64(len(inv.aliens))
	if size > 0 {
		inv.aliens[i] = inv.aliens[size-1]
		inv.aliens = inv.aliens[:size-1]
		return
	}

	inv.aliens = []*alien.Alien{}
}
