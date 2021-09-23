package murmur3

import (
	"bytes"
	"encoding/binary"
	"io"
)

type ByteReader struct {
	buf    []byte
	index  int
	length int
	bytes.Buffer
}

func NewByteReader(b []byte) *ByteReader {
	return &ByteReader{
		buf:    b,
		index:  0,
		length: len(b),
	}
}

func (b *ByteReader) ReadByte() (buf byte, err error) {
	if b.index >= b.length {
		return 0, io.EOF
	}
	buf, err = b.buf[b.index], nil
	b.index++
	return buf, err
}

func (b *ByteReader) ReadBytes(n int) (buf []byte, err error) {
	if b.index+n > b.length {
		return nil, io.EOF
	}
	buf, err = b.buf[b.index:b.index+n], nil
	b.index += n
	return buf, err
}

func (b *ByteReader) IsDrained() bool {
	return b.index >= b.length
}

func (b *ByteReader) ReadUint32Anyway() (u uint32, err error) {
	if b.IsDrained() {
		return 0, io.EOF
	}
	t := uint32(0)
	buf := b.buf[b.index:]
	for i := 0; i < len(buf); i++ {
		t |= (uint32(buf[i]) << (i * 8))
	}
	return t, nil
}

func (b *ByteReader) ReadUint32() (u uint32, err error) {
	buf, err := b.ReadBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(buf), nil
}
