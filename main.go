package main

import (
	"fmt"

	"github.com/robhor/ernis/lib"
)

func main() {
	meal, err := ernis.GetTodaysMeals()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(meal)
	}
}
