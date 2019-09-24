package world

import (
	"bytes"
	"testing"
)

func TestMapLoader(t *testing.T) {

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
			cities, err := loadMap(tc.data)
			if err != nil {
				t.Error(err)
			}

			for _, e := range tc.expected {
				var found bool
				for _, c := range cities {
					if e.name == c.name {
						found = true
					}
				}
				if !found {
					t.Errorf("city %s not found", e.name)
				}
			}

		})
	}
}
