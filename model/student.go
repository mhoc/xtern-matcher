package model

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Student struct {
	Name      string
	Companies []string
}

type Students []*Student

func LoadStudents(filename string) (Students, error) {
	ss := make([]*Student, 0)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range content {
		s := &Student{}
		s.Name = row[0]
		s.Companies = []string{}
		for _, companyName := range row[1:] {
			if companyName == "Choose a Company" || companyName == "" {
				continue
			}
			s.Companies = append(s.Companies, companyName)
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func (s Students) Find(name string) *Student {
	for _, student := range s {
		if student.Name == name {
			return student
		}
	}
	panic(fmt.Sprintf("no student found with name %v", name))
}

func (s Students) Validate() error {
	studentMap := make(map[string]*Student)
	for _, student := range s {
		if _, in := studentMap[student.Name]; in {
			return fmt.Errorf("student appears twice in csv input: %v", student.Name)
		}
	}
	return nil
}
