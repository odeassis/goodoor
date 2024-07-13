package handler

import "fmt"

func fieldfValidate(field, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", field, typ)
}

type CreateOpeningRequest struct {
	Title    string `json:"title"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {

	if c.Title == "" {
		return fieldfValidate("title", "string")
	}

	if c.Role == "" {
		return fieldfValidate("role", "string")
	}

	if c.Company == "" {
		return fieldfValidate("company", "string")
	}

	if c.Location == "" {
		return fieldfValidate("location", "string")
	}

	if c.Remote == nil {
		return fieldfValidate("remote", "bool")
	}

	if c.Salary <= 0 {
		return fieldfValidate("salary", "int64")
	}

	if c.Link == "" {
		return fieldfValidate("link", "string")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Title    string `json:"title"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (u *UpdateOpeningRequest) Validate() error {
	if u.Title != "" || u.Role != "" || u.Company != "" || u.Location != "" || u.Remote != nil || u.Link != "" || u.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at last one valid field must be provider")
}
