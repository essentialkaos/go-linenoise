package linenoise

// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type LinenoiseSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&LinenoiseSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

// This is not a real test, is just basic compilation test
func (s *LinenoiseSuite) TestSimple(c *C) {
	c.Assert(ErrKillSignal, NotNil)
}

// This is not a real test, is just basic compilation test
func (s *LinenoiseSuite) TestSetHintColor(c *C) {
	c.Assert(SetHintColor(1, false), NotNil)
	c.Assert(SetHintColor(38, false), NotNil)
	c.Assert(SetHintColor(120, false), NotNil)
	c.Assert(SetHintColor(31, false), IsNil)
	c.Assert(SetHintColor(92, false), IsNil)
}

func (s *LinenoiseSuite) TestNilHandler(c *C) {
	SetCompletionHandler(nil)
	c.Assert(complHandler, NotNil)

	SetHintHandler(nil)
	c.Assert(hintHandler, NotNil)
}
