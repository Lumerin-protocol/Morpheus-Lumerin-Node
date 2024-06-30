package proxyapi

import "net/http"

type ResponderFlusher interface {
	http.ResponseWriter
	http.Flusher
}
