package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"message": "hello"}
	s, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(s))
}
