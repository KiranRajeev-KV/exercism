package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var clean strings.Builder

	for _, v := range phoneNumber {
		if v >= '0' && v <= '9' {
			clean.WriteRune(v)
		} else if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			return "", errors.New("letters not permitted")
		}
	}

	number := clean.String()

	if len(number) == 11 {
		if number[0] != '1' {
			return "", errors.New("11 digits must start with 1")
		}
		number = number[1:]
	}

	if len(number) != 10 {
		return "", errors.New("incorrect number of digits")
	}

	if number[0] == '0' || number[0] == '1' {
		return "", errors.New("area code cannot start with 0 or 1")
	}

	if number[3] == '0' || number[3] == '1' {
		return "", errors.New("exchange code cannot start with 0 or 1")
	}

	return number, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}