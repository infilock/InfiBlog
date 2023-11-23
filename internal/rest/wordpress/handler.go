package wordpress

import (
	"encoding/json"
	"fmt"
	"github.com/infilock/InfiBlog/config"
	"github.com/infilock/InfiBlog/pkg/res"
	"io"
	"net/http"
)

func (h handler) HandlerListCategories() http.HandlerFunc {
	hh := func(w http.ResponseWriter, r *http.Request) {

		//begin: get list all categories from WordPress
		req, err := http.NewRequest(http.MethodGet, config.WordPressAPICategories.Get(), nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Add adds the key, value pair to the header.
		req.Header.Add(config.WordPressUsername.Get(), config.WordPressPassword.Get())

		client := &http.Client{}
		resx, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resx.Body.Close()

		body, err := io.ReadAll(resx.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var listCategories []ResponsCollection

		errUnmarshal := json.Unmarshal(body, &listCategories)
		if errUnmarshal != nil {
			res.Done(w, r, res.InternalServerError(errUnmarshal))

			return
		}
		//end: get list all categories from WordPress

		collectionCategories := make([]Collection, 0)
		for i := 0; i < len(listCategories); i++ {
			itemCategory := Collection{
				ID:   listCategories[i].Id,
				Name: listCategories[i].Name,
			}
			collectionCategories = append(collectionCategories, itemCategory)
		}

		res.Done(w, r, collectionCategories)

		return
	}

	return hh
}

func (h handler) HandlerListTags() http.HandlerFunc {
	hh := func(w http.ResponseWriter, r *http.Request) {

		//begin: get list all tags from WordPress
		req, err := http.NewRequest(http.MethodGet, config.WordPressAPITag.Get(), nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Add adds the key, value pair to the header.
		req.Header.Add(config.WordPressUsername.Get(), config.WordPressPassword.Get())

		client := &http.Client{}
		resx, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resx.Body.Close()

		body, err := io.ReadAll(resx.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var listTags []ResponsCollection

		errUnmarshal := json.Unmarshal(body, &listTags)
		if errUnmarshal != nil {
			res.Done(w, r, res.InternalServerError(errUnmarshal))

			return
		}
		//end: get list all tags from WordPress

		collectionTags := make([]Collection, 0)
		for i := 0; i < len(listTags); i++ {
			itemTag := Collection{
				ID:   listTags[i].Id,
				Name: listTags[i].Name,
			}
			collectionTags = append(collectionTags, itemTag)
		}

		res.Done(w, r, collectionTags)

		return
	}

	return hh
}
