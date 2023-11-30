package nhsnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_Validate(t *testing.T) {

	t.Run("Test validator with valid NHS number", func(t *testing.T) {
		v := NewValidator()

		validNHSNumber := "5990128088"
		err := v.Validate(validNHSNumber)
		assert.Nil(t, err)
	})

	t.Run("Test validator with invalid NHS number", func(t *testing.T) {
		v := NewValidator()

		invalidNHSNumber := "5990128087"
		err := v.Validate(invalidNHSNumber)
		assert.EqualError(t, err, ErrInvalidNHSNumber.Error())
	})

	t.Run("Test validator with NHS number that is not 10 digits long", func(t *testing.T) {
		v := NewValidator()

		invalidNHSNumber := "59901280"
		err := v.Validate(invalidNHSNumber)
		assert.EqualError(t, err, ErrInvalidInput.Error())
	})

	t.Run("Test validator with NHS number that is not a number", func(t *testing.T) {
		v := NewValidator()

		invalidNHSNumber := "123TEST123"
		err := v.Validate(invalidNHSNumber)
		assert.EqualError(t, err, ErrNotANumber.Error())
	})

	t.Run("Test validator with NHS number that has check digit 10", func(t *testing.T) {
		v := NewValidator()

		invalidNHSNumber := "1234567891"
		err := v.Validate(invalidNHSNumber)
		assert.EqualError(t, err, ErrCheckDigitIsTen.Error())
	})
}
