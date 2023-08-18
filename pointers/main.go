package main

import (
	"fmt"
)

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p\n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("address: %p\n", m)
	m.Name = name
}

func main() {
	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println(&i)
	fmt.Println(iP)
	fmt.Println(i)
	fmt.Println(*iP)

	*iP = 1

	fmt.Printf("value i: %d, value pointer i: %d\n", i, *iP)

	myVar := 30
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)
	increment(myVar)
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)
	incrementP(&myVar)
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("address: %p, values: %v\n", mySlice, mySlice)
	fmt.Printf("addr1: %p, addr2: %p, addr3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println(mySlice)

	myStruct := &MyStruct{ID: 1234, Name: "Test"}
	fmt.Printf("address: %p\n", myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	fmt.Println()
	myStruct.Set("test method")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	myStruct.SetP("test method 2")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

}

func increment(val int) {
	val++
	fmt.Println(&val)
	fmt.Println(val)
}

func incrementP(val *int) {
	*val++
	fmt.Println(&val)
	fmt.Println(val)
	fmt.Println(*val)
}

func updateSlice(v []int) {
	fmt.Printf("addrs: %p\n", v)
	fmt.Printf("addr1: %p, addr2: %p, addr3: %p\n", &v[0], &v[1], &v[2])
	v[0] = 12
}

func updateStruct(v *MyStruct) {
	fmt.Printf("address in function: %p\n", v)
	v.ID = 999
	v.Name = "Update"
}
