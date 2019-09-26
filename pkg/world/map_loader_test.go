package world

import (
	"bytes"
	"testing"

	"github.com/lucasdss/alieninvasion/pkg/world/city"
)

func TestMapLoader(t *testing.T) {

	tt := []struct {
		name     string
		data     *bytes.Buffer
		expected []*city.City
	}{
		{
			name: "multi cities",
			data: bytes.NewBufferString(`Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
LX east=Bar west=Foo north=Bee`),
			expected: []*city.City{
				city.New("Foo", []string{"north=Bar", "west=Baz", "south=Qu-ux"}),
				city.New("Bar", []string{"south=Foo", "west=Bee"}),
				city.New("LX", []string{"east=Bar", "west=Foo", "north=Bee"}),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cities, err := loadMap(tc.data)
			if err != nil {
				t.Error(err)
			}

			for _, e := range tc.expected {
				var found bool
				for _, c := range cities {
					if e.Name() == c.Name() {
						found = true
					}
				}
				if !found {
					t.Errorf("city %s not found", e.Name())
				}
			}

		})
	}
}
