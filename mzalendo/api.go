package mzalendo

import "fmt"

type PersonService struct {
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
	Result struct {
		Category       string        `json:"category"`
		Classification string        `json:"classification"`
		ContactDetails []interface{} `json:"contact_details"`
		HtmlURL        string        `json:"html_url"`
		ID             string        `json:"id"`
		Identifiers    []interface{} `json:"identifiers"`
		Images         []interface{} `json:"images"`
		Links          []interface{} `json:"links"`
		Memberships    []struct {
			ContactDetails []interface{} `json:"contact_details"`
			EndDate        string        `json:"end_date"`
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
		Posts      []interface{} `json:"posts"`
		Slug       string        `json:"slug"`
		URL        string        `json:"url"`
	} `json:"result"`
}

func (p *PersonService) GetPerson(id string) (Person, error) {
	result := new(Person)
	u := fmt.Sprintf("persons/core_person:%v", id)
	req, err := p.client.NewRequest("GET", u)
	if err != nil {
		return Person{}, err
	}

	_, err = p.client.Do(req, result)

	if err != nil {
		return Person{}, err
	}

	return *result, err

}
