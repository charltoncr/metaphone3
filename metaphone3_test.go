// metaphone3_test.go - test metaphone3/metaphone3.go by comparing lines from
// testInputData.txt.gz as input with testWantData.txt.gz
// that contains output from a test program using metaphone3.go.
// testInputData.txt.gz contains 171,109 words from an Aspell wordlist.
// testWantData.txt.gz contains the same number of lines output by
// metaphone3.go and metaphone3_Test.go.  The test data is not from an
// external reference source but rather from metaphone3.go results after
// metaphone3.go's near agreement with metaph.cpp output.  There were
// 27,481 differences between metaph.cpp and metaphone3.go output,
// with the differences looking like reasonable improvements in metaphone3.go.
// Author: Ron Charlton (copied from metaphone_test.go)
// Date:   2023-01-11
// This file is public domain.  Public domain is per CC0 1.0; see
// https://creativecommons.org/publicdomain/zero/1.0/ for information.

package metaphone3

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// $Id: metaphone3_test.go,v 1.14 2023-01-20 03:13:03-05 ron Exp $

const maxlength = 6

func TestMetaphone3(t *testing.T) {
	var words, want []string
	var err error

	if words, err = readFileLines("testInputData.txt.gz"); err == nil {
		want, err = readFileLines("testWantData.txt.gz")
	}
	if err != nil {
		t.Fatalf("%v", err)
	}

	idx := 0
	met := NewMetaphone3(maxlength)
	for _, word := range words {
		if len(word) > 0 {
			m, m2 := met.Encode(word)
			got := fmt.Sprintf("'%s' '%s' %s", m, m2, word)
			if got != want[idx] {
				t.Errorf("At line %d got: <%s>;  want: <%s>", idx+1, got, want[idx])
			}
		}
		idx++
	}
}

// readFileLines reads a (gzipped) text file and returns its lines.
// Parameter name is the file's name.  Carriage returns are skipped.
func readFileLines(name string) (lines []string, err error) {
	var b []byte
	var r io.Reader
	var fp *os.File

	if fp, err = os.Open(name); err != nil {
		err = fmt.Errorf("trying to open file %s: %v", name, err)
		return
	}
	defer fp.Close()
	r = fp
	if strings.HasSuffix(name, ".gz") {
		if r, err = gzip.NewReader(r); err != nil {
			err = fmt.Errorf(
				"trying to make a gzip reader for file %s: %v", name, err)
			return
		}
	}
	if b, err = io.ReadAll(r); err != nil {
		err = fmt.Errorf("trying to read word list file %s: %v", name, err)
		return
	}
	lines = strings.Split(string(noCRs(b)), "\n")
	return
}

func TestConvenience(t *testing.T) {
	metaph, err := NewMetaphMapFromFile("testInputData.txt.gz", maxlength)
	if err != nil {
		t.Fatalf("%v", err)
	}
	words := metaph.MatchWord("knewmoanya")
	if len(words) != 11 {
		t.Errorf("got: %d;  want: 11", len(words))
	}
}

func BenchmarkEncode(b *testing.B) {
	met := NewMetaphone3(6)
	met.SetMaxLength(52)
	str := "abcdefghijklmnopqrstuvwxyz"
	b.SetBytes(26)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		met.Encode(str)
	}
}
func BenchmarkDoFile(b *testing.B) {
	words, err := readFileLines("testInputData.txt.gz")
	var size int
	for _, word := range words {
		size += len(word)
	}
	if err != nil {
		b.Fatalf("%v", err)
	}
	met := NewMetaphone3(maxlength)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, word := range words {
			met.Encode(word)
		}
	}
}
func BenchmarkNewMetaphMap(b *testing.B) {
	words, err := readFileLines("testInputData.txt.gz")
	var size int
	for _, word := range words {
		size += len(word)
	}
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewMetaphMap(words, maxlength)
	}
}
