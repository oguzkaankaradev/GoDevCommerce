package shared

import (
	"log"
	"net/http"
)

type HttpContextAccessor struct {
	HttpContext *http.Request
}

func (httpContextAccessor *HttpContextAccessor) IsRequestAvailable() bool {

	if httpContextAccessor.HttpContext == nil {
		return false
	}

	if httpContextAccessor == nil {
		return false
	} // burayıı kontrol et.

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	return true
}
