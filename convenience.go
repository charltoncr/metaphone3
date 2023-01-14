// Convenience functions and methods that use Metaphone3.
// Created 2022-12-16 by Ron Charlton and placed in the public domain.
//
// $Id: convenience.go,v 1.37 2023-01-14 15:29:14-05 ron Exp $

package metaphone3

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

// MetaphMap defines a MetaphMap type to store a word list mapped by met to
// mapper.
type MetaphMap struct {
	mapper map[string][]string
	met    *Metaphone3
}

// NewMetaphMap returns a MetaphMap made from wordlist and a maximum
// length for the Metaphone3 return values.
// The MetaphMap can be used with MatchWord to find all words in the
// MetaphMap that sound like a given word or misspelling.
// Argument maxLen is 4 in the original Double Metaphone algorithm.
// Letter case is ignored in the words in wordlist, as are non-alphabetic
// characters.
func NewMetaphMap(wordlist []string, maxLen int) *MetaphMap {
	return NewMetaphMapExact(wordlist, maxLen, false, false)
}

// NewMetaphMapExact is like NewMetaphMap but allows control of whether
// vowels after the first character are encoded, and whether consonants are
// encoded more precisely.
func NewMetaphMapExact(wordlist []string, maxLen int,
	encodeVowels, encodeExact bool) *MetaphMap {
	MMap := make(map[string][]string)
	meta := NewMetaphone3(maxLen)
	meta.SetEncodeVowels(encodeVowels)
	meta.SetEncodeExact(encodeExact)
	for _, word := range wordlist {
		m, m2 := meta.Encode(word)
		if len(m) > 0 {
			MMap[m] = append(MMap[m], word)
		}
		if len(m2) > 0 {
			MMap[m2] = append(MMap[m2], word)
		}
	}
	return &MetaphMap{
		mapper: MMap,
		met:    meta,
	}
}

// AddWordsToMap adds the words in wordlist to an existing MetaphMap.  This
// can be useful if you have a general word list and a specific user word list
// to combine into one MetaphMap.
// Letter case is ignored in the words in wordlist, as are non-alphabetic
// characters.
func (metaph *MetaphMap) AddWordsToMap(wordlist []string) {
	for _, word := range wordlist {
		m, m2 := metaph.met.Encode(word)
		if len(m) > 0 {
			metaph.mapper[m] = append(metaph.mapper[m], word)
		}
		if len(m2) > 0 {
			metaph.mapper[m2] = append(metaph.mapper[m2], word)
		}
	}
}

// NewMetaphMapFromFile returns a MetaphMap made from a file containing a
// word list, and using a maximum length for the Metaphone3 return values.
// The file can be a gzipped file with its name ending with ".gz".
// The MetaphMap can be used with MatchWord to find all words in the
// MetaphMap that sound like a given word or misspelling.
// Argument maxLen is 4 in the original Double Metaphone algorithm.
// Letter case and non-alphabetic characters in the file are ignored.
func NewMetaphMapFromFile(fileName string, maxLen int) (
	metaph *MetaphMap, err error) {
	return NewMetaphMapFromFileExact(fileName, maxLen, false, false)
}

// NewMetaphMapFromFileExact is like NewMetaphMapFromFile but allows control
// of whether vowels after the first character are encoded, and whether
// consonants are encoded more precisely.
func NewMetaphMapFromFileExact(fileName string, maxLen int,
	encodeVowels, encodeExact bool) (metaph *MetaphMap, err error) {
	var lines []string
	lines, err = getWordsFromFile(fileName)
	if err != nil {
		return
	}
	return NewMetaphMapExact(lines, maxLen, encodeVowels, encodeExact), err
}

// AddWordsFromFile adds words from a file to an existing MetaphMap.  This
// can be useful if you have a general word list and a specific user word list
// to combine into one MetaphMap.
// Letter case is ignored in the file, as are non-alphabetic characters.
func (metaph *MetaphMap) AddWordsFromFile(fileName string) error {
	lines, err := getWordsFromFile(fileName)
	if err == nil {
		metaph.AddWordsToMap(lines)
	}
	return err
}

// Get a string slice of lines from a file.  Return the lines or an error code.
func getWordsFromFile(fileName string) (lines []string, err error) {
	var b []byte
	var r io.Reader
	var fp *os.File

	if fp, err = os.Open(fileName); err != nil {
		err = fmt.Errorf("trying to open file %s: %v", fileName, err)
		return
	}
	defer fp.Close()
	r = fp
	if strings.HasSuffix(fileName, ".gz") {
		if r, err = gzip.NewReader(r); err != nil {
			err = fmt.Errorf(
				"trying to make a gzip reader for file %s: %v", fileName, err)
			return
		}
	}
	if b, err = io.ReadAll(r); err != nil {
		err = fmt.Errorf("trying to read file %s: %v", fileName, err)
		return
	}
	lines = strings.Split(string(noCRs(b)), "\n")
	return
}

// Len returns the number of keys in metaph.
func (metaph *MetaphMap) Len() int {
	return len(metaph.mapper)
}

// MatchWord returns all words in metaph that sound like word.
// Letter case and non-alphabetic characters in word are ignored.
func (metaph *MetaphMap) MatchWord(word string) (output []string) {
	m, m2 := metaph.met.Encode(word)
	if len(m) > 0 {
		output = metaph.mapper[m]
	}
	if len(m2) > 0 {
		output = append(output, metaph.mapper[m2]...)
	}
	output = removeDups(output)
	return
}

// removeDups removes duplicate strings in s.
func removeDups(s []string) (out []string) {
	m := make(map[string]struct{})
	for _, w := range s {
		m[w] = struct{}{}
	}
	for o := range m {
		out = append(out, o)
	}
	return
}

// noCRs removes CRs. Assumes UTF-8, ANSI, iso-8859-n or ASCII encoding.
func noCRs(b []byte) []byte {
	// bytes.Map would use an additional len(b) buffer and be slower.
	from := bytes.IndexByte(b, '\r')
	if from < 0 {
		return b
	}
	to := from

	for from < len(b) {
		if b[from] != '\r' {
			b[to] = b[from]
			to++
		}
		from++
	}
	return b[:to]
}
