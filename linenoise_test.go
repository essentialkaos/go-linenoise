package linenoise

// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
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
