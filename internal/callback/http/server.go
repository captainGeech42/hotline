package http

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/captainGeech42/hotline/internal/config"
	"github.com/captainGeech42/hotline/internal/db"
	"github.com/gorilla/mux"
)

var httpResponse []byte
var callbackDomain string

// handler used to send message from the config back to the requesting server
func reqHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(httpResponse)
}

// middleware used to log any request that comes in
func writeReqToDb(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse out the source IP address, we don't care about the port
		srcIP := strings.Split(r.RemoteAddr, ":")[0]

		// parse out the callback name

		// identify how many labels are in the top-level callback domain
		numLabels := strings.Count(callbackDomain, ".") + 1
		// split the http host used in the request
		httpHost := strings.Split(r.Host, ":")[0] // used later too
		httpHostParts := strings.Split(httpHost, ".")
		// get the callback name
		numParts := len(httpHostParts)
		idx := numParts - numLabels - 1
		if idx < 0 {
			// not enough parts in the domain, probably wasn't using a subdomain
			// bail out
			log.Printf("couldn't parse callback domain from %s\n", httpHost)
			next.ServeHTTP(w, r)
			return
		}

		cbName := httpHostParts[idx]

		// build the base64 header block
		// headers are a map of strings -> slice of strings
		// note that the order of the headers may not align with the order
		// in which they were sent by the client, because maps are unordered
		headerString := ""
		for k, vals := range r.Header {
			// check if we got a X-Hotline-Real-IP header
			// this is set by the hotline nginx config so we can log the real source IP
			// this header shouldn't be exposed to the user
			if k == "X-Hotline-Real-Ip" {
				srcIP = vals[0]
				continue
			}

			for _, v := range vals {
				headerString += fmt.Sprintf("%s: %s\n", k, v)
			}
		}
		encodedHeaders := base64.StdEncoding.EncodeToString([]byte(headerString))

		// encode the body
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		encodedBody := base64.StdEncoding.EncodeToString(bodyBytes)

		db.AddHttpRequest(cbName, r.RequestURI, httpHost, r.Method, encodedHeaders, encodedBody, srcIP)

		next.ServeHTTP(w, r)
	})
}

func StartServer(cfg *config.Config) {
	// connect to database
	if !db.ConnectToDb(cfg.Server.Database) {
		return
	}

	// build router
	router := mux.NewRouter()
	router.Use(writeReqToDb)
	router.PathPrefix("/").HandlerFunc(reqHandler)

	// set package-level globals for the middleware and handler
	httpResponse = []byte(fmt.Sprintf(`{"message": "%s"}`, cfg.Server.Callback.Http.DefaultReponse))
	callbackDomain = cfg.Server.Callback.Domain

	// start the http listener
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Server.Callback.Http.Port)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
