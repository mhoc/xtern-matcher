package matcher

import (
	"github.com/mhoc/xtern-matcher/model"
)

func Simple(students model.Students, companies model.Companies) model.Matches {
	var matches model.Matches

	// The core of the matching algorithm works by company rank; starting with rank 0, going until
	// rank n, finding as many matches at each rank as possible. 12 is just a magic number here to
	// represent the globally maximum number of ranks any company could have.
	for rank := 0; rank < 12; rank++ {

		// Assemble a list of students who were ranked at this rank by any company.
		studentsAtThisRank := make(map[string][]*model.Company)
		for _, company := range companies {
			if len(company.Students) > rank {
				studentName := company.Students[rank]
				studentsAtThisRank[studentName] = append(studentsAtThisRank[studentName], company)
			}
		}

		// Iterate over every student.
		for studentName, companiesAtRank := range studentsAtThisRank {
			student := students.Find(studentName)

			if len(companiesAtRank) == 1 {
				// If only one company ranked this student at this rank, we assign the match.
				company := companiesAtRank[0]
				if matches.CompanyCanSupportMatch(company) && matches.FindByStudent(student) == nil {
					matches = matches.Add(student, company)
				}
			} else if len(companiesAtRank) > 1 {

				// If multiple companies want the student at this level, we resolve the match by using the
				// student's preferences
				for _, studentRankCompanyName := range student.Companies {
					for _, companyAtRank := range companiesAtRank {
						if studentRankCompanyName == companyAtRank.Name &&
							matches.CompanyCanSupportMatch(companyAtRank) && matches.FindByStudent(student) == nil {
							matches = matches.Add(student, companyAtRank)
						}
					}
				}

				// At this point, it is possible that multiple companies ranked this student but the student
				// didn't rank any of them; in this case, we just give them the first company that has
				// availability.
				for _, companyAtRank := range companiesAtRank {
					if matches.CompanyCanSupportMatch(companyAtRank) && matches.FindByStudent(student) == nil {
						matches = matches.Add(student, companyAtRank)
					}
				}

			} else {
				// No companies want the student at this rank; continue on to the next rank.
				continue
			}
		}

	}

	return matches
}
