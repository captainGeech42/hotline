package backend

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/captainGeech42/hotline/internal/db"
)

type newCallbackRequest struct {
	Name string `json:"name"`
}

type newCallbackResponse struct {
	response
}

// POST /api/v1/callback
// create the callback if it doesn't exist
// otherwise, confirm that it exists and send back to client
func newCallback(w http.ResponseWriter, r *http.Request) {
	// read body bytes
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	// parse json into struct
	var cbReq newCallbackRequest
	err = json.Unmarshal(body, &cbReq)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	var cbResp newCallbackResponse
	cbResp.Error = false

	// check if cb exists
	cb := db.GetCallback(cbReq.Name)
	if cb == nil {
		// cb doesn't exist, make it
		cb = db.CreateCallback(cbReq.Name)
		if cb == nil {
			cbResp.Error = true
			cbResp.Message = "failed to create callback"

		}
		cbResp.Message = "new callback created"
	} else {
		cbResp.Message = "using existing callback"
	}

	// send response
	b, err := json.Marshal(cbResp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}
