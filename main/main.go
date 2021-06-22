package main

import (
	"fmt"
	_ "log"
)

type House struct {
	address  string
	city     string
	numRooms int
	price    float64
}

func main() {
	rooms := 5
	address := "123 Blah St"
	city := "Blahville"
	price := 123456.0
	fmt.Printf("%s\t%s\t%d Rooms \t$%.0f\n", address, city, rooms, price)
}
