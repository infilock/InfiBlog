package res

import (
	"encoding/json"

	"net/http"

	"gopkg.in/yaml.v2"
)

func Done(w http.ResponseWriter, r *http.Request, rsp interface{}) {
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

	if r.Header.Get("Accept") == "application/x-yaml" {
		w.Header().Set("Content-Type", "application/x-yaml; charset=utf-8")
		_ = yaml.NewEncoder(w).Encode(rsp)

		return
	}

	_ = json.NewEncoder(w).Encode(rsp) //nolint:errchkjson
}
