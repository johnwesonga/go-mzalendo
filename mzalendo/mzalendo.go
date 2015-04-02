/*
Package mzalendo provides a client for using the Mzalendo API.
Access different parts of the  API using the various
services.
         apiToken := "api-token"
         client := mzalendo.NewClient(nil)
         // search for an interest
         results, _, err := client.Topics.SearchForInterest("Clojure")
The full Mzalendo API is documented at http://info.mzalendo.com/help/api.
*/
package mzalendo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://kenyan-politicians.popit.mysociety.org/api/v0.1"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL
	Persons *PersonService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Persons = &PersonService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlString string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil

}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
