package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

// make an HTTP request
// body gets marshalled to json
// expectedRet can be zero or one values, if not provided, 200 is assumed
func makeReq(urlFromCfg string, uri string, method string, body interface{}, expectedRet ...int) ([]byte, error) {
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

	// build the request
	req, err := http.NewRequest(method, url, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Hotline-client")

	// actually make the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check the error code
	var retShouldBe int
	if len(expectedRet) > 0 {
		retShouldBe = expectedRet[0]
	} else {
		retShouldBe = 200
	}

	if resp.StatusCode != retShouldBe {
		log.Printf("got an unexpected status code from the server, %v instead of %v\n", resp.Status, retShouldBe)
	}

	// parse the body bytes
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
