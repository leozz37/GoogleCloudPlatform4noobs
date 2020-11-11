package helloheart

import (
	"fmt"
	"net/http"
)

// HelloGet is an HTTP Cloud Function.
func HelloHeart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
