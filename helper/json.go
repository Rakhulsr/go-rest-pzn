package helper

import (
	"encoding/json"
	"net/http"
)

func ParseJson(r *http.Request, payload any) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(payload)
	PanicError(err)
}

func WriteJson(w http.ResponseWriter, payload any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(payload)
	PanicError(err)
}
