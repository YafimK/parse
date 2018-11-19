package parse

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tdewolff/test"
)

func TestPosition(t *testing.T) {
	var newlineTests = []struct {
		offset int
		buf    string
		line   int
		col    int
	}{
		{0, "x", 1, 1},
		{1, "xx", 1, 2},
		{2, "x\nx", 2, 1},
		{2, "\n\nx", 3, 1},
		{3, "\nxxx", 2, 3},
		{2, "\r\nx", 2, 1},
		{1, "\rx", 2, 1},

		// edge cases
		{0, "", 1, 1},
		{0, "\n", 1, 1},
		{1, "\r\n", 1, 2},
		{-1, "x", 1, 2}, // continue till the end
		{0, "\x00a", 1, 1},
	}
	for _, tt := range newlineTests {
		t.Run(fmt.Sprint(tt.buf, " ", tt.offset), func(t *testing.T) {
			b := bytes.NewBufferString(tt.buf)
			line, col, _ := Position(b, tt.offset)
			test.T(t, line, tt.line, "line")
			test.T(t, col, tt.col, "column")
		})
	}
}
