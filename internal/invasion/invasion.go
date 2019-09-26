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
			city := inv.alienCity[id]
			msg := fmt.Sprintf("Alien %d could not attack the planet.", id)
			if city != nil {
				msg = fmt.Sprintf("Alien %d leaved the planet from city %s.", id, city.Name())
			}
			fmt.Println(msg)
			inv.stop(id)
			continue
		}

		c, ok := inv.alienCity[id]
		if !ok {
			c = inv.world.RandomCity()
			if c == nil {
				continue
			}
			fmt.Printf("Alien %d deployed in city %s\n", id, c.Name())
			inv.alienCity[id] = c
			alien.MoveIn(c)
		}

		// If there is no more direction
		// the alien is trapped.
		city := inv.world.City(c.Next())
		if city == nil {
			city = c
		}

		alien.Leave(c)
		alien.MoveIn(city)
		inv.alienCity[id] = city

		aliens, destroyed := alien.Attack(city)
		if destroyed {
			fmt.Printf("%s has been destroyed by alien %d and alien %d.\n", city.Name(), aliens[0], aliens[1])
			inv.stop(aliens[0])
			inv.stop(aliens[1])
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
