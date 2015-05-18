package doc

import (
	"testing"

	"github.com/adams-sarah/prettytest"
)

type suite struct{ prettytest.Suite }

func TestSuite(t *testing.T) {
	prettytest.Run(
		t,
		new(suite),
	)
}
