package mzalendo

import "fmt"

type PersonService struct {
	client *Client
}

type Person struct {
	Result Result
}

type Result struct {
	Name        string       `json:"name"`
	Gender      string       `json:"gender"`
	BirthDate   string       `json:"birth_date"`
	Memberships []Membership `json:"memberships"`
}

type Membership struct {
	Role           string `json:"role"`
	OrganizationId string `json:"organization_id"`
}

type Organization struct {
	Result OrganizationResult
}

type OrganizationResult struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Id       string `json:"id"`
}

func (p *PersonService) GetPerson(id string) (Person, error) {
	result := new(Person)
	u := fmt.Sprintf("/persons/core_person:%v", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return Person{}, err
	}

	_, err = p.client.Do(req, result)

	if err != nil {
		return Person{}, err
	}

	return *result, err

}
