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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://kenyan-politicians.popit.mysociety.org/api/v0.1/"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL
	Api     *Api
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Api = &Api{client: c}

	return c
}

func (c *Client) NewRequest(method, urlString string) (*http.Request, error) {
	rel, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)

	log.Print(u.String())

	if err != nil {
		return nil, err
	}

	return req, nil

}

// Response is a Prismatic API response.  This wraps the standard http.Response
// returned from Prismatic.
type Response struct {
	*http.Response
}

// newResponse creats a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v: %v: %v", e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode)
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)

	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	err = json.NewDecoder(resp.Body).Decode(v)

	return response, err
}
