package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func getRequiredFields() []string {
	return []string{ "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" }
}

type Passport struct {
	fields map[string]string
}

func HasRequiredFields(passport Passport) bool {
	valid := true

	for _, req := range getRequiredFields() {
		_, ok := passport.fields[req]

		if !ok {
			valid = false
		}
	}

	return valid
}

func ValueInRange(val string, min int, max int) bool {
	iVal, err := strconv.Atoi(val)

	if err != nil {
		panic(err)
	}

	return iVal >= min && iVal <= max
}

func BirthYearValid(val string) bool {
	return ValueInRange(val, 1920, 2002)
}

func IssueYearValid(val string) bool {
	return ValueInRange(val, 2010, 2020)
}

func ExpirationYearValid(val string) bool {
	return ValueInRange(val, 2020, 2030)
}

func HeightValid(val string) bool {
	if len(val) < 3 {
		return false
	}

	measure := val[0:len(val)-2]
	units := val[len(val)-2:]

	if units == "cm" {
		return ValueInRange(measure, 150, 193)
	} else if units == "in" {
		return ValueInRange(measure, 59, 76)
	} else {
		return false
	}
}

func HairColorValid(val string) bool {
	re := regexp.MustCompile("^#([a-fA-F0-9]{6})$")
	return re.MatchString(val)
}

func EyeColorValid(val string) bool {
	switch val {
		case
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func PassportIdValid(val string) bool {
	if len(val) != 9 {
		return false
	}

	_, err := strconv.Atoi(val)

	if err != nil {
		return false
	}

	return true
}

func HasValidFields(passport Passport) bool {
	return BirthYearValid(passport.fields["byr"]) &&
		IssueYearValid(passport.fields["iyr"]) &&
		ExpirationYearValid(passport.fields["eyr"]) &&
		HeightValid(passport.fields["hgt"]) &&
		HairColorValid(passport.fields["hcl"]) &&
		EyeColorValid(passport.fields["ecl"]) &&
		PassportIdValid(passport.fields["pid"])
}

func main() {
	lines := ReadInput("day-04-input.txt")

	i := 0
	requiredFieldsCount := 0
	validFieldsCount := 0

	// parse passports
	for i < len(lines) {
		line := lines[i]

		passport := Passport{
			fields: make(map[string]string),
		}

		// parse a single passport
		for line != "" {
			fields := strings.Split(line, " ")

			for _, field := range fields {
				splitField := strings.Split(field, ":")
				passport.fields[splitField[0]] = splitField[1]
			}

			i++

			if i >= len(lines) {
				line = ""
			} else {
				line = lines[i]
			}
		}

		// count valid passport
		if HasRequiredFields(passport) {
			requiredFieldsCount++

			if HasValidFields(passport) {
				validFieldsCount++
			}
		}

		i++
	}

	fmt.Println("Passports w/ required fields (part 1):", requiredFieldsCount)
	fmt.Println("Passports w/ valid fields (part 2):", validFieldsCount)
}
