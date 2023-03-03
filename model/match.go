package model

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
)

type Match struct {
	Company *Company
	Student *Student
}

func NewMatch(c *Company, s *Student) *Match {
	return &Match{
		Company: c,
		Student: s,
	}
}

func (m Match) Equals(m2 Match) bool {
	return m.Company.Equals(*m2.Company) && m.Student.Equals(*m2.Student)
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

// CompanyCanSupportMatch returns true if a company has room to match
func (m Matches) CompanyCanSupportMatch(company *Company) bool {
	nMatches := 0
	for _, match := range m {
		if match.Company.Name == company.Name {
			nMatches++
		}
	}
	return nMatches < company.NumberHiring
}

// Equals returns true of two sets of matches are exactly equal to one-another.
func (m Matches) Equals(m2 Matches) bool {
	if len(m) != len(m2) {
		return false
	}
	for _, match := range m {
		match2 := m2.FindByStudent(match.Student)
		if match2 == nil {
			return false
		}
		if !match.Equals(*match2) {
			return false
		}
	}
	return true
}

// FindByStudent finds the student's match in the list of matches, or nil if they haven't been
// matched yet
func (m Matches) FindByStudent(student *Student) *Match {
	for _, match := range m {
		if match.Student.Name == student.Name {
			return match
		}
	}
	return nil
}

// FreeAgents filters the full student list provided by students who are not represented in this
// list of matches.
func (m Matches) FreeAgents(fullStudentList Students) Students {
	var unmatched Students
	for _, student := range fullStudentList {
		matched := false
		for _, match := range m {
			if match.Student.Equals(*student) {
				matched = true
			}
		}
		if !matched {
			unmatched = append(unmatched, student)
		}
	}
	return unmatched
}

// StudentHasMatch returns true if the given student has a match already assigned
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
		byCompany[match.Company.Name] = append(byCompany[match.Company.Name], match)
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

func (m Matches) WriteFreeAgents(fullStudentList Students) {
	fmt.Printf("free agents\n")
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	freeAgents := m.FreeAgents(fullStudentList)
	for _, freeAgent := range freeAgents {
		fmt.Fprintf(w, "%v\n", freeAgent.Name)
	}
	w.Flush()
}

func (m Matches) WriteCSVByCompany() {
	byCompany := make(map[string][]*Match)
	for _, match := range m {
		byCompany[match.Company.Name] = append(byCompany[match.Company.Name], match)
	}
	w := csv.NewWriter(os.Stdout)
	for companyName, matches := range byCompany {
		line := []string{companyName}
		for _, match := range matches {
			line = append(line, match.Student.Name)
		}
		_ = w.Write(line)
	}
	w.Flush()
}

func (m Matches) WriteCSVByStudent() {
	w := csv.NewWriter(os.Stdout)
	for _, match := range m {
		_ = w.Write([]string{match.Student.Name, match.Company.Name})
	}
	w.Flush()
}

func (m Matches) WriteCSVFreeAgents(fullStudentList Students) {
	w := csv.NewWriter(os.Stdout)
	freeAgents := m.FreeAgents(fullStudentList)
	for _, freeAgent := range freeAgents {
		_ = w.Write([]string{freeAgent.Name})
	}
	w.Flush()
}
