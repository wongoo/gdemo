// authors: wangoo
// created: 2018-05-25
// buffer demo

package basic

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestBuffer(t *testing.T) {
	NewBufferString()
	Grow()

	Write()
	WriteString()
	WriteByte()
	WriteRune()
	WriteTo()

	Read()
	ReadByte()
	ReadRune()
	ReadString()
	ReadBytes()
	ReadFrom()
	Unread()

	Next()

	Reset()
	Truncate()
}

func NewBufferString() {
	fmt.Println("----------------> NewBufferString")
	// the following three are equal
	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))
	buf3 := bytes.NewBuffer([]byte{'h', 'e', 'l', 'l', 'o'})
	fmt.Printf("buf1 cap=%v\n", buf1.Cap())
	fmt.Printf("buf1 len=%v\n", buf1.Len())
	fmt.Printf("buf1 len=%v\n", buf1.UnreadRune())

	fmt.Printf("buf1=%v\n", buf1.String())
	fmt.Printf("buf2=%v\n", buf2.String())
	fmt.Printf("buf3=%v\n", buf3.String())

	// the following two are equal
	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Printf("buf4 Len=%v\n", buf4.Len())
	fmt.Printf("buf5 Len=%v\n", buf5.Len())
}

// func (b *Buffer) Grow(n int)
func Grow() {
	fmt.Println("----------------> Grow")

	buf := bytes.NewBufferString("")

	// buf.Cap=0
	fmt.Printf("buf Cap=%v\n", buf.Cap())

	buf.Grow(8)
	buf.WriteByte('h')
	fmt.Printf("h buf Cap=%v\n", buf.Cap())
	buf.WriteByte('e')
	fmt.Printf("e buf Cap=%v\n", buf.Cap())
	buf.WriteByte('l')
	fmt.Printf("l buf Cap=%v\n", buf.Cap())
	buf.WriteByte('l')
	fmt.Printf("l buf Cap=%v\n", buf.Cap())
	buf.WriteByte('o')
	fmt.Printf("o buf Cap=%v\n", buf.Cap())

	buf.Grow(6)
	// buf.Cap=22= 8*2+6
	fmt.Printf("buf Cap=%v\n", buf.Cap())

	buf.Grow(11)
	// still buf.Cap=22 , NOT change
	fmt.Printf("buf Cap=%v\n", buf.Cap())

	buf.WriteString("123")      // 8
	buf.WriteString("12345678") // 16

	buf.Grow(11)
	// buf.Cap=55=22*2+11
	fmt.Printf("buf Cap=%v\n", buf.Cap())
}

// func (b *Buffer) Write(p []byte) (n int, err error)
func Write() {
	fmt.Println("----------------> Write")
	s := []byte(" world")
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	//write slice at then end of buffer
	buf.Write(s)

	fmt.Println(buf.String())
}

// func (b *Buffer) WriteString(s string) (n int, err error)
func WriteString() {
	fmt.Println("----------------> WriteString")
	s := " world"
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	//write string at then end of buffer
	buf.WriteString(s)

	fmt.Println(buf.String())
}

// func (b *Buffer) WriteByte(c byte) error
func WriteByte() {
	fmt.Println("----------------> WriteByte")
	var b byte = '!'
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	//write byte at then end of buffer
	buf.WriteByte(b)

	fmt.Println(buf.String())
}

// func (b *Buffer) WriteRune(r rune) (n int, err error)
func WriteRune() {
	fmt.Println("----------------> WriteRune")
	var r = '啊'
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	//write rune at then end of buffer
	buf.WriteRune(r)

	fmt.Println(buf.String())
}

// func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
func WriteTo() {
	fmt.Println("----------------> WriteTo")
	path := "text.txt"
	// create file
	file, _ := os.Create(path)

	buf := bytes.NewBufferString("hello")

	// file is an io.Writer
	buf.WriteTo(file)

	// this do same work as the above
	fmt.Fprintf(file, buf.String())

	// remove file
	os.Remove(path)
}

// func (b *Buffer) Read(p []byte) (n int, err error)
func Read() {
	fmt.Println("----------------> Read")
	s1 := []byte("hello")
	buff := bytes.NewBuffer(s1)
	s2 := []byte(" world")
	buff.Write(s2)
	fmt.Println(buff.String())

	// slice with length 3
	s3 := make([]byte, 3)

	// read 3 bytes from buffer into s3
	buff.Read(s3)

	// buffer="lo world"
	fmt.Println(buff.String())

	// s3="hel"
	fmt.Println(string(s3))

	// read 3 bytes from buffer into s3
	buff.Read(s3)

	//buff="world"
	fmt.Println(buff.String())

	//s3="lo "
	fmt.Println(string(s3))
}

// func (b *Buffer) ReadByte() (c byte, err error)
func ReadByte() {
	fmt.Println("----------------> ReadByte")
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())

	// read one byte assign to b
	b, _ := buf.ReadByte()

	// buf=ello
	fmt.Println(buf.String())

	// b=h
	fmt.Println(string(b))
}

// func (b *Buffer) ReadRune() (r rune, size int, err error)
func ReadRune() {
	fmt.Println("----------------> ReadRune")
	buf := bytes.NewBufferString("好hello")
	fmt.Println(buf.String())

	// read one rune from buf into b
	b, n, _ := buf.ReadRune()

	// buf=hello
	fmt.Println(buf.String())

	// b=好
	fmt.Println(string(b))

	// b=好 contains 3 bytes
	fmt.Println(n)

	// read one rune from buf into b
	b, n, _ = buf.ReadRune()

	// buf= ello
	fmt.Println(buf.String())

	// b=h
	fmt.Println(string(b))

	// b=h contains 1 byte
	fmt.Println(n)
}

// func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
// ReadBytes需要一个byte作为分隔符，读的时候从缓冲区里找第一个出现的分隔符（delim），
// 找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为byte类型的slice，返回后，缓冲区也会空掉一部分
func ReadBytes() {
	fmt.Println("----------------> ReadBytes")

	//delim='e'
	var d byte = 'e'

	buf := bytes.NewBufferString("hello")

	//buf=hello
	fmt.Println(buf.String())

	// read bytes till delim='e'
	b, _ := buf.ReadBytes(d)

	// buf=llo
	fmt.Println(buf.String())

	// b=he
	fmt.Println(string(b))
}

// func (b *Buffer) ReadString(delim byte) (line string, err error)
func ReadString() {
	fmt.Println("----------------> ReadString")

	//delim='e'
	var d byte = 'e'

	buf := bytes.NewBufferString("hello")

	//buf=hello
	fmt.Println(buf.String())

	// read bytes till delim='e'
	s, _ := buf.ReadString(d)

	// buf=llo
	fmt.Println(buf.String())

	// s=he
	fmt.Println(s)
}

// func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
func ReadFrom() {
	fmt.Println("----------------> ReadFrom")

	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBufferString(" world ")

	//read bytes from buf2(io.Reader) and append to the end of buf1
	buf1.ReadFrom(buf2)

	// buf1="hello world"
	fmt.Println(buf1.String())

	//buf2=""
	fmt.Println(buf2.String())
}

// func (b *Buffer) Next(n int) []byte
func Unread() {
	fmt.Println("----------------> Next")

	buf := bytes.NewBufferString("hello")

	//buf=hello
	fmt.Println(buf.String())

	b, _ := buf.ReadByte()
	fmt.Println(string(b))

	// buf=ello
	fmt.Println(buf.String())

	buf.UnreadByte()
	// buf=hello
	fmt.Println(buf.String())

	buf = bytes.NewBufferString("世界你好")

	//buf=世界你好
	fmt.Println(buf.String())

	r, _, _ := buf.ReadRune()
	// 世
	fmt.Println(string(r))

	// buf=界你好
	fmt.Println(buf.String())

	buf.UnreadRune()
	// buf=世界你好
	fmt.Println(buf.String())

}

// func (b *Buffer) Next(n int) []byte
func Next() {
	fmt.Println("----------------> Next")

	buf := bytes.NewBufferString("hello")

	// buf=hello
	fmt.Println(buf.String())

	// read 2 bytes into b
	b := buf.Next(2)

	//buf=llo
	fmt.Println(buf.String())

	//b=he
	fmt.Println(string(b))
}

// func (b *Buffer) Reset()
func Reset() {
	fmt.Println("----------------> Reset")

	buf := bytes.NewBufferString("hello")

	//buf=hello
	fmt.Println(buf.String())

	buf.Reset()

	//buf=
	fmt.Println(buf.String())
}

// func (b *Buffer) Truncate(n int)
func Truncate() {
	fmt.Println("----------------> Truncate")

	buf := bytes.NewBufferString("hello")

	//buf=hello
	fmt.Println(buf.String())

	buf.Truncate(3)
	//buf=hel
	fmt.Println(buf.String())

	buf.Truncate(2)
	//buf=he
	fmt.Println(buf.String())

	buf.Truncate(0)
	//buf=
	fmt.Println(buf.String())
}
