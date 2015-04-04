package mzalendo

import (
	"fmt"
	"log"
	"time"
)

type Api struct {
	client *Client
}

type Person struct {
	Result struct {
		BirthDate       string        `json:"birth_date"`
		ContactDetails  []interface{} `json:"contact_details"`
		Gender          string        `json:"gender"`
		HonorificSuffix string        `json:"honorific_suffix"`
		HtmlURL         string        `json:"html_url"`
		ID              string        `json:"id"`
		Identifiers     []interface{} `json:"identifiers"`
		Image           string        `json:"image"`
		Images          []struct {
			ProxyURL string `json:"proxy_url"`
			URL      string `json:"url"`
		} `json:"images"`
		Links       []interface{} `json:"links"`
		Memberships []struct {
			ContactDetails []interface{} `json:"contact_details"`
			HtmlURL        string        `json:"html_url"`
			ID             string        `json:"id"`
			Identifiers    []interface{} `json:"identifiers"`
			Images         []interface{} `json:"images"`
			Links          []interface{} `json:"links"`
			OrganizationID string        `json:"organization_id"`
			PersonID       string        `json:"person_id"`
			Role           string        `json:"role"`
			StartDate      string        `json:"start_date"`
			URL            string        `json:"url"`
		} `json:"memberships"`
		Name       string        `json:"name"`
		OtherNames []interface{} `json:"other_names"`
		ProxyImage string        `json:"proxy_image"`
		SortName   string        `json:"sort_name"`
		URL        string        `json:"url"`
	} `json:"result"`
}

type Organization struct {
	Result OrganizationResult `json:"result"`
}

type OrganizationResult struct {
	HTMLURL        string        `json:"html_url"`
	URL            string        `json:"url"`
	Category       string        `json:"category"`
	Classification string        `json:"classification"`
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Slug           string        `json:"slug"`
	Images         []interface{} `json:"images",omitempty`
	Posts          []interface{} `json:"posts",omitempty`
	Memberships    []struct {
		HTMLURL        string        `json:"html_url"`
		URL            string        `json:"url"`
		EndDate        string        `json:"end_date"`
		ID             string        `json:"id"`
		Identifiers    []interface{} `json:"identifiers"`
		OrganizationID string        `json:"organization_id"`
		PersonID       string        `json:"person_id"`
		Role           string        `json:"role"`
		StartDate      string        `json:"start_date"`
		Images         []interface{} `json:"images",omitempty`
		Links          []interface{} `json:"links",omitempty`
		ContactDetails []interface{} `json:"contact_details",omitempty`
	} `json:"memberships"`
	Links          []interface{} `json:"links, omitempty"`
	ContactDetails []interface{} `json:"contact_details,omitempty"`
	Identifiers    []interface{} `json:"identifiers,omitempty"`
	OtherNames     []interface{} `json:"other_names",omitempty`
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (a *Api) GetPerson(id string) (Person, error) {
	defer timeTrack(time.Now(), "GetPerson")

	result := new(Person)
	u := fmt.Sprintf("persons/core_person:%v", id)
	req, err := a.client.NewRequest("GET", u)
	if err != nil {
		return Person{}, err
	}

	_, err = a.client.Do(req, result)

	if err != nil {
		return Person{}, err
	}

	return *result, err

}

func (a *Api) GetOrganization(id string) (Organization, error) {
	defer timeTrack(time.Now(), "GetOrganization")

	result := new(Organization)
	u := fmt.Sprintf("organizations/core_organisation:%v", id)
	req, err := a.client.NewRequest("GET", u)

	if err != nil {
		return Organization{}, err
	}

	_, err = a.client.Do(req, result)

	if err != nil {
		return Organization{}, err
	}

	return *result, err

}
