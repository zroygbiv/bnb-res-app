package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key 	string
	value	string
}

var tests = []struct {
	name 								string
	url 								string
	method 							string
	params 							[]postData
	expectedStatusCode 	int
}{
	// GET
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"kitchen", "/kitchen", "GET", []postData{}, http.StatusOK},
	{"common", "/common", "GET", []postData{}, http.StatusOK},
	{"bedroom", "/bedroom", "GET", []postData{}, http.StatusOK},
	{"avail", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"makeres", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	// POST
	{"post-avail", "/search-availability", "POST", []postData{
		{key: "start", value: "01-01-2023"},
		{key: "end", value: "12-01-2023"},
	}, http.StatusOK},
	{"post-makeres", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Lennon"},
		{key: "email", value: "plastic_ono_band@gmail.com"},
		{key: "phone", value: "555-555-5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range tests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params{
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}