package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type House struct {
	address  string
	city     string
	numRooms int
	price    float64
}

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var houses []House

	for {
		// add is used to decide when to terminate house-appending loop
		if !add() {
			fmt.Println() // For terminal spacing
			break
		}

		fmt.Println() // For terminal spacing

		var address, city string
		var numRooms int
		var price float64

		// Stack Overflow solution to getting input with spaces:
		fmt.Print("Enter address -> ")
		if scanner.Scan() {
			address = scanner.Text()
		}

		fmt.Print("Enter city -> ")
		if scanner.Scan() {
			city = scanner.Text()
		}

		// fmt.Print("Enter address -> ")
		// fmt.Fscanf(reader, "%q", &address)

		// fmt.Print("Enter city -> ")
		// fmt.Fscanf(reader, "%q", &city)

		fmt.Print("Enter number of rooms -> ")
		if _, err := fmt.Fscanln(reader, &numRooms); err != nil {
			log.Fatal(err)
		}

		fmt.Print("Enter price -> ")
		if _, err := fmt.Fscanln(reader, &price); err != nil {
			log.Fatal(err)
		}

		houses = append(houses, House{address, city, numRooms, price})

		fmt.Println() // For terminal spacing
	}

	for _, house := range houses {
		fmt.Printf("%s\t\t%s\t\t%d Rooms\t\t$%.2f\n", house.address, house.city, house.numRooms, house.price)
	}
}

// add is used to decide when to terminate house-appending loop
func add() bool {
	var add string

	fmt.Print("Add a new house? (y/n) -> ")

	if _, err := fmt.Fscanln(reader, &add); err != nil {
		log.Fatal(err)
	}

	if add == "n" {
		return false
	}

	return true
}
