package response

import (
	"moviedle/src/core/domain/oscar"

	"github.com/google/uuid"
)

type Oscar struct {
	ID 	 *uuid.UUID `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type oscarBuilder struct{}

func NewOscarBuilder() *oscarBuilder {
    return &oscarBuilder{}
}

func (*oscarBuilder) BuildFromDomain(data oscar.Oscar) Oscar {
	return Oscar {
		data.ID(),
        data.Name(),
        data.Code(),
	}
}

func (*oscarBuilder) BuildFromDomainList(data []oscar.Oscar) []Oscar {
	var oscars []Oscar
    for _, oscarData := range data {
        oscars = append(oscars, NewOscarBuilder().BuildFromDomain(oscarData))
    }
    return oscars
}