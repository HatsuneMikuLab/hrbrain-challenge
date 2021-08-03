package models

import (
	"testing"
)

type TestCases map[*User]int

func TestValidateFunc(t *testing.T) {
	testCases := TestCases{
		&User{ ID: "Ophelia", Email: "ophelia@gmail.com" }: 0,
		&User{ ID: "", Email: "ophelia@gmail.com" }: 1,
		&User{ ID: "", Email: "ophelia@" }: 2,
	}
	for input, expectedOutput := range testCases {
		if validationErrors := input.Validate(); len(validationErrors) != expectedOutput {
			t.Fatalf("Input %v should result in %v validation erros, but got %v", input, expectedOutput, len(validationErrors))
		}
	}
}