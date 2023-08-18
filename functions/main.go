package main

import (
	"fmt"

	"github.com/digitalhouse-content/go-fundamentals-bases/functions/function"
)

func main() {
	var myIntVar int64
	var myIntVar1 int64 = 30
	var myIntVar2 int64 = 50
	var myIntVar3 int64 = 90

	function.Display(myIntVar)
	function.Display(myIntVar1)
	function.Display(myIntVar2)
	function.Display(myIntVar3)

	v := function.Add(4, 2)
	fmt.Println(v)

	function.RepeatString(10, "LA")

	value, err := function.Calc(function.SUM, 20.12, 13)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("value: ", value)

	xVal, yVal := function.Split(40)
	fmt.Println(xVal, yVal)

	val2 := function.MSum(1, 2, 3, 1, 2, 3, 4, 90, 100)
	fmt.Println(val2)

	mOperVal, err := function.MOperations(function.DIV, 8)
	fmt.Println("value: ", mOperVal, " - err: ", err)

	factOpFunc := function.FactoryOperation(function.DIV)
	fmt.Println(factOpFunc(10, 0))
}
