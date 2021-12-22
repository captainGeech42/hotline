package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/captainGeech42/hotline/internal/db"
	"github.com/captainGeech42/hotline/internal/web/schema"
)

// POST /api/callback
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

	var cbResp schema.NewCallbackResponse
	cbResp.Error = false
	cbResp.UsedExisting = false

	if len(body) > 0 {
		// there was a json body, parse out the name
		var cbReq schema.NewCallbackRequest
		err = json.Unmarshal(body, &cbReq)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}

		// check if that name exists
		if cb := db.GetCallback(cbReq.Name); cb != nil {
			// name exists, send it back to the client
			cbResp.Message = "using existing callback"
			cbResp.Name = cbReq.Name
			cbResp.UsedExisting = true
			sendResponse(w, cbResp)

			return
		}
	}

	// cb doesn't exist, generate a new one
	// intentionally not using the requested name,
	// using a high entropy name instead
	cbName := generateCallbackName()
	cb := db.CreateCallback(cbName)

	if cb == nil {
		cbResp.Error = true
		cbResp.Message = "failed to create callback"
	} else {
		cbResp.Message = "new callback created"
		cbResp.Name = cb.Name
	}

	cbResp.FullCbDomain = fmt.Sprintf("%s.%s", cbResp.Name, callbackDomain)

	// send response
	sendResponse(w, cbResp)
}
