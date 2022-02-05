package entities

import (
	"fmt"
	"unicode"
)

const (
	descriptionMinLength = 1
	descriptionMaxLength = 64
)

// ValidateDescription validates `Todo`'s description.
func ValidateDescription(description string) error {
	for _, c := range description {
		if unicode.IsSpace(c) && c != rune(' ') {
			return fmt.Errorf("description contains some prohibit charactor(tab, new line, carriage return, vertical tab, form feed)")
		}
	}

	descLen := len(description)
	if descLen < descriptionMinLength {
		return fmt.Errorf("description is too short. min: %d, actual: %d", descriptionMinLength, descLen)
	}

	if descLen > descriptionMaxLength {
		return fmt.Errorf("description is too long. max: %d, actual: %d", descriptionMaxLength, descLen)
	}

	return nil
}
