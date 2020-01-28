package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAbout tests GET / request handler
func TestIndex(t *testing.T) {

	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	Index(httprr, req)
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Index" {
		t.Errorf("want the body to contain the word %q", "Index")
	}
}

func TestAdminSearch(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/s", AdminSearch)
	testServ := httptest.NewTLSServer(mux)
	defer testServ.Close()

	testClient := testServ.Client()
	url := testServ.URL

	resp, err := testClient.Get(url + "/s")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Adminsearch" {
		t.Errorf("want the body to contain the word %q", "Adminsearch")
	}
}
