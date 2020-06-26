package interfaces

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hoge", r.FormValue("hoge"))
	fmt.Println("foo", r.FormValue("foo"))
}
