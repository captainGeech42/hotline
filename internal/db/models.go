package db

import "gorm.io/gorm"

type Callback struct {
	gorm.Model
	Name string
}

type HttpRequest struct {
	gorm.Model
	SourceIP   string
	Url        string
	Headers    string // base64
	Body       string // base64
	CallbackID int    // fkey to callback
	Callback   Callback
}

type DnsRequest struct {
	gorm.Model
	SourceIP    string
	RequestName string
	CallbackID  int // fkey to callback
	Callback    Callback
}
