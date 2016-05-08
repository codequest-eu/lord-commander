package lordcommander

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"

	"github.com/codequest-eu/ctxcmd"

	"golang.org/x/net/context"
)

type commanderImpl struct {
	cmd    *exec.Cmd
	ctx    context.Context
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	signal os.Signal
	env    []string
}

// New is a constructor for an actual implementation of the Commander interface,
// allowing it to be customized using Option functions.
func New(opts ...Option) Commander {
	ret := &commanderImpl{
		ctx:    context.Background(),
		signal: syscall.SIGTERM,
		stdout: bytes.NewBuffer(nil),
		stderr: bytes.NewBuffer(nil),
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func (c *commanderImpl) Run(name string, args ...string) error {
	if c.cmd != nil {
		return ErrAlreadyUsed
	}
	c.cmd = c.buildCommand(name, args...)
	if err := ctxcmd.NewCommand(c.ctx, c.cmd).RunWithSignal(c.signal); err != nil {
		return err
	}
	if !c.cmd.ProcessState.Success() {
		return ErrNonZero
	}
	return nil
}

func (c *commanderImpl) Stdout() *bytes.Buffer {
	if c.stdout == ioutil.Discard {
		return nil
	}
	return c.stdout.(*bytes.Buffer)
}

func (c *commanderImpl) Stderr() *bytes.Buffer {
	if c.stderr == ioutil.Discard {
		return nil
	}
	return c.stderr.(*bytes.Buffer)
}

func (c *commanderImpl) buildCommand(name string, args ...string) *exec.Cmd {
	ret := exec.Command(name, args...)
	ret.Stdin = c.stdin
	ret.Stdout = c.stdout
	ret.Stderr = c.stderr
	if c.env != nil {
		ret.Env = c.env
	}
	return ret
}
