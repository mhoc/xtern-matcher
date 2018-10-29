package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Company struct {
	Name         string
	NumberHiring int
	Students     []string
}

type Companies []*Company

func LoadCompanies(filename string) Companies {
	cs := make([]*Company, 0)
	f, err := os.Open(filename)
	check(err)
	reader := csv.NewReader(f)
	content, err := reader.ReadAll()
	check(err)
	for _, row := range content {
		c := &Company{}
		c.Name = row[0]
		numHiring, err := strconv.Atoi(row[1])
		check(err)
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
	return cs
}

func (c Companies) Find(companyName string) *Company {
	for _, company := range c {
		if company.Name == companyName {
			return company
		}
	}
	panic(fmt.Sprintf("no company found with name: %v", companyName))
}
