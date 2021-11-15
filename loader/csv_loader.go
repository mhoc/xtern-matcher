package loader

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/mhoc/xtern-matcher/model"
)

type CSV struct {
	companiesFile string
	studentsFile  string
}

func NewCSV(args map[string]string) *CSV {
	companiesFile := args["companiesFile"]
	studentsFile := args["studentsFile"]
	return &CSV{
		companiesFile: companiesFile,
		studentsFile:  studentsFile,
	}
}

func (c CSV) Companies() (model.Companies, error) {
	cs := make([]*model.Company, 0)
	f, err := os.Open(c.companiesFile)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range content {
		c := &model.Company{}
		c.Name = row[0]
		numHiring, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, err
		}
		c.NumberHiring = numHiring
		c.Students = []string{}
		for _, studentName := range row[2:] {
			studentName = strings.TrimFunc(studentName, func(r rune) bool {
				return !unicode.IsGraphic(r)
			})
			if studentName == "Choose Your Candidate" ||
				studentName == "First Name, Last Name" ||
				studentName == "Last Name, First Name" ||
				studentName == "" ||
				studentName == "Choose Your Candidate " {
				continue
			}
			c.Students = append(c.Students, studentName)
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func (c CSV) Students() (model.Students, error) {
	ss := make([]*model.Student, 0)
	f, err := os.Open(c.studentsFile)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range content {
		s := &model.Student{}
		s.Name = row[0]
		s.Companies = []string{}
		for _, companyName := range row[1:] {
			companyName = strings.TrimFunc(companyName, func(r rune) bool {
				return !unicode.IsGraphic(r)
			})
			if companyName == "Choose a Company" ||
				companyName == "" ||
				companyName == "Choose Your Company/Department" {
				continue
			}
			s.Companies = append(s.Companies, companyName)
		}
		ss = append(ss, s)
	}
	return ss, nil
}
