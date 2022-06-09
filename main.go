package main

import (
	"fmt"
	"reflect"
)

type Man struct{}

func (m Man) Sleep() {}
func (m Man) Eat()   {}
func (m Man) Work()  {}

// и тип Пёс, который у нас умеет только лаять
type Dog struct{}

func (d Dog) Bark() {}

// теперь мы хотим понять, кто может стать программистом,
// для этого мы определяем интерфейс Programmer,
// и определяем в нём метод Work
// этим мы ограничиваем количество объектов,
// которые смогут быть программистами
// так случилось, что в нашем коде, чтобы стать программистом,
// достаточно уметь работать :)
type Programmer interface {
	Work()
}

type request interface {
}

func main() {
	Vasiliy := Man{}
	fmt.Printf("%s\n", reflect.TypeOf(Vasiliy).String())

	Sharik := Dog{}
	fmt.Printf("%s\n", reflect.TypeOf(Sharik).String())

	var worker Programmer
	worker = Vasiliy

	fmt.Printf("%s\n", reflect.TypeOf(worker).String())
}
