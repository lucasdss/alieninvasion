package invasion

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lucasdss/alieninvasion/pkg/alien"
	"github.com/lucasdss/alieninvasion/pkg/world"
	"github.com/lucasdss/alieninvasion/pkg/world/city"
)

type Invasion struct {
	alienCity map[int64]*city.City
	aliens    []*alien.Alien
	world     *world.World
}

func New(numAliens int64, w *world.World) *Invasion {

	inv := Invasion{
		world:     w,
		alienCity: make(map[int64]*city.City),
	}

	for i := int64(0); i < numAliens; i++ {
		inv.aliens = append(inv.aliens, alien.New(i))
	}

	return &inv
}

func (inv *Invasion) Start() {

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var i int64

	for {

		size := int64(len(inv.aliens))
		if size == 0 {
			break
		}

		i = rd.Int63n(size)

		alien := inv.aliens[i]

		id := alien.ID()

		if !alien.Continue() {
			inv.stop(id)
			fmt.Printf("%d leaved the planed\n", id)
			continue
		}

		c, ok := inv.alienCity[id]
		if !ok {
			c = inv.world.City()

			inv.alienCity[id] = c
			invaders, destroyed := alien.Attack(c)
			if destroyed {
				fmt.Printf("%s has been destroyed by alien %d and alien %d.\n", c.Name(), invaders[0], invaders[1])
				inv.stop(invaders[0])
				inv.stop(invaders[1])
			}
			continue
		}

		city := inv.world.Travel(id, c)
		if city == nil {
			continue
		}

		inv.alienCity[id] = city
		invaders, destroyed := alien.Attack(city)
		if destroyed {
			fmt.Printf("%s has been destroyed by alien %d and alien %d.\n", city.Name(), invaders[0], invaders[1])
			inv.stop(invaders[0])
			inv.stop(invaders[1])
		}

	}
}

func (inv *Invasion) stop(id int64) {
	size := int64(len(inv.aliens))

	for i, invader := range inv.aliens {
		if invader.ID() == id {
			inv.aliens[i] = inv.aliens[size-1]
			inv.aliens = inv.aliens[:size-1]
			return
		}
	}

	inv.aliens = []*alien.Alien{}
}
