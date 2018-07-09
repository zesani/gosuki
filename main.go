package main

import (
	"fmt"
)

func init() {
	fmt.Println("init")
}
func main() {
	// name := ""
	// fmt.Scanf("%s", &name)
	// fmt.Printf("hello %s\n", name)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// var i int64
	// var name string
	// var arr []string{
	// 	"a",
	// 	"b",
	// }

	// slice make([]string)
	slice := []string{}
	// fmt.Println(slice)
	for i := range slice {
		fmt.Println(i)
	}
	mapVal := map[string]string{
		"key":  "value",
		"key2": "value2",
	}
	mapVal2 := map[string]interface{}{
		"key":  "value",
		"key2": 2,
		"nestedMap": map[string]interface{}{
			"childKey": "childVal",
		},
	}

	fmt.Println(mapVal)
	fmt.Println(mapVal2)
	Hello("kiti", "sukma")
	Display("kitisak", "sukma")
	a, b := ReturnVal("a")
	fmt.Println(a, b)
	average, sum := calculate(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(average, sum)
}

func Hello(name, surname string) {
	fmt.Println("Hello", name, surname)
}

func Display(str ...string) {
	for _, s := range str {
		fmt.Println(s)
	}
}

func ReturnVal(a string) (string, string) {
	return a, "b"
}

func calculate(numbers ...int) (float32, int) {
	sum := 0
	countP := 0
	for _, number := range numbers {
		count := 0
		for i := 2; i <= number; i++ {
			if number%i == 0 {
				count++
			}
		}
		if count == 1 {
			sum += number
			countP++
		}
	}
	return float32(sum) / float32(countP), sum
}
