package schema

import "time"

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// POST /api/callback request schema
type NewCallbackRequest struct {
	Name string `json:"name"`
}

// POST /api/callback response schema
type NewCallbackResponse struct {
	Response
	Name         string `json:"name"`
	UsedExisting bool   `json:"used_existing"`
	FullCbDomain string `json:"full_domain"`
}

// GET /api/callback/requests request schema
type GetCbRequestsRequest struct {
	// name of the callback to return results for
	Name string `json:"name"`

	// should all historical requests be returned?
	// if this is true, AfterTs is ignored
	All bool `json:"all_requests"`

	// timestamp to filter for new requests
	// if set, only cb requests that came in after this
	// will be returned to the client
	AfterTs time.Time `json:"after_ts"`
}

// GET /api/callback/requests response schema
type GetCbRequestsResponse struct {
	Response

	DnsRequests  []CbDnsRequest  `json:"dns_reqs"`
	HttpRequests []CbHttpRequest `json:"http_reqs"`
}

type CbRequest struct {
	// CreatedAt value from the GORM model, when it was
	// written to the db
	Timestamp time.Time `json:"ts"`

	// IP that generated the callback
	SourceIP string `json:"src_ip"`
}
type CbDnsRequest struct {
	CbRequest

	// DNS query name
	QueryName string `json:"qname"`

	// DNS query type, already parsed to string b/f written to DB
	QueryType string `json:"qtype"`
}

type CbHttpRequest struct {
	CbRequest

	// URL path/URI used for the request
	URI string `json:"uri"`

	// Host header value
	Host string `json:"host"`

	// HTTP method used for the request
	Method string `json:"method"`

	// base64 encoded of \n delimited HTTP headers
	Headers string `json:"headers"`

	// base64 encoded HTTP request body, unparsed
	Body string `json:"body"`
}
