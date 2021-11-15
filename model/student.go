package model

type Student struct {
	Name      string
	Companies []string
}

func NewStudent(name string, companies []string) *Student {
	return &Student{
		Name:      name,
		Companies: companies,
	}
}

func (s Student) Equals(s2 Student) bool {
	if s.Name != s2.Name {
		return false
	}
	if len(s.Companies) != len(s2.Companies) {
		return false
	}
	for i := range s.Companies {
		if s.Companies[i] != s2.Companies[i] {
			return false
		}
	}
	return true
}
