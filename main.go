package main

import "flag"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	output := flag.String("out", "pretty", "output format: pretty or csv")
	pivot := flag.String("pivot", "companies", "output pivot: companies or students")
	flag.Parse()

	companies := LoadCompanies("companies.csv")
	students := LoadStudents("students.csv")
	err := students.Validate()
	if err != nil {
		panic(err)
	}

	matches := FindMatches(students, companies)

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
