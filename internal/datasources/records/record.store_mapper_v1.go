package records

import V1Domains "github.com/snykk/go-rest-boilerplate/internal/business/domains/v1"

func (s *Stores) ToV1Domain() V1Domains.StoreDomain {
	return V1Domains.StoreDomain{
		Id:       s.Id,
		Name:     s.Name,
		Address:  s.Address,
		Contacts: s.Contacts,
	}
}

func FromStoresV1Domain(s *V1Domains.StoreDomain) Stores {
	return Stores{
		Id:       s.Id,
		Name:     s.Name,
		Address:  s.Address,
		Contacts: s.Contacts,
	}
}

func ToArrayOfStoresV1Domain(s *[]Stores) []V1Domains.StoreDomain {
	var result []V1Domains.StoreDomain

	for _, val := range *s {
		result = append(result, val.ToV1Domain())
	}

	return result
}
