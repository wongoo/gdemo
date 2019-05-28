package numbers

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func printInt8Bits(t *testing.T, name string, i int8) {
	t.Logf("%s: %b", name, byte(i))
}

func printInt16Bits(t *testing.T, name string, i int16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	t.Logf("%s: %b", name, b)
}

func TestAbs(t *testing.T) {
	doInt8AbsTest(t, 127)
	doInt8AbsTest(t, -127)
	doInt8AbsTest(t, -128) // overflow
	doInt16AbsTest(t, 127)
	doInt16AbsTest(t, -128)
	doInt16AbsTest(t, math.MinInt16) // overflow

	assert.Equal(t, float64(128), math.Abs(-128))
	assert.Equal(t, float64(math.MaxInt16+1), math.Abs(math.MinInt16))
	assert.Equal(t, float64(math.MaxInt64), math.Abs(math.MinInt64))
}

func doInt8AbsTest(t *testing.T, n int8) {
	t.Log("----------------")

	printInt8Bits(t, "n", n)

	y := n >> 7
	printInt8Bits(t, "y", y)

	xor := n ^ y
	printInt8Bits(t, "xor", xor)

	abs := xor - y
	printInt8Bits(t, "abs", abs)

	t.Logf("src int: %d, abs int: %d", int8(n), int8(abs))
}

func doInt16AbsTest(t *testing.T, n int16) {
	t.Log("----------------")

	printInt16Bits(t, "n", n)

	y := n >> 15
	printInt16Bits(t, "y", y)

	xor := n ^ y
	printInt16Bits(t, "xor", xor)

	abs := xor - y
	printInt16Bits(t, "abs", abs)

	t.Logf("src int: %d, abs int: %d", int16(n), int16(abs))
}
