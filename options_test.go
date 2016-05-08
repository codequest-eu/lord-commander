package lordcommander

import (
	"bytes"
	"io/ioutil"
	"syscall"
	"testing"

	"golang.org/x/net/context"

	"github.com/stretchr/testify/suite"
)

type optionsSuite struct {
	suite.Suite
	commander  *commanderImpl
	defaultCtx context.Context
}

func (s *optionsSuite) SetupTest() {
	s.commander = New(WithContext(s.defaultCtx)).(*commanderImpl)
}

func (s *optionsSuite) TestWithContext() {
	s.Equal(s.defaultCtx, s.commander.ctx)
	newCtx := context.Background()
	WithContext(newCtx)(s.commander)
	s.Equal(newCtx, s.commander.ctx)
}

func (s *optionsSuite) TestWithStdin() {
	s.Nil(s.commander.stdin)
	newStdin := bytes.NewBuffer(nil)
	WithStdin(newStdin)(s.commander)
	s.Equal(newStdin, s.commander.stdin)
}

func (s *optionsSuite) TestWithSignal() {
	s.Equal(syscall.SIGTERM, s.commander.signal)
	WithSignal(syscall.SIGHUP)(s.commander)
	s.Equal(syscall.SIGHUP, s.commander.signal)
}

func (s *optionsSuite) TestWithEnvironment() {
	newEnv := []string{"bacon"}
	s.Nil(s.commander.env)
	WithEnvironment(newEnv)(s.commander)
	s.Equal(newEnv, s.commander.env)
}

func (s *optionsSuite) TestDiscardStdout() {
	s.NotNil(s.commander.Stdout())
	s.NotEqual(ioutil.Discard, s.commander.stdout)
	DiscardStdout()(s.commander)
	s.Nil(s.commander.Stdout())
	s.Equal(ioutil.Discard, s.commander.stdout)
}

func (s *optionsSuite) TestDiscardStderr() {
	s.NotNil(s.commander.Stderr())
	s.NotEqual(ioutil.Discard, s.commander.stderr)
	DiscardStderr()(s.commander)
	s.Nil(s.commander.Stderr())
	s.Equal(ioutil.Discard, s.commander.stderr)
}

func TestOptions(t *testing.T) {
	suite.Run(t, new(optionsSuite))
}
