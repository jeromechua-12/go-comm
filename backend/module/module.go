package module

import "net/http"

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc	
	Protected   bool
}

type Module interface {
	Routes() []Route
}
