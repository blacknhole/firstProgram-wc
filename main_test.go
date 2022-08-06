package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word word word word\n")
	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word word word word\nline\nline")
	exp := 3

	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}
