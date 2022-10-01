package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

var (
	binName = "wordcounter"
	files   = []string{"file1", "file2", "file3"}
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	fmt.Println("Building tool...")
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()
	fmt.Println("Cleanning up...")
	os.Remove(binName)

	os.Exit(result)
}

func TestCountWordsFromFiles(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)
	args := append([]string{"-f"}, files...)
	cmd := exec.Command(cmdPath, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	expected := "1. file1: 5\n2. file2: 7\n3. file3: 6\n"
	if string(out) != expected {
		t.Errorf("Expected %q, got %q instead\n", expected, string(out))
	}
}
