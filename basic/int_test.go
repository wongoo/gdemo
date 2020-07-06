package basic

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestInt8Convert(t *testing.T) {
	var b1 = 0xd8
	var b2 = 0xe0
	var u8 = uint8(b1 - b2)
	t.Logf("u8=%d, %x, %x", u8, u8, byte(u8))

	var i8 = int8(u8)

	t.Logf("i8=%d, %x, %x", i8, i8, byte(i8))

	b3 := 0xF7
	t.Logf("b3=%d, %x, %x", int8(b3), b3, byte(b3))

}

var (
	_long3ByteValueMin  = -0x40000 // -262144
	_long3ByteValueMax  = 0x3ffff  // 262143
	_long3ByteZero      = byte(0x3c)
	_long3ByteZeroInt64 = int64(_long3ByteZero)
)

func TestLongToByteConvert(t *testing.T) {
	var i1 int64 = int64(_long3ByteValueMin)
	f1 := byte(i1 >> 16)
	t.Logf("f1=%d, %d, %x", f1, int8(f1), byte(f1)) // f1=252, -4, fc

	b1 := byte(_long3ByteZeroInt64 + (i1 >> 16))
	t.Logf("b1=%d, %d, %x", b1, int8(b1), byte(b1)) // b1=56, 56, 38

	var i2 int64 = int64(_long3ByteValueMax)
	b2 := byte(_long3ByteZeroInt64 + (i2 >> 16))
	t.Logf("b2=%d, %d, %x", b2, int8(b2), byte(b2)) // b2=63, 63, 3f

}

func TestBigInt(t *testing.T) {
	i := big.NewInt(1)
	t.Logf("big int i: %v, %p", i, i)

	i2, ok := i.SetString("1234", 10)
	assert.True(t, ok)

	t.Logf("big int i: %v, %p", i, i)
	t.Logf("big int i2: %v, %p", i2, i2)

	i.SetBytes()
}
