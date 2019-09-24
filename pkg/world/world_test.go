package world

import (
	"bytes"
	"testing"
)

func TestWorld(t *testing.T) {

	tt := []struct {
		name     string
		data     *bytes.Buffer
		expected []City
	}{
		{
			name: "multi cities",
			data: bytes.NewBufferString(`Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
LX east=Bar west=Foo noth=Bee`),
			expected: []City{
				City{
					name:       "Foo",
					directions: []string{"north=Bar", "west=Baz", "south=Qu-ux"},
				},
				City{
					name:       "Bar",
					directions: []string{"south=Foo", "west=Bee"},
				},
				City{
					name:       "LX",
					directions: []string{"east=Bar", "west=Foo", "noth=Bee"},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			world, err := New("X", tc.data)
			if err != nil {
				t.Error(err)
			}

			for _, c := range world.cities {
				t.Logf("\n%s - direction: %#v - referenced: ", c.name, c.directions)
				for _, ref := range c.referencedBy {
					t.Logf("city=%#v", ref.name)
				}
			}

		})
	}
}
