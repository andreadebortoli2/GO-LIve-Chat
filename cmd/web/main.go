package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := Router()

	fmt.Println("serving on port 8080")
	_ = http.ListenAndServe(":8080", router)

}
