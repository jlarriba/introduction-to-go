package main

import (
	"fmt"
)

func main() {
	str, err := power()
	if err != nil {
		fmt.Errorf("error")
	}
	fmt.Println(str)

}

func power() (string, error) {
	return "hola, mundo", nil
}
