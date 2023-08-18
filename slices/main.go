package main

import (
	"fmt"
)

func main() {

	myArrayVar := [5]int{3, 6, 9, 10, 16}
	fmt.Println("array: ", myArrayVar, " - len: ", len(myArrayVar))

	mySliceVar := []int{}

	mySliceVar = append(mySliceVar, 12, 34, 54, 31, 12)
	fmt.Println("slice: ", mySliceVar, " - len: ", len(mySliceVar))

	mySliceVar2 := myArrayVar[2:4]

	mySliceVar2[0] = 19
	fmt.Println("slice: ", mySliceVar2, " - len: ", len(mySliceVar2), " - address: ", &mySliceVar2[0])

	fmt.Println("array: ", myArrayVar, " - len: ", len(myArrayVar), " - address: ", &myArrayVar[2])

	mySliceVar3 := mySliceVar[2:]
	fmt.Println(mySliceVar3)

	mySliceVar4 := make([]int, 80)
	fmt.Println("slice: ", mySliceVar4, " - len: ", len(mySliceVar4))

	mySliceVar5 := []int{1, 2, 6, 11, 20, 5, 1, 0}
	fmt.Println(mySliceVar5)
}
