/*
    install to bin
    $ go install github.com/hello
    $ $GOPATH/bin/hello

    build to current location
    go build github.com/hello/
    ./hello

    simple run
    go run hello.go
*/

package main

import (
    // Package strconv implements conversions to and from string representations of basic data types.
    "strconv"
    // Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
    "fmt"
)

// func main() {
//     my_array := [ ]int{4, 3, 1, 7, 3, 9, 2}
//
//     converted, extra := convert_int_to_string_plus_extra(my_array)
//
//     fmt.Println("The value of sum : " + converted, extra)
// }

func main() {
    scores := []int{1, 2, 3, 4, 5}
    scores = removeAtIndex(scores, 2)
    fmt.Println(scores)
}
func removeAtIndex(source []int, index int) []int {
    lastIndex := len(source) - 1
    fmt.Println("lastindex: ", lastIndex)
    //swap the last value and the value we want to remove
    source[index], source[lastIndex] = source[lastIndex], source[index]
    return source[:lastIndex]
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

// func declarations(){
//     var x int = 404
//     var y uint32 = 234242
//     var xx float64 = 34.23
//     var yy complex64 = 3 + 2i
//     var weird string = "weird string"
//
//     for i := 0; i < 10; i++ {
//         fmt.Println("The value of i : ", i)
//     }
//
//     my_array := [ ]int{4, 3, 1, 7, 3, 9, 2}
//
// }
