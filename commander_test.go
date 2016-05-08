package lordcommander

import (
	"bytes"
	"syscall"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/stretchr/testify/assert"
)

func TestConstructorWithoutOptions(t *testing.T) {
	commander := New()
	assert.IsType(t, (*commanderImpl)(nil), commander)
}

func TestConstructorWithOptions(t *testing.T) {
	commander := New(WithSignal(syscall.SIGHUP))
	assert.Equal(t, syscall.SIGHUP, commander.(*commanderImpl).signal)
}

func TestStdout(t *testing.T) {
	const output = "bacon"
	commander := New()
	err := commander.Run("echo", output)
	assert.NoError(t, err)
	assert.Contains(t, commander.Stdout().String(), output)
	assert.Empty(t, commander.Stderr().String())
}

func TestStderr(t *testing.T) {
	commander := New()
	err := commander.Run("git", "clone", "idontexist")
	assert.Error(t, err)
	assert.Contains(t, commander.Stderr().String(), "does not exist")
	assert.Empty(t, commander.Stdout().String())
}

func TestTerminate(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Microsecond)
	commander := New(WithContext(ctx))
	err := commander.Run("sleep", "10")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "deadline exceeded")
	assert.Contains(t, err.Error(), "killed with terminated")
}

func TestStdin(t *testing.T) {
	const input = "bacon"
	commander := New(WithStdin(bytes.NewBufferString(input)))
	err := commander.Run("cat")
	assert.NoError(t, err)
	assert.Contains(t, commander.Stdout().String(), input)
}

func TestEnvironment(t *testing.T) {
	const envVariable = "BACON=bacon"
	commander := New(WithEnvironment([]string{envVariable}))
	err := commander.Run("env")
	assert.NoError(t, err)
	assert.Contains(t, commander.Stdout().String(), envVariable)
}

func TestDiscardStdout(t *testing.T) {
	commander := New(DiscardStdout())
	err := commander.Run("echo", "bacon")
	assert.NoError(t, err)
	assert.Nil(t, commander.Stdout())
}

func TestDiscardStderr(t *testing.T) {
	commander := New(DiscardStderr())
	err := commander.Run("git", "clone", "idontexist")
	assert.Error(t, err)
	assert.Nil(t, commander.Stderr())
}

func TestReuse(t *testing.T) {
	commander := New()
	assert.NoError(t, commander.Run("echo", "bacon"))
	assert.Equal(t, ErrAlreadyUsed, commander.Run("echo", "bacon"))
}
