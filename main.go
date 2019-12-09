package main

import (
	"fmt"
)

// Go syntax is awesome! Notice how you don't need explicit "enums" in Go since you are describing the object properties when building type structs (the building blocks for representing your data).
type Employee struct {
	department string
	position   string
}

type Planet struct {
	name string
	color  string
}

type Galaxy struct {
	name    string
	planets []Planet
}

func main() {

  // In Go we "initialize" a type struct
	newPerson := Employee{department: "Engineering", position: "ceo"}

  fmt.Println(newPerson)
  fmt.Println(newPerson.department)

  // An example of complex structs build from oter structs.
	milkyWay := Galaxy{
		name: "Milky Way",
		planets: []Planet{
			{name: "Earth",
				color: "blue and green",
			},
			{name: "Pluto",
				color: "purple",
			},
			{name: "Mars",
				color: "reddish orange",
			},
		},
	}

	fmt.Printf("The second thing is %s with the color %s.\n", milkyWay.planets[1].name, milkyWay.planets[1].color)


}
 
