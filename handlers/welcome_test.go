package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcome(t *testing.T) {
	router := Router()
	ts := httptest.NewServer(router)
	client, err := ts.Client().Get(ts.URL + "/")
	if err != nil {
		t.Error(err)
	}
	if client.StatusCode != http.StatusOK {
		t.Error("Wrong response")
	}
	body, err := ioutil.ReadAll(client.Body)

	if err != nil {
		t.Error(err)
	}
	want := "You are welcome"

	if string(body) != want {
		t.Error("Wrong response from server")
	}

	client, err = ts.Client().Post(ts.URL+"/", "application/json", nil)
	if err != nil {
		t.Error(err)
	}
	if client.StatusCode != http.StatusMethodNotAllowed {
		t.Error("Method checking failed")
	}
}
