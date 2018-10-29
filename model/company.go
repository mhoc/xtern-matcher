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
