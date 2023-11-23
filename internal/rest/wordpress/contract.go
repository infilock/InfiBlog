package wordpress

import "net/http"

type Contract interface {
	//TODO: per_page |	Maximum number of items to be returned in result set | Default: 10 | implement to code handler api
	HandlerListCategories() http.HandlerFunc
}
