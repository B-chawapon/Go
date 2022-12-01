package main

import (
	"demo1/demo2"
	"fmt"
)

func getALL() int {
	return 500
}
func getTwoValue() (int, string) {
	return 100, "HELLOW"
}

func main() {
	fmt.Println("Hellow2243")
	fmt.Println("test")
	name := "DOGE"
	var dog string = "doggy"
	var num float32 = 1.432

	fmt.Println(name)
	fmt.Println(dog)
	fmt.Println(num)
	fmt.Printf("my dog is  %v\n", dog)
	fmt.Printf("my dog height  %v\n", num)
	fmt.Printf("my dog height TYPE %T\n", num)

	number1, number2 := 10, 20
	fmt.Println(number1, number2)

	test := []int{10, 20, 30}
	size := len(test)
	fmt.Println(size)

	coin := map[string]string{}
	coin["ETH"] = "ETher"
	coin["BTC"] = "Bitcoin"

	coin2 := map[string]string{"ETH": "Ether", "BTC": "Bitcoin"}
	fmt.Println(coin2["ETH"])

	for index, value := range test {
		fmt.Print(index)
		fmt.Println(value)

	}
	fmt.Println(getALL())

	e1, e2 := getTwoValue()
	fmt.Println(e1, e2)

	fmt.Println(demo2.GetALLDemo1_test())

	// var score int
	// fmt.Scanf("%s", &score)

	// if score == 0 {
	// 	fmt.Println("MY score IS ", score)
	// } else {
	// 	fmt.Println("MY score == ", score)
	// }

	// var myName string
	// fmt.Scanf("%s", &myName)
	// fmt.Println("MY NAME IS ", myName)

}
