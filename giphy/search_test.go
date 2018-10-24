package giphy

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		q string
	}{
		{"foo"},
		{"bar"},
	} {
		requests := []*http.Request{}
		server, client := serverAndClient(200, `{
			"meta": {
				"status": 200,
				"msg": "OK"
			}
		}`, &requests)
		defer server.Close()

		_, err := client.Search(tt.q)
		if err != nil {
			t.Errorf(`unexpected error %v`, err)
		}

		if got := len(requests); got != 1 {
			t.Fatalf(`unexpected number of requests %d`, got)
		}

		r := requests[0]

		if got := r.URL.Path; got != "/v1/gifs/search" {
			t.Errorf(`unexpected path %#v`, got)
		}

		params := r.URL.Query()

		if got := params.Get("api_key"); got != PublicBetaAPIKey {
			t.Errorf(`unexpected api_key %#v`, got)
		}

		if got := params.Get("limit"); got != "25" {
			t.Errorf(`unexpected limit %#v`, got)
		}

		if got := params.Get("rating"); got != "g" {
			t.Errorf(`unexpected rating %#v`, got)
		}

		if got := params.Get("q"); got != tt.q {
			t.Errorf(`unexpected q %#v`, got)
		}
	}
}

func serverAndClient(code int, body string, requests ...*[]*http.Request) (*httptest.Server, *Client) {
	return testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		if len(requests) > 0 {
			*requests[0] = append(*requests[0], r)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write([]byte(body))
	})
}

func testServerAndClient(f func(http.ResponseWriter, *http.Request)) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(f))

	client := NewClient()
	client.baseURL, _ = url.Parse(server.URL)

	return server, client
}
