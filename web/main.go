package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	b64 "encoding/base64"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type Response struct {
	Took     int32           `json:"took"`
	TimedOut bool            `json:"timed_out"`
	Shards   json.RawMessage `json:"shards"`
	Hits     struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string            `json:"_index"`
			Type      string            `json:"_type"`
			Id        string            `json:"_id"`
			Score     float64           `json:"_score"`
			Timestamp string            `json:@timestamp`
			Source    map[string]string `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func make_search_request(term string) Response {
	const user = "admin"
	const pass = "Complexpass#123"
	const authString = user + ":" + pass
	authBytes := make([]byte, len(authString))
	copy(authBytes, authString)
	encodedCredentials := b64.StdEncoding.EncodeToString(authBytes)

	const zincSearchURL = "http://localhost:4080/api/enron_mails/_search"
	jsonBody := `{
    "search_type": "querystring",
    "query": {
        "term": "@term"
    },
    "sort_fields": ["-@timestamp"],
    "from": 0,
    "max_results": 100,
    "_source": []
	}`
	jsonRequest := strings.Replace(jsonBody, "@term", term, 1)
	req, err := http.NewRequest(http.MethodPost, zincSearchURL, strings.NewReader(jsonRequest))
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+encodedCredentials)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
	}

	//log.Println(string(resBody))

	var myStoredVariable Response
	if err := json.Unmarshal(resBody, &myStoredVariable); err != nil {
		log.Println("error", err)
	}

	//log.Println(myStoredVariable)

	return myStoredVariable
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	fs := http.FileServer(http.Dir("ui/dist"))

	// static files
	r.Handle("/*", http.StripPrefix("/", fs))

	// api routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/search/{searchTerm}", func(r chi.Router) {
			r.Get("/", GetSearch)
		})
	})

	http.ListenAndServe(":3000", r)
}

type Match struct {
	To      string
	From    string
	Subject string
	Body    string
}

type SearchResponse struct {
	SearchTerm string
	Matches    []Match
}

func GetSearch(w http.ResponseWriter, r *http.Request) {
	if searchTerm := chi.URLParam(r, "searchTerm"); searchTerm != "" {

		results := make_search_request(searchTerm)

		//log.Printf("took %d\n", results.Took)

		var response SearchResponse

		response.SearchTerm = searchTerm

		response.Matches = make([]Match, len(results.Hits.Hits))

		for i, hit := range results.Hits.Hits {
			response.Matches[i] = Match{
				To:      hit.Source["To"],
				From:    hit.Source["From"],
				Subject: hit.Source["Subject"],
				Body:    hit.Source["Body"],
			}
		}

		// jsonResponse, err := json.Marshal(response)

		// if err != nil {
		// 	log.Println("wat")
		// }

		render.JSON(w, r, response)
	}
}
