package model

import "testing"

func TestMatchesAddSuccessful(t *testing.T) {
	c1 := NewCompany("c1", 2, []string{"s1"})
	s1 := NewStudent("s1", []string{"c1"})
	matches := Matches([]*Match{
		{
			Company: c1,
			Student: s1,
		},
	})
	s2 := NewStudent("s2", []string{"c1"})
	matches = matches.Add(s2, c1)
}

func TestMatchesAddStudentTaken(t *testing.T) {
	defer func() {
		recover()
	}()
	c1 := NewCompany("c1", 2, []string{"s1"})
	c2 := NewCompany("c2", 2, []string{"s1"})
	s1 := NewStudent("s1", []string{"c1"})
	matches := Matches([]*Match{
		{
			Company: c1,
			Student: s1,
		},
	})
	matches = matches.Add(s1, c2)
	t.Fatalf("expected to fatally panic on s1-c2 match")
}

func TestMatchesAddStudentCompanyFull(t *testing.T) {
	defer func() {
		recover()
	}()
	c1 := NewCompany("c1", 1, []string{"s1", "s2"})
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s2", []string{"c1"})
	matches := Matches([]*Match{
		{
			Company: c1,
			Student: s1,
		},
	})
	matches = matches.Add(s2, c1)
	t.Fatalf("expected to fatally panic on s2-c1 match")
}

func TestMatchesFindByStudentExists(t *testing.T) {
	c1 := NewCompany("c1", 2, []string{"s1", "s2"})
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s2", []string{"c1"})
	matches := Matches([]*Match{
		{
			Company: c1,
			Student: s1,
		},
		{
			Company: c1,
			Student: s2,
		},
	})
	ms1f := matches.FindByStudent(s1)
	if !s1.Equals(*ms1f.Student) {
		t.Fatalf("expected to find s1 in matches")
	}
	ms2f := matches.FindByStudent(s2)
	if !s2.Equals(*ms2f.Student) {
		t.Fatalf("expected to find s2 in matches")
	}
}

func TestMatchesFindByStudentNoExist(t *testing.T) {
	c1 := NewCompany("c1", 2, []string{"s1", "s2"})
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s2", []string{"c1"})
	matches := Matches([]*Match{
		{
			Company: c1,
			Student: s1,
		},
	})
	ms2f := matches.FindByStudent(s2)
	if ms2f != nil {
		t.Fatalf("expected to find nil when searching for s2 in matches")
	}
}
