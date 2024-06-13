package handlers

import (
	"fmt"
	"net/http"
)

func HomeFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Arun")
}
