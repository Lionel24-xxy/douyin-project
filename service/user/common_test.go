package user

import (
	"testing"

)

func TestIsValidUser(t *testing.T) {
	if err := IsValidUser("xxy", "aA12345678"); err != nil {
		t.Errorf(err.Error())
	}
}
