package model

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

/**
 * GlobalFitness calculates the "accuracy" of the set of matches. We consider an ideal global match
 * to mean that every company gets their first N picks, where N is the number of students they want.
 * In reality, this is impossible to reach because of competition between companies, but we'd want
 * to optimize it regardlesss.
 *
 * Lets say we have a ranking like C1(1): S1 S2 S3 S4
 * The global fitnesses would look like:
 *   S1:   1.00
 *   S2:   0.75
 *   S3:   0.50
 *   S4:   0.25
 *   none: 0.00
 *
 * Each match is calculated independently.
 */
func (m Matches) GlobalFitness() float64 {
	return 1
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
