package city

import (
	"testing"
)

func TestCityInvadersInOut(t *testing.T) {

	tt := []struct {
		name     string
		city     City
		moveIn   []int64
		moveOut  []int64
		expected []int64
	}{
		{
			name:     "zero invaders",
			city:     City{},
			expected: []int64{},
		},
		{
			name:     "one invader",
			city:     City{},
			moveIn:   []int64{1},
			expected: []int64{1},
		},
		{
			name:     "two invaders",
			city:     City{},
			moveIn:   []int64{1, 2},
			expected: []int64{1, 2},
		},
		{
			name:     "two invaders in one out",
			city:     City{},
			moveIn:   []int64{1, 2},
			moveOut:  []int64{1},
			expected: []int64{2},
		},
		{
			name:     "two invaders in zero out",
			city:     City{},
			moveIn:   []int64{1, 2},
			moveOut:  []int64{1, 2},
			expected: []int64{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			for _, id := range tc.moveIn {
				tc.city.MoveIn(id)
			}

			for _, id := range tc.moveOut {
				tc.city.MoveOut(id)
			}

			if len(tc.expected) != len(tc.city.invaders) {
				t.Errorf("expected %d; got %d", len(tc.expected), len(tc.city.invaders))
			}

			for _, id := range tc.expected {
				var ok bool
				for _, invader := range tc.city.Invaders() {
					if invader == id {
						ok = true
					}
				}
				if !ok {
					t.Errorf("expected %d; got %#v", id, tc.city.Invaders())
				}
			}

		})
	}
}

func TestCityDirections(t *testing.T) {

	tt := []struct {
		name          string
		city          City
		directionsOut []string
		expected      []string
	}{
		{
			name:     "zero directions",
			city:     City{},
			expected: []string{},
		},
		{
			name:     "one direction",
			city:     City{directions: []string{"south=LX"}},
			expected: []string{"south=LX"},
		},
		{
			name:     "two directions",
			city:     City{directions: []string{"south=LX", "north=SP"}},
			expected: []string{"south=LX", "north=SP"},
		},
		{
			name:          "two directions one direction out",
			city:          City{directions: []string{"south=LX", "north=SP"}},
			directionsOut: []string{"south=LX"},
			expected:      []string{"north=SP"},
		},
		{
			name:          "two directions zero direction out",
			city:          City{directions: []string{"south=LX", "north=SP"}},
			directionsOut: []string{"south=LX", "north=SP"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			for _, d := range tc.directionsOut {
				tc.city.deleteDirection(getCityName(d))
			}

			if len(tc.expected) != len(tc.city.directions) {
				t.Errorf("expected %d; got %d", len(tc.expected), len(tc.city.directions))
			}

			for _, direction := range tc.expected {
				var ok bool
				for _, d := range tc.city.directions {
					if direction == d {
						ok = true
					}
				}
				if !ok {
					t.Errorf("expected %s; got %#v", direction, tc.city.directions)
				}
			}

		})
	}
}
