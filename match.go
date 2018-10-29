package main

import (
	"fmt"
	"os"
	"text/tabwriter"
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

func (m Matches) WriteByCompany() {
	fmt.Printf("by company\n")
	byCompany := make(map[string][]*Match)
	for _, match := range m {
		if _, in := byCompany[match.Company.Name]; in {
			byCompany[match.Company.Name] = append(byCompany[match.Company.Name], match)
		} else {
			byCompany[match.Company.Name] = []*Match{match}
		}
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.DiscardEmptyColumns)
	for companyName, matches := range byCompany {
		str := fmt.Sprintf("%v \t%v ", companyName, matches[0].Company.NumberHiring)
		for _, match := range matches {
			str += fmt.Sprintf("\t%v ", match.Student.Name)
		}
		str += "\n"
		fmt.Fprint(w, str)
	}
	w.Flush()
}

func (m Matches) WriteByStudent() {
	fmt.Printf("by student\n")
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for _, match := range m {
		fmt.Fprintf(w, "%v\t%v\n", match.Student.Name, match.Company.Name)
	}
	w.Flush()
}
