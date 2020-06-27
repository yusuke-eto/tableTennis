package interfaces

import (
	"encoding/json"
	"net/http"
)

type hello struct {
	Name string
	Age  int8
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	hello := hello{Name: "こんにちは", Age: 32}
	json.NewEncoder(w).Encode(hello)
}
