package interfaces

import (
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/hello_world", helloWorldHandler)
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
