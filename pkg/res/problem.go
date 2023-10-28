package res

import (
	"encoding/json"
	"net/http"
)

type Problem struct {
	Extensions map[string]interface{}
	Type       string
	Title      string
	Detail     string
	Instance   string
	Status     int
}

func (p Problem) MarshalJSON() ([]byte, error) {
	c := make(map[string]interface{})

	c["type"] = "about:blank"
	if p.Type != "" {
		c["type"] = p.Type
	}

	c["status"] = http.StatusInternalServerError
	if p.Status != 0 {
		c["status"] = p.Status
	}

	c["title"] = http.StatusText(c["status"].(int))
	if p.Title != "" {
		c["title"] = p.Title
	}

	c["detail"] = p.Detail

	if p.Instance != "" {
		c["instance"] = p.Instance
	}

	for k, v := range p.Extensions {
		switch k {
		case "type", "status", "title", "detail", "instance":
			c["_"+k] = v
		default:
			c[k] = v
		}
	}

	return json.Marshal(c) //nolint:wrapcheck
}

func (p Problem) StatusCode() int {
	if p.Status == 0 {
		return http.StatusInternalServerError
	}

	return p.Status
}

func (p Problem) Header() http.Header {
	res := make(http.Header)
	res.Set("Content-Type", "application/problem+json")

	return res
}

type Option func(e *Problem)

func WithExtension(key string, val interface{}) Option {
	return func(e *Problem) {
		if e.Extensions == nil {
			e.Extensions = make(map[string]interface{})
		}

		e.Extensions[key] = val
	}
}

func InternalServerError(err error, options ...Option) Problem {
	id := "undefined" // FIXME

	e := Problem{
		Status: http.StatusInternalServerError,
		Detail: err.Error(),
		Extensions: map[string]interface{}{
			"tracking_code": id,
		},
	}

	for i := range options {
		options[i](&e)
	}

	return e
}

func BadRequest(detail string, options ...Option) Problem {
	e := Problem{
		Status:     http.StatusBadRequest,
		Detail:     detail,
		Extensions: map[string]interface{}{},
	}

	for i := range options {
		options[i](&e)
	}

	return e
}
