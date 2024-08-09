package main

import (
	"flag"
	"fmt"
)

func application() {
	type1 := flag.Bool("type1", false, "Basic Calculator")
	type2 := flag.Bool("type2", false, "Scientific Calculator")

	flag.Parse()

	if *type1 {
		bscCalc()

	} else if *type2 {

		sciCalc()

	} else {
		fmt.Println("Invalid operation. Use -type1 for basic or -type2 for scientific calculator.")
	}

}
