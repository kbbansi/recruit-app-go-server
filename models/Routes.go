package models

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunction http.HandlerFunc
}

type Routes []Route
