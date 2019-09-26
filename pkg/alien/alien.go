package alien

const (
	// MaxInteractions was defined by the the exercise
	// Every time an Alien achive this number of interactions
	// it must stop and leave the planet.
	MaxInteractions = 10000
)

type city interface {
	MoveIn(id int64)
	Invaders() []int64
	MoveOut(id int64)
	Destroy()
}

// Alien is used to wrapping some logic and also
// define some behavior.
type Alien struct {
	id int64

	interactions int64
}

// New is helper funntion
func New(id int64) *Alien {
	return &Alien{id: id}
}

// Leave is wrapping city.MoveOut()
func (a *Alien) Leave(c city) {
	c.MoveOut(a.id)
}

// MoveIn is wrapping city.MoveIn()
func (a *Alien) MoveIn(c city) {
	c.MoveIn(a.id)
}

// Attack is used to check the number of invaders inside a city.
// If there are two aliens inside the city the c.Destroy() is called.
func (a *Alien) Attack(c city) (invaders []int64, destroyed bool) {

	invaders = c.Invaders()
	if len(invaders) == 2 {
		c.Destroy()
		destroyed = true
	}

	return invaders, destroyed
}

// ID returns the alien ID
func (a *Alien) ID() int64 {
	return a.id
}

// Continue is used to validate it the MaxInteractions has been achieved.
// It also increments the alien interaction counter.
func (a *Alien) Continue() bool {
	a.interactions++
	return a.interactions < MaxInteractions
}
