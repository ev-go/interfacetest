package main

import "fmt"

type WayToMove interface {
	move()
}

// структура "Автомобиль"
type Car struct{}

// структура "Самолет"
type Aircraft struct{}

type Boat struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}
func (a Boat) move() {
	fmt.Println("Лодка идет")
}

func main() {

	var tesla WayToMove = Car{}
	var boing WayToMove = Aircraft{}
	var admiral WayToMove = Boat{}
	tesla.move()
	boing.move()
	admiral.move()
}
