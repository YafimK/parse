package parse

import (
	"fmt"
	"strings"

	"github.com/tdewolff/parse/v2/buffer"
)

// Position returns the line and column number for a certain position in a file. It is useful for recovering the position in a file that caused an error.
// It only treates \n, \r, and \r\n as newlines, which might be different from some languages also recognizing \f, \u2028, and \u2029 to be newlines.
func Position(l *buffer.Lexer, offset int) (line, col int, context string) {
	line = 1
	for {
		c := l.Peek(0)
		if c == 0 && l.IsEOF() || offset == l.Pos() {
			col = l.Pos() + 1
			context = positionContext(l, line, col)
			return
		}

		nNewline := 0
		if c == '\n' {
			nNewline = 1
		} else if c == '\r' {
			if l.Peek(1) == '\n' {
				nNewline = 2
			} else {
				nNewline = 1
			}
		} else if c >= 0xC0 {
			if r, n := l.PeekRune(0); r == '\u2028' || r == '\u2029' {
				nNewline = n
			}
		} else {
			l.Move(1)
		}

		if nNewline > 0 {
			if offset < l.Pos()+nNewline {
				// move onto offset position, let next iteration handle it
				l.Move(offset - l.Pos())
				continue
			}
			l.Move(nNewline)
			line++
			offset -= l.Pos()
			l.Skip()
		}
	}
}

func positionContext(l *buffer.Lexer, line, col int) (context string) {
	for {
		c := l.Peek(0)
		if c == 0 && l.IsEOF() || c == '\n' || c == '\r' {
			break
		}
		l.Move(1)
	}

	// replace unprintable characters by a space
	b := l.Lexeme()
	for i, c := range b {
		if c < 0x20 || c == 0x7F {
			b[i] = ' '
		}
	}

	context += fmt.Sprintf("%5d: %s\n", line, string(b))
	context += fmt.Sprintf("%s^", strings.Repeat(" ", col+6))
	return
}
