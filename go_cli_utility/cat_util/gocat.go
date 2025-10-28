package gocat

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// RunCat returns file content as a string.
// Options: show line numbers (-n) and limit lines (-l)
func RunCat(fs *flag.FlagSet, number bool, limit int) (string, error) {
	if fs.NArg() == 0 {
		return "", errors.New("usage: gocat cat [-n] [-l num] <file>")
	}

	file, err := os.Open(fs.Arg(0))
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		text := scanner.Text()
		if number {
			builder.WriteString(fmt.Sprintf("%6d  %s\n", lineNum, text))
		} else {
			builder.WriteString(text + "\n")
		}
		lineNum++
		if limit > 0 && lineNum > limit {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return builder.String(), nil
}

// RunHead returns the first N lines of a file as a string.
func RunHead(fs *flag.FlagSet, lines int) (string, error) {
	if fs.NArg() == 0 {
		return "", errors.New("usage: gocat head [-n num] <file>")
	}

	file, err := os.Open(fs.Arg(0))
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		builder.WriteString(scanner.Text() + "\n")
		count++
		if count >= lines {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return builder.String(), nil
}

// RunTail returns the last N lines of a file as a string.
func RunTail(fs *flag.FlagSet, lines int) (string, error) {
	if fs.NArg() == 0 {
		return "", errors.New("usage: gocat tail [-n num] <file>")
	}

	data, err := os.ReadFile(fs.Arg(0))
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	allLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	total := len(allLines)
	start := 0
	if total > lines {
		start = total - lines
	}

	return strings.Join(allLines[start:], "\n") + "\n", nil
}
