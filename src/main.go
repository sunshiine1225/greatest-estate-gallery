package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
)

type SanityResult struct {
	Query  string          `json:"query"`
	Result json.RawMessage `json:"result"`
}

var (
	SANITY_PROJECT_ID string
	DATASET           string
	SANITY_AUTH_TOKEN string
)

func init() {
	SANITY_PROJECT_ID = os.Getenv("SANITY_PROJECT_ID")
	DATASET = os.Getenv("SANITY_DATASET")
	SANITY_AUTH_TOKEN = os.Getenv("SANITY_AUTH_TOKEN")
}

func sanityQuery(query string) (SanityResult, error) {
	query = url.QueryEscape(query)

	endpoint := fmt.Sprintf(`https://%s.apicdn.sanity.io/v2021-10-21/data/query/%s?query=%s`, SANITY_PROJECT_ID, DATASET, query)

	request, _ := http.NewRequest("GET", endpoint, bytes.NewBuffer([]byte{}))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", SANITY_AUTH_TOKEN))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return SanityResult{}, nil
	}

	var response SanityResult
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return SanityResult{}, err
	}

	return response, nil
}

func Projects(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/featured-container.html"))

	resp, _ := sanityQuery(`*[_type == 'project'] {
  'name': project_name,
  'type': type,
  'lower_limit': lower_limit,
  'upper_limit': upper_limit,
  'image_url': images[0].asset->url,
  'area': location.area,
  'city': location.city,
  'id': _id
} | order(_createdAt) [0...6]`)

	type ProjectThumbnail struct {
		Name       string `json:"name"`
		Type       string `json:"type"`
		ImageUrl   string `json:"image_url"`
		LowerLimit struct {
			LowerLimitPrice float64 `json:"lower_limit_price"`
			Denomination    string  `json:"denomination"`
		} `json:"lower_limit"`
		UpperLimit struct {
			UpperLimitPrice float64 `json:"upper_limit_price"`
			Denomination    string  `json:"denomination"`
		} `json:"upper_limit"`
		Area string `json:"area"`
		City string `json:"city"`
		Id   string `json:"id"`
	}

	var projectThumbnails []ProjectThumbnail
	_ = json.Unmarshal(resp.Result, &projectThumbnails)

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, struct {
		ProjectThumbnails []ProjectThumbnail
	}{projectThumbnails})
}

func ResaleProperties(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/resale-properties.html"))

	resp, _ := sanityQuery(`*[_type == 'property'] {
  'property_name': property_name,
  'area': location.area,
  'city': location.city,
  'num_bedrooms': num_bedrooms,
  'num_bathrooms': num_bathrooms,
  'image': images[0].asset->url,
  'price': price.price,
  'denomination': price.denomination,
  'size': size,
  'property_type': type,
} | order(_createdAt) [0...4]`)

	type PropertyThumbnail struct {
		Name         string  `json:"property_name"`
		Area         string  `json:"area"`
		City         string  `json:"city"`
		Bedrooms     int     `json:"num_bedrooms"`
		Bathrooms    int     `json:"num_bathrooms"`
		ImageUrl     string  `json:"image"`
		Price        float64 `json:"price"`
		Denomination string  `json:"denomination"`
		Type         string  `json:"property_type"`
		Size         float64 `json:"size"`
	}

	var propertyThumbnails []PropertyThumbnail
	_ = json.Unmarshal(resp.Result, &propertyThumbnails)

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, struct {
		PropertyThumbnails []PropertyThumbnail
	}{propertyThumbnails})
}

func FeaturedPlots(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/featured-plots.html"))

	resp, _ := sanityQuery(`*[_type == 'plot'] {
  'plot_name': project_name,
  'price': {
    'lower_limit': string(lower_limit.lower_limit_price) + " " + lower_limit.denomination,
    'upper_limit': string(upper_limit.upper_limit_price) + " " + upper_limit.denomination,
  },
  'sizes': {
    'lower_limit_size': lower_limit_size,
    'upper_limit_size': upper_limit_size,
  },
  'area': location.area,
  'city': location.city,
  'image': images[0].asset->url,
}[0...6]`)

	type PlotThumbnail struct {
		Name  string `json:"plot_name"`
		Price struct {
			LowerLimit string `json:"lower_limit"`
			UpperLimit string `json:"upper_limit"`
		} `json:"price"`
		Sizes struct {
			LowerLimit float64 `json:"lower_limit_size"`
			UpperLimit float64 `json:"upper_limit_size"`
		} `json:"sizes"`
		Area     string `json:"area"`
		City     string `json:"city"`
		ImageUrl string `json:"image"`
	}

	var plotThumbnails []PlotThumbnail
	_ = json.Unmarshal(resp.Result, &plotThumbnails)

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, struct {
		PlotThumbnails []PlotThumbnail
	}{plotThumbnails})
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
	fmt.Println("Email: ", r.FormValue("email"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Homepage).Methods(http.MethodGet)
	r.HandleFunc("/contact", ContactPost).Methods(http.MethodPost)
	r.HandleFunc("/section/featured-projects", Projects).Methods(http.MethodGet)
	r.HandleFunc("/section/resale-container", ResaleProperties).Methods(http.MethodGet)
	r.HandleFunc("/section/featured-plots", FeaturedPlots).Methods(http.MethodGet)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", r)
}
