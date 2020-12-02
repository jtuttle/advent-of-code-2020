package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PasswordEntry struct {
	first int
	second int
	letter string
	password string
}

func ParsePasswordEntry(str string) PasswordEntry {
	fields := strings.FieldsFunc(str, Split)

	first, err := strconv.Atoi(fields[0])

	if err != nil {
		panic(err)
	}

	second, err := strconv.Atoi(fields[1])

	if err != nil {
		panic(err)
	}

	return PasswordEntry{
		first: first,
		second: second,
		letter: fields[2],
		password: fields[3],
	}
}

func Split(r rune) bool {
	return r == '-' || r == ':' || r == ' '
}

func IsValidEntry1(entry PasswordEntry) bool {
	occurrences := strings.Count(entry.password, entry.letter)
	return occurrences >= entry.first && occurrences <= entry.second
}

func IsValidEntry2(entry PasswordEntry) bool {
	charAtFirst := entry.password[entry.first - 1]
	charAtSecond := entry.password[entry.second - 1]

	firstMatches := (string(charAtFirst) == entry.letter)
	secondMatches := (string(charAtSecond) == entry.letter)

	return firstMatches != secondMatches
}

func main() {
	lines := ReadInput("day-02-input.txt")

	var entries []PasswordEntry

	for _, line := range lines {
		entry := ParsePasswordEntry(line)
		entries = append(entries, entry)
	}

	validCount := 0

	for _, entry := range entries {
		if IsValidEntry1(entry) {
			validCount++
		}
	}

	fmt.Println("Valid passwords (part 1):", validCount)

	validCount = 0

	for _, entry := range entries {
		if IsValidEntry2(entry) {
			validCount++
		}
	}

	fmt.Println("Valid passwords (part 2):", validCount)
}
