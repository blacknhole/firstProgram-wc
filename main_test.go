package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word word word word\n")
	exp := 4

	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word word word word\nline\nline\n")
	exp := 3

	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word word word word\nline\nline\n")
	exp := 30

	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}
