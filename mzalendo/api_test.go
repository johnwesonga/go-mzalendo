package mzalendo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Prismatic client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url

}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
}

func TestGetPerson(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/persons/core_person:", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintln(w, `{"result":{"name": "Foo", "gender": "male", "birth_date": "0000-00-00", "memberships": [{"role": "coalition member", "organization_id": "1"}]}}`)
	})

	results, err := client.Persons.GetPerson("1000")
	if err != nil {
		t.Errorf("GetPerson returned error: %v", err)
	}

	want := Person{
		Result{Name: "Foo", Gender: "male", BirthDate: "0000-00-00", Memberships: []Membership{{Role: "coalition member", OrganizationId: "1"}}},
	}

	if !reflect.DeepEqual(results, want) {
		t.Errorf("GetPerson returned %+v, want %+v", results, want)
	}

}
