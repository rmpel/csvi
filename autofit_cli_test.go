package csvi_test

import (
	"testing"

	"github.com/hymkor/csvi/uncsv"
)

// -autofit reads all data before the first draw; editing must still work.
func TestAutoFitOption(t *testing.T) {
	src := "a,bb\nccc,d\n"
	op := "r|x"
	exp := "x,bb\nccc,d" + uncsv.OsNewline
	testCase(t, src, op, exp, "-autofit")
}
