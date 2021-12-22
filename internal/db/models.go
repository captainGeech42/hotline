package db

import "gorm.io/gorm"

type Callback struct {
	gorm.Model
	Name string
}

type HttpRequest struct {
	gorm.Model
	SourceIP   string
	URI        string
	Host       string
	Method     string
	Headers    string // base64
	Body       string // base64
	CallbackID int    // fkey to callback
	Callback   Callback
}

type DnsRequest struct {
	gorm.Model
	SourceIP   string
	QueryName  string
	QueryType  string
	CallbackID int // fkey to callback
	Callback   Callback
}
