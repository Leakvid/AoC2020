package main

import "fmt"

func main() {
	input := getInputFromFile()

	var passports []Passport
	current := Passport{fields: make(map[string]string)}

	for _, str := range input {
		if str == "" {
			passports = append(passports, current)
			current = Passport{fields: make(map[string]string)}
		} else {
			current.parseString(str)
		}
	}

	passports = append(passports, current)

	count := 0
	for _, passport := range passports {
		if passport.isValid() {
			count++
		}
	}
	fmt.Print(count, " are valid")
}
