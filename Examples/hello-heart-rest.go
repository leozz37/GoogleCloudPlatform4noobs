package helloworld

import (
	"fmt"
	"net/http"
)

// HelloHeart is an HTTP Cloud Function.
func HelloHeart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, He4rt!")
}
