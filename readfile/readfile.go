package readfile

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// func init() {
// 	fmt.Println("=== INIT OF READFILE PACKAGE ===")
// }

func read(filename string) {
	fmt.Println("=== READ THE FILE ===");
	dat, err := ioutil.ReadFile(filename)
	fmt.Println("=== CHECK ERRORS ===");
	check(err)
	fmt.Println("=== FILE READED ===");
    fmt.Print("Readed file conten: ",string(dat))
	fmt.Println("=== FILE READ OVER ===");
}

func Start() {
	fmt.Println("=== MYPACKAGE/READFILE START ===")
	read("/tmp/readfile")
	fmt.Println("=== MYPACKAGE/READFILE END ===")
}
