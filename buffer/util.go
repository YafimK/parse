package buffer

// NullTerminator will append a NULL character to the end of the byte slice.
// If there is capacity, it will expand the length by one and replace the byte past the end of the buffer by NULL.
// By calling the returned function, the replaced byte past the end of the buffer by NULL will get restored.
func NullTerminator(b []byte) ([]byte, func()) {
    n := len(b)
    if cap(b) > n {
        // There is room in the capacity to add a NULL character
        b = b[:n+1]
        c := b[n]
        b[n] = 0
        return b, func() {
            b[n] = c
        }
    } else {
        // No room, move data to a bigger buffer
        return append(b, 0), func(){}
    }
}
