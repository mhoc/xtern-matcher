package model

import "testing"

func TestStudentEquals1(t *testing.T) {
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s2", []string{"c2"})
	if s1.Equals(*s2) {
		t.Fatalf("expected s1 != s2")
	}
}

func TestStudentEquals2(t *testing.T) {
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s1", []string{"c2"})
	if s1.Equals(*s2) {
		t.Fatalf("expected s1 != s2")
	}
}

func TestStudentEquals3(t *testing.T) {
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s1", []string{"c1"})
	if !s1.Equals(*s2) {
		t.Fatalf("expected s1 == s2")
	}
}

func TestStudentEquals4(t *testing.T) {
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s1", []string{"c1", "c1"})
	if s1.Equals(*s2) {
		t.Fatalf("expected s1 != s2")
	}
}

func TestStudentFindExists(t *testing.T) {
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s2", []string{"c2"})
	s3 := NewStudent("s3", []string{"c3"})
	students := Students([]*Student{s1, s2, s3})
	s1f, err := students.Find("s1")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !s1.Equals(*s1f) {
		t.Fatalf("expected find to return s1")
	}
	s2f, err := students.Find("s2")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !s2.Equals(*s2f) {
		t.Fatalf("expected find to return s2")
	}
	s3f, err := students.Find("s3")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !s3.Equals(*s3f) {
		t.Fatalf("expected find to return s3")
	}
}

func TestStudentFindNoExists(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	s1 := NewStudent("s1", []string{"c1"})
	s2 := NewStudent("s2", []string{"c2"})
	s3 := NewStudent("s3", []string{"c3"})
	students := Students([]*Student{s1, s2, s3})
	students.Find("s4")
	t.Fatalf("expected find to fatally panic on s4")
}
