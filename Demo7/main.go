package main

import "fmt"

func main() {
	var array [10]int
	array[0] = 65

	fmt.Printf("Length: %v\n", len(array))

	for i, val := range array {
		fmt.Printf("Index: %v, Value: %v\n", i, val)
	}

	slice := make([]int, 3)
	slice[0] = 20
	slice[1] = 21
	slice[2] = 22
	slice = append(slice, 23)

	fmt.Println(slice)
	fmt.Println(slice[2:4])

	m := make(map[string]string)

	m["indice"] = "valor"

	fmt.Println(m["indice"])
}
