package world

import "fmt"

// PrintMap is used to display what is remaining of the World.
func (w *World) PrintMap() {

	fmt.Printf("\n\n########## REMAINING WORLD ##########\n\n")
	for _, c := range w.worldMap {
		if !c.Destroyed() {
			fmt.Printf("%s\n", c)
		}
	}
}
