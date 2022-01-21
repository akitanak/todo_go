package entities

import "fmt"

func ValidateDescription(description string) error {
	maxLength := 64
	if len(description) > maxLength {
		return fmt.Errorf("description is too long. max: %d, actual: %d", maxLength, len(description))
	}

	return nil
}
