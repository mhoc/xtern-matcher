package model

import "testing"

func TestCompanyEquals1(t *testing.T) {
	c1 := NewCompany("c1", 1, []string{"s1"})
	c2 := NewCompany("c1", 1, []string{"s1"})
	if !c1.Equals(*c2) {
		t.Fatalf("expected c1 == c2")
	}
}

func TestCompanyEquals2(t *testing.T) {
	c1 := NewCompany("c1", 1, []string{"s1"})
	c2 := NewCompany("c1", 2, []string{"s1"})
	if c1.Equals(*c2) {
		t.Fatalf("expected c1 != c2")
	}
}

func TestCompanyEquals3(t *testing.T) {
	c1 := NewCompany("c1", 1, []string{"s1"})
	c2 := NewCompany("c1", 1, []string{"s1", "s2"})
	if c1.Equals(*c2) {
		t.Fatalf("expected c1 != c2")
	}
}

func TestCompanyEquals4(t *testing.T) {
	c1 := NewCompany("c1", 1, []string{"s1"})
	c2 := NewCompany("c2", 1, []string{"s1"})
	if c1.Equals(*c2) {
		t.Fatalf("expected c1 != c2")
	}
}

func TestCompanyEquals5(t *testing.T) {
	c1 := NewCompany("c1", 1, []string{"s1", "s2"})
	c2 := NewCompany("c1", 1, []string{"s1", "s3"})
	if c1.Equals(*c2) {
		t.Fatalf("expected c1 != c2")
	}
}

func TestCompaniesFindExists(t *testing.T) {
	c1 := NewCompany("c1", 1, []string{"s1"})
	c2 := NewCompany("c2", 1, []string{"s4"})
	c3 := NewCompany("c3", 1, []string{"s5"})
	companies := Companies([]*Company{c1, c2, c3})
	c1f := companies.Find("c1")
	if c1f == nil || !c1.Equals(*c1f) {
		t.Fatalf("expected to find c1")
	}
	c2f := companies.Find("c2")
	if c2f == nil || !c2.Equals(*c2f) {
		t.Fatalf("expected to find c1")
	}
	c3f := companies.Find("c3")
	if c3f == nil || !c3.Equals(*c3f) {
		t.Fatalf("expected to find c1")
	}
}

func TestCompaniesFindNoExists(t *testing.T) {
	defer func() {
		recover()
	}()
	c1 := NewCompany("c1", 1, []string{"s1"})
	c2 := NewCompany("c2", 1, []string{"s4"})
	c3 := NewCompany("c3", 1, []string{"s5"})
	companies := Companies([]*Company{c1, c2, c3})
	companies.Find("c4")
	t.Fatalf("expected to fatally panic on find c4")
}
