package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	fields map[string]string
}

type PassportBehaviour interface {
	parseFields(str string)
	isValid() bool
}

func (p Passport) parseString(str string) {
	for _, field := range strings.Split(str, " ") {
		splitted := strings.Split(field, ":")
		p.fields[splitted[0]] = splitted[1]
	}
}

type Rule struct {
	field   string
	isValid func(val string) bool
}

func (p Passport) isValid() bool {
	rules := []Rule{
		{"byr", func(val string) bool { return isStringBetween(1920, 2002, val) }},
		{"iyr", func(val string) bool { return isStringBetween(2010, 2020, val) }},
		{"eyr", func(val string) bool { return isStringBetween(2020, 2030, val) }},
		{"hgt", isHeightValid},
		{"hcl", func(val string) bool { return isRegexMatch("#[a-f0-9]{6}", val) }},
		{"ecl", isInList},
		{"pid", func(val string) bool { return isRegexMatch("[0-9]{9}", val) && len(val) == 9 }},
	}

	for _, rule := range rules {
		value, ok := p.fields[rule.field]
		if !ok || !rule.isValid(value) {
			return false
		}
	}

	return true
}

func isInList(val string) bool {
	validValues := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	for _, valid := range validValues {
		if val == valid {
			return true
		}
	}

	return false
}

func isHeightValid(val string) bool {

	if strings.Contains(val, "cm") {
		height := strings.TrimSuffix(val, "cm")
		return isStringBetween(150, 193, height)
	}

	if strings.Contains(val, "in") {
		height := strings.TrimSuffix(val, "in")
		return isStringBetween(59, 76, height)
	}

	return false
}

func isStringBetween(min int, max int, val string) bool {
	i, err := strconv.Atoi(val)
	return err == nil && (i <= max && i >= min)
}

func isRegexMatch(pattern string, val string) bool {
	match, err := regexp.MatchString(pattern, val)
	return err == nil && match
}
