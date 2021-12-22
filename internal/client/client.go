package client

import (
	"encoding/json"
	"log"

	"github.com/captainGeech42/hotline/internal/config"
	"github.com/captainGeech42/hotline/internal/web/schema"
)

// get a callback name to use
// if cbName is non-empty, server will try to use it
// if server does use it, second ret is true
// otherwise, second ret is false
// first ret always has the name of the callback that should be used
func getCallbackName(cbName string, cfg *config.Config) (string, bool) {
	newCbReq := schema.NewCallbackRequest{Name: cbName}

	cbRespBytes, err := makeReq(cfg.Client.ServerUrl, "/api/callback", "POST", newCbReq)
	if err != nil {
		log.Println("failed to setup callback with the server")
		log.Panicln(err)
	}

	newCbResp := schema.NewCallbackResponse{}
	err = json.Unmarshal(cbRespBytes, &newCbResp)
	if err != nil {
		log.Println("failed to parse POST /api/callback response from the server")
		log.Panicln(err)
	}

	log.Println(newCbResp.Message)

	return newCbResp.FullCbDomain, newCbResp.UsedExisting
}

// run the hotline client
// prefCbName: the preferred callback name from the cli args
// showHistorical: true if client wants all callbacks for an existing cb logged
// cfg: the config object
func StartClient(prefCbname string, showHistorical bool, cfg *config.Config) {
	cbDomain, usedExisting := getCallbackName(prefCbname, cfg)

	if usedExisting {
		log.Println("Hotline is now using your previous callback")
	} else {
		log.Println("Hotline is now active using your new callback")
	}

	log.Printf("Start making requests!\n\n\t\t$ curl http://%[1]s/test\n\n\t\t$ dig +short TXT %[1]s\n", cbDomain)
}
