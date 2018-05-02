package model

import (
	"github.com/acassio/base-app-go/enum"
)

//UserResponse representa os dados da resposta
type UserResponse struct {
	Success   bool   `json:"success"`
	Data      User   `json:"data,omitempty"`
	ErrorCode int    `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
}

//User representa os dados do usuario
type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

//FillStruct - preenche os campos da struct
func (s *UserResponse) FillStruct(m map[string]interface{}) error {

	s.Success = true

	for k, v := range m {

		switch k {
		case "message":
			if v != nil {
				s.Data = v.(User)
				s.ErrorCode = enum.ERRORS[s.Message]
				s.Success = false
			}

		}

	}
	return nil
}
