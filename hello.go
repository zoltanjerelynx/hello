package main

import (
    "strconv"
    "fmt"
)

func main() {

    // var x int = 404
    // var y uint32 = 234242
    // var xx float64 = 34.23
    // var yy complex64 = 3 + 2i
    // var weird string = "weird string"

    // for i := 0; i < 10; i++ {
	//        fmt.Println("The value of i : ", i)
    // }

    my_array := [ ]int{4, 3, 1, 7, 3, 9, 2}

    converted, extra := convert_int_to_string_plus_extra(my_array)

    fmt.Println("The value of sum : " + converted, extra)
}


func convert_int_to_string_plus_extra(a []int) (convert_to_string string, extra int) {
    var sum int
    for _, i := range a {
        sum = sum + i
    }
    convert_to_string = strconv.Itoa(sum)
    extra =  10

    return convert_to_string, extra
}
