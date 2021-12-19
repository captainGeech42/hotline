package dns

// based on https://jameshfisher.com/2017/08/04/golang-dns-server/

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/captainGeech42/hotline/internal/config"
	"github.com/captainGeech42/hotline/internal/db"
	"github.com/miekg/dns"
)

var callbackDomain string
var defaultAResponse net.IP
var defaultTXTResponse string

type dnsHandler struct{}

func (handler *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	reqDomain := r.Question[0].Name
	qtype := qtypeMapping[r.Question[0].Qtype]

	srcIP := strings.Split(w.RemoteAddr().String(), ":")[0]
	log.Printf("got a DNS %s request from %v for %s\n", qtype, srcIP, reqDomain)

	// parse out the callback name
	// identify how many labels are in the top-level callback domain
	numLabels := strings.Count(callbackDomain, ".")
	// split the dns request. the last entry is "" because there is a trailing .
	reqParts := strings.Split(reqDomain, ".")
	// get the callback name
	numParts := len(reqParts)
	idx := numParts - numLabels - 3
	if idx >= 0 {
		// get the callback name
		cbName := reqParts[idx]

		// log the request to the db
		db.AddDnsRequest(cbName, reqDomain, qtype, srcIP)
	} else {
		log.Println("couldn't parse the callback name out of the DNS request")
	}

	// generate the response answer
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   defaultAResponse,
		})
	case dns.TypeTXT:
		msg.Answer = append(msg.Answer, &dns.TXT{
			Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 60},
			Txt: []string{defaultTXTResponse},
		})
	default:
		log.Printf("got an unsupported query type request, not sending a response: %v\n", qtype)
	}

	w.WriteMsg(&msg)
}

func StartServer(cfg *config.Config) {
	// connect to database
	if !db.ConnectToDb(cfg.Server.Database) {
		return
	}

	// setup dns server struct
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Server.Callback.Dns.Port)
	srv := &dns.Server{Addr: addr, Net: "udp"}
	srv.Handler = &dnsHandler{}

	// define global vars for responses
	callbackDomain = cfg.Server.Callback.Domain
	defaultAResponse = net.ParseIP(cfg.Server.Callback.Dns.DefaultAResponse)
	defaultTXTResponse = cfg.Server.Callback.Dns.DefaultTXTResponse

	// start the server
	log.Printf("starting dns callback listener on %s\n", addr)
	log.Fatalln(srv.ListenAndServe())
}
