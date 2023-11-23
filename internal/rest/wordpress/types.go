package wordpress

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResponsCategories struct {
	Id            int           `json:"id"`
	Count         int           `json:"count"`
	Description   string        `json:"description"`
	Link          string        `json:"link"`
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Taxonomy      string        `json:"taxonomy"`
	Parent        int           `json:"parent"`
	Meta          []interface{} `json:"meta"`
	YoastHead     string        `json:"yoast_head"`
	YoastHeadJson YoastHeadJson `json:"yoast_head_json"`
	Links         struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		WpPostType []struct {
			Href string `json:"href"`
		} `json:"wp:post_type"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
}

type Robots struct {
	Index           string `json:"index"`
	Follow          string `json:"follow"`
	MaxSnippet      string `json:"max-snippet"`
	MaxImagePreview string `json:"max-image-preview"`
	MaxVideoPreview string `json:"max-video-preview"`
}

type IsPartOf struct {
	Id string `json:"@id"`
}

type Breadcrumb struct {
	Id string `json:"@id"`
}

type ItemListElement struct {
	Type     string `json:"@type"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	Item     string `json:"item,omitempty"`
}
type Publisher struct {
	Id string `json:"@id"`
}
type Target struct {
	Type        string `json:"@type"`
	UrlTemplate string `json:"urlTemplate"`
}

type Logo struct {
	Type       string `json:"@type"`
	InLanguage string `json:"inLanguage"`
	Id         string `json:"@id"`
	Url        string `json:"url"`
	ContentUrl string `json:"contentUrl"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Caption    string `json:"caption"`
}
type Image struct {
	Id string `json:"@id"`
}

type PotentialAction struct {
	Type       string `json:"@type"`
	Target     Target `json:"target"`
	QueryInput string `json:"query-input"`
}

type Graph struct {
	Type            string            `json:"@type"`
	Id              string            `json:"@id"`
	Url             string            `json:"url,omitempty"`
	Name            string            `json:"name,omitempty"`
	IsPartOf        IsPartOf          `json:"isPartOf,omitempty"`
	Breadcrumb      Breadcrumb        `json:"breadcrumb,omitempty"`
	InLanguage      string            `json:"inLanguage,omitempty"`
	ItemListElement []ItemListElement `json:"itemListElement,omitempty"`
	Description     string            `json:"description,omitempty"`
	Publisher       Publisher         `json:"publisher,omitempty"`
	PotentialAction []PotentialAction `json:"potentialAction,omitempty"`
	Logo            Logo              `json:"logo,omitempty"`
	Image           Image             `json:"image,omitempty"`
	SameAs          []string          `json:"sameAs,omitempty"`
}
type Schema struct {
	Context string  `json:"@context"`
	Graph   []Graph `json:"@graph"`
}
type YoastHeadJson struct {
	Title         string `json:"title"`
	Robots        Robots `json:"robots"`
	Canonical     string `json:"canonical"`
	OgLocale      string `json:"og_locale"`
	OgType        string `json:"og_type"`
	OgTitle       string `json:"og_title"`
	OgUrl         string `json:"og_url"`
	OgSiteName    string `json:"og_site_name"`
	TwitterCard   string `json:"twitter_card"`
	TwitterSite   string `json:"twitter_site"`
	Schema        Schema `json:"schema"`
	Description   string `json:"description,omitempty"`
	OgDescription string `json:"og_description,omitempty"`
}

//type
//type
//type
//type
//type
//type
//type
