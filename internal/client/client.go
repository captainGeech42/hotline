package client

import (
	"encoding/json"
	"log"
	"time"

	"github.com/captainGeech42/hotline/internal/config"
	"github.com/captainGeech42/hotline/internal/web/schema"
)

// get a callback name to use
// if cbName is non-empty, server will try to use it
// if server does use it, second ret is true
// otherwise, second ret is false
// first ret always has the name of the callback that should be used
func getCallbackName(cbName string, cfg *config.Config) (string, string, bool) {
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

	return newCbResp.Name, newCbResp.FullCbDomain, newCbResp.UsedExisting
}

// get requests for the callback
func getRequests(cbName string, since time.Time, cfg *config.Config) (*[]schema.CbHttpRequest, *[]schema.CbDnsRequest) {
	req := schema.GetCbRequestsRequest{Name: cbName, All: false, AfterTs: since}

	respBytes, err := makeReq(cfg.Client.ServerUrl, "/api/callback/requests", "GET", req)
	if err != nil {
		log.Println("failed to get requests from the server")
		log.Panicln(err)
	}

	resp := schema.GetCbRequestsResponse{}
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		log.Println("failed to parse GET /api/callback/requests response from the server")
		log.Panicln(err)
	}

	return &resp.HttpRequests, &resp.DnsRequests
}

// run the hotline client
// prefCbName: the preferred callback name from the cli args
// showHistorical: true if client wants all callbacks for an existing cb logged
// cfg: the config object
func StartClient(prefCbname string, showHistorical bool, cfg *config.Config) {
	cbName, cbDomain, usedExisting := getCallbackName(prefCbname, cfg)

	if usedExisting {
		log.Printf("Hotline is now using your previous callback: %s\n", cbName)
	} else {
		log.Printf("Hotline is now active using your new callback: %s\n", cbName)
	}

	log.Printf("Start making requests!\n\n\t\t$ curl http://%[1]s/test\n\n\t\t$ dig +short TXT %[1]s\n\n", cbDomain)

	for {
		since := time.Now()
		time.Sleep(1 * time.Second)

		httpReqs, dnsReqs := getRequests(cbName, since, cfg)

		for _, req := range *dnsReqs {
			log.Printf("DNS Request: %s request for %s from %s\n", req.QueryType, req.QueryName, req.SourceIP)
		}

		for _, req := range *httpReqs {
			log.Printf("HTTP Request: %s request for %s from %s\n", req.Method, req.URI, req.SourceIP)
		}
	}
}
