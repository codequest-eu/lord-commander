package lordcommander

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mockCommanderSuite struct {
	suite.Suite
	commander *MockCommander
}

func (s *mockCommanderSuite) SetupTest() {
	s.commander = new(MockCommander)
}

func (s *mockCommanderSuite) TestRun() {
	s.commander.On("Run", "send", "more", "bacon").Return(nil)
	s.Nil(s.commander.Run("send", "more", "bacon"))
}

func (s *mockCommanderSuite) TestStdout() {
	ret := bytes.NewBuffer(nil)
	s.commander.On("Stdout").Return(ret)
	s.Equal(ret, s.commander.Stdout())
}

func (s *mockCommanderSuite) TestSterr() {
	ret := bytes.NewBuffer(nil)
	s.commander.On("Stderr").Return(ret)
	s.Equal(ret, s.commander.Stderr())
}

func TestMockCommander(t *testing.T) {
	suite.Run(t, new(mockCommanderSuite))
}
