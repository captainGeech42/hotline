package db

import (
	"log"
)

func GetCallback(cbName string) *Callback {
	if dbHandle == nil {
		log.Panicf("dbHandle is nil!")
	}

	var cb Callback

	result := dbHandle.Where("name = ?", cbName).First(&cb)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}

	return &cb
}

func CreateCallback(cbName string) *Callback {
	if dbHandle == nil {
		log.Panicf("dbHandle is nil!")
	}

	cb := Callback{Name: cbName}

	result := dbHandle.Create(&cb)

	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}

	return &cb
}

func AddDnsRequest(cbName string, request string, srcIP string) {
	if dbHandle == nil {
		log.Panicf("dbHandle is nil!")
	}

	// get the associated callback record
	cb := GetCallback(cbName)
	if cb == nil {
		return
	}

	dnsReq := DnsRequest{SourceIP: srcIP, RequestName: request, Callback: *cb}

	result := dbHandle.Create(&dnsReq)

	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Printf("added DNS request from %s for %s to database\n", srcIP, cbName)
	}
}

// these should be base64 encoded before coming into this function
func AddHttpRequest(cbName string, url string, headers string, body string, srcIP string) {
	if dbHandle == nil {
		log.Panicf("dbHandle is nil!")
	}

	// get the associated callback record
	cb := GetCallback(cbName)
	if cb == nil {
		return
	}

	httpReq := HttpRequest{SourceIP: srcIP, Url: url, Headers: headers, Body: body, Callback: *cb}

	result := dbHandle.Create(&httpReq)

	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Printf("added HTTP request from %s for %s to database\n", srcIP, cbName)
	}
}
