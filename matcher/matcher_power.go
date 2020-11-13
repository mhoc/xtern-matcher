// This system will work similar to simple.  It will assemble the students by rank
// and add inverse power relative to the current rank.  Rank 1 == 7 power.
// Companies use this power to "battle" for students.  The company with the
// highest power wins the battle with ties going to the student preference.
// The winning company then resets their power level to 0 while the losing
// company takes their power to the next battle. This repeats for all battles.
// The goal is to lead to happier matches and fewer "losses"

// Question - should power reset or subtract? Reset is easy and "fair"
// while subtraction carries "memory" of previously lost values.
// ex --> rank 1 == 7 power: loss. Rank 2 == 7+6 power: win. Rank 3 == 7+5: win

package matcher

import (
	"github.com/mhoc/xtern-matcher/model"
)

func Simple(students model.Students, companies model.Companies) model.Matches {
	var matches model.Matches

	// The core of the matching algorithm works by company rank; starting with rank 0, going until
	// rank n, finding as many matches at each rank as possible. 8 is just a magic number here to
	// represent the globally maximum number of ranks any company could have.
	for rank := 0; rank < 8; rank++ {

		// TODO: Add Power rank to companies by inverse (0-8)
		// ************************************


		// Assemble a list of students who were ranked at this rank by any company.
		studentsAtThisRank := make(map[string][]*model.Company)
		for _, company := range companies {
			if len(company.Students) > rank {
				studentName := company.Students[rank]
				if _, in := studentsAtThisRank[studentName]; in {
					studentsAtThisRank[studentName] = append(studentsAtThisRank[studentName], company)
				} else {
					studentsAtThisRank[studentName] = []*model.Company{company}
				}
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
					// TODO: Reset company power
					// *************************
				}



			} else if len(companiesAtRank) > 1 {
				// Here is where we need to get creative.
				// if multiple companies want the student we evaluate in this order.

				// TODO:
				// If company power is the same, give preference to student
				// then remove company power.



				// TODO:
				// else we need to give to company with the highest power.
				// When selected, reset winning company power to 0.
				// (((mike if you're ambitious try subtracting this rounds power value instead)))



				// TODO:
				// At this point, it is possible that multiple companies tied for power in this rank
				// but the student did not preference any of them.  Here we give preference to
				// the company with the least remaining matches (this could be hard)


			} else {
				// No companies want the student at this rank; continue on to the next rank.
				continue
			}
		}

	}

	return matches
}
