package main

import "fmt"

type Bitmap struct {
	b      []byte
	length int
}

func NewBitmap(length int) *Bitmap {
	if length < 1 {
		panic(fmt.Sprintf("size %d must be greater than 0", length))
	}
	size := length/8 + 1
	return &Bitmap{
		b:      make([]byte, size),
		length: size,
	}
}

func NewBitmapAllOn(length int) *Bitmap {
	b := NewBitmap(length)
	for i := 0; i < len(b.b); i++ {
		b.b[i] = 255
	}
	return b
}

func (b *Bitmap) Length() int {
	return b.length
}

func (b *Bitmap) Set(idx int) {
	bidx, num, mask := b.seek(idx)
	b.b[bidx] = num | mask
}

func (b *Bitmap) Complement() {
	for i := 0; i < len(b.b); i++ {
		b.b[i] = ^b.b[i]
	}
}

func (b *Bitmap) IsSet(idx int) bool {
	_, num, mask := b.seek(idx)
	return num&mask == mask
}

func (b *Bitmap) Union(in *Bitmap) {
	if b.Length() != in.Length() {
		panic(fmt.Sprintf("sizes dont match %d %d", b.Length(), in.Length()))
	}
	for i := 0; i < b.Length(); i++ {
		b.b[i] = b.b[i] | in.b[i]
	}
}

func (b *Bitmap) Intersect(in *Bitmap) {
	if b.Length() != in.Length() {
		panic(fmt.Sprintf("sizes dont match %d %d", b.Length(), in.Length()))
	}
	for i := 0; i < b.Length(); i++ {
		b.b[i] = b.b[i] & in.b[i]
	}
}

func (b *Bitmap) seek(idx int) (int, uint8, uint8) {
	if idx < 0 {
		panic(fmt.Sprintf("index %d is out of bounds %d", idx, b.Length()))
	}
	bidx := idx / 8
	num := uint8(b.b[bidx])
	off := uint8(7 - (idx % 8))
	mask := uint8(1) << off

	return bidx, num, mask
}
