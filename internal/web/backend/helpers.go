package backend

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func sendResponse(w http.ResponseWriter, resp interface{}) {
	// send response
	b, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

// based on https://stackoverflow.com/a/31832326
func generateCallbackName() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
