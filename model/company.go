package model

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

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

type Companies []*Company

func LoadCompanies(filename string) (Companies, error) {
	cs := make([]*Company, 0)
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
		c := &Company{}
		c.Name = row[0]
		numHiring, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, err
		}
		c.NumberHiring = numHiring
		c.Students = []string{}
		for _, studentName := range row[2:] {
			if studentName == "Choose Your Candidate" || studentName == "" {
				continue
			}
			c.Students = append(c.Students, studentName)
		}
		cs = append(cs, c)
	}
	return cs, nil
}

func (c Companies) Find(companyName string) *Company {
	for _, company := range c {
		if company.Name == companyName {
			return company
		}
	}
	panic(fmt.Sprintf("no company found with name: %v", companyName))
}

func (c Companies) Random() *Company {
	return c[rand.Intn(len(c))]
}
