package main

import (
	"fmt"
)

func main() {

	yearsOld := 33

	fmt.Println(yearsOld > 30)  // true
	fmt.Println(yearsOld < 33)  // false
	fmt.Println(yearsOld <= 33) // true
	fmt.Println(yearsOld >= 40) // false
	fmt.Println(yearsOld == 33) // true

	fmt.Println(yearsOld < 33 || yearsOld == 33) // (false or true) = true
	fmt.Println(yearsOld < 33 || yearsOld == 34) // (false or false) = false
	fmt.Println(yearsOld < 40 || yearsOld == 33) // (true or true ) = true

	fmt.Println(yearsOld < 33 && yearsOld == 33) // (false or true) = false
	fmt.Println(yearsOld < 33 && yearsOld == 34) // (false or false) = false
	fmt.Println(yearsOld < 40 && yearsOld == 33) // (true or true ) = true

	fmt.Println(true)
	fmt.Println(!true)
	fmt.Println(yearsOld < 40)    // true
	fmt.Println(!(yearsOld < 40)) // false

	fmt.Println(yearsOld < 25 && yearsOld == 33 || yearsOld < 40)   // (false and true  or true) = true
	fmt.Println(yearsOld < 25 && (yearsOld == 33 || yearsOld < 40)) // (false and (true or true)) = false

}
