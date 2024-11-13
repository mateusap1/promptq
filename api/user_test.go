package api

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	assert.Equal(t, true, validateEmail("a@b"))
	assert.Equal(t, false, validateEmail("@ba"))
	assert.Equal(t, false, validateEmail("ab@"))
	assert.Equal(t, false, validateEmail("a@"+strings.Repeat("a", 256)))
}

func TestValidatePassword(t *testing.T) {
	assert.Equal(t, true, validatePassword("P@ssw0rd"))
	assert.Equal(t, true, validatePassword("Pass1-_*)!"))
	assert.Equal(t, false, validatePassword("password"))
	assert.Equal(t, false, validatePassword("pWd0!"))
	assert.Equal(t, false, validatePassword("pAssw0rd"))
	assert.Equal(t, false, validatePassword("12345a678@"))
	assert.Equal(t, false, validatePassword("12345A678!"))
	assert.Equal(t, false, validatePassword(strings.Repeat("a", 64)+"A@0"))
}
