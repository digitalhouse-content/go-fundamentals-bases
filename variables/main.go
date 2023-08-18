package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string
	myStringVar = "my string variable"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("my variable is: ", myBoolVar)

	fmt.Println("my variable address is: ", &myStringVar)

	myIntVar2 := 12
	fmt.Println("my variable is: ", myIntVar2)

	myStringVar2 := "My string variable with :="
	fmt.Println("my variable is: ", myStringVar2)

	myBoolVar2 := true
	fmt.Println("my variable is: ", myBoolVar2)

	const myIntConst int = 12
	fmt.Println("Mi constante es: ", myIntConst)

	const myFirstStringConst = "a12"
	fmt.Println("Mi constante es:", myFirstStringConst)

	const myStringConst string = "50"
	fmt.Println("Mi constante es: ", myStringConst)

	myOtherScopeVariable := 50

	{
		myScopeVariable := 40
		{
			fmt.Println("Mi variable: ", myOtherScopeVariable)
			fmt.Println("Mi variable: ", myScopeVariable)
		}

	}

	fmt.Println("Mi variable: ", myOtherScopeVariable)
	// fmt.Println("Mi variable: ", myScopeVariable)

	var myIntVarSize int = 23

	fmt.Printf("%d \n", myIntVarSize)
	fmt.Printf("%T \n", myIntVarSize)

	/*
		int8	8 bits	-128 to 127
		int16	16 bits	-2^15 to 2^15 -1
		int32	32 bits	-2^31 to 2^31 -1
		int64	64 bits	-2^63 to 2^63 -1
		int	Platform dependent	Platform dependent
	*/

	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", myIntVarSize, myIntVarSize, unsafe.Sizeof(myIntVarSize), unsafe.Sizeof(myIntVarSize)*8)

	// int8	8 bits	-128 to 127
	var my8BitsIntVar int8 = 127
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsIntVar, my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16 = 127
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my16BitsIntVar, my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32 = 200
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my32BitsIntVar, my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64 = 400
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my64BitsIntVar, my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	/*
		Type	Size	Range
		uint8	8 bits	0 to 255
		uint16	16 bits	0 to 2^16 -1
		uint32	32 bits	0 to 2^32 -1
		uint64	64 bits	0 to 2^64 -1
		uint	Platform dependent	Platform dependent
	*/

	var my8BitsUintVar uint8 = 20
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsUintVar, my8BitsUintVar, unsafe.Sizeof(my8BitsUintVar), unsafe.Sizeof(my8BitsUintVar)*8)

	var my16BitsUintVar uint16 = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my16BitsUintVar, my16BitsUintVar, unsafe.Sizeof(my16BitsUintVar), unsafe.Sizeof(my16BitsUintVar)*8)

	var my32BitsUint uint32 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my32BitsUint, my32BitsUint, unsafe.Sizeof(my32BitsUint), unsafe.Sizeof(my32BitsUint)*8)

	var my64BitsUint uint64 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my64BitsUint, my64BitsUint, unsafe.Sizeof(my64BitsUint), unsafe.Sizeof(my64BitsUint)*8)

	var myUintSizeVar uint = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", myUintSizeVar, myUintSizeVar, unsafe.Sizeof(myUintSizeVar), unsafe.Sizeof(myUintSizeVar)*8)

	var myFloat32Var float32 = 3.14
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64 = 590.12435
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	myOtherFloat := 7654.123
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myOtherFloat, myOtherFloat, unsafe.Sizeof(myOtherFloat), unsafe.Sizeof(myOtherFloat)*8)

	var myStringVar3 string = "test1234"
	fmt.Printf("mi valor %s \n", myStringVar3)

	myStringVar4 := `Mi variable de tipo texto en go
	   con multiples
	   lienas
	   :)`
	fmt.Printf("mi valor es %s \n", myStringVar4)

	{

		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)

		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)

		intStrVar := fmt.Sprintf("%d - %d", intVar, 23)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		intVar2 := 342
		intStrVar2 := strconv.Itoa(intVar2)
		fmt.Printf("type: %T, value: %s\n", intStrVar2, intStrVar2)

		strIntVar2, err := strconv.Atoi("15")
		fmt.Printf("type: %T, value: %d, err: %v\n", strIntVar2, strIntVar2, err)

		strInvVar3, err := strconv.ParseInt("10", 10, 64)
		fmt.Printf("type: %T, value: %d, err: %v\n", strInvVar3, strInvVar3, err)

		strFloatVar, err := strconv.ParseFloat("-11.2", 64)
		fmt.Printf("type: %T, value: %f, err: %v\n", strFloatVar, strFloatVar, err)
	}
}
