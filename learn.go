package main

import (
	"encoding/json"
	"fmt"
)

//func addition(num1, num2 int) int {
//	return num1 + num2
//}
//
//func Division(numerator, denominator int) (float64, int, error) {
//
//	var err error
//	if denominator == 0 {
//		err = errors.New("can't divide by 0")
//		return 0, 0, err
//	}
//	result := float64(numerator / denominator)
//	var deno = numerator % denominator
//	return result, deno, nil
//}

//var division, reminder, err = Division(4, 0)
//if err != nil {
//fmt.Print("Something is wrong, ", err.Error())
//} else if reminder == 0 {
//fmt.Printf("%f, %d", division, reminder)
//} else {
//fmt.Printf("%f, %d", division, reminder)
//}

func array() {
	var arr [4]int
	arr[0] = 1
	arr[2] = 5
	arr[1] = 11
	arr[3] = 10
	fmt.Println(arr[0:2])
	arr2 := [2]int{1, 2}
	arr3 := [...]int{1, 2, 7} // go will infer size at run time
	fmt.Println(arr2)
	fmt.Println(arr3)
	var sum int = 0
	for _, val := range arr {
		sum += val
	}
	fmt.Println(sum)
}

func slice() {
	var intSlice []int
	fmt.Printf("Capacity %v, Size: %v", cap(intSlice), len(intSlice))
	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 82)
	intSlice = append(intSlice, 22)
	var slice2 []int = []int{2, 3, 4}
	fmt.Printf("\nCapacity %v, Size: %v\n", cap(intSlice), len(intSlice))
	appendSlice := append(intSlice, slice2...)
	var newSlice []int32 = make([]int32, 2, 4)
	fmt.Println(intSlice)
	fmt.Println(appendSlice)
	fmt.Println(newSlice)
}

func maps() {
	var myMap map[string]int = make(map[string]int)
	myMap["Chandrakant"] = 24
	myMap["Go"] = 124

	for key, val := range myMap {
		fmt.Println(key, val)
	}

	//	init oneline
	var inline map[string]int = map[string]int{"Name": 34, "afda": 4}

	value, status := inline["Name"]

	fmt.Println(len(inline))
	delete(inline, "afda")
	fmt.Println(len(inline))

	fmt.Println(value, status)
}

func String(name string) {
	var myString = "resume"

	println(myString)
}

type PertolEngine struct {
	millage  int
	capacity int
}

type DieselEngine struct {
	millage  int
	capacity int
}

func (pe PertolEngine) calculateMillage() int {
	return pe.millage * pe.capacity
}

func (pe DieselEngine) calculateMillage() int {
	return pe.millage * pe.capacity
}

type Engine interface {
	calculateMillage() int
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` // remove the field while removing
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() ([]byte, error) {
	courses := []course{
		{"C++", 233, "Udemy", "Ok", []string{"web dev"}},
		{"Node js", 299, "Udemy", "node ", []string{"web dev"}},
		{"Node js", 299, "Udemy", "node ", []string{}},
	}
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	} else {
		//fmt.Printf("%s", finalJson)
	}
	return finalJson, nil
}

func DecodeJson() {
	//jsonDataFromWeb, err := EncodeJson()
	//if err != nil {
	//	panic(err)
	//}

	jsonDataFromWeb := []byte(`{
		"name": {
		"courseName": "C++",
		"price": 233,
		"platform": "Udemy",
		"tags": [
			"web dev"
		]
	}`)

	//var mycourse []course
	isValid := json.Valid(jsonDataFromWeb)
	if !isValid {
		fmt.Printf("Json is not valid")
		return
	}

	//err = json.Unmarshal(jsonDataFromWeb, &mycourse)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%#v\n", mycourse)
	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

}

func learn() {

	DecodeJson()

	//var c = make(chan int, 5)
	//go process(c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	//var petrolEngine = PertolEngine{20, 24}
	//var dieselEngin = DieselEngine{20, 24}
	//
	//fmt.Println(petrolEngine.calculateMillage())
	//fmt.Println(dieselEngin.calculateMillage())

	//maps()
	//
	//app := fiber.New()
	//
	//app.Use(func(c *fiber.Ctx) error {
	//	fmt.Println("Inside the middleware")
	//	return c.Next()
	//})
	//
	//app.Get("/*", func(c *fiber.Ctx) error {
	//	pathVal := c.Params("*")
	//	return c.SendString("Hello String\n" + pathVal)
	//})
	//
	//app.Listen(":3000")
}
