package matcher

import (
	"testing"

	"github.com/mhoc/xtern-matcher/model"
)

// TestSimpleS1C1A1 tests a simple case where one company matches with one student and they both
// rank each other.
func TestSimpleS1C1A1(t *testing.T) {
	s1 := model.NewStudent("s1", []string{"c1"})
	students := model.Students([]*model.Student{s1})
	c1 := model.NewCompany("c1", 1, []string{"s1"})
	companies := model.Companies([]*model.Company{c1})
	expectedMatches := model.Matches([]*model.Match{
		{
			Company: c1,
			Student: s1,
		},
	})
	actualMatches := Simple(students, companies)
	if !expectedMatches.Equals(actualMatches) {
		t.Fatalf("expected matches to be equal")
	}
}

// TestSimpleS1C1A1NoStudentRank tests a case where a company ranks a given student but the student
// doesn't rank the company back; in this case we expect the match to still happen.
func TestSimpleS1C1A1NoStudentRank(t *testing.T) {
	s1 := model.NewStudent("s1", []string{})
	students := model.Students([]*model.Student{s1})
	c1 := model.NewCompany("c1", 1, []string{"s1"})
	companies := model.Companies([]*model.Company{c1})
	expectedMatches := model.Matches([]*model.Match{
		{
			Company: c1,
			Student: s1,
		},
	})
	actualMatches := Simple(students, companies)
	if !expectedMatches.Equals(actualMatches) {
		t.Fatalf("expected matches to be equal")
	}
}

// TestSimpleS1C1A1NoCompanyRank tests a case where a student ranks a company which doesn't rank
// them back; in this case we do not expect a match to happen.
func TestSimpleS1C1A1NoCompanyRank(t *testing.T) {
	s1 := model.NewStudent("s1", []string{"c1"})
	students := model.Students([]*model.Student{s1})
	c1 := model.NewCompany("c1", 1, []string{})
	companies := model.Companies([]*model.Company{c1})
	expectedMatches := model.Matches([]*model.Match{})
	actualMatches := Simple(students, companies)
	if !expectedMatches.Equals(actualMatches) {
		t.Fatalf("expected matches to be equal")
	}
}

// TestSimpleS2C2A1 is a basic 2x2 case to test introducing more variables into the matching
// process.
func TestSimpleS2C2A1(t *testing.T) {
	s1 := model.NewStudent("s1", []string{"c1"})
	s2 := model.NewStudent("s2", []string{"c2"})
	students := model.Students([]*model.Student{s1, s2})
	c1 := model.NewCompany("c1", 1, []string{"s1"})
	c2 := model.NewCompany("c2", 1, []string{"s2"})
	companies := model.Companies([]*model.Company{c1, c2})
	expectedMatches := model.Matches([]*model.Match{
		{
			Company: c1,
			Student: s1,
		},
		{
			Company: c2,
			Student: s2,
		},
	})
	actualMatches := Simple(students, companies)
	if !expectedMatches.Equals(actualMatches) {
		t.Fatalf("expected matches to be equal")
	}
}

// TestSimpleS2C2A1Tiebreaker extends the S2C2A1 test by modifying company preferences to want
// both students equally; in this case we expect the tie to be broken by the student preferences.
func TestSimpleS2C2A1Tiebreaker(t *testing.T) {
	s1 := model.NewStudent("s1", []string{"c1"})
	s2 := model.NewStudent("s2", []string{"c2"})
	students := model.Students([]*model.Student{s1, s2})
	c1 := model.NewCompany("c1", 1, []string{"s1", "s2"})
	c2 := model.NewCompany("c2", 1, []string{"s1", "s2"})
	companies := model.Companies([]*model.Company{c1, c2})
	expectedMatches := model.Matches([]*model.Match{
		{
			Company: c1,
			Student: s1,
		},
		{
			Company: c2,
			Student: s2,
		},
	})
	actualMatches := Simple(students, companies)
	if !expectedMatches.Equals(actualMatches) {
		t.Fatalf("expected matches to be equal")
	}
}

// TestSimpleS2C2A2Tiebreaker scales up the availability of each company to 2. In this test we have
// 4 available spots but only two students to fill them; we expect each company to get one student.
func TestSimpleS2C2A2Tiebreaker(t *testing.T) {
	s1 := model.NewStudent("s1", []string{"c1"})
	s2 := model.NewStudent("s2", []string{"c2"})
	students := model.Students([]*model.Student{s1, s2})
	c1 := model.NewCompany("c1", 2, []string{"s1", "s2"})
	c2 := model.NewCompany("c2", 2, []string{"s1", "s2"})
	companies := model.Companies([]*model.Company{c1, c2})
	expectedMatches := model.Matches([]*model.Match{
		{
			Company: c1,
			Student: s1,
		},
		{
			Company: c2,
			Student: s2,
		},
	})
	actualMatches := Simple(students, companies)
	if !expectedMatches.Equals(actualMatches) {
		t.Fatalf("expected matches to be equal")
	}
}

// TestSimpleS4C4A2Tiebreaker scales up the number of companies and students to 4 each, with one
// tiebreaker during R0.
func TestSimpleS4C4A2Tiebreaker(t *testing.T) {
	s1 := model.NewStudent("s1", []string{"c1"})
	s2 := model.NewStudent("s2", []string{"c4"})
	s3 := model.NewStudent("s3", []string{"c3"})
	s4 := model.NewStudent("s4", []string{"c4"})
	students := model.Students([]*model.Student{s1, s2, s3, s4})
	c1 := model.NewCompany("c1", 2, []string{"s1", "s2", "s3"})
	c2 := model.NewCompany("c2", 2, []string{"s1", "s2", "s4"})
	c3 := model.NewCompany("c3", 2, []string{"s2", "s3", "s1"})
	c4 := model.NewCompany("c4", 2, []string{"s2", "s4"})
	companies := model.Companies([]*model.Company{c1, c2, c3, c4})
	expectedMatches := model.Matches([]*model.Match{
		{
			Company: c1,
			Student: s1,
		},
		{
			Company: c4,
			Student: s2,
		},
		{
			Company: c3,
			Student: s3,
		},
		{
			Company: c4,
			Student: s4,
		},
	})
	actualMatches := Simple(students, companies)
	if !expectedMatches.Equals(actualMatches) {
		t.Fatalf("expected matches to be equal")
	}
}
