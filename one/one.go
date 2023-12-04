package one

import (
	"fmt"
	"strconv"
	"strings"
)

func isDigit(s string) bool {
	switch s {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	}
	return false
}

func ParseStringForDigit(s string) (int, error) {
	var digits []string

	for _, l := range strings.Split(s, "") {

		if isDigit(l) {
			digits = append(digits, l)
		}
	}

	if len(digits) < 1 {
		return 0, fmt.Errorf("string:'' contains less than 1 digit")
	}

	digit := digits[0] + digits[len(digits)-1]

	i, err := strconv.Atoi(digit)
	if err != nil {
		return 0, err
	}

	return i, nil
}
