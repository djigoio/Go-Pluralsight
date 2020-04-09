package learning

import (
	"fmt"
)

func learning() {
	fmt.Println("Hello Djigo")

	//Variable types
	var integer int
	integer = 12
	//Shorthand to init variables
	shortInteger := 12

	fmt.Println(integer, shortInteger)

	//Pointers, * means pointer
	var firstName *string = new(string)
	*firstName = "Tonio"
	fmt.Println(*firstName)

	secondName := "Tonio"
	ptr := &secondName
	fmt.Println(ptr, *ptr)

	//Constants
	const pi = 3.1415
	fmt.Println(pi)

	const (
		first  = iota
		second = iota + 2
		third
	)

	fmt.Println(first, second, third)

	//Arrays
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers, numbers[2])
	//Shorthand for arrays
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2, arr2[2])
	//Slices
	slice := arr2[:]
	fmt.Println(slice)

	coolSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(coolSlice)
	coolSlice = append(coolSlice, 6)
	fmt.Println(coolSlice)
	sliceOfSlice := coolSlice[1:4]
	fmt.Println(sliceOfSlice)

	//Maps

	m := map[string]int{"foo": 42}
	fmt.Println(m["foo"])
	delete(m, "foo")
	fmt.Println(m)

	//Structs

	type Estructura struct {
		ID       int
		Nombre   string
		Material string
	}

	var estruct Estructura
	estruct.ID = 1
	estruct.Nombre = "Edificio"
	estruct.Material = "Hormigón"

	estructDos := Estructura{ID: 2, Nombre: "Pared", Material: "Ladrillo"}
	fmt.Println(estruct)
	fmt.Println(estructDos)

	//Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Infinite Loop
	var i int
	for {
		if i < 5 {
			break
		}
		println(i)
		i++
	}

	//Loop collections

	numberos := []int{1, 2, 3}
	for i, v := range numberos {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, port := range wellKnownPorts {
		println(port)
	}

	// Panics

	// panic("We're fokd m8")

	//If else
	if true {
		println("Of course")
	} else if true && !false || true {
		println("lol")
	} else {
		panic("How did this happen")
	}

	// Switch case
	type HTTPRequest struct {
		Method string
	}

	r := HTTPRequest{Method: "GET"}
	switch r.Method {
	case "GET":
		println("GET")
		fallthrough // Next case executed
	case "DELETE":
		println("DELETE")
	default:
		println("Unhandled method")
	}

}
