package buffer

import "testing"
import "github.com/tdewolff/test"

func TestNullTerminator(t *testing.T) {
	b := []byte{'a', 'b', 'c', 'd'}
	z, restore := NullTerminator(b[:2])

	test.T(t, len(z), 3, "must have terminating NULL")
	test.T(t, z[2], byte(0), "must have terminating NULL")
	test.Bytes(t, b, []byte{'a', 'b', 0, 'd'}, "terminating NULL overwrites underlying buffer")

	restore()
	test.Bytes(t, b, []byte{'a', 'b', 'c', 'd'}, "terminating NULL has been restored")
}
