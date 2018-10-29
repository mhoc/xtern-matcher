package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/mhoc/xtern-matcher/matcher"
	"github.com/mhoc/xtern-matcher/model"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	alg := flag.String("alg", "simple", "matching algorithm: simple, brute, or evolve")
	inCompanies := flag.String("in-companies", "", "input csv file with company data")
	inStudents := flag.String("in-students", "", "input csv file with student data")
	output := flag.String("out", "pretty", "output format: pretty or csv")
	pivot := flag.String("pivot", "companies", "output pivot: companies or students")
	flag.Parse()

	companies, err := model.LoadCompanies(*inCompanies)
	if err != nil {
		panic(err)
	}
	students, err := model.LoadStudents(*inStudents)
	if err != nil {
		panic(err)
	}
	err = students.Validate()
	if err != nil {
		panic(err)
	}

	var matches model.Matches
	switch *alg {
	case "simple":
		matches = matcher.Simple(students, companies)
	case "brute":
		matches = matcher.Brute(students, companies)
	case "evolve":
		matches = matcher.Evolve(students, companies)
	default:
		panic("alg not recognized")
	}

	if *output == "pretty" {
		if *pivot == "students" {
			matches.WriteByStudent()
		} else if *pivot == "companies" {
			matches.WriteByCompany()
		} else {
			panic("pivot not recognized")
		}
	} else if *output == "csv" {
		panic("csv output not implemented yet")
	} else {
		panic("output format not recognized")
	}
}
