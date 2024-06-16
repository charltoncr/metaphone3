// Convenience functions and methods that use Metaphone3.
// Created 2022-12-16 by Ron Charlton.
// This file is public domain per CC0 1.0, see
// https://creativecommons.org/publicdomain/mark/1.0/
//
// $Id: convenience.go,v 3.10 2024-06-16 15:15:17-04 ron Exp $

package metaphone3

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// MetaphMap stores a word list as a Go map of words indexed by keys from
// metaphone3.Encode.
type MetaphMap struct {
	mapper map[string][]string
	met    *Metaphone3
}

// NewMetaphMap returns a MetaphMap made from wordlist and a maximum
// length for metaphone3.Encode return values.
// The MetaphMap can be used with MatchWord to find all words in the
// MetaphMap that sound like a given word or misspelling.
// Argument maxLen is 4 in the Double Metaphone algorithm.
// Letter case is ignored in mapping the words in wordlist, as are
// non-alphabetic characters.
func NewMetaphMap(wordlist []string, maxLen int) *MetaphMap {
	return NewMetaphMapExact(wordlist, maxLen, false, false)
}

// NewMetaphMapExact is like NewMetaphMap but allows control of whether
// vowels after the first character are encoded, and whether consonants are
// encoded more selectively.
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
// Letter case is ignored in mapping the words in wordlist, as are
// non-alphabetic characters.
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
// Argument maxLen is 4 in the Double Metaphone algorithm.
// Letter case is ignored in mapping the words in the file, as are
// non-alphabetic characters.  The default values of encodeVowels and
// encodeExact are false.
func NewMetaphMapFromFile(fileName string, maxLen int) (
	metaph *MetaphMap, err error) {
	return NewMetaphMapFromFileExact(fileName, maxLen, false, false)
}

// NewMetaphMapFromFileExact is like NewMetaphMapFromFile but allows control
// of whether vowels after the first character are encoded, and whether
// consonants are encoded more selectively.
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
// Letter case is ignored in mapping the words in the file, as are
// non-alphabetic characters.
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
	var gr *gzip.Reader
	var fp *os.File

	if fp, err = os.Open(fileName); err != nil {
		err = fmt.Errorf("trying to open file %s: %v", fileName, err)
		return
	}
	defer func() {
		err = errors.Join(err, fp.Close())
	}()
	r = fp
	if strings.HasSuffix(fileName, ".gz") {
		if gr, err = gzip.NewReader(fp); err != nil {
			err = fmt.Errorf(
				"trying to make a gzip reader for file %s: %v", fileName, err)
			return
		}
		defer func() {
			err = errors.Join(err, gr.Close())
		}()
		r = gr
	}
	if b, err = io.ReadAll(r); err != nil {
		err = fmt.Errorf("trying to read file %s: %v", fileName, err)
		return
	}
	s := strings.ReplaceAll(string(b), "\r", "")
	if len(s) > 0 {
		lines = strings.Split(s, "\n")
		if s[len(s)-1] == '\n' {
			lines = lines[:len(lines)-1]
		}
	}
	return
}

// Len returns the number of keys in metaph.
func (metaph *MetaphMap) Len() int {
	return len(metaph.mapper)
}

// MatchWord returns all words in metaph that sound like word.
// The returned words are sorted by order of their approximate frequency of
// occurrence in English, so more likely choices appear earlier.
// Letter case and non-alphabetic characters in word are ignored.
func (metaph *MetaphMap) MatchWord(word string) []string {
	var output []string
	m, m2 := metaph.met.Encode(word)
	if len(m) > 0 {
		output = append(output, metaph.mapper[m]...) // copy of metaph.mapper[m]
	}
	if len(m2) > 0 {
		output = append(output, metaph.mapper[m2]...)
	}
	return RankWords(removeDups(output))
}

// mySort stable sorts words into alphabetical order while ignoring case.
func mySort(words []string) (output []string) {
	LC := strings.ToLower // alias
	output = append(output, words...)
	less := func(i, j int) bool {
		return LC(output[i]) < LC(output[j])
	}
	sort.SliceStable(output, less)
	return
}

// gunzipBytes accepts a gzip compressed byte slice and returns a gunzip'ed
// byte slice in b.  Any error is returned in err.
func gunzipBytes(g []byte) (b []byte, err error) {
	gzr, err1 := gzip.NewReader(bytes.NewReader(g))
	if err1 != nil {
		err = fmt.Errorf("making a gzip reader: %v", err1)
		return
	}
	defer gzr.Close()
	if b, err = io.ReadAll(gzr); err != nil {
		err = fmt.Errorf("reading gzip'ed bytes: %v", err)
	}
	return
}

// the frequency of occurrence as integer for each word: map[word]frequency
var freqs = map[string]uint8{}

func init() {
	// get ready for RankWords (gzWordFrequencies is in wordFreq.go)
	b, err := gunzipBytes(gzWordFrequencies)
	if err != nil {
		panic(err)
	}
	s := strings.ReplaceAll(string(b), "\r", "")
	lines := strings.Split(s, "\n")
	if s[len(s)-1] == '\n' {
		lines = lines[:len(lines)-1]
	}
	var fr uint8 = 200
	for _, line := range lines {
		if strings.HasPrefix(line, ".COMMENT") {
			continue
		}
		if strings.HasPrefix(line, ".FREQ ") {
			t := strings.Split(line, " ")
			if len(t) == 2 {
				f, err := strconv.Atoi(t[1])
				if err != nil || f < 0 || f > 255 {
					f = 200
				}
				fr = uint8(f)
			}
			continue
		}
		freqs[line] = fr
	}
}

// RankWords returns words sorted by order of their approximate frequency of
// occurrence in English, so more common words appear earlier in output.
// The sort is stable.
func RankWords(words []string) (output []string) {
	LC := strings.ToLower  // alias
	output = mySort(words) // for consistent output order
	less := func(i, j int) bool {
		ia := freqs[output[i]]
		ja := freqs[output[j]]
		if ia == ja {
			return LC(output[i]) < LC(output[j])
		}
		if ia == 0 { // output[i] not found in freqs, therefore not common word
			ia = 100
		}
		if ja == 0 { // ditto
			ja = 100
		}
		return ia < ja
	}
	sort.SliceStable(output, less)
	return
}

// GetAllWords returns a string slice of the American English words used
// by RankWords for evaluation.  The slice contains the most commonly
// occurring 70% of American English words, taken from the SCOWL word lists.
func GetAllWords() (output []string) {
	for w := range freqs {
		output = append(output, w)
	}
	output = mySort(output)
	return
}

// removeDups creates a new string slice from s without duplicated strings.
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
