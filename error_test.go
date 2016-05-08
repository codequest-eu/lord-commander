package lordcommander

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrAlreadyUsed(t *testing.T) {
	assert.Error(t, ErrAlreadyUsed)
	assert.Equal(t, "commander already used", ErrAlreadyUsed.Error())
}

func TestErrNonZero(t *testing.T) {
	assert.Error(t, ErrNonZero)
	assert.Equal(t, "command exited with non-zero status", ErrNonZero.Error())
}
