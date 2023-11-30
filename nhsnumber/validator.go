package nhsnumber

import (
	"errors"
	"strconv"
)

type validator struct {
}

func NewValidator() *validator {
	return &validator{}
}

var (
	ErrInvalidInput     = errors.New("NHS number is not 10 digits long")
	ErrNotANumber       = errors.New("NHS number is not a number")
	ErrCheckDigitIsTen  = errors.New("Check digit is 10")
	ErrInvalidNHSNumber = errors.New("NHS number is invalid")
)

func (v *validator) Validate(nhsNumber string) error {
	if len(nhsNumber) != 10 {
		return ErrInvalidInput
	}

	if _, err := strconv.Atoi(nhsNumber); err != nil {
		return ErrNotANumber
	}

	sum := 0
	for i, c := range nhsNumber[:9] {
		sum += int(c-'0') * (10 - i) // we subtract the ascii value of '0' to get the actual number
	}

	remainder := sum % 11
	checkDigit := 11 - remainder
	if checkDigit == 11 {
		checkDigit = 0
	} else if checkDigit == 10 {
		return ErrCheckDigitIsTen
	}

	lastDigit := int(nhsNumber[len(nhsNumber)-1] - '0')
	if checkDigit != lastDigit {
		return ErrInvalidNHSNumber
	}

	return nil
}
