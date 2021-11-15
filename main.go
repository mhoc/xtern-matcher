package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/mhoc/xtern-matcher/loader"
	"github.com/mhoc/xtern-matcher/matcher"
	"github.com/mhoc/xtern-matcher/model"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	alg := flag.String("alg", "simple", "matching algorithm: simple, brute, or evolve")
	inCompanies := flag.String("in-companies", "", "input csv file with company data")
	inStudents := flag.String("in-students", "", "input csv file with student data")
	output := flag.String("out", "pretty", "output format: pretty or csv")
	pivot := flag.String("pivot", "companies", "output pivot: companies or students")
	version := flag.Bool("v", false, "output binary version")
	flag.Parse()

	if *version {
		fmt.Printf("v1.3.0\n")
		return
	}

	loader := loader.Get("csv", map[string]string{
		"companiesFile": *inCompanies,
		"studentsFile":  *inStudents,
	})

	companies, err := loader.Companies()
	if err != nil {
		panic(err)
	}
	students, err := loader.Students()
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
		if *pivot == "students" {
			matches.WriteCSVByStudent()
		} else if *pivot == "companies" {
			matches.WriteCSVByCompany()
		} else {
			panic("pivot not recognized")
		}
	} else {
		panic("output format not recognized")
	}
}
