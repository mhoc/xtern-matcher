package main

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	companies := LoadCompanies("companies.csv")
	students := LoadStudents("students.csv")
	err := students.Validate()
	if err != nil {
		panic(err)
	}
	FindMatches(students, companies)
}
