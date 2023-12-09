package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Howdy, %s!", name)
}

func WorldGreeter(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(WorldGreeter))
}
