package main

import "fmt"

func main() {

	map1 := make(map[int]string)

	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "C"
	map1[99] = "Z"
	map1[-5] = "a"

	fmt.Println(map1)
	fmt.Println(map1[-5])

	map2 := make(map[int][]string)
	map2[1] = []string{"A", "B"}
	map2[3] = []string{"Z", "A", "9"}

	fmt.Println(map2)
	fmt.Println(map2[3])

	map1[1] = "LA"

	delete(map1, 2)
	delete(map1, 3)
	map1[9] = ""
	fmt.Println(map1)
	fmt.Println(map1[8])

	v, ok := map1[9]
	fmt.Println("El valor es: ", v, " - existe: ", ok)

	v, ok = map1[8]
	fmt.Println("El valor es: ", v, " - existe: ", ok)

	map3 := map[int]string{
		4: "A",
		1: "N",
		8: "90",
	}
	fmt.Println(map3)

	for k, v := range map1 {
		fmt.Println("key: ", k, " - value: ", v)
	}

	for key, value := range map2 {
		fmt.Println("key: ", key)

		for _, v := range value {
			fmt.Println("Value: ", v)
		}

		fmt.Println()
	}
}
