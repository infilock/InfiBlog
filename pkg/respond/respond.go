package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

func Done(w http.ResponseWriter, _ *http.Request, rsp interface{}) {
	if rsp == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	if withHeader, ok := rsp.(interface{ Header() http.Header }); ok {
		for k, vv := range withHeader.Header() {
			for _, v := range vv {
				w.Header().Set(k, v)
			}
		}
	}

	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	if withStatusCode, ok := rsp.(interface{ StatusCode() int }); ok {
		statusCode := withStatusCode.StatusCode()
		w.WriteHeader(statusCode)

		if statusCode == http.StatusNoContent {
			return
		}
	}

	errEncode := json.NewEncoder(w).Encode(rsp)
	log.Println(errEncode)
}

type WithStatusCreated struct{}

func (WithStatusCreated) StatusCode() int {
	return http.StatusCreated
}
