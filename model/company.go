package model

type Company struct {
	Name         string
	NumberHiring int
	Students     []string
}

func NewCompany(name string, numberHiring int, students []string) *Company {
	return &Company{
		Name:         name,
		NumberHiring: numberHiring,
		Students:     students,
	}
}

func (c Company) Equals(c2 Company) bool {
	if c.Name != c2.Name || c.NumberHiring != c2.NumberHiring || len(c.Students) != len(c2.Students) {
		return false
	}
	for i := range c.Students {
		if c.Students[i] != c2.Students[i] {
			return false
		}
	}
	return true
}
