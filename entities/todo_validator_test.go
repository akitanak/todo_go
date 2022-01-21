package entities

import (
	"testing"
)

func TestValidateDescription(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"normal case": {
			input: "build a todo app.",
			want:  "",
		},
		"too short": {
			input: "",
			want:  "description is too short. min: 1, actual: 0",
		},
		"too long": {
			input: "12345678901234567890123456789012345678901234567890123456789012345",
			want:  "description is too long. max: 64, actual: 65",
		},
	}

	for name, test := range tests {
		err := ValidateDescription(test.input)

		if test.want == "" {
			if err != nil {
				t.Errorf("%v - ValidateDescription(%v): got: %v, want: %v", name, test.input, err.Error(), nil)
			}
		} else {
			switch {
			case err == nil:
				t.Errorf("%v - ValidateDescription(%v): got: %v, want: %v", name, test.input, err, test.want)
			case err.Error() != test.want:
				t.Errorf("%v - ValidateDescription(%v): got: %v, want: %v", name, test.input, err.Error(), test.want)
			}
		}
	}
}
