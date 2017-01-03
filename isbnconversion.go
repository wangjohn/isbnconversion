package isbnconversion

import (
	"fmt"
	"strconv"
	"strings"
)

func isbn10CheckDigit(digits []int) (string, error) {
	if len(digits) != 9 {
		return "", fmt.Errorf("Invalid length for isbn10 digits length=%v expected=%v", len(digits), 9)
	}
	sum := 0
	for i, digit := range digits {
		sum += (i + 1) * digit
	}
	modSum := sum % 11
	if modSum == 10 {
		return "X", nil
	} else {
		return strconv.Itoa(modSum), nil
	}
}

func isbn13CheckDigit(digits []int) (string, error) {
	if len(digits) != 12 {
		return "", fmt.Errorf("Invalid length for isbn13 digits length=%v expected=%v", len(digits), 12)
	}
	sum := 0
	for i, digit := range digits {
		if i%2 == 0 {
			sum += digit
		} else {
			sum += 3 * digit
		}
	}
	r := 10 - (sum % 10)
	if r == 10 {
		return "0", nil
	} else {
		return strconv.Itoa(r), nil
	}
}

func IsISBN13(isbn string) bool {
	return checkISBN(isbn, 13, isbn13CheckDigit)
}

func IsISBN10(isbn string) bool {
	return checkISBN(isbn, 10, isbn10CheckDigit)
}

func checkISBN(isbn string, size int, checkFunc func([]int) (string, error)) bool {
	if len(isbn) != size {
		return false
	}
	digits := make([]int, size-1)
	for i, r := range isbn {
		if i < (size - 1) {
			char := string(r)
			val, err := strconv.Atoi(char)
			if err != nil {
				return false
			}
			digits[i] = val
		} else {
			char := string(r)
			checkDigit, err := checkFunc(digits)
			if err != nil {
				return false
			}
			if char != checkDigit {
				return false
			}
		}
	}
	return true
}

func ISBN10to13(isbn10 string) (string, error) {
	if !IsISBN10(isbn10) {
		return "", fmt.Errorf("Invalid ISBN10 input: %v", isbn10)
	}
	baseDigits := make([]int, 9)
	for i := 0; i < 9; i++ {
		val, err := strconv.Atoi(string(isbn10[i]))
		if err != nil {
			return "", err
		}
		baseDigits[i] = val
	}
	intermedDigits := append([]int{9, 7, 8}, baseDigits...)
	checkDigit, err := isbn13CheckDigit(intermedDigits)
	if err != nil {
		return "", err
	}
	finalStr := make([]string, 13)
	for i, digit := range intermedDigits {
		finalStr[i] = strconv.Itoa(digit)
	}
	finalStr[12] = checkDigit
	return strings.Join(finalStr, ""), nil
}

func ISBN13to10(isbn13 string) (string, error) {
	if !IsISBN13(isbn13) {
		return "", fmt.Errorf("Invalid ISBN13 input: %v", isbn13)
	}
	baseDigits := make([]int, 9)
	for i := 3; i < 12; i++ {
		val, err := strconv.Atoi(string(isbn13[i]))
		if err != nil {
			return "", err
		}
		baseDigits[i-3] = val
	}
	checkDigit, err := isbn10CheckDigit(baseDigits)
	if err != nil {
		return "", err
	}
	finalStr := make([]string, 10)
	for i, digit := range baseDigits {
		finalStr[i] = strconv.Itoa(digit)
	}
	finalStr[9] = checkDigit
	return strings.Join(finalStr, ""), nil
}
