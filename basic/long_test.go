package basic

import "testing"

func TestLong3Convert(t *testing.T) {
	printLongHex(t, 0xFFFF0000)
	printLongHex(t, -0x40000)
	printLongHex(t, 0x3ffff)
}

//     -2048 2047 : 0xfffff800 0x000007ff
// -262144 262143 : 0xfffc0000 0x0003ffff
//
func printLongHex(t *testing.T, b int) {
	t.Log("-----------------")
	t.Logf("b=%d, %x, %x", int32(b), b, uint32(b))
	b1 := b - 1
	t.Logf("b1=%d, %x, %x", int32(b1), b1, uint32(b1))
	b2 := b + 1
	t.Logf("b2=%d, %x, %x", int32(b2), b2, uint32(b2))
}
