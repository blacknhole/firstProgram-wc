package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word word word word\n")
	exp := 4

	res := count(b)

	if res != exp {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}
