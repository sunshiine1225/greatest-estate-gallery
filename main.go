package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type SanityResult struct {
	Query  string          `json:"query"`
	Result json.RawMessage `json:"result"`
}

type PropertyThumbnail struct {
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	ImageUrl   string  `json:"image_url"`
	LowerLimit float64 `json:"lower_limit"`
	UpperLimit float64 `json:"upper_limit"`
	Id         string  `json:"id"`
}

var (
	SANITY_PROJECT_ID string
	DATASET           string
	SANITY_AUTH_TOKEN string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SANITY_PROJECT_ID = os.Getenv("SANITY_PROJECT_ID")
	DATASET = os.Getenv("SANITY_DATASET")
	SANITY_AUTH_TOKEN = os.Getenv("SANITY_AUTH_TOKEN")
}

func propertyThumbnailsGet() []PropertyThumbnail {
	query := url.QueryEscape(`*[_type == 'property'] {
  'name': property_name,
  'type': type,
  'lower_limit': price_range.lower_limit,
  'upper_limit': price_range.upper_limit,
  'image_url': images[0].asset->url,
  'id': _id
} | order(_createdAt) [0...6]`)

	endpoint := fmt.Sprintf(`https://%s.api.sanity.io/v2021-10-21/data/query/%s?query=%s`, SANITY_PROJECT_ID, DATASET, query)

	request, _ := http.NewRequest("GET", endpoint, bytes.NewBuffer([]byte{}))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", SANITY_AUTH_TOKEN))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return []PropertyThumbnail{}
	}

	var response SanityResult
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err.Error())
		return []PropertyThumbnail{}
	}

	var propertyThumbnails []PropertyThumbnail
	err = json.Unmarshal(response.Result, &propertyThumbnails)
	if err != nil {
		fmt.Println(err.Error())
		return []PropertyThumbnail{}
	}

	return propertyThumbnails
}

func FeaturedContainer(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/featured-container.html"))

	w.WriteHeader(http.StatusOK)

	tmpl.Execute(w, struct {
		PropertyThumbnails []PropertyThumbnail
	}{propertyThumbnailsGet()})
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html", "./templates/contact-form.html"))

	w.WriteHeader(http.StatusOK)

	tmpl.Execute(w, nil)
}

func ContactPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Name: ", r.FormValue("name"))
	fmt.Println("Message: ", r.FormValue("message"))
	fmt.Println("Subject: ", r.FormValue("subject"))
	fmt.Print("Email: ", r.FormValue("email"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Homepage).Methods(http.MethodGet)
	r.HandleFunc("/contact", ContactPost).Methods(http.MethodPost)
	r.HandleFunc("/section/featured-container", FeaturedContainer).Methods(http.MethodGet)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", r)
}
