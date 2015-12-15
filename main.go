package main

import (
	"fmt"
	"os"
	// "github.com/hello/readfile"
	"github.com/hello/httpservice"
)

func main() {
	fmt.Println("=== CORE MAIN START ===")
	// readfile.Start()

	// get the argument
	arg := os.Args[1]
	// build a map with functions
	funcs := map[string]func() {"httpservice":httpservice.Start}

	// check if the argument function exist
	if val, ok := funcs[arg]; ok {
		// if the function exist call it
		val()
	}else{
		// give exception if not
		fmt.Println("=== INVALID ARGUMENT ===")
	}

	// httpservice.Start()
	fmt.Println("=== CORE MAIN END ===")
}

// func somethingELse() {
// 	fmt.Println("=== THIS IS SOMETHING ELSE ===")
// }
