package test

import (
	"testing"

	"github.com/everytv/prettytest"
)

type suite struct{ prettytest.Suite }

func TestSuite(t *testing.T) {
	prettytest.Run(
		t,
		new(suite),
	)
}
