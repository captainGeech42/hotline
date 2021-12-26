package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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

	if newCbResp.Error {
		log.Printf("got an error from the server when registering callback: %s\n", newCbResp.Message)
	}

	return newCbResp.Name, newCbResp.FullCbDomain, newCbResp.UsedExisting
}

// get requests for the callback
func getRequests(cbName string, since *time.Time, cfg *config.Config) (*[]schema.CbHttpRequest, *[]schema.CbDnsRequest) {
	req := schema.GetCbRequestsRequest{Name: cbName}
	if since == nil {
		req.All = true
	} else {
		req.AfterTs = *since
	}

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

func printTime(ts time.Time) {
	t := ts.Local()
	fmt.Printf("%d/%02d/%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// pretty print a DNS request
func printDnsRequest(req schema.CbDnsRequest) {
	printTime(req.Timestamp)

	fmt.Printf("\t\033[1mNew DNS Query\033[0m\n\n\t Source IP:\t%s\n\tQuery Type:\t%s\n\tQuery Name:\t%s\n\n\n", req.SourceIP, req.QueryType, req.QueryName)
}

// https://gist.github.com/kennwhite/306317d81ab4a885a965e25aa835b8ef
func wrapStringOnWord(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n                        " + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped
}

// pretty print a HTTP request
func printHttpRequest(req schema.CbHttpRequest) {
	printTime(req.Timestamp)

	headerBytes, err := base64.StdEncoding.DecodeString(req.Headers)
	if err != nil {
		log.Println("Failed to decode header string")
		log.Panicln(err)
	}

	headerStr := string(headerBytes)
	headerStr = strings.ReplaceAll(headerStr, "\n", "\n                        ")

	bodyBytes, err := base64.StdEncoding.DecodeString(req.Body)
	if err != nil {
		log.Println("Failed to decode body string")
		log.Panicln(err)
	}

	bodyStr := string(bodyBytes)
	bodyStr = wrapStringOnWord(bodyStr, 50)

	fmt.Printf("\t\033[1mNew HTTP Request\033[0m\n\n\t Source IP:\t%s\n\t   Request:\t%s %s\n\t      Host:\t%s\n\n\t   Headers:\t%s\n\t      Body:\t%s\n\n\n", req.SourceIP, req.Method, req.URI, req.Host, headerStr, bodyStr)
}

func retrieveAndDisplayRequests(cbName string, since *time.Time, cfg *config.Config) {
	httpReqs, dnsReqs := getRequests(cbName, since, cfg)

	for _, req := range *dnsReqs {
		printDnsRequest(req)
	}

	for _, req := range *httpReqs {
		printHttpRequest(req)
	}
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
		if prefCbname != "" {
			log.Printf("Your specified callback wasn't found, a new one has been generated instead")
		}
		log.Printf("Hotline is now active using your new callback: %s\n", cbName)
	}

	log.Printf("Start making requests!\n\n\t$ curl http://%[1]s\n\n\t$ dig +short TXT %[1]s\n\n", cbDomain)
	fmt.Println("===========================================================================")
	fmt.Println("")

	if showHistorical {
		retrieveAndDisplayRequests(cbName, nil, cfg)
	}

	for {
		since := time.Now()
		time.Sleep(1 * time.Second)

		httpReqs, dnsReqs := getRequests(cbName, &since, cfg)

		for _, req := range *dnsReqs {
			printDnsRequest(req)
		}

		for _, req := range *httpReqs {
			printHttpRequest(req)
		}
	}
}
