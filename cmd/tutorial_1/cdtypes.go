package main

import (
	"fmt"
)

func main3() {

	var intArr [3]int32
	fmt.Print(intArr[0])
	fmt.Print((intArr[0:3]))

	var intArr2 [3]int = [3]int{1, 2, 3}
	fmt.Print(intArr2)

	intArr3 := [...]int{1, 2, 3}
	fmt.Println(intArr3)

	// Slices: Wrappers around arrays. analogous to vectors.
	var intSlice []int = []int{4, 5, 6}
	fmt.Println(intSlice, len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice, len(intSlice), cap(intSlice))

	intSlice2 := make([]int, 3, 8)
	fmt.Println(intSlice2)

	var myMap map[string]int = make(map[string]int)
	fmt.Println(myMap)

	myMap["harsh"] = 2
	myMap["abc"] = 3
	fmt.Println(myMap)
	fmt.Println(myMap["harsh"])

	value, existFlag := myMap["abc"]
	fmt.Println(value, existFlag)

	myMap2 := map[string]int{"adam": 1, "eve": 2}
	delete(myMap2, "eve")
	fmt.Println(myMap2, len(myMap2))

	// Only way to iterate.
	for name, value := range myMap {
		fmt.Printf("Name: %v, Age: %v\n", name, value)
	}

	// for var name string = range myMap2 {
	// 	fmt.Printf("Name: %v\n", name)
	// }

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}
