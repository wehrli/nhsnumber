package main

import (
	"fmt"

	"nhs/nhsnumber"
)

func main() {
	nhsValidator := nhsnumber.NewValidator()
	validNHSNumber := [3]string{"5990128088", "1275988113", "4536026665"}
	invalidNHSNumber := [2]string{"5990128087", "4536016660"}

	for _, nhsNumber := range validNHSNumber {
		err := nhsValidator.Validate(nhsNumber)
		if err != nil {
			fmt.Printf("NHS number %s is invalid and sould be valid\n", nhsNumber)
			fmt.Printf("Error: %s\n", err)
			return
		}
	}

	fmt.Println("All valid NHS numbers are valid")

	for _, nhsNumber := range invalidNHSNumber {
		err := nhsValidator.Validate(nhsNumber)
		if err == nil {
			fmt.Printf("NHS number %s is valid and sould be invalid\n", nhsNumber)
			return
		}
	}

	fmt.Println("All invalid NHS numbers are invalid")
}
