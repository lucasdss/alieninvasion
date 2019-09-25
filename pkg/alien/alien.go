package alien

const (
	MaxInteractions = 10000
)

type city interface {
	MoveIn(id int64)
	Invaders() []int64
	MoveOut(id int64)
	Destroy()
}

type Alien struct {
	id int64

	interactions int64
}

func New(id int64) *Alien {
	return &Alien{id: id}
}

func (a *Alien) Leave(c city) {
	c.MoveOut(a.id)
}

func (a *Alien) MoveIn(c city) {
	c.MoveIn(a.id)
}

func (a *Alien) Attack(c city) (invaders []int64, destroyed bool) {

	invaders = c.Invaders()
	if len(invaders) == 2 {
		c.Destroy()
		destroyed = true
	}

	return invaders, destroyed
}

func (a *Alien) ID() int64 {
	return a.id
}

func (a *Alien) Continue() bool {
	a.interactions++
	return a.interactions < MaxInteractions
}
