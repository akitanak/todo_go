package entities

import "fmt"

const (
	descriptionMinLength = 1
	descriptionMaxLength = 64
)

// ValidateDescription validates `Todo`'s description.
func ValidateDescription(description string) error {
	descLen := len(description)
	if descLen < descriptionMinLength {
		return fmt.Errorf("description is too short. min: %d, actual: %d", descriptionMinLength, descLen)
	}

	if descLen > descriptionMaxLength {
		return fmt.Errorf("description is too long. max: %d, actual: %d", descriptionMaxLength, descLen)
	}

	return nil
}
