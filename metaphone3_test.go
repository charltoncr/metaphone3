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
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// $Id: metaphone3_test.go,v 1.39 2024-06-16 15:13:33-04 ron Exp $

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

	met := NewMetaphone3(maxlength)
	for idx, word := range words {
		if len(word) > 0 {
			m, m2 := met.Encode(word)
			got := fmt.Sprintf("'%s' '%s' %s", m, m2, word)
			if got != want[idx] {
				t.Errorf("At line %d got: <%s>;  want: <%s>", idx+1, got, want[idx])
			}
		}
	}
}

// readFileLines reads a (gzipped) text file and returns its lines.
// Carriage returns are omitted.
func readFileLines(fileName string) (lines []string, err error) {
	var b []byte
	var r io.Reader
	var gr *gzip.Reader
	var fp *os.File

	if fp, err = os.Open(fileName); err != nil {
		err = fmt.Errorf("opening file %s: %v", fileName, err)
		return
	}
	defer func() {
		err = errors.Join(err, fp.Close())
	}()
	r = fp
	if strings.HasSuffix(fileName, ".gz") {
		if gr, err = gzip.NewReader(fp); err != nil {
			err = fmt.Errorf(
				"making a gzip reader for file %s: %v", fileName, err)
			return
		}
		defer func() {
			err = errors.Join(err, gr.Close())
		}()
		r = gr
	}
	if b, err = io.ReadAll(r); err != nil {
		err = fmt.Errorf("reading word list file %s: %v", fileName, err)
		return
	}
	lines = strings.Split(strings.ReplaceAll(string(b), "\r", ""), "\n")
	if len(b) > 0 && b[len(b)-1] == '\n' {
		lines = lines[:len(lines)-1]
	}
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
	found := false
	for _, w := range words {
		if w == "pneumonia" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("matching 'knewmoanya' didn't find 'pneumonia'")
	}
}

// 'go test -fuzz=FuzzEncode -fuzztime 30s' to run for 30 seconds.
// Ctrl+C to stop test if fuzztime is not specified.
func FuzzEncode(f *testing.F) {
	testCases := []string{"pneumonia", "knewmoanya", "ß", "Ç", "Ð", "Ñ",
		"Þ", "\u008a", "\u008e", "12345!"}
	for _, tc := range testCases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	meta := NewMetaphone3(200)
	f.Fuzz(func(t *testing.T, orig string) {
		meta.Encode(orig)
	})
}

func BenchmarkEncode(b *testing.B) {
	met := NewMetaphone3(maxlength)
	met.SetMaxLength(52)
	str := "abcdefghijklmnopqrstuvwxyz"
	b.SetBytes(int64(len(str)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		met.Encode(str)
	}
}

func getWords(b *testing.B) (words []string, size int) {
	var err error
	words, err = readFileLines("testInputData.txt.gz")
	if err != nil {
		b.Fatalf("%v", err)
	}
	for _, word := range words {
		size += len(word)
	}
	return words, size
}

func BenchmarkDoFile(b *testing.B) {
	words, size := getWords(b)
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
	words, size := getWords(b)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewMetaphMap(words, maxlength)
	}
}

func BenchmarkLookupWord(b *testing.B) {
	words, _ := getWords(b)
	m := NewMetaphMap(words, maxlength)
	s := "knewmoanya"
	b.SetBytes(int64(len(s)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.MatchWord(s)
	}
}
