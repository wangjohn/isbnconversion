package isbnconversion

import (
	"testing"
)

func TestISBN10(t *testing.T) {
	dataset := []struct {
		ISBN10  string
		IsValid bool
	}{
		{"080442957X", true},
		{"0851310419", true},
		{"0943396042", true},
		{"097522980X", true},
		{"1843560283", true},
		{"0684843285", true},
		{"0684843284", false},
		{"184356028X", false},
		{"0804429578", false},
		{"0804429", false},
		{"0804429578234", false},
	}

	for _, data := range dataset {
		reported := IsISBN10(data.ISBN10)
		if reported != data.IsValid {
			t.Fatalf("IsISBN10 did not match expected value reported=%v expected=%v", reported, data.IsValid)
		}
	}
}

func TestISBN13(t *testing.T) {
	dataset := []struct {
		ISBN13  string
		IsValid bool
	}{
		{"9781566199094", true},
		{"9781402894626", true},
		{"9780804429573", true},
		{"9780851310411", true},
		{"9780943396040", true},
		{"9780975229804", true},
		{"9781843560289", true},
		{"9780684843285", true},
		{"9780684843284", false},
		{"9780684843281", false},
		{"9780943396049", false},
		{"080442957X", false},
		{"0851310419", false},
		{"0943396042", false},
	}
	for _, data := range dataset {
		reported := IsISBN13(data.ISBN13)
		if reported != data.IsValid {
			t.Fatalf("IsISBN13 did not match expected value reported=%v expected=%v", reported, data.IsValid)
		}
	}
}

func TestISBNConversion(t *testing.T) {
	dataset := []struct {
		ISBN10 string
		ISBN13 string
	}{
		{"080442957X", "9780804429573"},
		{"0851310419", "9780851310411"},
		{"0943396042", "9780943396040"},
		{"097522980X", "9780975229804"},
		{"1843560283", "9781843560289"},
		{"0684843285", "9780684843285"},
	}
	for _, data := range dataset {
		reported13, err := ISBN10to13(data.ISBN10)
		if err != nil {
			t.Fatalf("ISBN10to13 Error=%v", err)
		}
		if reported13 != data.ISBN13 {
			t.Fatalf("ISBN10to13 Incorrect conversion reported=%v expected=%v", reported13, data.ISBN13)
		}
		reported10, err := ISBN13to10(data.ISBN13)
		if err != nil {
			t.Fatalf("ISBN13to10 Error=%v", err)
		}
		if reported10 != data.ISBN10 {
			t.Fatalf("ISBN13to10 Incorrect conversion reported=%v expected=%v", reported10, data.ISBN10)
		}
	}
}
