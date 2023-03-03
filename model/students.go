package model

import "fmt"

type Students []*Student

func (s Students) Find(name string) (*Student, error) {
	for _, student := range s {
		if student.Name == name {
			return student, nil
		}
	}
	return nil, fmt.Errorf("no student found with name '%v'", name)
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
