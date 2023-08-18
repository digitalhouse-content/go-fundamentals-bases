package main

import (
	"fmt"
)

func main() {

	yearsOld := 16

	if yearsOld > 18 {
		fmt.Printf("%d es mayor a 18\n", yearsOld)
	}

	boolVar := true

	if boolVar {
		fmt.Println("es verdadero")
	} else {
		fmt.Println("es false")
	}

	if value := true; value {
		fmt.Println("es verdadero")
	}

	number := 0

	if number == 1 {
		fmt.Println("uno")
	} else if number == 2 {
		fmt.Println("dos")
	} else if number == 3 {
		fmt.Println("tres")
	} else {
		fmt.Println("ninguna es valida")
	}

	switch number {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	default:
		fmt.Println("ninguna es valida")
	}

	switch number := 1; number {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	default:
		fmt.Println("ninguna es valida")
	}

	switch {
	case number > 0:
		fmt.Println("positivo")
	case number < 0:
		fmt.Println("negativo")
	case number == 0:
		fmt.Println("es cero")
	}
}
