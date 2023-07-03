package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTest = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availability", "/search-availability", "POST", []postData{
		{ key: "start", value: "2022-01-01"},
		{ key: "end", value: "2022-01-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{ key: "start", value: "2022-01-01"},
		{ key: "end", value: "2022-01-02"},
	}, http.StatusOK},
	{"make-reservation-post", "/make-reservation", "POST", []postData{
		{ key: "first_name", value: "John"},
		{ key: "last_name", value: "Smith"},
		{ key: "email", value: "test@test.com"},
		{ key: "phone", value: "444-444-444"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	// create a web server in the test env to return a status code.
	testServer := httptest.NewTLSServer(routes)

	defer testServer.Close()

	for _, e := range theTest {
		if e.method == "GET" {
			resp, err := testServer.Client().Get(testServer.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but god %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x :=range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := testServer.Client().PostForm(testServer.URL + e.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but god %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
