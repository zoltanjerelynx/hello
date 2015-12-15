package httpservice

import (
	"fmt"
	"encoding/json"
	"net/http"
	"html"
)

func init() {
	fmt.Println("=== INIT OF HTTPSERVICE PACKAGE ===")
}

// func Start(){
// 	http.Handle("/static", http.FileServer(http.Dir("/home/jerezoltan/static")))
// 	http.ListenAndServe(":3000", nil)
// }

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRestFood(writer http.ResponseWriter, request *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(writer, string(response))
}

func serveRestSomethingElse(writer http.ResponseWriter, request *http.Request) {
	response, err := getJsonResponseSomethingElse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(writer, string(response))
}
func invalidUrlRest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, this request does not exist: %q", html.EscapeString(request.URL.Path))
}

func Start(){
	http.HandleFunc("/", invalidUrlRest)
	http.HandleFunc("/getFruitsAndVegetables", serveRestFood)
	http.HandleFunc("/getSomethingElse", serveRestSomethingElse)
	http.ListenAndServe("localhost:1337", nil)
}

func getJsonResponse() ([]byte, error){
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits,vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}

func getJsonResponseSomethingElse() ([]byte, error){
	fruits := make(map[string]int)
	fruits["Apples"] = 99
	fruits["Oranges"] = 99

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 99
	vegetables["Peppers"] = 99

	d := Data{fruits,vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
