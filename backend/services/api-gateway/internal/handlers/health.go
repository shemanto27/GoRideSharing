package handlers

import (
	"fmt"
	"net/http"
)

func Healthhandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "API Gateway is running")
}