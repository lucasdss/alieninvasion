package world

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func loadMap(fd io.Reader) (cities []City, err error) {

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()

		cityData := strings.Split(strings.Trim(line, "\n"), " ")
		if len(cityData) < 1 {
			return nil, fmt.Errorf("could not load city from line; %s", line)
		}

		name := strings.Trim(cityData[0], "\t\n")
		directions := cityData[1:]

		cities = append(
			cities,
			City{
				name:       name,
				directions: directions,
			},
		)
	}

	err = scanner.Err()
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return cities, nil
}
