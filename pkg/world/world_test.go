package world

import (
	"bytes"
	"testing"

	"github.com/lucasdss/alieninvasion/pkg/world/city"
)

func TestWorld(t *testing.T) {

	tt := []struct {
		name     string
		data     *bytes.Buffer
		expected []*city.City
	}{
		{
			name: "multi cities",
			data: bytes.NewBufferString(`Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
LX east=Bar west=Foo noth=Bee`),
			expected: []*city.City{
				city.New("Foo", []string{"north=Bar", "west=Baz", "south=Qu-ux"}),
				city.New("Bar", []string{"south=Foo", "west=Bee"}),
				city.New("LX", []string{"east=Bar", "west=Foo", "noth=Bee"}),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			world, err := New("X", tc.data)
			if err != nil {
				t.Error(err)
			}

			if len(tc.expected) != len(world.worldMap) {
				t.Errorf("expected %d; got %d", len(tc.expected), len(world.worldMap))
			}

			for _, e := range tc.expected {
				var found bool
				for _, c := range world.worldMap {
					if c.Name() == e.Name() {
						if c.String() != e.String() {
							t.Errorf("expected %s; got %s", e.String(), c.String())
						}
						found = true
					}
				}
				if !found {
					t.Errorf("expected %s; got %#v", e.Name(), world.worldMap)
				}
			}

		})
	}
}
