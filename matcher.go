package main

func FindMatches(students Students, companies Companies) Matches {
	var matches Matches
	for _, student := range students {

		// Iterate over each "placement" a company could have ranked a student, starting at 0
		// 12 is just a magic number here; it represents the maximum ranking a company could have
		// placed a student at.
		for i := 0; i < 12; i++ {
			companiesRankedStudentAtI := []*Company{}
			for _, company := range companies {
				if len(company.Students) > i && company.Students[i] == student.Name {
					companiesRankedStudentAtI = append(companiesRankedStudentAtI, company)
				}
			}

			if len(companiesRankedStudentAtI) == 1 {

				// If only one company ranked the student at #i, and that company has room, assign the match.
				matches = matches.Add(student, companiesRankedStudentAtI[0])

			} else if len(companiesRankedStudentAtI) > 1 {

				// If multiple companies ranked the student at #i, use the student's rankings to break the
				// tie.
				for _, studentCompanyRanking := range student.Companies {
					for _, eligibleCompany := range companiesRankedStudentAtI {
						if studentCompanyRanking == eligibleCompany.Name && matches.FindByStudent(student) == nil {
							matches = matches.Add(student, eligibleCompany)
						}
					}
				}

				// At this point, we can assume that the student hasn't even ranked the company. Too bad,
				// you get matched anyway. Instead of using some consistent way to decide this, we just
				// select the first company in the CSV. Hopefully this happens rarely.
				if matches.FindByStudent(student) == nil {
					matches = matches.Add(student, companiesRankedStudentAtI[0])
				}

			} else {

				// If no one ranked the student at this rank, we skip this rank and move on to the next
				// one
				continue

			}
		}
	}
	return matches
}
