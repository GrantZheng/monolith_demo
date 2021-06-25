// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "github.com/GrantZheng/monolith_demo/pay/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makePayHandler(m, endpoints, options["Pay"])
	return m
}
