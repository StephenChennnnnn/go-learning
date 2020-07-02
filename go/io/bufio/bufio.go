package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Reader struct {
	buf          []byte
	rd           io.Reader
	r, w         int
	err          error
	lastByte     int
	lastRuneSize int
}

type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

func (b *Reader) Read(p []byte) (n int, err error) {
	n = len(p)
	if n == 0 {
		return 0, b.readErr()
	}
	if b.r == b.w {
		if b.err != nil {
			return 0, b.readErr()
		}
		if len(p) >= len(b.buf) {
			n, b.err = b.rd.Read(p)
			if n < 0 {
				panic("")
			}
			if n > 0 {
				b.lastByte = int(p[n-1])
				b.lastRuneSize = -1
			}
			return n, b.readErr()
		}
		b.r = 0
		b.w = 0
		n, b.err = b.rd.Read(b.buf)
		if n < 0 {
			panic("")
		}
		if n == 0 {
			return 0, b.readErr()
		}
		b.w += n
	}

	n = copy(p, b.buf[b.r:b.w])
	b.r += n
	b.lastByte = int(b.buf[b.r-1])
	b.lastRuneSize = -1
	return n, nil
}

func (b *Reader) readErr() error {
	err := b.err
	b.err = nil
	return err
}

func (b *Writer) Write(p []byte) (nn int, err error) {
	for len(p) > b.Available() && b.err == nil {
		var n int
		if b.Buffered() == 0 {
			n, b.err = b.wr.Write(p)
		} else {
			n = copy(b.buf[b.n:], p)
			b.n += n
			b.Flush()
		}
		nn += n
		p = p[n:]
	}
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, b.err
}

func (b *Writer) Available() int {
	return len(b.buf) - b.n
}

func (b *Writer) Buffered() int {
	return b.n
}

func (b *Writer) Flush() error {
	if b.err != nil {
		return b.err
	}
	if b.n == 0 {
		return nil
	}
	n, err := b.wr.Write(b.buf[0:b.n])
	if n < b.n && err == nil {
		err = io.ErrShortWrite
	}
	if err != nil {
		if n > 0 && n < b.n {
			copy(b.buf[0:b.n-n], b.buf[n:b.n])
		}
		b.n -= n
		b.err = err
		return err
	}
	b.n = 0
	return nil
}

func main() {
	fileName := "/Users/chenxinyuan/TTTT.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	b2 := bufio.NewReader(os.Stdin)
	s2, _ := b2.ReadString('\n')
	fmt.Println(s2)
}
