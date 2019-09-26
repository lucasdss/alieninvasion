package goinvasion

import (
	"fmt"
	"sync"

	"github.com/lucasdss/alieninvasion/pkg/alien"
	"github.com/lucasdss/alieninvasion/pkg/world"
	"github.com/lucasdss/alieninvasion/pkg/world/city"
)

type GoInvasion struct {
	world     *world.World
	numAliens int64
}

type (
	AttackRequest struct {
		alien    *alien.Alien
		prevCity *city.City
		respCh   chan<- AttackResponse
	}

	AttackResponse struct {
		cityName  string
		aliens    []int64
		destroyed bool
		city      *city.City
	}
)

func New(numAliens int64, w *world.World) *GoInvasion {

	inv := GoInvasion{
		world:     w,
		numAliens: numAliens,
	}

	return &inv
}

func (g *GoInvasion) Start() {

	attackCh := make(chan AttackRequest)
	defer close(attackCh)

	var stomp sync.WaitGroup
	var wg sync.WaitGroup

	go func() {
		goInvasion(g.world, attackCh)
	}()

	stomp.Add(1)

	for i := int64(0); i < g.numAliens; i++ {
		wg.Add(1)

		a := alien.New(i)
		go func(a *alien.Alien) {
			defer wg.Done()
			stomp.Wait()

			goAlien(a, attackCh)
		}(a)
	}

	stomp.Done()
	wg.Wait()
}

func goAlien(a *alien.Alien, attackCh chan<- AttackRequest) {

	var c *city.City
	attackResponseCh := make(chan AttackResponse)

	for {

		if !a.Continue() {
			msg := fmt.Sprintf("Alien %d could not attack the planet.", a.ID())
			if c != nil {
				msg = fmt.Sprintf("Alien %d leaved the planet from city %s.", a.ID(), c.Name())
			}
			fmt.Println(msg)
			return
		}

		if c != nil {
			if c.Destroyed() {
				return
			}
		}

		attack := AttackRequest{
			alien:    a,
			prevCity: c,
			respCh:   attackResponseCh,
		}

		attackCh <- attack

		attackRes := <-attackResponseCh

		aliens := attackRes.aliens
		if attackRes.destroyed {
			fmt.Printf("%s has been destroyed by alien %d and alien %d.\n", attackRes.cityName, aliens[0], aliens[1])
			return
		}

		c = attackRes.city
	}

}

func goInvasion(w *world.World, attackCh <-chan AttackRequest) {

	for {

		attack, ok := <-attackCh
		if !ok {
			return
		}
		var cityName string
		var aliens []int64
		var destroyed bool

		var attackCity *city.City

		if attack.prevCity != nil {
			prevCity := attack.prevCity

			destroyed = prevCity.Destroyed()
			cityName = prevCity.Name()
			aliens = prevCity.Invaders()

			if destroyed {
				attack.respCh <- AttackResponse{cityName, aliens, false, prevCity}
				continue
			}

			attackCity = w.City(prevCity.Next())
			if attackCity == nil {
				attack.respCh <- AttackResponse{cityName, aliens, destroyed, prevCity}
				continue
			}

			attack.alien.Leave(attack.prevCity)

			cityName = attackCity.Name()
			attack.alien.MoveIn(attackCity)
			aliens, destroyed = attack.alien.Attack(attackCity)

			attack.respCh <- AttackResponse{cityName, aliens, destroyed, attackCity}
			continue
		}

		attackCity = w.RandomCity()
		if attackCity != nil {
			cityName = attackCity.Name()
			attack.alien.MoveIn(attackCity)
			aliens, destroyed = attack.alien.Attack(attackCity)

			attack.respCh <- AttackResponse{cityName, aliens, destroyed, attackCity}
			continue
		}

		attack.respCh <- AttackResponse{cityName, aliens, false, nil}
	}

}
