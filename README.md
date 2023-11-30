# NHS Number Validator

The goal of this exercise is to validate if a number is a NHS number.

# Instructions

The NHS NUMBER, the primary identifier of a PERSON, is a unique identifier for a PATIENT within the NHS in England and Wales.

The NHS NUMBER is 10 numeric digits in length. The tenth digit is a check digit used to confirm its validity.  The check digit is validated using the Modulus 11 algorithm and the use of this algorithm is mandatory.  There are 5 steps in the validation of the check digit:

(The following specification is taken straight from an NHS specification of the NHS NUMBER):

## Step 1
Multiply each of the first nine digits by a weighting factor as follows:

| Digit position | Factor |
| --- | --- |
| 1 | 10 |
| 2 | 9 |
| 3 | 8 |
| 4 | 7 |
| 5 | 6 |
| 6 | 5 |
| 7 | 4 |
| 8 | 3 |
| 9 | 2 |

## Step 2

Add the results of each multiplication together.

## Step 3

Divide the total by 11 and establish the remainder.

## Step 4

Subtract the remainder from 11 to give the check digit.
If the result is 11 then a check digit of 0 is used. If the result is 10 then the NHS NUMBER is invalid and not used.

## Step 5

Check the remainder matches the check digit. If it does not, the NHS NUMBER is invalid. (I suspect an error in this section, as the remainder will never match the check digit, while the last digit of the NHS number should correspond to the check digit.)

# Requirements

Your Go environment is set up, and I am currently using Go version 1.20.

# Make it, Run it
I let a `main.go` in order to be able to test it manually

- Make
- ./nhsnumber_bin

# Test it
- go test ./nhsnumber
