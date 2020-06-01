package main

import (
	"io"
)

// Writer implements Writer interface and converts input to Pig Latin on Write.
type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	pl, err := ConvertSentence(string(p))
	n, err = w.w.Write([]byte(pl))
	return
}

func (w *Writer) WriteString(s string) (n int, err error) {
	pl, err := ConvertSentence(s)
	n, err = w.w.Write([]byte(pl))
	return
}
