package db

import (
	"log"
)

func GetCallback(cbName string) *Callback {
	if dbHandle == nil {
		log.Panicln("dbHandle is nil!")
	}

	var cb Callback

	result := dbHandle.Where("name = ?", cbName).Limit(1).Find(&cb)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}

	if result.RowsAffected == 1 {
		return &cb
	} else {
		return nil
	}
}

func CreateCallback(cbName string) *Callback {
	if dbHandle == nil {
		log.Panicln("dbHandle is nil!")
	}

	cb := Callback{Name: cbName}

	result := dbHandle.Create(&cb)

	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}

	return &cb
}

func AddDnsRequest(cbName string, request string, queryType string, srcIP string) {
	if dbHandle == nil {
		log.Panicln("dbHandle is nil!")
	}

	// get the associated callback record
	cb := GetCallback(cbName)
	if cb == nil {
		log.Printf("tried to save DNS request for callback that doesn't exist: %s\n", cbName)
		return
	}

	dnsReq := DnsRequest{SourceIP: srcIP, RequestName: request, QueryType: queryType, Callback: *cb}

	result := dbHandle.Create(&dnsReq)

	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Printf("added DNS request from %s for %s to database\n", srcIP, cbName)
	}
}

// these should be base64 encoded before coming into this function
func AddHttpRequest(cbName string, uri string, host string, method string, headers string, body string, srcIP string) {
	if dbHandle == nil {
		log.Panicln("dbHandle is nil!")
	}

	// get the associated callback record
	cb := GetCallback(cbName)
	if cb == nil {
		log.Printf("tried to save HTTP request for callback that doesn't exist: %s\n", cbName)
		return
	}

	httpReq := HttpRequest{SourceIP: srcIP, URI: uri, Host: host, Method: method, Headers: headers, Body: body, Callback: *cb}

	result := dbHandle.Create(&httpReq)

	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Printf("added HTTP request from %s for %s to database\n", srcIP, cbName)
	}
}
