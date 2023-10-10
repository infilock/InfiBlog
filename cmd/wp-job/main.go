package main

import (
	"bytes"
	"context"
	cfg "crawler/internal/pkg/postgresql"
	repoArticle "crawler/internal/repository/postgresql/article"
	repoQuestion "crawler/internal/repository/postgresql/question"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/cavaliergopher/grab/v3"
	"github.com/go-co-op/gocron"
	"github.com/gosimple/slug"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq" // import postgres driver
	"github.com/nasermirzaei89/env"
	"github.com/pkg/errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	location, err := time.LoadLocation(env.MustGetString("JOB_LOCATION_WP"))
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(location) //time.UTC

	//Every schedules a new periodic Job with an interval.
	_, errJobOne := scheduler.Every(1).Day().At(env.MustGetString("JOB_TIME_WP")).Do(func() {
		log.Println("- start job Wordpress ")

		postgresql := cfg.GetPostgresqlConfig()
		defer func() { _ = postgresql.Close() }()

		articleRepo := repoArticle.NewArticleRepository(postgresql)
		qr := repoQuestion.NewQuestionRepository(postgresql)
		items, errListStatus := articleRepo.ListStatus(context.Background(), "draft")
		if errListStatus != nil {
			log.Fatal(errListStatus)
		}

		for i := 0; i < len(items); i++ {
			//begin: check folder picture
			_, errStat := os.Stat(env.MustGetString("UPLOAD_PATH"))
			if errStat != nil {
				errMkdir := os.Mkdir("picture", os.ModePerm)
				if errMkdir != nil {
					log.Fatal("Mkdir: ", errMkdir)
				}
			}
			//end: check folder picture

			//begin: download picture order title question form googleapis
			fileDownload := downloadPicture(items[i].Title)
			//end: download picture order title question form googleapis

			if fileDownload == "" {
				type Article struct {
					Title         string `json:"title"`
					Content       string `json:"content"`
					Format        string `json:"format"`
					CommentStatus string `json:"comment_status"`
					PingStatus    string `json:"ping_status"`
					Slug          string `json:"slug"`
					Categories    []int  `json:"categories"`
					Tags          []int  `json:"tags"`
					FeaturedMedia int    `json:"featured_media"`
					Status        string `json:"status"`
				}

				//find tag is and category id
				resFind, errFind := qr.Find(context.Background(), items[i].QuestionID)
				if errFind != nil {
					log.Println("Find: ", errFind)

					continue
				}

				t := Article{
					Title:         items[i].Title,
					Content:       items[i].Content,
					Format:        "standard",
					CommentStatus: "open",
					PingStatus:    "open",
					Slug:          slug.Make(items[i].Title),
					Categories:    convertToIntArray(strings.Split(resFind.CategoryID, ",")),
					Tags:          convertToIntArray(strings.Split(resFind.TagID, ",")),
					Status:        "draft",
				}

				c, errMarshal := json.Marshal(t)
				if errMarshal != nil {
					log.Fatal("Marshal: ", errMarshal)
				}

				payload := strings.NewReader(string(c))

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
				defer cancel()

				req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, env.MustGetString("WP_URL_CREATE_ARTICLE"), payload)
				if errRequest != nil {
					log.Fatal("error on new request with context: ", errRequest)
				}

				req.SetBasicAuth(env.MustGetString("WP_USERNAME"), env.MustGetString("WP_PASSWORD"))
				req.Header.Add("Content-Type", "application/json")

				client := &http.Client{}
				res, errDo := client.Do(req)
				if errDo != nil {
					log.Fatal(errDo)
				}
				defer res.Body.Close()

				if strings.Contains(res.Status, "201") {
					errUpdateArticle := articleRepo.UpdateArticleStatus(context.Background(), items[i].ID)
					if errUpdateArticle != nil {
						log.Println("UpdateArticleStatus: ", errUpdateArticle)
						continue //nolint:nlreturn
					}
				}

				continue
			}

			//begin: create media for article
			mediaID := mediaWordPress(fileDownload)
			//end: create media for article

			//begin: remove picture file
			errRemove := os.Remove(fileDownload)
			if errRemove != nil {
				log.Fatal("error on remove file:", errRemove)
			}
			//end: remove picture file

			type Article struct {
				Title         string `json:"title"`
				Content       string `json:"content"`
				Format        string `json:"format"`
				CommentStatus string `json:"comment_status"`
				PingStatus    string `json:"ping_status"`
				Slug          string `json:"slug"`
				Categories    []int  `json:"categories"`
				Tags          []int  `json:"tags"`
				FeaturedMedia int    `json:"featured_media"`
				Status        string `json:"status"`
			}

			//find tag is and category id
			resFind, errFind := qr.Find(context.Background(), items[i].QuestionID)
			if errFind != nil {
				log.Println("Find: ", errFind)

				continue
			}

			t := Article{
				Title:         items[i].Title,
				Content:       items[i].Content,
				Format:        "standard",
				CommentStatus: "open",
				PingStatus:    "open",
				Slug:          slug.Make(items[i].Title),
				Categories:    convertToIntArray(strings.Split(resFind.CategoryID, ",")),
				Tags:          convertToIntArray(strings.Split(resFind.TagID, ",")),
				FeaturedMedia: mediaID,
				Status:        "publish",
			}

			c, errMarshal := json.Marshal(t)
			if errMarshal != nil {
				log.Fatal("Marshal: ", errMarshal)
			}

			payload := strings.NewReader(string(c))

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
			defer cancel()

			req, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, env.MustGetString("WP_URL_CREATE_ARTICLE"), payload)
			if errRequest != nil {
				log.Fatal("error on new request with context: ", errRequest)
			}

			req.SetBasicAuth(env.MustGetString("WP_USERNAME"), env.MustGetString("WP_PASSWORD"))
			req.Header.Add("Content-Type", "application/json")

			client := &http.Client{}
			res, errDo := client.Do(req)
			if errDo != nil {
				log.Fatal(errDo)
			}
			defer res.Body.Close()

			if strings.Contains(res.Status, "201") {
				errUpdateArticle := articleRepo.UpdateArticleStatus(context.Background(), items[i].ID)
				if errUpdateArticle != nil {
					log.Println("UpdateArticleStatus: ", errUpdateArticle)
					continue //nolint:nlreturn
				}
			}
		}
	})

	if errJobOne != nil {
		log.Fatal(errors.Wrap(errJobOne, "error on job wordpress"))
	}

	//StartBlocking starts all jobs and blocks the current thread.
	scheduler.StartBlocking()
}

func convertToIntArray(array []string) []int {
	var items = []int{}

	for _, i := range array {
		item, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}

type MediaWordPress struct {
	Id      int    `json:"id"`
	Date    string `json:"date"`
	DateGmt string `json:"date_gmt"`
	Guid    struct {
		Rendered string `json:"rendered"`
		Raw      string `json:"raw"`
	} `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGmt string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Title       struct {
		Raw      string `json:"raw"`
		Rendered string `json:"rendered"`
	} `json:"title"`
	Author            int           `json:"author"`
	CommentStatus     string        `json:"comment_status"`
	PingStatus        string        `json:"ping_status"`
	Template          string        `json:"template"`
	Meta              []interface{} `json:"meta"`
	PermalinkTemplate string        `json:"permalink_template"`
	GeneratedSlug     string        `json:"generated_slug"`
	Description       struct {
		Raw      string `json:"raw"`
		Rendered string `json:"rendered"`
	} `json:"description"`
	Caption struct {
		Raw      string `json:"raw"`
		Rendered string `json:"rendered"`
	} `json:"caption"`
	AltText      string `json:"alt_text"`
	MediaType    string `json:"media_type"`
	MimeType     string `json:"mime_type"`
	MediaDetails struct {
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		File     string `json:"file"`
		Filesize int    `json:"filesize"`
		Sizes    struct {
			Medium struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"medium"`
			Thumbnail struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"thumbnail"`
			MediumLarge struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"medium_large"`
			BetterAmpSmall struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"better-amp-small"`
			BetterAmpNormal struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"better-amp-normal"`
			BetterAmpLarge struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"better-amp-large"`
			PublisherTb1 struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-tb1"`
			PublisherSm struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-sm"`
			PublisherMg2 struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-mg2"`
			PublisherMd struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-md"`
			PublisherLg struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-lg"`
			PublisherFull struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-full"`
			PublisherTallSm struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-tall-sm"`
			PublisherTallLg struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-tall-lg"`
			PublisherTallBig struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filesize  int    `json:"filesize"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"publisher-tall-big"`
			Full struct {
				File      string `json:"file"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				MimeType  string `json:"mime_type"`
				SourceUrl string `json:"source_url"`
			} `json:"full"`
		} `json:"sizes"`
		ImageMeta struct {
			Aperture         string        `json:"aperture"`
			Credit           string        `json:"credit"`
			Camera           string        `json:"camera"`
			Caption          string        `json:"caption"`
			CreatedTimestamp string        `json:"created_timestamp"`
			Copyright        string        `json:"copyright"`
			FocalLength      string        `json:"focal_length"`
			Iso              string        `json:"iso"`
			ShutterSpeed     string        `json:"shutter_speed"`
			Title            string        `json:"title"`
			Orientation      string        `json:"orientation"`
			Keywords         []interface{} `json:"keywords"`
		} `json:"image_meta"`
	} `json:"media_details"`
	Post              interface{}   `json:"post"`
	SourceUrl         string        `json:"source_url"`
	MissingImageSizes []interface{} `json:"missing_image_sizes"`
	Links             struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		Author []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"author"`
		Replies []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"replies"`
		WpActionUnfilteredHtml []struct {
			Href string `json:"href"`
		} `json:"wp:action-unfiltered-html"`
		WpActionAssignAuthor []struct {
			Href string `json:"href"`
		} `json:"wp:action-assign-author"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
}

func mediaWordPress(fileNamePath string) int {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, errOpen := os.Open(fileNamePath)
	if errOpen != nil {
		log.Fatal("error on open the named file for reading:", errOpen)
	}

	defer file.Close()

	formFile, errCreateFormFile := writer.CreateFormFile("file", filepath.Base(fileNamePath))
	if errCreateFormFile != nil {
		log.Fatal("error on convenience wrapper around CreatePart:", errCreateFormFile)
	}

	_, errCopy := io.Copy(formFile, file)
	if errCopy != nil {
		log.Fatal("error on Copy copies from src to dst until either EOF is reached on src or an error occurs:", errCopy)
	}

	errClose := writer.Close()
	if errClose != nil {
		log.Fatal("error on Close finishes the multipart message: ", errClose)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	req, errNewRequest := http.NewRequestWithContext(ctx, http.MethodPost, env.MustGetString("WP_URL_CREATE_MEDIA"), payload)
	if errNewRequest != nil {
		log.Fatal("error on new request: ", errNewRequest)
	}

	req.SetBasicAuth(env.MustGetString("WP_USERNAME"), env.MustGetString("WP_PASSWORD"))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	res, errDo := client.Do(req)
	if errDo != nil {
		log.Fatal("error on Do sends an HTTP request and returns an HTTP response", errDo)
	}

	defer res.Body.Close()

	body, errReadAll := io.ReadAll(res.Body)
	if errReadAll != nil {
		log.Fatal("error on ReadAll reads from r until an error or EOF", errReadAll)
	}

	var wp = &MediaWordPress{}

	errUnmarshal := json.Unmarshal(body, wp)
	if errUnmarshal != nil {
		log.Fatal("error on Unmarshal parses the JSON-encoded data: ", errUnmarshal)
	}

	return wp.Id
}

type GoogleSearch struct {
	Kind string `json:"kind"`
	Url  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
			ImgSize        string `json:"imgSize"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
			ImgSize        string `json:"imgSize"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind        string `json:"kind"`
		Title       string `json:"title"`
		HtmlTitle   string `json:"htmlTitle"`
		Link        string `json:"link"`
		DisplayLink string `json:"displayLink"`
		Snippet     string `json:"snippet"`
		HtmlSnippet string `json:"htmlSnippet"`
		Mime        string `json:"mime"`
		FileFormat  string `json:"fileFormat"`
		Image       struct {
			ContextLink     string `json:"contextLink"`
			Height          int    `json:"height"`
			Width           int    `json:"width"`
			ByteSize        int    `json:"byteSize"`
			ThumbnailLink   string `json:"thumbnailLink"`
			ThumbnailHeight int    `json:"thumbnailHeight"`
			ThumbnailWidth  int    `json:"thumbnailWidth"`
		} `json:"image"`
	} `json:"items"`
}

func downloadPicture(title string) string {
	var res *http.Response

	var gs = &GoogleSearch{}

	var filename string

	errRetry := retry.Do(
		func() error {
			var errRequest error
			//TODO: bodyclose
			res, errRequest = http.Get( //nolint:bodyclose,noctx
				"https://www.googleapis.com/customsearch/v1?" +
					"cx=" + env.MustGetString("GOOGLE_SEARCH_ID") +
					"&q=" + strings.Replace(title, " ", "%20", -1) + "site%3www.pixabay.com+OR+site%3www.freeimages.com" +
					"&searchType=image" +
					"&num=1" +
					"&imgSize=medium" +
					"&key=" + env.MustGetString("GOOGLE_API_KEY"))

			if errRequest != nil {
				log.Fatal("error on Get issues a GET to the specified URL", errRequest)
			}

			if res.Status != "200 OK" {
				return errors.New(fmt.Sprint("status code => ", res.Status))
			}

			body, errReadAll := io.ReadAll(res.Body)
			if errReadAll != nil {
				log.Println("error on ReadAll reads from r until an error or EOF", errReadAll)
			}

			errUnmarshal := json.Unmarshal(body, gs)
			if errUnmarshal != nil {
				log.Println("error on Unmarshal parses the JSON-encoded data: ", errUnmarshal)
			}

			fileDir, errGrab := grab.Get(env.MustGetString("UPLOAD_PATH"), gs.Items[0].Link)
			if errGrab != nil {
				log.Println("error on Get sends a HTTP request and downloads the content: ", errGrab)
			}

			filename = fileDir.Filename

			return nil
		},
		//Attempts set count of retry.
		retry.Attempts(3),
		//Delay set delay between retry.
		retry.Delay(10*time.Second),
		//OnRetry function callback are called each retry log each retry.
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retrying request after error: %v", err)
		}),
	)
	if errRetry != nil {
		log.Println("error retry", errRetry)
		return "" //nolint:nlreturn
	}

	defer res.Body.Close()

	return filename
}
