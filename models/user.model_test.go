package models

import (
	"testing"
)

func TestValidateFunc(t *testing.T) {
	user := &User{ ID: "Ophelia", Email: "ophelia@gmail.com" }
	if validationErrors := user.Validate(); len(validationErrors) > 0 {
		t.Fatalf("Data %v should pass validation, but got validation errors %v", user, validationErrors)
	}
	user = &User{ ID: "", Email: "ophelia@gmail.com" }
	if validationErrors := user.Validate(); len(validationErrors) != 1 {
		t.Fatalf("Request with data %v should trigger 1 error, but got different number of errors %v", user, validationErrors)
	}
	user = &User{ ID: "", Email: "ophelia@" }
	if validationErrors := user.Validate(); len(validationErrors) != 2 {
		t.Fatalf("Request with data %v should trigger 2 errors, but got different number of errors %v", user, validationErrors)
	}
}