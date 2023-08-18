package main

import (
	"fmt"

	"github.com/digitalhouse-content/go-fundamentals-bases/interface/vehicles"
)

func main() {

	Display(123)
	Display("hola")
	Display(true)
	Display(123.78)

	vArray := []string{"CAR", "MOTORCYCLE", "PLANE", "TRUCK", "MOTORCYCLE", "TRUCK", "CAR"}

	var d float64
	for _, v := range vArray {
		fmt.Printf("Vehicle: %s\n", v)

		veh, err := vehicles.New(v, 400)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}

		distance := veh.Distance()
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()
		d += distance
	}

	fmt.Printf("Total distance: %.2f", d)
	fmt.Println()
}

func Display(value interface{}) {
	fmt.Println(value)
}
