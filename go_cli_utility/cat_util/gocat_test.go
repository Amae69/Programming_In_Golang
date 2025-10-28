package gocat_test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"gocat"
)

func createTempFile(t *testing.T, content string) *os.File {
	t.Helper()
	tmp, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}
	tmp.WriteString(content)
	tmp.Close()
	return tmp
}

func TestRunCat(t *testing.T) {
	file := createTempFile(t, "line1\nline2\nline3\n")
	defer os.Remove(file.Name())

	fs := flag.NewFlagSet("cat", flag.ExitOnError)
	fs.Parse([]string{file.Name()})

	out, err := gocat.RunCat(fs, true, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "line1") {
		t.Errorf("expected 'line1' in output, got %q", out)
	}
}

func TestRunHead(t *testing.T) {
	file := createTempFile(t, "a\nb\nc\nd\n")
	defer os.Remove(file.Name())

	fs := flag.NewFlagSet("head", flag.ExitOnError)
	fs.Parse([]string{file.Name()})

	out, err := gocat.RunHead(fs, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "a\nb") {
		t.Errorf("expected 'a' and 'b' in output, got %q", out)
	}
}

func TestRunTail(t *testing.T) {
	file := createTempFile(t, "1\n2\n3\n4\n5\n")
	defer os.Remove(file.Name())

	fs := flag.NewFlagSet("tail", flag.ExitOnError)
	fs.Parse([]string{file.Name()})

	out, err := gocat.RunTail(fs, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "3\n4\n5") {
		t.Errorf("expected last three lines in output, got %q", out)
	}
}
