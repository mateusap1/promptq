package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	assert.Equal(t, true, ValidateEmail("a@b"))
	assert.Equal(t, false, ValidateEmail("@ba"))
	assert.Equal(t, false, ValidateEmail("ab@"))
	assert.Equal(t, false, ValidateEmail("a@"+strings.Repeat("a", 256)))
}

func TestValidatePassword(t *testing.T) {
	assert.Equal(t, true, ValidatePassword("P@ssw0rd"))
	assert.Equal(t, true, ValidatePassword("Pass1-_*)!"))
	assert.Equal(t, false, ValidatePassword("password"))
	assert.Equal(t, false, ValidatePassword("pWd0!"))
	assert.Equal(t, false, ValidatePassword("pAssw0rd"))
	assert.Equal(t, false, ValidatePassword("12345a678@"))
	assert.Equal(t, false, ValidatePassword("12345A678!"))
	assert.Equal(t, false, ValidatePassword(strings.Repeat("a", 64)+"A@0"))
}
