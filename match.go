package main

import (
	"fmt"
)

type Match struct {
	Company *Company
	Student *Student
}

type Matches []*Match

func (m Matches) Add(student *Student, company *Company) Matches {
	if m.StudentHasMatch(student) {
		panic(fmt.Sprintf("attempted to match already matched student: %v -> %v", student.Name, company.Name))
	}
	if !m.CompanyCanSupportMatch(company) {
		panic(fmt.Sprintf("attempted to match company beyond matching limit: %v (%v) <- %v", company.NumberHiring, company.NumberHiring, student.Name))
	}
	fmt.Printf("%v -> %v (%v)\n", student.Name, company.Name, company.NumberHiring)
	return append(m, &Match{
		Company: company,
		Student: student,
	})
}

/** CompanyCanSupportMatch returns true if a company has room to match */
func (m Matches) CompanyCanSupportMatch(company *Company) bool {
	nMatches := 0
	for _, match := range m {
		if match.Company.Name == company.Name {
			nMatches++
		}
	}
	return nMatches < company.NumberHiring
}

/** Finds the student's match in the list of matches, or nil if they haven't been matched yet */
func (m Matches) FindByStudent(student *Student) *Match {
	for _, match := range m {
		if match.Student.Name == student.Name {
			return match
		}
	}
	return nil
}

func (m Matches) StudentHasMatch(student *Student) bool {
	for _, match := range m {
		if match.Student.Name == student.Name {
			return true
		}
	}
	return false
}
