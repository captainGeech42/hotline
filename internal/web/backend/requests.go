package backend

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/captainGeech42/hotline/internal/db"
	"github.com/captainGeech42/hotline/internal/web/schema"
)

// GET /api/callback/requests
func getCbRequests(w http.ResponseWriter, r *http.Request) {
	// read body bytes
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	reqResp := schema.GetCbRequestsResponse{}

	// make sure there was a body
	if len(body) == 0 {
		reqResp.Error = true
		reqResp.Message = "no JSON body provided"
		w.WriteHeader(400)
		sendResponse(w, reqResp)
		return
	}

	// unmarshal the body
	cbReqBody := schema.GetCbRequestsRequest{}
	if err := json.Unmarshal(body, &cbReqBody); err != nil {
		log.Println(err)

		reqResp.Error = true
		reqResp.Message = "failed to unmarshal request body"
		w.WriteHeader(500)
		sendResponse(w, reqResp)
		return
	}

	var afterTs *time.Time
	if cbReqBody.All {
		afterTs = nil
	} else {
		afterTs = &cbReqBody.AfterTs
	}

	// get the dns requests from the db
	dnsDbReqs := db.GetDnsRequests(cbReqBody.Name, afterTs)
	dnsReqs := ConvertDnsDbToJson(&dnsDbReqs)

	// get the http requests from the db
	httpDbReqs := db.GetHttpRequests(cbReqBody.Name, afterTs)
	httpReqs := ConvertHttpDbToJson(&httpDbReqs)

	// build resp object
	respBody := schema.GetCbRequestsResponse{}
	respBody.Error = false
	respBody.DnsRequests = *dnsReqs
	respBody.HttpRequests = *httpReqs

	sendResponse(w, respBody)
}

func ConvertDnsDbToJson(dnsReqs *[]db.DnsRequest) *[]schema.CbDnsRequest {
	outArray := []schema.CbDnsRequest{}

	for _, dbReq := range *dnsReqs {
		outArray = append(outArray, schema.CbDnsRequest{
			// https://github.com/golang/go/issues/9859
			// if ^^ ever becomes a lang feature, this can be simplified
			CbRequest: schema.CbRequest{
				Timestamp: dbReq.CreatedAt,
				SourceIP:  dbReq.SourceIP,
			},

			QueryName: dbReq.QueryName,
			QueryType: dbReq.QueryType,
		})
	}

	return &outArray
}

func ConvertHttpDbToJson(httpReqs *[]db.HttpRequest) *[]schema.CbHttpRequest {
	outArray := []schema.CbHttpRequest{}

	for _, dbReq := range *httpReqs {
		outArray = append(outArray, schema.CbHttpRequest{
			// https://github.com/golang/go/issues/9859
			// if ^^ ever becomes a lang feature, this can be simplified
			CbRequest: schema.CbRequest{
				Timestamp: dbReq.CreatedAt,
				SourceIP:  dbReq.SourceIP,
			},

			URI:     dbReq.URI,
			Host:    dbReq.Host,
			Method:  dbReq.Method,
			Headers: dbReq.Headers,
			Body:    dbReq.Body,
		})
	}

	return &outArray
}
