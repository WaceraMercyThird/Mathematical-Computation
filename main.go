package main

import (
	"example.com/greatings/service"
	"fmt"
)

func main() {
	name, _ := service.Hello("Jane Doe")
	results := service.Sum(2, 8)
	fmt.Println(results)
	fmt.Println(name)

}
