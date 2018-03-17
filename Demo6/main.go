package main

import "fmt"

type Structure struct {
	Name    string
	Counter int64
}

func main() {
	i1 := Structure{}

	i1.Counter = 22
	i1.Name = "i1"

	fmt.Println("Name: " + i1.Name)

	incCounter(i1)

	fmt.Printf("Counter: %v\n", i1.Counter)

	i2 := &Structure{}

	i2.Counter = 22
	i2.Name = "i2"

	fmt.Println("Name: " + i2.Name)

	incCounterPointer(i2)

	fmt.Printf("Counter: %v\n", i2.Counter)

	i3 := &Structure{}

	i3.Counter = 22
	i3.Name = "i3"

	fmt.Println("Name: " + i3.Name)

	i3.method()

	fmt.Printf("Counter: %v\n", i3.Counter)

}

func incCounter(s Structure) {
	s.Counter += 100
}

func incCounterPointer(s *Structure) {
	s.Counter += 200
}

func (s *Structure) method() {
	s.Counter += 1000
}
