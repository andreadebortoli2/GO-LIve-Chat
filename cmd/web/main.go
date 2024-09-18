package main

import (
	"fmt"
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HomePage)

	fmt.Println("serving on port 8080")
	_ = http.ListenAndServe(":8080", nil)

}
