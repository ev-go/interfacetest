package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type animaler interface {
	makeSound()
	run()
}

type cat struct {
	color string
	class string
	speed string
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
	fmt.Println("Paw", c.speed)

}

func (d *dog) makeSound() {
	fmt.Println("moo")
}

func (d *dog) run() {
	fmt.Println("PawPaw")
}

type Gettokenanswerstruct struct {
	TokenRequestAt string
	User           string
	Login          string
	Password       string
	DataAnswer     string
	Token          string
}

type httpRequestMessageStruct struct {
	requestUseLogin    string
	requestUsePassword string
	requestUseData     string
}

type httpRequestStruct struct {
	requestUseUrl      string
	requestUsePort     string
	requestUseRout     string
	httpRequestMessage httpRequestMessageStruct
}

type httpRequester interface {
	getToken()
	getPrivateRout()
}

func (req *httpRequestStruct) getPrivateRout() {
	httpRequestString := "http://localhost:3000/products"

	bearer := "Bearer " //+ tokenFromRedis
	cli := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest("GET", httpRequestString, nil)
	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", `application/json`)
	if err != nil {
		panic(err)
	}

	response, err := cli.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nresponse?:", string(responseData))
	//return responseData
}

func (req *httpRequestStruct) getToken() {
	httpRequestString := req.httpReqStructToString()
	cli := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest("GET", httpRequestString, nil)
	//request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", `application/json`)
	if err != nil {
		panic(err)
	}
	//defer request.Body.Close()

	response, err := cli.Do(request)
	if err != nil {
		panic(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nresponse?:", string(responseData))

	var Gettokenanswer = &Gettokenanswerstruct{}
	json.Unmarshal([]byte(responseData), Gettokenanswer)
	fmt.Println("\ntoken from struct:", Gettokenanswer.Token)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	newTokenToRedis := rdb.Set(ctx, req.httpRequestMessage.requestUseLogin, Gettokenanswer.Token, 0).Err()
	if newTokenToRedis != nil {
		panic(newTokenToRedis)
	}
}

func (req *httpRequestStruct) httpReqStructToString() string {
	String := "http://" +
		req.requestUseUrl +
		":" +
		req.requestUsePort +
		"/" +
		req.requestUseRout +
		"?login=" +
		req.httpRequestMessage.requestUseLogin +
		"&password=" +
		req.httpRequestMessage.requestUsePassword +
		"&data=" +
		req.httpRequestMessage.requestUseData

	return String
}

func main() {
	var c, d animaler = &cat{"2", "3", "4"}, &dog{}
	c.makeSound()
	d.makeSound()
	c.run()

	requestUseUrl := "localhost"
	requestUsePort := "3000"
	requestUseRout := "get-token"
	requestUseLogin := "root2"
	requestUsePassword := "1"
	requestUseData := "21"

	var httpReq httpRequester = &httpRequestStruct{
		requestUseUrl,
		requestUsePort,
		requestUseRout,
		httpRequestMessageStruct{
			requestUseLogin,
			requestUsePassword,
			requestUseData}}
	httpReq.getToken()
	httpReq.getPrivateRout()
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
