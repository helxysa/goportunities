package handler

import "fmt"

// Create Opening
func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("O parâmetro %s é obrigatório e deve ser do tipo %s", name, typ)
}


type CreateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("request body is empty")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if !r.Remote && r.Remote {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

