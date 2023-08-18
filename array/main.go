package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var myIntVar int = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar1 [5]int
	fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1))

	myArrayVar2 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar2, " - size: ", len(myArrayVar2))

	myArrayVar1[0] = 2
	myArrayVar1[1] = 5
	myArrayVar1[2] = 9
	fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1), " - get index 1: ", myArrayVar1[1])

	fmt.Printf("type: %T, value: %v, bytes: %d, bits: %d \n", myArrayVar1, myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myArrayVar1)*8)

	for i := range myArrayVar1 {
		fmt.Println("index: ", i, " - value: ", myArrayVar1[i])
	}

	for _, v := range myArrayVar1 {
		fmt.Println("value: ", v)
	}
}
