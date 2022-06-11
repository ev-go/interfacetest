package main

import "fmt"

type animaler interface {
	makeSound()
	run()
}

type cat struct {
	color string
	class string
}

type dog struct{}

func (c *cat) makeSound() {
	switch c.color {
	case "2":
		fmt.Println("meow")
	case "1":
		fmt.Println("moew")

	}
	fmt.Println("hi")
}

func (c *cat) run() {
	fmt.Println("Paw")
}

func (d *dog) makeSound() {
	fmt.Println("moo")
}

func (d *dog) run() {
	fmt.Println("PawPaw")
}

func main() {
	var c, d animaler = &cat{"2", "3"}, &dog{}
	c.makeSound()
	d.makeSound()
	c.run()
}

// package main

// import "fmt"

// type WayToMove interface {
// 	move()
// }

// // структура "Автомобиль"
// type Car struct{}

// // структура "Самолет"
// type Aircraft struct{}

// type Boat struct{}

// func (c Car) move() {
// 	fmt.Println("Автомобиль едет")
// }
// func (a Aircraft) move() {
// 	fmt.Println("Самолет летит")
// }
// func (a Boat) move() {
// 	fmt.Println("Лодка идет")
// }

// func main() {

// 	var tesla WayToMove = Car{}
// 	var boing WayToMove = Aircraft{}
// 	var admiral WayToMove = Boat{}
// 	tesla.move()
// 	boing.move()
// 	admiral.move()
// }
