package main

import (
	"fmt"
	"log"

	"github.com/robhor/ernis/lib"
)

func main() {
	meal, err := ernis.GetTodaysMeals()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(meal)
}
