package db

import (
	"log"
	"time"
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

func AddDnsRequest(cbName string, queryName string, queryType string, srcIP string) {
	if dbHandle == nil {
		log.Panicln("dbHandle is nil!")
	}

	// get the associated callback record
	cb := GetCallback(cbName)
	if cb == nil {
		log.Printf("tried to save DNS request for callback that doesn't exist: %s\n", cbName)
		return
	}

	dnsReq := DnsRequest{SourceIP: srcIP, QueryName: queryName, QueryType: queryType, Callback: *cb}

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

func GetHttpRequests(cbName string, since *time.Time) []HttpRequest {
	var httpReqs []HttpRequest

	query := dbHandle.Table("http_requests").Joins("Callback").Where("Callback.name = ?", cbName)
	if since != nil {
		query = query.Where("`http_requests`.`created_at` >= ?", since)
	}

	results := query.Find(&httpReqs)
	if results.Error != nil {
		log.Println(results.Error)
	}

	return httpReqs
}

func GetDnsRequests(cbName string, since *time.Time) []DnsRequest {
	var dnsReqs []DnsRequest

	query := dbHandle.Table("dns_requests").Joins("Callback").Where("Callback.name = ?", cbName)
	if since != nil {
		query = query.Where("`dns_requests`.`created_at` >= ?", since)
	}

	results := query.Find(&dnsReqs)
	if results.Error != nil {
		log.Println(results.Error)
	}

	return dnsReqs
}
