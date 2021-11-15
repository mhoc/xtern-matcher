package model

import (
	"fmt"
	"math/rand"
)

type Companies []*Company

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
