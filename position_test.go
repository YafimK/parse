package parse

import (
	"fmt"
	"testing"

	"github.com/tdewolff/parse/v2/buffer"
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
		{3, "\u2028x", 2, 1},
		{3, "\u2029x", 2, 1},

		// edge cases
		{0, "", 1, 1},
		{0, "\nx", 1, 1},
		{1, "\r\nx", 1, 2},
		{-1, "x", 1, 2}, // continue till the end
		{0, "\x00a", 1, 1},
		{1, "x\u2028x", 1, 2},
		{2, "x\u2028x", 1, 3},
		{3, "x\u2028x", 1, 4},
	}
	for _, tt := range newlineTests {
		t.Run(fmt.Sprint(tt.buf, " ", tt.offset), func(t *testing.T) {
			l := buffer.NewLexerString(tt.buf)
			line, col, _ := Position(l, tt.offset)
			test.T(t, line, tt.line, "line")
			test.T(t, col, tt.col, "column")
		})
	}
}
