package lordcommander

import (
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/context"
)

// Option is a function that can be passed to a Commander's constructor to
// change it's default behavior.
type Option func(*commanderImpl)

// WithContext allows changing Commander's default context.
func WithContext(ctx context.Context) Option {
	return func(c *commanderImpl) {
		c.ctx = ctx
	}
}

// WithStdin allows changing Commander's default standard input.
func WithStdin(stdin io.Reader) Option {
	return func(c *commanderImpl) {
		c.stdin = stdin
	}
}

// WithSignal allows changing Commander's default termination signal.
func WithSignal(signal os.Signal) Option {
	return func(c *commanderImpl) {
		c.signal = signal
	}
}

// WithEnvironment allows changing Commander's default environment.
func WithEnvironment(env []string) Option {
	return func(c *commanderImpl) {
		c.env = env
	}
}

// DiscardStdout tells the Commander not to capture standard output.
func DiscardStdout() Option {
	return func(c *commanderImpl) {
		c.stdout = ioutil.Discard
	}
}

// DiscardStderr tells the Commander not to capture standard error.
func DiscardStderr() Option {
	return func(c *commanderImpl) {
		c.stderr = ioutil.Discard
	}
}
