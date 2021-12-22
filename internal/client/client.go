package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/captainGeech42/hotline/internal/config"
	"github.com/captainGeech42/hotline/internal/web/schema"
)

// make an HTTP request
// body gets marshalled to json
func makeReq(urlFromCfg string, uri string, method string, body interface{}) ([]byte, error) {
	// build the url
	// https://stackoverflow.com/a/34668130
	u, err := url.Parse(urlFromCfg)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, uri)
	url := u.String()

	// build the request body
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Hotline-client")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func StartClient(cbName string, showHistorical bool, cfg *config.Config) {
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
}
