// metaphone3.go - ported to the Go programming language by Ron Charlton
// on 2023-01-05 from the original Java code at
// https://github.com/OpenRefine/OpenRefine/blob/master/main/src/com/google/refine/clustering/binning/Metaphone3.java
//
// $Id: metaphone3.go,v 4.46 2023-01-25 10:36:42-05 ron Exp $
//
// This open source Go file is based on Metaphone3.java 2.1.3 that is
// copyright 2010 by Laurence Philips, and is also open source.
// Ron Charlton asserts no additional copyright to this Go implementation
// of Metaphone3.
//
// A later metaphone 3 version with better proper name encoding and other
// corrections, and in a variety of programming languages (but not Go), is
// available online for $240 as of 2023-01-10.

// Metaphone3.java copyright and header comments follow:
/*
Copyright 2010, Lawrence Philips
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

    * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
    * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

============================================================================

Metaphone 3
VERSION 2.1.3

by Lawrence Philips

Metaphone 3 is designed to return an *approximate* phonetic key (and an alternate
approximate phonetic key when appropriate) that should be the same for English
words, and most names familiar in the United States, that are pronounced *similarly*.
The key value is *not* intended to be an *exact* phonetic, or even phonemic,
representation of the word. This is because a certain degree of 'fuzziness' has
proven to be useful in compensating for variations in pronunciation, as well as
misheard pronunciations. For example, although americans are not usually aware of it,
the letter 's' is normally pronounced 'z' at the end of words such as "sounds".

The 'approximate' aspect of the encoding is implemented according to the following rules:

(1) All vowels are encoded to the same value - 'A'. If the parameter m.encodeVowels
is set to false, only *initial* vowels will be encoded at all. If m.encodeVowels is set
to true, 'A' will be encoded at all places in the word that any vowels are normally
pronounced. 'W' as well as 'Y' are treated as vowels. Although there are differences in
the pronunciation of 'W' and 'Y' in different circumstances that lead to their being
classified as vowels under some circumstances and as consonants in others, for the purposes
of the 'fuzziness' component of the Soundex and Metaphone family of algorithms they will
be always be treated here as vowels.

(2) Voiced and un-voiced consonant pairs are mapped to the same encoded value. This
means that:
'D' and 'T' -> 'T'
'B' and 'P' -> 'P'
'G' and 'K' -> 'K'
'Z' and 'S' -> 'S'
'V' and 'F' -> 'F'

- In addition to the above voiced/unvoiced rules, 'CH' and 'SH' -> 'X', where 'X'
represents the "-SH-" and "-CH-" sounds in Metaphone 3 encoding.

- Also, the sound that is spelled as "TH" in English is encoded to '0' (zero symbol). (Although
Americans are not usually aware of it, "TH" is pronounced in a voiced (e.g. "that") as
well as an unvoiced (e.g. "theater") form, which are naturally mapped to the same encoding.)

The encodings in this version of Metaphone 3 are according to pronunciations common in the
United States. This means that they will be inaccurate for consonant pronunciations that
are different in the United Kingdom, for example "tube" -> "CHOOBE" -> XAP rather than american TAP.

Metaphone 3 was preceded by by Soundex, patented in 1919, and Metaphone and Double Metaphone,
developed by Lawrence Philips. All of these algorithms resulted in a significant number of
incorrect encodings. Metaphone3 was tested against a database of about 100 thousand English words,
names common in the United States, and non-English words found in publications in the United States,
with an emphasis on words that are commonly mispronounced, prepared by the Moby Words website,
but with the Moby Words 'phonetic' encodings algorithmically mapped to Double Metaphone encodings.
Metaphone3 increases the accuracy of encoding of english words, common names, and non-English
words found in American publications from the 89% for Double Metaphone, to over 98%.

DISCLAIMER:
Anthropomorphic Software LLC claims only that Metaphone 3 will return correct encodings,
within the 'fuzzy' definition of correct as above, for a very high percentage of correctly
spelled English and commonly recognized non-English words. Anthropomorphic Software LLC
warns the user that a number of words remain incorrectly encoded, that misspellings may not
be encoded 'properly', and that people often have differing ideas about the pronunciation
of a word. Therefore, Metaphone 3 is not guaranteed to return correct results every time, and
so a desired target word may very well be missed. Creators of commercial products should
keep in mind that systems like Metaphone 3 produce a 'best guess' result, and should
condition the expectations of end users accordingly.

METAPHONE3 IS PROVIDED "AS IS" WITHOUT
WARRANTY OF ANY KIND. LAWRENCE PHILIPS AND ANTHROPOMORPHIC SOFTWARE LLC
MAKE NO WARRANTIES, EXPRESS OR IMPLIED, THAT IT IS FREE OF ERROR,
OR ARE CONSISTENT WITH ANY PARTICULAR STANDARD OF MERCHANTABILITY,
OR THAT IT WILL MEET YOUR REQUIREMENTS FOR ANY PARTICULAR APPLICATION.
LAWRENCE PHILIPS AND ANTHROPOMORPHIC SOFTWARE LLC DISCLAIM ALL LIABILITY
FOR DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES RESULTING FROM USE
OF THIS SOFTWARE.

@author Lawrence Philips

Metaphone 3 is designed to return an *approximate* phonetic key (and an alternate
approximate phonetic key when appropriate) that should be the same for English
words, and most names familiar in the United States, that are pronounced "similarly".
The key value is *not* intended to be an exact phonetic, or even phonemic,
representation of the word. This is because a certain degree of 'fuzziness' has
proven to be useful in compensating for variations in pronunciation, as well as
misheard pronunciations. For example, although Americans are not usually aware of it,
the letter 's' is normally pronounced 'z' at the end of words such as "sounds".

The 'approximate' aspect of the encoding is implemented according to the following rules:

(1) All vowels are encoded to the same value - 'A'. If the parameter m.encodeVowels
is set to false, only *initial* vowels will be encoded at all. If m.encodeVowels is set
to true, 'A' will be encoded at all places in the word that any vowels are normally
pronounced. 'W' as well as 'Y' are treated as vowels. Although there are differences in
the pronunciation of 'W' and 'Y' in different circumstances that lead to their being
classified as vowels under some circumstances and as consonants in others, for the purposes
of the 'fuzziness' component of the Soundex and Metaphone family of algorithms they will
be always be treated here as vowels.

(2) Voiced and un-voiced consonant pairs are mapped to the same encoded value. This
means that:
'D' and 'T' -> 'T'
'B' and 'P' -> 'P'
'G' and 'K' -> 'K'
'Z' and 'S' -> 'S'
'V' and 'F' -> 'F'

- In addition to the above voiced/unvoiced rules, 'CH' and 'SH' -> 'X', where 'X'
represents the "-SH-" and "-CH-" sounds in Metaphone 3 encoding.

- Also, the sound that is spelled as "TH" in English is encoded to '0' (zero symbol). (Although
Americans are not usually aware of it, "TH" is pronounced in a voiced (e.g. "that") as
well as an unvoiced (e.g. "theater") form, which are naturally mapped to the same encoding.)

In the "Exact" encoding, voiced/unvoiced pairs are *not* mapped to the same encoding, except
for the voiced and unvoiced versions of 'TH', sounds such as 'CH' and 'SH', and for 'S' and 'Z',
so that the words whose metaph keys match will in fact be closer in pronunciation that with the
more approximate setting. Keep in mind that encoding settings for search strings should always
be exactly the same as the encoding settings of the stored metaph keys in your database!
Because of the considerably increased accuracy of Metaphone3, it is now possible to use this
setting and have a very good chance of getting a correct encoding.

In the Encode Vowels encoding, all non-initial vowels and diphthongs will be encoded to
'A', and there will only be one such vowel encoding character between any two consonants.
It turns out that there are some surprising wrinkles to encoding non-initial vowels in
practice, pre-eminently in inversions between spelling and pronunciation such as e.g.
"wrinkle" => 'RANKAL', where the last two sounds are inverted when spelled.

The encodings in this version of Metaphone 3 are according to pronunciations common in the
United States. This means that they will be inaccurate for consonant pronunciations that
are different in the United Kingdom, for example "tube" -> "CHOOBE" -> XAP rather than American TAP.
*/
// End of Metaphone3.java copyright and header comments.

package metaphone3

import (
	"strings"
)

var ()

// Metaphone3 defines a type and return length for Metaphone3, as well as
// recording whether vowels after the first character are encoded and whether
// consonants are encoded more exactly.
type Metaphone3 struct {
	/** Length of word sent in to be encoded, as
	* measured at beginning of encoding. */
	length int

	/** Maximum length of encoded key string. */
	maxLength int

	/** Flag whether or not to encode non-initial vowels. */
	encodeVowels bool

	/** Flag whether or not to encode consonants as exactly
	* as possible. */
	encodeExact bool

	/** Internal copy of word to be encoded, allocated separately
	* from string pointed to in incoming parameter. */
	inWord []rune

	/** Running copy of primary key. */
	primary []rune

	/** Running copy of secondary key. */
	secondary []rune

	/** Index of character in m.inWord currently being
	* encoded. */
	current int

	/** Index of last character in m.inWord. */
	last int

	/** Flag that an AL inversion has already been done. */
	flag_AL_inversion bool
}

// NewMetaphone3 returns a Metaphone3 instance with a maximum
// length for the Metaphone3 return values.
// Argument maxLen is 4 in the Double Metaphone algorithm.
func NewMetaphone3(maxLen int) *Metaphone3 {
	return &Metaphone3{
		maxLength: maxLen,
	}
}

// SetEncodeVowels determines whether or not vowels after the first character
// are encoded.
func (m *Metaphone3) SetEncodeVowels(b bool) {
	m.encodeVowels = b
}

// SetEncodeExact determines whether or not consonants are encoded more exactly.
func (m *Metaphone3) SetEncodeExact(b bool) {
	m.encodeExact = b
}

// SetMaxLength sets the maximum length for each of the return values from
// Encode.  If max is less than 1 the maximum length is set to 4.
func (m *Metaphone3) SetMaxLength(max int) {
	if max < 1 {
		max = 4
	}
	m.maxLength = max
}

// Encode returns main and alternate keys for word.  It honors
// the values set by SetEncodeVowels, SetEncodeExact and
// SetMaxLength.  One of a word's keys will match one of the other
// word's keys for similar sounding words.
func (m *Metaphone3) Encode(word string) (metaph, metaph2 string) {
	m.flag_AL_inversion = false

	m.current = 0

	m.inWord = []rune(strings.ToUpper(word))
	m.length = len(m.inWord)

	m.primary = []rune{}
	m.secondary = []rune{}

	if m.length < 1 {
		return
	}

	// zero based index
	m.last = m.length - 1

	///////////main loop//////////////////////////
	for len(m.primary) < m.maxLength || len(m.secondary) < m.maxLength {
		if m.current >= m.length {
			break
		}
		switch m.charAt(m.current) {
		case 'B':
			m.encodeB()
		case 'ß':
		case 'Ç':
			m.metaphAdd("S")
			m.current++
		case 'C':
			m.encodeC()
		case 'D':
			m.encodeD()
		case 'F':
			m.encodeF()
		case 'G':
			m.encodeG()
		case 'H':
			m.encodeH()
		case 'J':
			m.encodeJ()
		case 'K':
			m.encodeK()
		case 'L':
			m.encodeL()
		case 'M':
			m.encodeM()
		case 'N':
			m.encodeN()
		case 'Ñ':
			m.metaphAdd("N")
			m.current++
		case 'P':
			m.encodeP()
		case 'Q':
			m.encodeQ()
		case 'R':
			m.encodeR()
		case 'S':
			m.encodeS()
		case 'T':
			m.encodeT()
		case 'Ð', 'Þ': // eth, thorn
			m.metaphAdd("0")
			m.current++
		case 'V':
			m.encodeV()
		case 'W':
			m.encodeW()
		case 'X':
			m.encodeX()
		case '':
			m.metaphAdd("X")
			m.current++
		case '':
			m.metaphAdd("S")
			m.current++
		case 'Z':
			m.encodeZ()
		default:
			if m.isVowel(m.current) {
				m.encodeVowel()
			} else {
				m.current++
			}
		}
	}

	// only give back m.maxLength number of chars in m.primary/m.secondary
	if len(m.primary) > m.maxLength {
		m.primary = m.primary[:m.maxLength]
	}

	if len(m.secondary) > m.maxLength {
		m.secondary = m.secondary[:m.maxLength]
	}

	metaph = string(m.primary)
	metaph2 = string(m.secondary)

	// it is possible for the two metaphs to be the same
	// after truncation. lose the second one if so
	if metaph2 == metaph {
		metaph2 = ""
	}

	return
}

// metaphAdd adds a string to primary and secondary.  Call it with 1 or 2
// arguments.  The first argument is appended to primary (and to
// secondary if a second argument is not provided).  Any second
// argument is appended to secondary.  But don't append an 'A' after
// another 'A'.
func (m *Metaphone3) metaphAdd(s ...string) {
	if len(s) < 1 || len(s) > 2 {
		panic("metaphAdd requires one or two arguments")
	}
	// main: s[0]   alt: s[1]
	primaryLen := len(m.primary)
	secondaryLen := len(m.secondary)
	if !(s[0] == "A" &&
		(primaryLen > 0) &&
		(m.primary[primaryLen-1] == 'A')) {
		m.primary = append(m.primary, []rune(s[0])...)
	}
	if len(s) == 1 {
		if !(s[0] == "A" &&
			(secondaryLen > 0) &&
			(m.secondary[secondaryLen-1] == 'A')) {
			m.secondary = append(m.secondary, []rune(s[0])...)
		}
	} else {
		if !(s[1] == "A" &&
			(secondaryLen > 0) &&
			(m.secondary[secondaryLen-1] == 'A')) && len(s[1]) > 0 {
			m.secondary = append(m.secondary, []rune(s[1])...)
		}
	}
}

/**
 * Adds an encoding character to the encoded key value string - Exact/Approx version
 *
 * @param mainExact primary encoding character to be added to encoded key string if
 * m.encodeExact is set
 *
 * @param altExact alternative encoding character to be added to encoded alternative
 * key string if m.encodeExact is set
 *
 * @param main primary encoding character to be added to encoded key string
 *
 * @param alt alternative encoding character to be added to encoded alternative key string
 *
 */
func (m *Metaphone3) metaphAddExactApprox(s ...string) {
	// args are either mainExact/main or mainExact/altExact/main/alt.
	if len(s) != 2 && len(s) != 4 {
		panic("metaphAddExactApprox requires 2 or 4 arguments")
	}
	if len(s) == 2 {
		if m.encodeExact {
			m.metaphAdd(s[0])
		} else {
			m.metaphAdd(s[1])
		}
	} else {
		if m.encodeExact {
			m.metaphAdd(s[0], s[1])
		} else {
			m.metaphAdd(s[2], s[3])
		}
	}
}

/**
 * Subscript safe charAt()
 *
 * @param at index of character to access
 * @return 0 if index out of bounds, rune at index at otherwise
 */
func (m *Metaphone3) charAt(at int) rune {
	// check substring bounds
	if at >= 0 && at < m.length {
		return m.inWord[at]
	}
	return 0
}

// stringAtPos determines if any of the strings in s are
// in m.inWord at position pos.  The strings in s must be in order by
// increasing length, shortest first.
func (m *Metaphone3) stringAtPos(pos int, s ...string) bool {
	if pos >= 0 && pos < m.length && len(s) > 0 && len(s[0]) <= m.length {
	outerForLoop:
		for _, str := range s {
			if (pos + len(str)) > m.length {
				break
			}
			for i, r := range str {
				if r != m.inWord[pos+i] {
					continue outerForLoop
				}
			}
			return true
		}
	}
	return false
}

// stringEqual determines if any of the strings in s are
// equal to m.inWord.  The strings in s must be in order by
// increasing length, shortest first.
func (m *Metaphone3) stringEqual(s ...string) bool {
outerForLoop:
	for _, str := range s {
		if len(str) < m.length {
			continue
		}
		if len(str) > m.length {
			break
		}
		for i, r := range str {
			if r != m.inWord[i] {
				continue outerForLoop
			}
		}
		return true
	}
	return false
}

// stringAt determines if any of the strings in s are
// in m.inWord at m.current+offset.  The strings in s must be in order by
// increasing length, shortest first.
func (m *Metaphone3) stringAt(offset int, s ...string) bool {
	return m.stringAtPos(m.current+offset, s...)
}

// stringAtStart determines if any of the strings in s are
// in m.inWord at its beginning.  The strings in s must be in order by
// increasing length, shortest first.
func (m *Metaphone3) stringAtStart(s ...string) bool {
	return m.stringAtPos(0, s...)
}

// stringAtEnd determines if any of the strings in s are
// in m.inWord at its end.  The strings in s must be in order by
// increasing length, shortest first.
func (m *Metaphone3) stringAtEnd(s ...string) bool {
outerForLoop:
	for _, str := range s {
		i := m.length - len(str)
		if i < 0 {
			break
		}
		for _, r := range str {
			if r != m.inWord[i] {
				continue outerForLoop
			}
			i++
		}
		return true
	}
	return false
}

/**
 * Test for close front vowels
 *
 * @return true if close front vowel
 */
func (m *Metaphone3) frontVowel(at int) bool {
	c := m.charAt(at)
	return c == 'E' || c == 'I' || c == 'Y'
}

/**
 * Detect names or words that begin with spellings
 * typical of german or slavic words, for the purpose
 * of choosing alternate pronunciations correctly
 *
 */
func (m *Metaphone3) slavoGermanic() bool {
	return m.stringAtStart("SCH") ||
		m.stringAtStart("SW") ||
		(m.charAt(0) == 'J') ||
		(m.charAt(0) == 'W')
}

/**
 * Tests if character is a vowel
 *
 * @param at integer location of character to be tested
 * @return true if character is a vowel, false if not
 *
 */
// this code is faster than strings.ContainsRune("AEIOUYÀ...", m.charAt(at))
func (m *Metaphone3) isVowel(at int) bool {
	switch m.charAt(at) {
	case 'A', 'E', 'I', 'O', 'U', 'Y', 'À', 'Á', 'Â', 'Ã',
		'Ä', 'Å', 'Æ', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï',
		'Ò', 'Ó', 'Ô', 'Õ', 'Ö', '', 'Ø', 'Ù', 'Ú', 'Û', 'Ü', 'Ý', '':
		return true
	default:
		return false
	}
}

/**
 * Skips over vowels in a string. Has exceptions for skipping consonants that
 * will not be encoded.
 *
 * @param at position, in string to be encoded, of character to start skipping from
 *
 * @return position of next consonant in string to be encoded
 */
func (m *Metaphone3) skipVowels(at int) int {
	if at < 0 {
		return 0
	}
	if at >= m.length {
		return m.length
	}
	for m.isVowel(at) || (m.charAt(at) == 'W') {
		if m.stringAtPos(at, "WICZ", "WITZ", "WIAK") ||
			m.stringAtPos(at-1, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
			m.stringAtEnd("WICKI", "WACKI") {
			break
		}
		at++
		if ((m.charAt(at-1) == 'W') && (m.charAt(at) == 'H')) &&
			!(m.stringAtPos(at, "HOP") ||
				m.stringAtPos(at, "HIDE", "HARD", "HEAD", "HAWK", "HERD", "HOOK", "HAND", "HOLE") ||
				m.stringAtPos(at, "HEART", "HOUSE", "HOUND") ||
				m.stringAtPos(at, "HAMMER")) {
			at++
		}
		if at >= m.length {
			break
		}
	}
	return at
}

/**
 * Advanced counter m.current so that it indexes the next character to be encoded
 *
 * @param ifNotEncodeVowels number of characters to advance if not encoding internal vowels
 * @param ifEncodeVowels number of characters to advance if encoding internal vowels
 *
 */
func (m *Metaphone3) advanceCounter(ifNotEncodeVowels, ifEncodeVowels int) {
	if !m.encodeVowels {
		m.current += ifNotEncodeVowels
	} else {
		m.current += ifEncodeVowels
	}
}

/**
 * Tests whether the word is the root or a regular english inflection
 * of it, e.g. "ache", "achy", "aches", "ached", "aching", "achingly"
 * This is for cases where we want to match only the root and corresponding
 * inflected forms, and not completely different words which may have the
 * same substring in them.
 */
func (m *Metaphone3) rootOrInflections(InWord []rune, root string) bool {
	inWord := string(InWord)
	rootrune := []rune(root)
	len := len(rootrune)
	lastrune := rootrune[len-1]
	var test string

	if inWord == root {
		return true
	}

	if inWord == root+"S" {
		return true
	}

	if lastrune != 'E' && inWord == root+"ES" {
		return true
	}

	if lastrune != 'E' {
		test = root + "ED"
	} else {
		test = root + "D"
	}

	if inWord == test {
		return true
	}

	if lastrune == 'E' {
		root = string(rootrune[:len-1])
	}

	if inWord == root+"ING" {
		return true
	}

	if inWord == root+"INGLY" {
		return true
	}

	return inWord == root+"Y"
}

/**
 * Tests for cases where non-initial 'o' is not pronounced
 * Only executed if non initial vowel encoding is turned on
 *
 * @return true if encoded as silent - no addition to m.metaph key
 *
 */
func (m *Metaphone3) oSilent() bool {
	// if "iron" at beginning or end of word and not "irony"
	if (m.charAt(m.current) == 'O') && m.stringAt(-2, "IRON") {
		if (m.stringAtStart("IRON") ||
			(m.stringAt(-2, "IRON") &&
				(m.last == (m.current + 1)))) &&
			!m.stringAt(-2, "IRONIC") {
			return true
		}
	}

	return false
}

/**
 * Detect conditions required
 * for the 'E' not to be pronounced
 *
 */
func (m *Metaphone3) eSilentSuffix(at int) bool {
	return (m.current == (at - 1)) &&
		(m.length > (at + 1)) &&
		(m.isVowel((at + 1)) ||
			(m.stringAtPos(at, "ST", "SL") &&
				(m.length > (at + 2))))
}

/**
 * Detect endings that will
 * cause the 'e' to be pronounced
 *
 */
func (m *Metaphone3) ePronouncingSuffix(at int) bool {
	// e.g. 'bridgewood' - the other vowels will get eaten
	// up so we need to put one in here
	if m.stringAtEnd("WOOD") {
		return true
	}

	// same as above
	if m.stringAtEnd("WATER", "WORTH") {
		return true
	}

	// e.g. 'bridgette'
	if m.stringAtEnd("TTE", "LIA", "NOW", "ROS", "RAS") {
		return true
	}

	// e.g. 'olena'
	if m.stringAtEnd("TA", "TT", "NA", "NO", "NE",
		"RS", "RE", "LA", "AU", "RO", "RA") {
		return true
	}

	// e.g. 'bridget'
	if m.stringAtEnd("T", "R") {
		return true
	}

	return false
}

/**
 * Detect internal silent 'E's e.g. "roseman",
 * "firestone"
 *
 */
func (m *Metaphone3) silentInternalE() bool {
	// 'olesen' but not 'olen'	RAKE BLAKE
	return (m.stringAtStart("OLE") &&
		m.eSilentSuffix(3) && !m.ePronouncingSuffix(3)) ||
		(m.stringAtStart("BARE", "FIRE", "FORE", "GATE", "HAGE", "HAVE",
			"HAZE", "HOLE", "CAPE", "HUSE", "LACE", "LINE",
			"LIVE", "LOVE", "MORE", "MOSE", "MORE", "NICE",
			"RAKE", "ROBE", "ROSE", "SISE", "SIZE", "WARE",
			"WAKE", "WISE", "WINE") &&
			m.eSilentSuffix(4) && !m.ePronouncingSuffix(4)) ||
		(m.stringAtStart("BLAKE", "BRAKE", "BRINE", "CARLE", "CLEVE", "DUNNE",
			"HEDGE", "HOUSE", "JEFFE", "LUNCE", "STOKE", "STONE",
			"THORE", "WEDGE", "WHITE") &&
			m.eSilentSuffix(5) && !m.ePronouncingSuffix(5)) ||
		(m.stringAtStart("BRIDGE", "CHEESE") &&
			m.eSilentSuffix(6) && !m.ePronouncingSuffix(6)) ||
		m.stringAt(-5, "CHARLES")
}

/**
 * Tests for words where an 'E' at the end of the word
 * is pronounced
 *
 * special cases, mostly from the greek, spanish, japanese,
 * italian, and french words normally having an acute accent.
 * also, pronouns and articles
 *
 * Many Thanks to ali, QuentinCompson, JeffCO, ToonScribe, Xan,
 * Trafalz, and VictorLaszlo, all of them atriots from the Eschaton,
 * for all their fine contributions!
 *
 * @return true if 'E' at end is pronounced
 *
 */
func (m *Metaphone3) ePronouncedAtEnd() bool {
	return (m.current == m.last) &&
		(m.stringAt(-6, "STROPHE") ||
			// if a vowel is before the 'E', vowel eater will have eaten it.
			//otherwise, consonant + 'E' will need 'E' pronounced
			(m.length == 2) ||
			((m.length == 3) && !m.isVowel(0)) ||
			// these german name endings can be relied on to have the 'e' pronounced
			(m.stringAtEnd("BKE", "DKE", "FKE", "KKE", "LKE",
				"NKE", "MKE", "PKE", "TKE", "VKE", "ZKE") &&
				!m.stringAtStart("FINKE", "FUNKE") &&
				!m.stringAtStart("FRANKE")) ||
			m.stringAtEnd("SCHKE") ||
			m.stringEqual("ACME", "NIKE", "CAFE", "RENE", "LUPE", "JOSE", "ESME",
				"LETHE", "CADRE", "TILDE", "SIGNE", "POSSE", "LATTE", "ANIME", "DOLCE", "CROCE",
				"ADOBE", "OUTRE", "JESSE", "JAIME", "JAFFE", "BENGE", "RUNGE",
				"CHILE", "DESME", "CONDE", "URIBE", "LIBRE", "ANDRE",
				"HECATE", "PSYCHE", "DAPHNE", "PENSKE", "CLICHE", "RECIPE",
				"TAMALE", "SESAME", "SIMILE", "FINALE", "KARATE", "RENATE", "SHANTE",
				"OBERLE", "COYOTE", "KRESGE", "STONGE", "STANGE", "SWAYZE", "FUENTE",
				"SALOME", "URRIBE",
				"ECHIDNE", "ARIADNE", "MEINEKE", "PORSCHE", "ANEMONE", "EPITOME",
				"SYNCOPE", "SOUFFLE", "ATTACHE", "MACHETE", "KARAOKE", "BUKKAKE",
				"VICENTE", "ELLERBE", "VERSACE",
				"PENELOPE", "CALLIOPE", "CHIPOTLE", "ANTIGONE", "KAMIKAZE", "EURIDICE",
				"YOSEMITE", "FERRANTE",
				"HYPERBOLE", "GUACAMOLE", "XANTHIPPE",
				"SYNECDOCHE"))
}

/**
 * Encodes "-UE".
 *
 * @return true if encoding handled in this routine, false if not
 */
func (m *Metaphone3) skipSilentUE() bool {
	// always silent except for cases listed below
	if (m.stringAt(-1, "QUE", "GUE") &&
		!m.stringAtStart("BARBEQUE", "PALENQUE", "APPLIQUE") &&
		// '-que' cases usually french but missing the acute accent
		!m.stringAtStart("RISQUE") &&
		!m.stringAt(-3, "ARGUE", "SEGUE") &&
		!m.stringAtStart("PIROGUE", "ENRIQUE") &&
		!m.stringAtStart("COMMUNIQUE")) &&
		(m.current > 1) &&
		(((m.current + 1) == m.last) ||
			m.stringAtStart("JACQUES")) {
		m.current = m.skipVowels(m.current)
		return true
	}

	return false
}

/**
 * Tests and encodes cases where non-initial 'e' is never pronounced
 * Only executed if non initial vowel encoding is turned on
 *
 * @return true if encoded as silent - no addition to m.metaph key
 *
 */
func (m *Metaphone3) eSilent() bool {
	if m.ePronouncedAtEnd() {
		return false
	}

	// 'e' silent when last letter, altho
	return (m.current == m.last) ||
		// also silent if before plural 's'
		// or past tense or participle 'd', e.g.
		// 'grapes' and 'banished' => PNXT
		(m.stringAtEnd("S", "D") &&
			(m.current > 1) &&
			((m.current + 1) == m.last) &&
			// and not e.g. "nested", "rises", or "pieces" => RASAS
			!(m.stringAt(-1, "TED", "SES", "CES") ||
				m.stringAtStart("ANTIPODES", "ANOPHELES") ||
				m.stringAtStart("MOHAMMED", "MUHAMMED", "MOUHAMED") ||
				m.stringAtStart("MOHAMED") ||
				m.stringAtStart("NORRED", "MEDVED", "MERCED", "ALLRED", "KHALED", "RASHED", "MASJED") ||
				m.stringAtStart("JARED", "AHMED", "HAMED", "JAVED") ||
				m.stringAtStart("ABED", "IMED"))) ||
		// e.g.  'wholeness', 'boneless', 'barely'
		(m.stringAt(1, "NESS", "LESS") && ((m.current + 4) == m.last)) ||
		(m.stringAt(1, "LY") && ((m.current + 2) == m.last) &&
			!m.stringAtStart("CICELY"))
}

/**
 * Exceptions where 'E' is pronounced where it
 * usually wouldn't be, and also some cases
 * where 'LE' transposition rules don't apply
 * and the vowel needs to be encoded here
 *
 * @return true if 'E' pronounced
 *
 */
func (m *Metaphone3) ePronouncedExceptions() bool {
	// greek names e.g. "herakles" or hispanic names e.g. "robles", where 'e' is pronounced, other exceptions
	return (((m.current + 1) == m.last) &&
		(m.stringAt(-3, "OCLES", "ACLES", "AKLES") ||
			m.stringAtStart("INES") ||
			m.stringAtStart("LOPES", "ESTES", "GOMES", "NUNES", "ALVES", "ICKES",
				"INNES", "PERES", "WAGES", "NEVES", "BENES", "DONES") ||
			m.stringAtStart("CORTES", "CHAVES", "VALDES", "ROBLES", "TORRES", "FLORES", "BORGES",
				"NIEVES", "MONTES", "SOARES", "VALLES", "GEDDES", "ANDRES", "VIAJES",
				"CALLES", "FONTES", "HERMES", "ACEVES", "BATRES", "MATHES") ||
			m.stringAtStart("DELORES", "MORALES", "DOLORES", "ANGELES", "ROSALES", "MIRELES", "LINARES",
				"PERALES", "PAREDES", "BRIONES", "SANCHES", "CAZARES", "REVELES", "ESTEVES",
				"ALVARES", "MATTHES", "SOLARES", "CASARES", "CACERES", "STURGES", "RAMIRES",
				"FUNCHES", "BENITES", "FUENTES", "PUENTES", "TABARES", "HENTGES", "VALORES") ||
			m.stringAtStart("GONZALES", "MERCEDES", "FAGUNDES", "JOHANNES", "GONSALES", "BERMUDES",
				"CESPEDES", "BETANCES", "TERRONES", "DIOGENES", "CORRALES", "CABRALES",
				"MARTINES", "GRAJALES") ||
			m.stringAtStart("CERVANTES", "FERNANDES", "GONCALVES", "BENEVIDES", "CIFUENTES", "SIFUENTES",
				"SERVANTES", "HERNANDES", "BENAVIDES") ||
			m.stringAtStart("ARCHIMEDES", "CARRIZALES", "MAGALLANES"))) ||
		m.stringAt(-2, "FRED", "DGES", "DRED", "GNES") ||
		m.stringAt(-5, "PROBLEM", "RESPLEN") ||
		m.stringAt(-4, "REPLEN") ||
		m.stringAt(-3, "SPLE")
}

/**
 * Encodes cases where non-initial 'e' is pronounced, taking
 * care to detect unusual cases from the greek.
 *
 * Only executed if non initial vowel encoding is turned on
 *
 *
 */
func (m *Metaphone3) encodeEPronounced() {
	// special cases with two pronunciations
	// 'agape' 'lame' 'resume'
	if m.stringEqual("LAME", "SAKE", "PATE",
		"AGAPE") ||
		((m.current == 5) && m.stringAtStart("RESUME")) {
		m.metaphAdd("", "A")
		return
	}

	// special case "inge" => 'INGA', 'INJ'
	if m.stringEqual("INGE") {
		m.metaphAdd("A", "")
		return
	}

	// special cases with two pronunciations
	// special handling due to the difference in
	// the pronunciation of the '-D'
	if (m.current == 5) && m.stringAtStart("BLESSED", "LEARNED") {
		m.metaphAddExactApprox("D", "AD", "T", "AT")
		m.current += 2
		return
	}

	// encode all vowels and diphthongs to the same value
	if (!m.eSilent() && !m.flag_AL_inversion && !m.silentInternalE()) ||
		m.ePronouncedExceptions() {
		m.metaphAdd("A")
	}

	// now that we've visited the vowel in question
	m.flag_AL_inversion = false
}

/**
 * Encodes silent 'B' for cases not covered under "-mb-"
 *
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentB() bool {
	//'debt', 'doubt', 'subtle'
	if m.stringAt(-2, "DEBT") ||
		m.stringAt(-2, "SUBTL") ||
		m.stringAt(-2, "SUBTIL") ||
		m.stringAt(-3, "DOUBT") {
		m.metaphAdd("T")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes all initial vowels to A.
 *
 * Encodes non-initial vowels to A if m.encodeVowels is true
 *
 *
 */
func (m *Metaphone3) encodeVowel() {
	if m.current == 0 {
		// all init vowels map to 'A'
		// as of Double Metaphone
		m.metaphAdd("A")
	} else if m.encodeVowels {
		if m.charAt(m.current) != 'E' {
			if m.skipSilentUE() {
				return
			}

			if m.oSilent() {
				m.current++
				return
			}

			// encode all vowels and
			// diphthongs to the same value
			m.metaphAdd("A")
		} else {
			m.encodeEPronounced()
		}
	}

	if !(!m.isVowel(m.current-2) && m.stringAt(-1, "LEWA", "LEWO", "LEWI")) {
		m.current = m.skipVowels(m.current + 1)
	} else {
		m.current++
	}
}

/**
 * Encodes 'B'
 *
 *
 */
func (m *Metaphone3) encodeB() {
	if m.encodeSilentB() {
		return
	}

	// "-mb", e.g", "dumb", already skipped over under
	// 'M', altho it should really be handled here...
	m.metaphAddExactApprox("B", "P")

	if (m.charAt(m.current+1) == 'B') ||
		((m.charAt(m.current+1) == 'P') &&
			((m.current+1 < m.last) && (m.charAt(m.current+2) != 'H'))) {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encodes transliterations from the hebrew where the
 * sound 'kh' is represented as "-CH-". The normal pronounciation
 * of this in english is either 'h' or 'kh', and alternate
 * spellings most often use "-H-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCHToH() bool {
	// hebrew => 'H', e.g. 'channukah', 'chabad'
	if ((m.current == 0) &&
		(m.stringAt(2, "AIM", "ETH", "ELM") ||
			m.stringAt(2, "ASID", "AZAN") ||
			m.stringAt(2, "UPPAH", "UTZPA", "ALLAH", "ALUTZ", "AMETZ") ||
			m.stringAt(2, "ESHVAN", "ADARIM", "ANUKAH") ||
			m.stringAt(2, "ALLLOTH", "ANNUKAH", "AROSETH"))) ||
		// and an irish name with the same encoding
		m.stringAt(-3, "CLACHAN") {
		m.metaphAdd("H")
		m.advanceCounter(3, 2)
		return true
	}

	return false
}

/**
 * Encodes cases where 'C' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentCAtBeginning() bool {
	//skip these when at start of word
	if (m.current == 0) && m.stringAt(0, "CT", "CN") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encodes exceptions where "-CA-" should encode to S
 * instead of K including cases where the cedilla has not been used
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCAToS() bool {
	// Special case: 'caesar'.
	// Also, where cedilla not used, as in "linguica" => LNKS
	if ((m.current == 0) && m.stringAt(0, "CAES", "CAEC", "CAEM")) ||
		m.stringAtStart("FRANCAIS", "FRANCAIX", "LINGUICA") ||
		m.stringAtStart("FACADE") ||
		m.stringAtStart("GONCALVES", "PROVENCAL") {
		m.metaphAdd("S")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encodes exceptions where "-CO-" encodes to S instead of K
 * including cases where the cedilla has not been used
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCOToS() bool {
	// e.g. 'coelecanth' => SLKN0
	if (m.stringAt(0, "COEL") &&
		(m.isVowel(m.current+4) || ((m.current + 3) == m.last))) ||
		m.stringAt(0, "COENA", "COENO") ||
		m.stringAtStart("FRANCOIS", "MELANCON") ||
		m.stringAtStart("GARCON") {
		m.metaphAdd("S")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encodes "-CHAE-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCHAE() bool {
	// e.g. 'michael'
	if (m.current > 0) && m.stringAt(2, "AE") {
		if m.stringAtStart("RACHAEL") {
			m.metaphAdd("X")
		} else if !m.stringAt(-1, "C", "K", "G", "Q") {
			m.metaphAdd("K")
		}

		m.advanceCounter(4, 2)
		return true
	}

	return false
}

/**
 * Encodes cases where "-CH-" is not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentCH() bool {
	// '-ch-' not pronounced
	if m.stringAt(-2, "FUCHSIA") ||
		m.stringAt(-2, "YACHT") ||
		m.stringAtStart("STRACHAN") ||
		m.stringAtStart("CRICHTON") ||
		(m.stringAt(-3, "DRACHM")) &&
			!m.stringAt(-3, "DRACHMA") {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes "-CH-" to X
 * English language patterns
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCHToX() bool {
	// e.g. 'approach', 'beach'
	if (m.stringAt(-2, "OACH", "EACH", "EECH", "OUCH", "OOCH", "MUCH", "SUCH") &&
		!m.stringAt(-3, "JOACH")) ||
		// e.g. 'dacha', 'macho'
		(((m.current + 2) == m.last) && m.stringAt(-1, "ACHA", "ACHO")) ||
		(m.stringAt(0, "CHOT", "CHOD", "CHAT") && ((m.current + 3) == m.last)) ||
		((m.stringAt(-1, "OCHE") && ((m.current + 2) == m.last)) &&
			!m.stringAt(-2, "DOCHE")) ||
		m.stringAt(-4, "ATTACH", "DETACH", "KOVACH") ||
		m.stringAt(-5, "SPINACH") ||
		m.stringAtStart("MACHAU") ||
		m.stringAt(-4, "PARACHUT") ||
		m.stringAt(-5, "MASSACHU") ||
		(m.stringAt(-3, "THACH") && !m.stringAt(-1, "ACHE")) ||
		m.stringAt(-2, "VACHON") {
		m.metaphAdd("X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes "-CH-" to K in contexts of
 * initial "A" or "E" follwed by "CH"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeEnglishCHToK() bool {
	//'ache', 'echo', alternate spelling of 'michael'
	if ((m.current == 1) && m.rootOrInflections(m.inWord, "ACHE")) ||
		(((m.current > 3) && m.rootOrInflections(m.inWord[m.current-1:], "ACHE")) &&
			(m.stringAtStart("EAR") ||
				m.stringAtStart("HEAD", "BACK") ||
				m.stringAtStart("HEART", "BELLY", "TOOTH"))) ||
		m.stringAt(-1, "ECHO") ||
		m.stringAt(-2, "MICHEAL") ||
		m.stringAt(-4, "JERICHO") ||
		m.stringAt(-5, "LEPRECH") {
		m.metaphAdd("K", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes "-CH-" to K in mostly germanic context
 * of internal "-ACH-", with exceptions
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGermanicCHToK() bool {
	// various germanic
	// "<consonant><vowel>CH-" implies a german word where 'ch' => K
	if ((m.current > 1) &&
		!m.isVowel(m.current-2) &&
		m.stringAt(-1, "ACH") &&
		!m.stringAt(-2, "MACHADO", "MACHUCA", "LACHANC", "LACHAPE", "KACHATU") &&
		!m.stringAt(-3, "KHACHAT") &&
		((m.charAt(m.current+2) != 'I') &&
			((m.charAt(m.current+2) != 'E') ||
				m.stringAt(-2, "BACHER", "MACHER", "MACHEN", "LACHER"))) ||
		// e.g. 'brecht', 'fuchs'
		(m.stringAt(2, "T", "S") &&
			!(m.stringAtStart("WHICHSOEVER") || m.stringAtStart("LUNCHTIME"))) ||
		// e.g. 'andromache'
		m.stringAtStart("SCHR") ||
		((m.current > 2) && m.stringAt(-2, "MACHE")) ||
		((m.current == 2) && m.stringAt(-2, "ZACH")) ||
		m.stringAt(-4, "SCHACH") ||
		m.stringAt(-1, "ACHEN") ||
		m.stringAt(-3, "SPICH", "ZURCH", "BUECH") ||
		(m.stringAt(-3, "KIRCH", "JOACH", "BLECH", "MALCH") &&
			// "kirch" and "blech" both get 'X'
			!(m.stringAt(-3, "KIRCHNER") || ((m.current + 1) == m.last))) ||
		(((m.current + 1) == m.last) && m.stringAt(-2, "NICH", "LICH", "BACH")) ||
		(((m.current + 1) == m.last) &&
			m.stringAt(-3, "URICH", "BRICH", "ERICH", "DRICH", "NRICH") &&
			!m.stringAt(-5, "ALDRICH") &&
			!m.stringAt(-6, "GOODRICH") &&
			!m.stringAt(-7, "GINGERICH"))) ||
		(((m.current + 1) == m.last) && m.stringAt(-4, "ULRICH", "LFRICH", "LLRICH",
			"EMRICH", "ZURICH", "EYRICH")) ||
		// e.g., 'wachtler', 'wechsler', but not 'tichner'
		((m.stringAt(-1, "A", "O", "U", "E") || (m.current == 0)) &&
			m.stringAt(2, "L", "R", "N", "M", "B", "H", "F", "V", "W", " ")) {
		// "CHR/L-" e.g. 'chris' do not get
		// alt pronunciation of 'X'
		if m.stringAt(2, "R", "L") || m.slavoGermanic() {
			m.metaphAdd("K")
		} else {
			m.metaphAdd("K", "X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-ARCH-". Some occurrences are from greek roots and therefore encode
 * to 'K', others are from english words and therefore encode to 'X'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeARCH() bool {
	if m.stringAt(-2, "ARCH") {
		// "-ARCH-" has many combining forms where "-CH-" => K because of its
		// derivation from the greek
		if ((m.isVowel(m.current+2) && m.stringAt(-2, "ARCHA", "ARCHI", "ARCHO", "ARCHU", "ARCHY")) ||
			m.stringAt(-2, "ARCHEA", "ARCHEG", "ARCHEO", "ARCHET", "ARCHEL", "ARCHES", "ARCHEP",
				"ARCHEM", "ARCHEN") ||
			(m.stringAt(-2, "ARCH") && ((m.current + 1) == m.last)) ||
			m.stringAtStart("MENARCH")) &&
			(!m.rootOrInflections(m.inWord, "ARCH") &&
				!m.stringAt(-4, "SEARCH", "POARCH") &&
				!m.stringAtStart("ARCHENEMY", "ARCHIBALD", "ARCHULETA", "ARCHAMBAU") &&
				!m.stringAtStart("ARCHER", "ARCHIE") &&
				!((((m.stringAt(-3, "LARCH", "MARCH", "PARCH") ||
					m.stringAt(-4, "STARCH")) &&
					!(m.stringAtStart("EPARCH") ||
						m.stringAtStart("NOMARCH") ||
						m.stringAtStart("EXILARCH", "HIPPARCH", "MARCHESE") ||
						m.stringAtStart("ARISTARCH") ||
						m.stringAtStart("MARCHETTI"))) ||
					m.rootOrInflections(m.inWord, "STARCH")) &&
					(!m.stringAt(-2, "ARCHU", "ARCHY") ||
						m.stringAtStart("STARCHY")))) {
			m.metaphAdd("K", "X")
		} else {
			m.metaphAdd("X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-CH-" to K when from greek roots
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGreekCHInitial() bool {
	// greek roots e.g. 'chemistry', 'chorus', ch at beginning of root
	if (m.stringAt(0, "CHAMOM", "CHARAC", "CHARIS", "CHARTO", "CHARTU", "CHARYB", "CHRIST", "CHEMIC", "CHILIA") ||
		(m.stringAt(0, "CHEMI", "CHEMO", "CHEMU", "CHEMY", "CHOND", "CHONA", "CHONI", "CHOIR", "CHASM",
			"CHARO", "CHROM", "CHROI", "CHAMA", "CHALC", "CHALD", "CHAET", "CHIRO", "CHILO", "CHELA", "CHOUS",
			"CHEIL", "CHEIR", "CHEIM", "CHITI", "CHEOP") &&
			!(m.stringAt(0, "CHEMIN") || m.stringAt(-2, "ANCHONDO"))) ||
		(m.stringAt(0, "CHISM", "CHELI") &&
			// exclude spanish "machismo"
			!(m.stringAtStart("MACHISMO") ||
				// exclude some french words
				m.stringAtStart("REVANCHISM") ||
				m.stringAtStart("RICHELIEU") ||
				m.stringEqual("CHISM") ||
				m.stringAtStart("MICHEL"))) ||
		// include e.g. "chorus", "chyme", "chaos"
		(m.stringAt(0, "CHOR", "CHOL", "CHYM", "CHYL", "CHLO", "CHOS", "CHUS", "CHOE") &&
			!m.stringAtStart("CHOLLO", "CHOLLA", "CHORIZ")) ||
		// "chaos" => K but not "chao"
		(m.stringAt(0, "CHAO") && ((m.current + 3) != m.last)) ||
		// e.g. "abranchiate"
		(m.stringAt(0, "CHIA") && !(m.stringAtStart("APPALACHIA") || m.stringAtStart("CHIAPAS"))) ||
		// e.g. "chimera"
		m.stringAt(0, "CHIMERA", "CHIMAER", "CHIMERI") ||
		// e.g. "chameleon"
		((m.current == 0) && m.stringAt(0, "CHAME", "CHELO", "CHITO")) ||
		// e.g. "spirochete"
		((((m.current + 4) == m.last) || ((m.current + 5) == m.last)) && m.stringAt(-1, "OCHETE"))) &&
		// more exceptions where "-CH-" => X e.g. "chortle", "crocheter"
		!(m.stringEqual("CHORE", "CHOLO", "CHOLA") ||
			m.stringAt(0, "CHORT", "CHOSE") ||
			m.stringAt(-3, "CROCHET") ||
			m.stringAtStart("CHEMISE", "CHARISE", "CHARISS", "CHAROLE")) {
		// "CHR/L-" e.g. 'christ', 'chlorine' do not get
		// alt pronunciation of 'X'
		if m.stringAt(2, "R", "L") {
			m.metaphAdd("K")
		} else {
			m.metaphAdd("K", "X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode a variety of greek and some german roots where "-CH-" => K
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGreekCHNonInitial() bool {
	//greek & other roots e.g. 'tachometer', 'orchid', ch in middle or end of root
	if m.stringAt(-2, "ORCHID", "NICHOL", "MECHAN", "LICHEN", "MACHIC", "PACHEL", "RACHIF", "RACHID",
		"RACHIS", "RACHIC", "MICHAL") ||
		m.stringAt(-3, "MELCH", "GLOCH", "TRACH", "TROCH", "BRACH", "SYNCH", "PSYCH",
			"STICH", "PULCH", "EPOCH") ||
		(m.stringAt(-3, "TRICH") && !m.stringAt(-5, "OSTRICH")) ||
		(m.stringAt(-2, "TYCH", "TOCH", "BUCH", "MOCH", "CICH", "DICH", "NUCH", "EICH", "LOCH",
			"DOCH", "ZECH", "WYCH") &&
			!(m.stringAt(-4, "INDOCHINA") || m.stringAt(-2, "BUCHON"))) ||
		m.stringAt(-2, "LYCHN", "TACHO", "ORCHO", "ORCHI", "LICHO") ||
		(m.stringAt(-1, "OCHER", "ECHIN", "ECHID") && ((m.current == 1) || (m.current == 2))) ||
		m.stringAt(-4, "BRONCH", "STOICH", "STRYCH", "TELECH", "PLANCH", "CATECH", "MANICH", "MALACH",
			"BIANCH", "DIDACH") ||
		(m.stringAt(-1, "ICHA", "ICHN") && (m.current == 1)) ||
		m.stringAt(-2, "ORCHESTR") ||
		m.stringAt(-4, "BRANCHIO", "BRANCHIF") ||
		(m.stringAt(-1, "ACHAB", "ACHAD", "ACHAN", "ACHAZ") &&
			!m.stringAt(-2, "MACHADO", "LACHANC")) ||
		m.stringAt(-1, "ACHISH", "ACHILL", "ACHAIA", "ACHENE") ||
		m.stringAt(-1, "ACHAIAN", "ACHATES", "ACHIRAL", "ACHERON") ||
		m.stringAt(-1, "ACHILLEA", "ACHIMAAS", "ACHILARY", "ACHELOUS", "ACHENIAL", "ACHERNAR") ||
		m.stringAt(-1, "ACHALASIA", "ACHILLEAN", "ACHIMENES") ||
		m.stringAt(-1, "ACHIMELECH", "ACHITOPHEL") ||
		// e.g. 'inchoate'
		(((m.current - 2) == 0) && (m.stringAt(-2, "INCHOA") ||
			// e.g. 'ischemia'
			m.stringAtStart("ISCH"))) ||
		// e.g. 'ablimelech', 'antioch', 'pentateuch'
		(((m.current + 1) == m.last) && m.stringAt(-1, "A", "O", "U", "E") &&
			!(m.stringAtStart("DEBAUCH") ||
				m.stringAt(-2, "MUCH", "SUCH", "KOCH") ||
				m.stringAt(-5, "OODRICH", "ALDRICH"))) {
		m.metaphAdd("K", "X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-CH-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCH() bool {
	if m.stringAt(0, "CH") {
		if m.encodeCHAE() ||
			m.encodeCHToH() ||
			m.encodeSilentCH() ||
			m.encodeARCH() ||
			// m.encodeCHToX() should be
			// called before the germanic
			// and greek encoding functions
			m.encodeCHToX() ||
			m.encodeEnglishCHToK() ||
			m.encodeGermanicCHToK() ||
			m.encodeGreekCHInitial() ||
			m.encodeGreekCHNonInitial() {
			return true
		}

		if m.current > 0 {
			if m.stringAtStart("MC") && (m.current == 1) {
				//e.g., "McHugh"
				m.metaphAdd("K")
			} else {
				m.metaphAdd("X", "K")
			}
		} else {
			m.metaphAdd("X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes reliably italian "-CCIA-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCCIA() bool {
	//e.g., 'focaccia'
	if m.stringAt(1, "CIA") {
		m.metaphAdd("X", "S")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-CC-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCC() bool {
	//double 'C', but not if e.g. 'McClellan'
	if m.stringAt(0, "CC") && !((m.current == 1) && (m.charAt(0) == 'M')) {
		// exception
		if m.stringAt(-3, "FLACCID") {
			m.metaphAdd("S")
			m.advanceCounter(3, 2)
			return true
		}

		//'bacci', 'bertucci', other italian
		if (((m.current + 2) == m.last) && m.stringAt(2, "I")) ||
			m.stringAt(2, "IO") ||
			(((m.current + 4) == m.last) && m.stringAt(2, "INO", "INI")) {
			m.metaphAdd("X")
			m.advanceCounter(3, 2)
			return true
		}

		//'accident', 'accede' 'succeed'
		if m.stringAt(2, "I", "E", "Y") &&
			//except 'bellocchio','bacchus', 'soccer' get K
			!((m.charAt(m.current+2) == 'H') ||
				m.stringAt(-2, "SOCCER")) {
			m.metaphAdd("KS")
			m.advanceCounter(3, 2)
			return true

		} else {
			//Pierce's rule
			m.metaphAdd("K")
			m.current += 2
			return true
		}
	}

	return false
}

/**
 * Encode cases where the consonant following "C" is redundant
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCKCGCQ() bool {
	if m.stringAt(0, "CK", "CG", "CQ") {
		// eastern european spelling e.g. 'gorecki' == 'goresky'
		if m.stringAt(0, "CKI", "CKY") &&
			((m.current + 2) == m.last) &&
			(m.length > 6) {
			m.metaphAdd("K", "SK")
		} else {
			m.metaphAdd("K")
		}
		m.current += 2

		if m.stringAt(0, "K", "G", "Q") {
			m.current++
		}
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeBritishSilentCE() bool {
	// english place names like e.g.'gloucester' pronounced glo-ster
	return (m.stringAt(1, "ESTER") && ((m.current + 5) == m.last)) ||
		m.stringAt(1, "ESTERSHIRE")
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCE() bool {
	// 'ocean', 'commercial', 'provincial', 'cello', 'fettucini', 'medici'
	if (m.stringAt(1, "EAN") && m.isVowel(m.current-1)) ||
		// e.g. 'rosacea'
		(m.stringAt(-1, "ACEA") &&
			((m.current + 2) == m.last) &&
			!m.stringAtStart("PANACEA")) ||
		// e.g. 'botticelli', 'concerto'
		m.stringAt(1, "ELLI", "ERTO", "EORL") ||
		// some italian names familiar to americans
		(m.stringAt(-3, "CROCE") && ((m.current + 1) == m.last)) ||
		m.stringAt(-3, "DOLCE") ||
		// e.g. 'cello'
		(m.stringAt(1, "ELLO") &&
			((m.current + 4) == m.last)) {
		m.metaphAdd("X", "S")
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCI() bool {
	// with consonant before C
	// e.g. 'fettucini', but exception for the americanized pronunciation of 'mancini'
	if ((m.stringAt(1, "INI") && !m.stringAtStart("MANCINI")) && ((m.current + 3) == m.last)) ||
		// e.g. 'medici'
		(m.stringAt(-1, "ICI") && ((m.current + 1) == m.last)) ||
		// e.g. "commercial', 'provincial', 'cistercian'
		m.stringAt(-1, "RCIAL", "NCIAL", "RCIAN", "UCIUS") ||
		// special cases
		m.stringAt(-3, "MARCIA") ||
		m.stringAt(-2, "ANCIENT") {
		m.metaphAdd("X", "S")
		return true
	}

	// with vowel before C (or at beginning?)
	if ((m.stringAt(0, "CIO", "CIE", "CIA") &&
		m.isVowel(m.current-1)) ||
		// e.g. "ciao"
		m.stringAt(1, "IAO")) &&
		!m.stringAt(-4, "COERCION") {
		if (m.stringAt(0, "CIAN", "CIAL", "CIAO", "CIES", "CIOL", "CION") ||
			// exception - "glacier" => 'X' but "spacier" = > 'S'
			m.stringAt(-3, "GLACIER") ||
			m.stringAt(0, "CIENT", "CIENC", "CIOUS", "CIATE", "CIATI", "CIATO", "CIABL", "CIARY") ||
			(((m.current + 2) == m.last) && m.stringAt(0, "CIA", "CIO")) ||
			(((m.current + 3) == m.last) && m.stringAt(0, "CIAS", "CIOS"))) &&
			// exceptions
			!(m.stringAt(-4, "ASSOCIATION") ||
				m.stringAtStart("OCIE") ||
				// exceptions mostly because these names are usually from
				// the spanish rather than the italian in america
				m.stringAt(-2, "LUCIO") ||
				m.stringAt(-2, "MACIAS") ||
				m.stringAt(-3, "GRACIE", "GRACIA") ||
				m.stringAt(-2, "LUCIANO") ||
				m.stringAt(-3, "MARCIANO") ||
				m.stringAt(-4, "PALACIO") ||
				m.stringAt(-4, "FELICIANO") ||
				m.stringAt(-5, "MAURICIO") ||
				m.stringAt(-7, "ENCARNACION") ||
				m.stringAt(-4, "POLICIES") ||
				m.stringAt(-2, "HACIENDA") ||
				m.stringAt(-6, "ANDALUCIA") ||
				m.stringAt(-2, "SOCIO", "SOCIE")) {
			m.metaphAdd("X", "S")
		} else {
			m.metaphAdd("S", "X")
		}

		return true
	}

	// exception
	if m.stringAt(-4, "COERCION") {
		m.metaphAdd("J")
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeLatinateSuffixes() bool {
	if m.stringAt(1, "EOUS", "IOUS") {
		m.metaphAdd("X", "S")
		return true
	}

	return false
}

/**
 * Encode cases where "C" preceeds a front vowel such as "E", "I", or "Y".
 * These cases most likely => S or X
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCFrontVowel() bool {
	if m.stringAt(0, "CI", "CE", "CY") {
		if m.encodeBritishSilentCE() ||
			m.encodeCE() ||
			m.encodeCI() ||
			m.encodeLatinateSuffixes() {
			m.advanceCounter(2, 1)
			return true
		}

		m.metaphAdd("S")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encodes some exceptions where "C" is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentC() bool {
	if m.stringAt(1, "T", "S") {
		if m.stringAtStart("CONNECTICUT") ||
			m.stringAtStart("INDICT", "TUCSON") {
			m.current++
			return true
		}
	}

	return false
}

/**
 * Encodes slavic spellings or transliterations
 * written as "-CZ-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCZ() bool {
	if m.stringAt(1, "Z") &&
		!m.stringAt(-1, "ECZEMA") {
		if m.stringAt(0, "CZAR") {
			m.metaphAdd("S")
		} else {
			// otherwise most likely a czech word...
			m.metaphAdd("X")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * "-CS" special cases
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCS() bool {
	// give an 'etymological' 2nd
	// encoding for "kovacs" so
	// that it matches "kovach"
	if m.stringAtStart("KOVACS") {
		m.metaphAdd("KS", "X")
		m.current += 2
		return true
	}

	if m.stringAt(-1, "ACS") &&
		((m.current + 1) == m.last) &&
		!m.stringAt(-4, "ISAACS") {
		m.metaphAdd("X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes 'C'
 *
 */
func (m *Metaphone3) encodeC() {
	if m.encodeSilentCAtBeginning() ||
		m.encodeCAToS() ||
		m.encodeCOToS() ||
		m.encodeCH() ||
		m.encodeCCIA() ||
		m.encodeCC() ||
		m.encodeCKCGCQ() ||
		m.encodeCFrontVowel() ||
		m.encodeSilentC() ||
		m.encodeCZ() ||
		m.encodeCS() {
		return
	}

	//else
	if !m.stringAt(-1, "C", "K", "G", "Q") {
		m.metaphAdd("K")
	}

	//name sent in 'mac caffrey', 'mac gregor
	if m.stringAt(1, " C", " Q", " G") {
		m.current += 2
	} else {
		if m.stringAt(1, "C", "K", "Q") &&
			!m.stringAt(1, "CE", "CI") {
			m.current += 2
			// account for combinations such as Ro-ckc-liffe
			if m.stringAt(0, "C", "K", "Q") &&
				!m.stringAt(1, "CE", "CI") {
				m.current++
			}
		} else {
			m.current++
		}
	}
}

/**
 * Encode "-DG-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeDG() bool {
	if m.stringAt(0, "DG") {
		// excludes exceptions e.g. 'edgar',
		// or cases where 'g' is first letter of combining form
		// e.g. 'handgun', 'waldglas'
		if m.stringAt(2, "A", "O") ||
			// e.g. "midgut"
			m.stringAt(1, "GUN", "GUT") ||
			// e.g. "handgrip"
			m.stringAt(1, "GEAR", "GLAS", "GRIP", "GREN", "GILL", "GRAF") ||
			// e.g. "mudgard"
			m.stringAt(1, "GUARD", "GUILT", "GRAVE", "GRASS") ||
			// e.g. "woodgrouse"
			m.stringAt(1, "GROUSE") {
			m.metaphAddExactApprox("DG", "TK")
		} else {
			//e.g. "edge", "abridgment"
			m.metaphAdd("J")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-DJ-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeDJ() bool {
	// e.g. "adjacent"
	if m.stringAt(0, "DJ") {
		m.metaphAdd("J")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-DD-" and "-DT-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeDTDD() bool {
	// eat redundant 'T' or 'D'
	if m.stringAt(0, "DT", "DD") {
		if m.stringAt(0, "DTH") {
			m.metaphAddExactApprox("D0", "T0")
			m.current += 3
		} else {
			if m.encodeExact {
				// devoice it
				if m.stringAt(0, "DT") {
					m.metaphAdd("T")
				} else {
					m.metaphAdd("D")
				}
			} else {
				m.metaphAdd("T")
			}
			m.current += 2
		}
		return true
	}

	return false
}

/**
 * Encode cases where "-DU-" "-DI-", and "-DI-" => J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeDToJ() bool {
	// e.g. "module", "adulate"
	if (m.stringAt(0, "DUL") &&
		(m.isVowel(m.current-1) && m.isVowel(m.current+3))) ||
		// e.g. "soldier", "grandeur", "procedure"
		(((m.current + 3) == m.last) &&
			m.stringAt(-1, "LDIER", "NDEUR", "EDURE", "RDURE")) ||
		m.stringAt(-3, "CORDIAL") ||
		// e.g.  "pendulum", "education"
		m.stringAt(-1, "NDULA", "NDULU", "EDUCA") ||
		// e.g. "individual", "individual", "residuum"
		m.stringAt(-1, "ADUA", "IDUA", "IDUU") {
		m.metaphAddExactApprox("J", "D", "J", "T")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode latinate suffix "-DOUS" where 'D' is pronounced as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeDOUS() bool {
	// e.g. "assiduous", "arduous"
	if m.stringAt(1, "UOUS") {
		m.metaphAddExactApprox("J", "D", "J", "T")
		m.advanceCounter(4, 1)
		return true
	}

	return false
}

/**
 * Encode silent "-D-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentD() bool {
	// silent 'D' e.g. 'wednesday', 'handsome'
	if m.stringAt(-2, "WEDNESDAY") ||
		m.stringAt(-3, "HANDKER", "HANDSOM", "WINDSOR") ||
		// french silent D at end in words or names familiar to americans
		m.stringAt(-5, "PERNOD", "ARTAUD", "RENAUD") ||
		m.stringAt(-6, "RIMBAUD", "MICHAUD", "BICHAUD") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-D-"
 *
 */
func (m *Metaphone3) encodeD() {
	if m.encodeDG() ||
		m.encodeDJ() ||
		m.encodeDTDD() ||
		m.encodeDToJ() ||
		m.encodeDOUS() ||
		m.encodeSilentD() {
		return
	}

	if m.encodeExact {
		// "final de-voicing" in this case
		// e.g. 'missed' == 'mist'
		if (m.current == m.last) &&
			m.stringAt(-3, "SSED") {
			m.metaphAdd("T")
		} else {
			m.metaphAdd("D")
		}
	} else {
		m.metaphAdd("T")
	}
	m.current++
}

/**
 * Encode "-F-"
 *
 */
func (m *Metaphone3) encodeF() {
	// Encode cases where "-FT-" => "T" is usually silent
	// e.g. 'often', 'soften'
	// This should really be covered under "T"!
	if m.stringAt(-1, "OFTEN") {
		m.metaphAdd("F", "FT")
		m.current += 2
		return
	}

	// eat redundant 'F'
	if m.charAt(m.current+1) == 'F' {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAdd("F")

}

/**
 * Encode cases where 'G' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentGAtBeginning() bool {
	//skip these when at start of word
	if (m.current == 0) &&
		m.stringAt(0, "GN") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode "-GG-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGG() bool {
	if m.charAt(m.current+1) == 'G' {
		// italian e.g, 'loggia', 'caraveggio', also 'suggest' and 'exaggerate'
		if m.stringAt(-1, "AGGIA", "OGGIA", "AGGIO", "EGGIO", "EGGIA", "IGGIO") ||
			// 'ruggiero' but not 'snuggies'
			(m.stringAt(-1, "UGGIE") && !(((m.current + 3) == m.last) || ((m.current + 4) == m.last))) ||
			(((m.current + 2) == m.last) && m.stringAt(-1, "AGGI", "OGGI")) ||
			m.stringAt(-2, "SUGGES", "XAGGER", "REGGIE") {
			// expection where "-GG-" => KJ
			if m.stringAt(-2, "SUGGEST") {
				m.metaphAddExactApprox("G", "K")
			}

			m.metaphAdd("J")
			m.advanceCounter(3, 2)
		} else {
			m.metaphAddExactApprox("G", "K")
			m.current += 2
		}
		return true
	}

	return false
}

/**
 * Encode "-GK-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGK() bool {
	// 'gingko'
	if m.charAt(m.current+1) == 'K' {
		m.metaphAdd("K")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGHAfterConsonant() bool {
	// e.g. 'burgher', 'bingham'
	if (m.current > 0) &&
		!m.isVowel(m.current-1) &&
		// not e.g. 'greenhalgh'
		!(m.stringAt(-3, "HALGH") &&
			((m.current + 1) == m.last)) {
		m.metaphAddExactApprox("G", "K")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeInitialGH() bool {
	if m.current < 3 {
		// e.g. "ghislane", "ghiradelli"
		if m.current == 0 {
			if m.charAt(m.current+2) == 'I' {
				m.metaphAdd("J")
			} else {
				m.metaphAddExactApprox("G", "K")
			}
			m.current += 2
			return true
		}
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGHToJ() bool {
	// e.g., 'greenhalgh', 'dunkenhalgh', english names
	if m.stringAt(-2, "ALGH") && ((m.current + 1) == m.last) {
		m.metaphAdd("J", "")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGHToH() bool {
	// special cases
	// e.g., 'donoghue', 'donaghy'
	if (m.stringAt(-4, "DONO", "DONA") && m.isVowel(m.current+2)) ||
		m.stringAt(-5, "CALLAGHAN") {
		m.metaphAdd("H")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeUGHT() bool {
	//e.g. "ought", "aught", "daughter", "slaughter"
	if m.stringAt(-1, "UGHT") {
		if (m.stringAt(-3, "LAUGH") &&
			!(m.stringAt(-4, "SLAUGHT") ||
				m.stringAt(-3, "LAUGHTO"))) ||
			m.stringAt(-4, "DRAUGH") {
			m.metaphAdd("FT")
		} else {
			m.metaphAdd("T")
		}
		m.current += 3
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGHHPartOfOtherWord() bool {
	// if the 'H' is the beginning of another word or syllable
	if m.stringAt(1, "HOUS", "HEAD", "HOLE", "HORN", "HARN") {
		m.metaphAddExactApprox("G", "K")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentGH() bool {
	//Parker's rule (with some further refinements) - e.g., 'hugh'
	if ((((m.current > 1) && m.stringAt(-2, "B", "H", "D", "G", "L")) ||
		//e.g., 'bough'
		((m.current > 2) &&
			m.stringAt(-3, "B", "H", "D", "K", "W", "N", "P", "V") &&
			!m.stringAtStart("ENOUGH")) ||
		//e.g., 'broughton'
		((m.current > 3) && m.stringAt(-4, "B", "H")) ||
		//'plough', 'slaugh'
		((m.current > 3) && m.stringAt(-4, "PL", "SL")) ||
		((m.current > 0) &&
			// 'sigh', 'light'
			((m.charAt(m.current-1) == 'I') ||
				m.stringAtStart("PUGH") ||
				// e.g. 'MCDONAGH', 'MURTAGH', 'CREAGH'
				(m.stringAt(-1, "AGH") &&
					((m.current + 1) == m.last)) ||
				m.stringAt(-4, "GERAGH", "DRAUGH") ||
				(m.stringAt(-3, "GAUGH", "GEOGH", "MAUGH") &&
					!m.stringAtStart("MCGAUGHEY")) ||
				// exceptions to 'tough', 'rough', 'lough'
				(m.stringAt(-2, "OUGH") &&
					(m.current > 3) &&
					!m.stringAt(-4, "CCOUGH", "ENOUGH", "TROUGH", "CLOUGH"))))) &&
		// suffixes starting w/ vowel where "-GH-" is usually silent
		(m.stringAt(-3, "VAUGH", "FEIGH", "LEIGH") ||
			m.stringAt(-2, "HIGH", "TIGH") ||
			((m.current + 1) == m.last) ||
			(m.stringAt(2, "IE", "EY", "ES", "ER", "ED", "TY") &&
				((m.current + 3) == m.last) &&
				!m.stringAt(-5, "GALLAGHER")) ||
			(m.stringAt(2, "Y") && ((m.current + 2) == m.last)) ||
			(m.stringAt(2, "ING", "OUT") && ((m.current + 4) == m.last)) ||
			(m.stringAt(2, "ERTY") && ((m.current + 5) == m.last)) ||
			(!m.isVowel(m.current+2) ||
				m.stringAt(-3, "GAUGH", "GEOGH", "MAUGH") ||
				m.stringAt(-4, "BROUGHAM")))) &&
		// exceptions where '-g-' pronounced
		!(m.stringAtStart("BALOGH", "SABAGH") ||
			m.stringAt(-2, "BAGHDAD") ||
			m.stringAt(-3, "WHIGH") ||
			m.stringAt(-5, "SABBAGH", "AKHLAGH")) {
		// silent - do nothing
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGHSpecialCases() bool {
	var handled = false

	// special case: 'hiccough' == 'hiccup'
	if m.stringAt(-6, "HICCOUGH") {
		m.metaphAdd("P")
		handled = true
	} else if m.stringAtStart("LOUGH") {
		// special case: 'lough' alt spelling for scots 'loch'
		m.metaphAdd("K")
		handled = true
	} else if m.stringAtStart("BALOGH") {
		// hungarian
		m.metaphAddExactApprox("G", "", "K", "")
		handled = true
	} else if m.stringAt(-3, "LAUGHLIN", "COUGHLAN", "LOUGHLIN") {
		// "maclaughlin"
		m.metaphAdd("K", "F")
		handled = true
	} else if m.stringAt(-3, "GOUGH") ||
		m.stringAt(-7, "COLCLOUGH") {
		m.metaphAdd("", "F")
		handled = true
	}

	if handled {
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGHToF() bool {
	// the cases covered here would fall under
	// the GH_To_F rule below otherwise
	if m.encodeGHSpecialCases() {
		return true
	} else {
		//e.g., 'laugh', 'cough', 'rough', 'tough'
		if (m.current > 2) &&
			(m.charAt(m.current-1) == 'U') &&
			m.isVowel(m.current-2) &&
			m.stringAt(-3, "C", "G", "L", "R", "T", "N", "S") &&
			!m.stringAt(-4, "BREUGHEL", "FLAUGHER") {
			m.metaphAdd("F")
			m.current += 2
			return true
		}
	}

	return false
}

/**
 * Encode "-GH-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGH() bool {
	if m.charAt(m.current+1) == 'H' {
		if m.encodeGHAfterConsonant() ||
			m.encodeInitialGH() ||
			m.encodeGHToJ() ||
			m.encodeGHToH() ||
			m.encodeUGHT() ||
			m.encodeGHHPartOfOtherWord() ||
			m.encodeSilentGH() ||
			m.encodeGHToF() {
			return true
		}

		m.metaphAddExactApprox("G", "K")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode some contexts where "g" is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentG() bool {
	// e.g. "phlegm", "apothegm", "voigt"
	if (((m.current + 1) == m.last) &&
		(m.stringAt(-1, "EGM", "IGM", "AGM") ||
			m.stringAt(0, "GT"))) ||
		m.stringEqual("HUGES") {
		m.current++
		return true
	}

	// vietnamese names e.g. "Nguyen" but not "Ng"
	if m.stringAtStart("NG") && (m.current != m.last) {
		m.current++
		return true
	}

	return false
}

/**
 * ENcode "-GN-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGN() bool {
	if m.charAt(m.current+1) == 'N' {
		// 'align' 'sign', 'resign' but not 'resignation'
		// also 'impugn', 'impugnable', but not 'repugnant'
		if ((m.current > 1) &&
			((m.stringAt(-1, "I", "U", "E") ||
				m.stringAt(-3, "LORGNETTE") ||
				m.stringAt(-2, "LAGNIAPPE") ||
				m.stringAt(-2, "COGNAC") ||
				m.stringAt(-3, "CHAGNON") ||
				m.stringAt(-5, "COMPAGNIE") ||
				m.stringAt(-4, "BOLOGN")) &&
				// Exceptions: following are cases where 'G' is pronounced
				// in "assign" 'g' is silent, but not in "assignation"
				!(m.stringAt(2, "ATION") ||
					m.stringAt(2, "ATOR") ||
					m.stringAt(2, "ATE", "ITY") ||
					// exception to exceptions, not pronounced:
					(m.stringAt(2, "AN", "AC", "IA", "UM") &&
						!(m.stringAt(-3, "POIGNANT") ||
							m.stringAt(-2, "COGNAC"))) ||
					m.stringAtStart("SPIGNER", "STEGNER") ||
					m.stringEqual("SIGNE") ||
					m.stringAt(-2, "LIGNI", "LIGNO", "REGNA", "DIGNI", "WEGNE",
						"TIGNE", "RIGNE", "REGNE", "TIGNO") ||
					m.stringAt(-2, "SIGNAL", "SIGNIF", "SIGNAT") ||
					m.stringAt(-1, "IGNIT")) &&
				!m.stringAt(-2, "SIGNET", "LIGNEO"))) ||
			//not e.g. 'cagney', 'magna'
			(((m.current + 2) == m.last) &&
				m.stringAt(0, "GNE", "GNA") &&
				!m.stringAt(-2, "SIGNA", "MAGNA", "SIGNE")) {
			m.metaphAddExactApprox("N", "GN", "N", "KN")
		} else {
			m.metaphAddExactApprox("GN", "KN")
		}
		m.current += 2
		return true
	}
	return false
}

/**
 * Encode "-GL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGL() bool {
	//'tagliaro', 'puglia' BUT add K in alternative
	// since americans sometimes do this
	if m.stringAt(1, "LIA", "LIO", "LIE") &&
		m.isVowel(m.current-1) {
		m.metaphAddExactApprox("L", "GL", "L", "KL")
		m.current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) initialGSoft() bool {
	if ((m.stringAt(1, "EL", "EM", "EN", "EO", "ER", "ES", "IA", "IN", "IO", "IP", "IU", "YM", "YN", "YP", "YR", "EE") ||
		m.stringAt(1, "IRA", "IRO")) &&
		// except for smaller set of cases where => K, e.g. "gerber"
		!(m.stringAt(1, "ELD", "ELT", "ERT", "INZ", "ERH", "ITE", "ERD", "ERL", "ERN",
			"INT", "EES", "EEK", "ELB", "EER") ||
			m.stringAt(1, "ERSH", "ERST", "INSB", "INGR", "EROW", "ERKE", "EREN") ||
			m.stringAt(1, "ELLER", "ERDIE", "ERBER", "ESUND", "ESNER", "INGKO", "INKGO",
				"IPPER", "ESELL", "IPSON", "EEZER", "ERSON", "ELMAN") ||
			m.stringAt(1, "ESTALT", "ESTAPO", "INGHAM", "ERRITY", "ERRISH", "ESSNER", "ENGLER") ||
			m.stringAt(1, "YNAECOL", "YNECOLO", "ENTHNER", "ERAGHTY") ||
			m.stringAt(1, "INGERICH", "EOGHEGAN"))) ||
		(m.isVowel(m.current+1) &&
			(m.stringAt(1, "EE ", "EEW") ||
				(m.stringAt(1, "IGI", "IRA", "IBE", "AOL", "IDE", "IGL") &&
					!m.stringAt(1, "IDEON")) ||
				m.stringAt(1, "ILES", "INGI", "ISEL") ||
				(m.stringAt(1, "INGER") && !m.stringAt(1, "INGERICH")) ||
				m.stringAt(1, "IBBER", "IBBET", "IBLET", "IBRAN", "IGOLO", "IRARD", "IGANT") ||
				m.stringAt(1, "IRAFFE", "EEWHIZ") ||
				m.stringAt(1, "ILLETTE", "IBRALTA"))) {
		return true
	}

	return false
}

/**
 * Encode cases where 'G' is at start of word followed
 * by a "front" vowel e.g. 'E', 'I', 'Y'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeInitialGFrontVowel() bool {
	// 'g' followed by vowel at beginning
	if (m.current == 0) && m.frontVowel(m.current+1) {
		// special case "gila" as in "gila monster"
		if m.stringAt(1, "ILA") && (m.length == 4) {
			m.metaphAdd("H")
		} else if m.initialGSoft() {
			m.metaphAddExactApprox("J", "G", "J", "K")
		} else {
			// only code alternate 'J' if front vowel
			if (m.inWord[m.current+1] == 'E') || (m.inWord[m.current+1] == 'I') {
				m.metaphAddExactApprox("G", "J", "K", "J")
			} else {
				m.metaphAddExactApprox("G", "K")
			}
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-NGER-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeNGER() bool {
	if (m.current > 1) &&
		m.stringAt(-1, "NGER") {
		// default 'G' => J  such as 'ranger', 'stranger', 'manger', 'messenger', 'orangery', 'granger'
		// 'boulanger', 'challenger', 'danger', 'changer', 'harbinger', 'lounger', 'ginger', 'passenger'
		// except for these the following
		if !(m.rootOrInflections(m.inWord, "ANGER") ||
			m.rootOrInflections(m.inWord, "LINGER") ||
			m.rootOrInflections(m.inWord, "MALINGER") ||
			m.rootOrInflections(m.inWord, "FINGER") ||
			(m.stringAt(-3, "HUNG", "FING", "BUNG", "WING", "RING", "DING", "ZENG",
				"ZING", "JUNG", "LONG", "PING", "CONG", "MONG", "BANG",
				"GANG", "HANG", "LANG", "SANG", "SING", "WANG", "ZANG") &&
				// exceptions to above where 'G' => J
				!(m.stringAt(-6, "BOULANG", "SLESING", "KISSING", "DERRING") ||
					m.stringAt(-8, "SCHLESING") ||
					m.stringAt(-5, "SALING", "BELANG") ||
					m.stringAt(-6, "BARRING") ||
					m.stringAt(-6, "PHALANGER") ||
					m.stringAt(-4, "CHANG"))) ||
			m.stringAt(-4, "STING", "YOUNG") ||
			m.stringAt(-5, "STRONG") ||
			m.stringAtStart("UNG", "ENG", "ING") ||
			m.stringAt(0, "GERICH") ||
			m.stringAtStart("SENGER") ||
			m.stringAt(-3, "WENGER", "MUNGER", "SONGER", "KINGER") ||
			m.stringAt(-4, "FLINGER", "SLINGER", "STANGER", "STENGER", "KLINGER", "CLINGER") ||
			m.stringAt(-5, "SPRINGER", "SPRENGER") ||
			m.stringAt(-3, "LINGERF") ||
			m.stringAt(-2, "ANGERLY", "ANGERBO", "INGERSO")) {
			m.metaphAddExactApprox("J", "G", "J", "K")
		} else {
			m.metaphAddExactApprox("G", "J", "K", "J")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-GER-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGER() bool {
	if (m.current > 0) && m.stringAt(1, "ER") {
		// Exceptions to 'GE' where 'G' => K
		// e.g. "JAGER", "TIGER", "LIGER", "LAGER", "LUGER", "AUGER", "EAGER", "HAGER", "SAGER"
		if (((m.current == 2) && m.isVowel(m.current-1) && !m.isVowel(m.current-2) &&
			!(m.stringAt(-2, "PAGER", "WAGER", "NIGER", "ROGER", "LEGER", "CAGER")) ||
			m.stringAt(-2, "AUGER", "EAGER", "INGER", "YAGER")) ||
			m.stringAt(-3, "SEEGER", "JAEGER", "GEIGER", "KRUGER", "SAUGER", "BURGER",
				"MEAGER", "MARGER", "RIEGER", "YAEGER", "STEGER", "PRAGER", "SWIGER",
				"YERGER", "TORGER", "FERGER", "HILGER", "ZEIGER", "YARGER",
				"COWGER", "CREGER", "KROGER", "KREGER", "GRAGER", "STIGER", "BERGER") ||
			// 'berger' but not 'bergerac'
			(m.stringAt(-3, "BERGER") && ((m.current + 2) == m.last)) ||
			m.stringAt(-4, "KREIGER", "KRUEGER", "METZGER", "KRIEGER", "KROEGER", "STEIGER",
				"DRAEGER", "BUERGER", "BOERGER", "FIBIGER") ||
			// e.g. 'harshbarger', 'winebarger'
			(m.stringAt(-3, "BARGER") && (m.current > 4)) ||
			// e.g. 'weisgerber'
			(m.stringAt(0, "GERBER") && (m.current > 0)) ||
			m.stringAt(-5, "SCHWAGER", "LYBARGER", "SPRENGER", "GALLAGER", "WILLIGER") ||
			m.stringAtStart("HARGER") ||
			m.stringEqual("AGER", "EGER") ||
			m.stringAt(-1, "YGERNE") ||
			m.stringAt(-6, "SCHWEIGER")) &&
			!(m.stringAt(-5, "BELLIGEREN") ||
				m.stringAtStart("MARGERY") ||
				m.stringAt(-3, "BERGERAC")) {
			if m.slavoGermanic() {
				m.metaphAddExactApprox("G", "K")
			} else {
				m.metaphAddExactApprox("G", "J", "K", "J")
			}
		} else {
			m.metaphAddExactApprox("J", "G", "J", "K")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * ENcode "-GEL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGEL() bool {
	// more likely to be "-GEL-" => JL
	if m.stringAt(1, "EL") && (m.current > 0) {
		// except for
		// "BAGEL", "HEGEL", "HUGEL", "KUGEL", "NAGEL", "VOGEL", "FOGEL", "PAGEL"
		if ((m.length == 5) &&
			m.isVowel(m.current-1) &&
			!m.isVowel(m.current-2) &&
			!m.stringAt(-2, "NIGEL", "RIGEL")) ||
			// or the following as combining forms
			m.stringAt(-2, "ENGEL", "HEGEL", "NAGEL", "VOGEL") ||
			m.stringAt(-3, "MANGEL", "WEIGEL", "FLUGEL", "RANGEL", "HAUGEN", "RIEGEL", "VOEGEL") ||
			m.stringAt(-4, "SPEIGEL", "STEIGEL", "WRANGEL", "SPIEGEL") ||
			m.stringAt(-4, "DANEGELD") {
			if m.slavoGermanic() {
				m.metaphAddExactApprox("G", "K")
			} else {
				m.metaphAddExactApprox("G", "J", "K", "J")
			}
		} else {
			m.metaphAddExactApprox("J", "G", "J", "K")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/*
 * Detect german names and other words that have
 * a 'hard' 'g' in the context of "-ge" at end
 *
 * @return true if encoding handled in this routine, false if not
 */
func (m *Metaphone3) hardGEAtEnd() bool {
	return m.stringAtStart("RENEGE", "STONGE", "STANGE", "PRANGE", "KRESGE") ||
		m.stringAtStart("BYRGE", "BIRGE", "BERGE", "HAUGE") ||
		m.stringAtStart("HAGE") ||
		m.stringAtStart("LANGE", "SYNGE", "BENGE", "RUNGE", "HELGE") ||
		m.stringAtStart("INGE", "LAGE")
}

/**
 * Detect words where "-ge-" or "-gi-" get a 'hard' 'g'
 * even though this is usually a 'soft' 'g' context
 *
 * @return true if 'hard' 'g' detected
 *
 */
func (m *Metaphone3) internalHardGOther() bool {
	return (m.stringAt(0, "GETH", "GEAR", "GEIS", "GIRL", "GIVI", "GIVE", "GIFT",
		"GIRD", "GIRT", "GILV", "GILD", "GELD") &&
		!m.stringAt(-3, "GINGIV")) ||
		// "gish" but not "largish"
		(m.stringAt(1, "ISH") && (m.current > 0) && !m.stringAtStart("LARG")) ||
		(m.stringAt(-2, "MAGED", "MEGID") && !((m.current + 2) == m.last)) ||
		m.stringAt(0, "GEZ") ||
		m.stringAtStart("WEGE", "HAGE") ||
		(m.stringAt(-2, "ONGEST", "UNGEST") &&
			((m.current + 3) == m.last) &&
			!m.stringAt(-3, "CONGEST")) ||
		m.stringAtStart("VOEGE", "BERGE", "HELGE") ||
		m.stringEqual("ENGE", "BOGY") ||
		m.stringAt(0, "GIBBON") ||
		m.stringAtStart("CORREGIDOR") ||
		m.stringAtStart("INGEBORG") ||
		(m.stringAt(0, "GILL") &&
			(((m.current + 3) == m.last) || ((m.current + 4) == m.last)) &&
			!m.stringAtStart("STURGILL"))
}

/**
 * Detect words where "-gy-", "-gie-", "-gee-",
 * or "-gio-" get a 'hard' 'g' even though this is
 * usually a 'soft' 'g' context
 *
 * @return true if 'hard' 'g' detected
 *
 */
func (m *Metaphone3) internalHardGOpenSyllable() bool {
	return m.stringAt(1, "EYE") ||
		m.stringAt(-2, "FOGY", "POGY", "YOGI") ||
		m.stringAt(-2, "MAGEE", "MCGEE", "HAGIO") ||
		m.stringAt(-1, "RGEY", "OGEY") ||
		m.stringAt(-3, "HOAGY", "STOGY", "PORGY") ||
		m.stringAt(-5, "CARNEGIE") ||
		(m.stringAt(-1, "OGEY", "OGIE") && ((m.current + 2) == m.last))
}

/**
 * Detect a number of contexts, mostly german names, that
 * take a 'hard' 'g'.
 *
 * @return true if 'hard' 'g' detected, false if not
 *
 */
func (m *Metaphone3) internalHardGENGINGETGIT() bool {
	return (m.stringAt(-3, "FORGET", "TARGET", "MARGIT", "MARGET", "TURGEN",
		"BERGEN", "MORGEN", "JORGEN", "HAUGEN", "JERGEN",
		"JURGEN", "LINGEN", "BORGEN", "LANGEN", "KLAGEN", "STIGER", "BERGER") &&
		!m.stringAt(0, "GENETIC", "GENESIS") &&
		!m.stringAt(-4, "PLANGENT")) ||
		(m.stringAt(-3, "BERGIN", "FEAGIN", "DURGIN") && ((m.current + 2) == m.last)) ||
		(m.stringAt(-2, "ENGEN") && !m.stringAt(3, "DER", "ETI", "ESI")) ||
		m.stringAt(-4, "JUERGEN") ||
		m.stringAtStart("NAGIN", "MAGIN", "HAGIN") ||
		m.stringEqual("ENGIN", "DEGEN", "LAGEN", "MAGEN", "NAGIN") ||
		(m.stringAt(-2, "BEGET", "BEGIN", "HAGEN", "FAGIN",
			"BOGEN", "WIGIN", "NTGEN", "EIGEN",
			"WEGEN", "WAGEN") &&
			!m.stringAt(-5, "OSPHAGEN"))
}

/**
 * Detect a number of contexts of '-ng-' that will
 * take a 'hard' 'g' despite being followed by a
 * front vowel.
 *
 * @return true if 'hard' 'g' detected, false if not
 *
 */
func (m *Metaphone3) internalHardNG() bool {
	return (m.stringAt(-3, "DANG", "FANG", "SING") &&
		// exception to exception
		!m.stringAt(-5, "DISINGEN")) ||
		m.stringAtStart("INGEB", "ENGEB") ||
		(m.stringAt(-3, "RING", "WING", "HANG", "LONG") &&
			!(m.stringAt(-4, "CRING", "FRING", "ORANG", "TWING", "CHANG", "PHANG") ||
				m.stringAt(-5, "SYRING") ||
				m.stringAt(-3, "RINGENC", "RINGENT", "LONGITU", "LONGEVI") ||
				// e.g. 'longino', 'mastrangelo'
				(m.stringAt(0, "GELO", "GINO") && ((m.current + 3) == m.last)))) ||
		(m.stringAt(-1, "NGY") &&
			// exceptions to exception
			!(m.stringAt(-3, "RANGY", "MANGY", "MINGY") ||
				m.stringAt(-4, "SPONGY", "STINGY")))
}

/**
 * Exceptions to default encoding to 'J':
 * encode "-G-" to 'G' in "-g<frontvowel>-" words
 * where we are not at "-GE" at the end of the word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) internalHardG() bool {
	// if not "-GE" at end
	return !(((m.current + 1) == m.last) && (m.charAt(m.current+1) == 'E')) &&
		(m.internalHardNG() ||
			m.internalHardGENGINGETGIT() ||
			m.internalHardGOpenSyllable() ||
			m.internalHardGOther())
}

/**
 * Encode "-G-" followed by a vowel when non-initial leter.
 * Default for this is a 'J' sound, so check exceptions where
 * it is pronounced 'G'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeNonInitialGFrontVowel() bool {
	// -gy-, gi-, ge-
	if m.stringAt(1, "E", "I", "Y") {
		// '-ge' at end
		// almost always 'j 'sound
		if m.stringAt(0, "GE") && (m.current == (m.last - 1)) {
			if m.hardGEAtEnd() {
				if m.slavoGermanic() {
					m.metaphAddExactApprox("G", "K")
				} else {
					m.metaphAddExactApprox("G", "J", "K", "J")
				}
			} else {
				m.metaphAdd("J")
			}
		} else {
			if m.internalHardG() {
				// don't encode KG or KK if e.g. "mcgill"
				if !((m.current == 2) && m.stringAtStart("MC")) ||
					((m.current == 3) && m.stringAtStart("MAC")) {
					if m.slavoGermanic() {
						m.metaphAddExactApprox("G", "K")
					} else {
						m.metaphAddExactApprox("G", "J", "K", "J")
					}
				}
			} else {
				m.metaphAddExactApprox("J", "G", "J", "K")
			}
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode special case where "-GA-" => J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGAToJ() bool {
	// 'margary', 'margarine'
	if (m.stringAt(-3, "MARGARY", "MARGARI") &&
		// but not in spanish forms such as "margatita"
		!m.stringAt(-3, "MARGARIT")) ||
		m.stringAtStart("GAOL") ||
		m.stringAt(-2, "ALGAE") {
		m.metaphAddExactApprox("J", "G", "J", "K")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-G-"
 *
 */
func (m *Metaphone3) encodeG() {
	if m.encodeSilentGAtBeginning() ||
		m.encodeGG() ||
		m.encodeGK() ||
		m.encodeGH() ||
		m.encodeSilentG() ||
		m.encodeGN() ||
		m.encodeGL() ||
		m.encodeInitialGFrontVowel() ||
		m.encodeNGER() ||
		m.encodeGER() ||
		m.encodeGEL() ||
		m.encodeNonInitialGFrontVowel() ||
		m.encodeGAToJ() {
		return
	}

	if !m.stringAt(-1, "C", "K", "G", "Q") {
		m.metaphAddExactApprox("G", "K")
	}

	m.current++
}

/**
 * Encode cases where initial 'H' is not pronounced (in American)
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeInitialSilentH() bool {
	//'hour', 'herb', 'heir', 'honor'
	if m.stringAt(1, "OUR", "ERB", "EIR") ||
		m.stringAt(1, "ONOR") ||
		m.stringAt(1, "ONOUR", "ONEST") {
		// british pronounce H in this word
		// americans give it 'H' for the name,
		// no 'H' for the plant
		if (m.current == 0) && m.stringAt(0, "HERB") {
			if m.encodeVowels {
				m.metaphAdd("HA", "A")
			} else {
				m.metaphAdd("H", "A")
			}
		} else if (m.current == 0) || m.encodeVowels {
			m.metaphAdd("A")
		}

		// don't encode vowels twice
		m.current = m.skipVowels(m.current + 1)
		return true
	}

	return false
}

/**
 * Encode "HS-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeInitialHS() bool {
	// old chinese pinyin transliteration
	// e.g., 'HSIAO'
	if (m.current == 0) && m.stringAtStart("HS") {
		m.metaphAdd("X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode cases where "HU-" is pronounced as part of a vowel dipthong
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeInitialHUHW() bool {
	// spanish spellings and chinese pinyin transliteration
	if m.stringAtStart("HUA", "HUE", "HWA") {
		if !m.stringAt(0, "HUEY") {
			m.metaphAdd("A")

			if !m.encodeVowels {
				m.current += 3
			} else {
				m.current++
				// don't encode vowels twice
				for m.isVowel(m.current) || (m.charAt(m.current) == 'W') {
					m.current++
				}
			}
			return true
		}
	}

	return false
}

/**
 * Encode cases where 'H' is silent between vowels
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeNonInitialSilentH() bool {
	//exceptions - 'h' not pronounced
	// "PROHIB" BUT NOT "PROHIBIT"
	if m.stringAt(-2, "NIHIL", "VEHEM", "LOHEN", "NEHEM",
		"MAHON", "MAHAN", "COHEN", "GAHAN") ||
		m.stringAt(-3, "GRAHAM", "PROHIB", "FRAHER",
			"TOOHEY", "TOUHEY") ||
		m.stringAt(-3, "TOUHY") ||
		m.stringAtStart("CHIHUAHUA") {
		if !m.encodeVowels {
			m.current += 2
		} else {
			// don't encode vowels twice
			m.current = m.skipVowels(m.current + 1)
		}
		return true
	}

	return false
}

/**
 * Encode cases where 'H' is pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeHPronounced() bool {
	if (((m.current == 0) ||
		m.isVowel(m.current-1) ||
		((m.current > 0) &&
			(m.charAt(m.current-1) == 'W'))) &&
		m.isVowel(m.current+1)) ||
		// e.g. 'alWahhab'
		((m.charAt(m.current+1) == 'H') && m.isVowel(m.current+2)) {
		m.metaphAdd("H")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode 'H'
 *
 *
 */
func (m *Metaphone3) encodeH() {
	if m.encodeInitialSilentH() ||
		m.encodeInitialHS() ||
		m.encodeInitialHUHW() ||
		m.encodeNonInitialSilentH() {
		return
	}

	//only keep if first & before vowel or btw. 2 vowels
	if !m.encodeHPronounced() {
		//also takes care of 'HH'
		m.current++
	}
}

/**
 * Encode cases where initial or medial "j" is in a spanish word or name
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSpanishJ() bool {
	//obvious spanish, e.g. "jose", "san jacinto"
	if (m.stringAt(1, "UAN", "ACI", "ALI", "EFE", "ICA", "IME", "OAQ", "UAR") &&
		!m.stringAt(0, "JIMERSON", "JIMERSEN")) ||
		(m.stringAt(1, "OSE") && ((m.current + 3) == m.last)) ||
		m.stringAt(1, "EREZ", "UNTA", "AIME", "AVIE", "AVIA") ||
		m.stringAt(1, "IMINEZ", "ARAMIL") ||
		(((m.current + 2) == m.last) && m.stringAt(-2, "MEJIA")) ||
		m.stringAt(-2, "TEJED", "TEJAD", "LUJAN", "FAJAR", "BEJAR", "BOJOR", "CAJIG",
			"DEJAS", "DUJAR", "DUJAN", "MIJAR", "MEJOR", "NAJAR",
			"NOJOS", "RAJED", "RIJAL", "REJON", "TEJAN", "UIJAN") ||
		m.stringAt(-3, "ALEJANDR", "GUAJARDO", "TRUJILLO") ||
		(m.stringAt(-2, "RAJAS") && (m.current > 2)) ||
		(m.stringAt(-2, "MEJIA") && !m.stringAt(-2, "MEJIAN")) ||
		m.stringAt(-1, "OJEDA") ||
		m.stringAt(-3, "LEIJA", "MINJA") ||
		m.stringAt(-3, "VIAJES", "GRAJAL") ||
		m.stringAt(0, "JAUREGUI") ||
		m.stringAt(-4, "HINOJOSA") ||
		m.stringAtStart("SAN ") ||
		(((m.current + 1) == m.last) &&
			(m.charAt(m.current+1) == 'O') &&
			// exceptions
			!(m.stringAtStart("TOJO") ||
				m.stringAtStart("BANJO") ||
				m.stringAtStart("MARYJO"))) {
		// americans pronounce "juan" as 'wan'
		// and "marijuana" and "tijuana" also
		// do not get the 'H' as in spanish, so
		// just treat it like a vowel in these cases
		if !(m.stringAt(0, "JUAN") || m.stringAt(0, "JOAQ")) {
			m.metaphAdd("H")
		} else {
			if m.current == 0 {
				m.metaphAdd("A")
			}
		}
		m.advanceCounter(2, 1)
		return true
	}

	// Jorge gets 2nd HARHA. also JULIO, JESUS
	if m.stringAt(1, "ORGE", "ULIO", "ESUS") &&
		!m.stringAtStart("JORGEN") {
		// get both consonants for "jorge"
		if ((m.current + 4) == m.last) && m.stringAt(1, "ORGE") {
			if m.encodeVowels {
				m.metaphAdd("JARJ", "HARHA")
			} else {
				m.metaphAdd("JRJ", "HRH")
			}
			m.advanceCounter(5, 5)
			return true
		}

		m.metaphAdd("J", "H")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode cases where 'J' is clearly in a german word or name
 * that americans pronounce in the german fashion
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGermanJ() bool {
	if m.stringAt(1, "AH") ||
		(m.stringAt(1, "OHANN") && ((m.current + 5) == m.last)) ||
		(m.stringAt(1, "UNG") && !m.stringAt(1, "UNGL")) ||
		m.stringAt(1, "UGO") {
		m.metaphAdd("A")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-JOJ-" and "-JUJ-" as spanish words
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSpanishOJUJ() bool {
	if m.stringAt(1, "OJOBA", "UJUY ") {
		if m.encodeVowels {
			m.metaphAdd("HAH")
		} else {
			m.metaphAdd("HH")
		}

		m.advanceCounter(4, 3)
		return true
	}

	return false
}

/**
 * Encode 'J' toward end in spanish words
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSpanishJ2() bool {
	// spanish forms e.g. "brujo", "badajoz"
	if (((m.current - 2) == 0) &&
		m.stringAt(-2, "BOJA", "BAJA", "BEJA", "BOJO", "MOJA", "MOJI", "MEJI")) ||
		(((m.current - 3) == 0) &&
			m.stringAt(-3, "FRIJO", "BRUJO", "BRUJA", "GRAJE", "GRIJA", "LEIJA", "QUIJA")) ||
		(((m.current + 3) == m.last) &&
			m.stringAt(-1, "AJARA")) ||
		(((m.current + 2) == m.last) &&
			m.stringAt(-1, "AJOS", "EJOS", "OJAS", "OJOS", "UJON", "AJOZ", "AJAL", "UJAR", "EJON", "EJAN")) ||
		(((m.current + 1) == m.last) &&
			(m.stringAt(-1, "OJA", "EJA") && !m.stringAtStart("DEJA"))) {
		m.metaphAdd("H")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode 'J' as vowel in some exception cases
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeJAsVowel() bool {
	if m.stringAt(0, "JEWSK") {
		m.metaphAdd("J", "")
		return true
	}

	// e.g. "stijl", "sejm" - dutch, scandanavian, and eastern european spellings
	return (m.stringAt(1, "L", "T", "K", "S", "N", "M") &&
		// except words from hindi and arabic
		!m.stringAt(2, "A")) ||
		m.stringAtStart("HALLELUJA", "LJUBLJANA") ||
		m.stringAtStart("LJUB", "BJOR") ||
		m.stringAtStart("HAJEK") ||
		m.stringAtStart("WOJ") ||
		// e.g. 'fjord'
		m.stringAtStart("FJ") ||
		// e.g. 'rekjavik', 'blagojevic'
		m.stringAt(0, "JAVIK", "JEVIC") ||
		(((m.current + 1) == m.last) && m.stringAtStart("SONJA", "TANJA", "TONJA"))
}

/**
 * Test whether the word in question
 * is a name starting with 'J' that should
 * match names starting with a 'Y' sound.
 * All forms of 'John', 'Jane', etc, get
 * and alt to match e.g. 'Ian', 'Yana'. Joelle
 * should match 'Yael', 'Joseph' should match
 * 'Yusef'. German and slavic last names are
 * also included.
 *
 * @return true if name starting with 'J' that
 * should get an alternate encoding as a vowel
 */
func (m *Metaphone3) namesBeginningWithJThatGetAltY() bool {
	if m.stringAtStart("JAN", "JON", "JAN", "JIN", "JEN") ||
		m.stringAtStart("JUHL", "JULY", "JOEL", "JOHN", "JOSH",
			"JUDE", "JUNE", "JONI", "JULI", "JENA",
			"JUNG", "JINA", "JANA", "JENI", "JOEL",
			"JANN", "JONA", "JENE", "JULE", "JANI",
			"JONG", "JOHN", "JEAN", "JUNG", "JONE",
			"JARA", "JUST", "JOST", "JAHN", "JACO",
			"JANG", "JUDE", "JONE") ||
		m.stringAtStart("JOANN", "JANEY", "JANAE", "JOANA", "JUTTA",
			"JULEE", "JANAY", "JANEE", "JETTA", "JOHNA",
			"JOANE", "JAYNA", "JANES", "JONAS", "JONIE",
			"JUSTA", "JUNIE", "JUNKO", "JENAE", "JULIO",
			"JINNY", "JOHNS", "JACOB", "JETER", "JAFFE",
			"JESKE", "JANKE", "JAGER", "JANIK", "JANDA",
			"JOSHI", "JULES", "JANTZ", "JEANS", "JUDAH",
			"JANUS", "JENNY", "JENEE", "JONAH", "JONAS",
			"JACOB", "JOSUE", "JOSEF", "JULES", "JULIE",
			"JULIA", "JANIE", "JANIS", "JENNA", "JANNA",
			"JEANA", "JENNI", "JEANE", "JONNA") ||
		m.stringAtStart("JORDAN", "JORDON", "JOSEPH", "JOSHUA", "JOSIAH",
			"JOSPEH", "JUDSON", "JULIAN", "JULIUS", "JUNIOR",
			"JUDITH", "JOESPH", "JOHNIE", "JOANNE", "JEANNE",
			"JOANNA", "JOSEFA", "JULIET", "JANNIE", "JANELL",
			"JASMIN", "JANINE", "JOHNNY", "JEANIE", "JEANNA",
			"JOHNNA", "JOELLE", "JOVITA", "JOSEPH", "JONNIE",
			"JANEEN", "JANINA", "JOANIE", "JAZMIN", "JOHNIE",
			"JANENE", "JOHNNY", "JONELL", "JENELL", "JANETT",
			"JANETH", "JENINE", "JOELLA", "JOEANN", "JULIAN",
			"JOHANA", "JENICE", "JANNET", "JANISE", "JULENE",
			"JOSHUA", "JANEAN", "JAIMEE", "JOETTE", "JANYCE",
			"JENEVA", "JORDAN", "JACOBS", "JENSEN", "JOSEPH",
			"JANSEN", "JORDON", "JULIAN", "JAEGER", "JACOBY",
			"JENSON", "JARMAN", "JOSLIN", "JESSEN", "JAHNKE",
			"JACOBO", "JULIEN", "JOSHUA", "JEPSON", "JULIUS",
			"JANSON", "JACOBI", "JUDSON", "JARBOE", "JOHSON",
			"JANZEN", "JETTON", "JUNKER", "JONSON", "JAROSZ",
			"JENNER", "JAGGER", "JASMIN", "JEPSEN", "JORDEN",
			"JANNEY", "JUHASZ", "JERGEN") ||
		m.stringAtStart("JAKOB") ||
		m.stringAtStart("JOHNSON", "JOHNNIE", "JASMINE", "JEANNIE", "JOHANNA",
			"JANELLE", "JANETTE", "JULIANA", "JUSTINA", "JOSETTE",
			"JOELLEN", "JENELLE", "JULIETA", "JULIANN", "JULISSA",
			"JENETTE", "JANETTA", "JOSELYN", "JONELLE", "JESENIA",
			"JANESSA", "JAZMINE", "JEANENE", "JOANNIE", "JADWIGA",
			"JOLANDA", "JULIANE", "JANUARY", "JEANICE", "JANELLA",
			"JEANETT", "JENNINE", "JOHANNE", "JOHNSIE", "JANIECE",
			"JOHNSON", "JENNELL", "JAMISON", "JANSSEN", "JOHNSEN",
			"JARDINE", "JAGGERS", "JURGENS", "JOURDAN", "JULIANO",
			"JOSEPHS", "JHONSON", "JOZWIAK", "JANICKI", "JELINEK",
			"JANSSON", "JOACHIM", "JANELLE", "JACOBUS", "JENNING",
			"JANTZEN", "JOHNNIE") ||
		m.stringAtStart("JOSEFINA", "JEANNINE", "JULIANNE", "JULIANNA", "JONATHAN",
			"JONATHON", "JEANETTE", "JANNETTE", "JEANETTA", "JOHNETTA",
			"JENNEFER", "JULIENNE", "JOSPHINE", "JEANELLE", "JOHNETTE",
			"JULIEANN", "JOSEFINE", "JULIETTA", "JOHNSTON", "JACOBSON",
			"JACOBSEN", "JOHANSEN", "JOHANSON", "JAWORSKI", "JENNETTE",
			"JELLISON", "JOHANNES", "JASINSKI", "JUERGENS", "JARNAGIN",
			"JEREMIAH", "JEPPESEN", "JARNIGAN", "JANOUSEK") ||
		m.stringAtStart("JOHNATHAN", "JOHNATHON", "JORGENSEN", "JEANMARIE", "JOSEPHINA",
			"JEANNETTE", "JOSEPHINE", "JEANNETTA", "JORGENSON", "JANKOWSKI",
			"JOHNSTONE", "JABLONSKI", "JOSEPHSON", "JOHANNSEN", "JURGENSEN",
			"JIMMERSON", "JOHANSSON") ||
		m.stringAtStart("JAKUBOWSKI") {
		return true
	}

	return false
}

/**
 * Encode 'J' => J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeJToJ() bool {
	if m.isVowel(m.current + 1) {
		if (m.current == 0) &&
			m.namesBeginningWithJThatGetAltY() {
			// 'Y' is a vowel so encode
			// is as 'A'
			if m.encodeVowels {
				m.metaphAdd("JA", "A")
			} else {
				m.metaphAdd("J", "A")
			}
		} else {
			if m.encodeVowels {
				m.metaphAdd("JA")
			} else {
				m.metaphAdd("J")
			}
		}

		m.current = m.skipVowels(m.current + 1)
		return false
	} else {
		m.metaphAdd("J")
		m.current++
		return true
	}

	//		return false
}

/**
 * Call routines to encode 'J', in proper order
 *
 */
func (m *Metaphone3) encodeOtherJ() {
	if m.current == 0 {
		if m.encodeGermanJ() {
			return
		} else {
			if m.encodeJToJ() {
				return
			}
		}
	} else {
		if m.encodeSpanishJ2() {
			return
		} else if !m.encodeJAsVowel() {
			m.metaphAdd("J")
		}

		//it could happen! e.g. "hajj"
		// eat redundant 'J'
		if m.charAt(m.current+1) == 'J' {
			m.current += 2
		} else {
			m.current++
		}
	}
}

/**
 * Encode 'J'
 *
 */
func (m *Metaphone3) encodeJ() {
	if m.encodeSpanishJ() || m.encodeSpanishOJUJ() {
		return
	}

	m.encodeOtherJ()
}

/**
 * Encode cases where 'K' is not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentK() bool {
	//skip this except for special cases
	if (m.current == 0) && m.stringAt(0, "KN") {
		if !(m.stringAt(2, "ESSET", "IEVEL") || m.stringAt(2, "ISH")) {
			m.current += 1
			return true
		}
	}

	// e.g. "know", "knit", "knob"
	if (m.stringAt(1, "NOW", "NIT", "NOT", "NOB") &&
		// exception, "slipknot" => SLPNT but "banknote" => PNKNT
		!m.stringAtStart("BANKNOTE")) ||
		m.stringAt(1, "NOCK", "NUCK", "NIFE", "NACK") ||
		m.stringAt(1, "NIGHT") {
		// N already encoded before
		// e.g. "penknife"
		if (m.current > 0) && m.charAt(m.current-1) == 'N' {
			m.current += 2
		} else {
			m.current++
		}

		return true
	}

	return false
}

/**
 * Encode 'K'
 *
 *
 */
func (m *Metaphone3) encodeK() {
	if !m.encodeSilentK() {
		m.metaphAdd("K")

		// eat redundant 'K's and 'Q's
		if (m.charAt(m.current+1) == 'K') ||
			(m.charAt(m.current+1) == 'Q') {
			m.current += 2
		} else {
			m.current++
		}
	}
}

/**
 * Cases where an L follows D, G, or T at the
 * end have a schwa pronounced before the L
 *
 */
func (m *Metaphone3) interpolateVowelWhenConsLAtEnd() {
	if m.encodeVowels == true {
		// e.g. "ertl", "vogl"
		if (m.current == m.last) &&
			m.stringAt(-1, "D", "G", "T") {
			m.metaphAdd("A")
		}
	}
}

/**
 * Catch cases where 'L' spelled twice but pronounced
 * once, e.g., 'DOCILELY' => TSL
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeLELYToL() bool {
	// e.g. "agilely", "docilely"
	if m.stringAt(-1, "ILELY") &&
		((m.current + 3) == m.last) {
		m.metaphAdd("L")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode special case "colonel" => KRNL. Can somebody tell
 * me how this pronounciation came to be?
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCOLONEL() bool {
	if m.stringAt(-2, "COLONEL") {
		m.metaphAdd("R")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-AULT-", found in a french names
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeFrenchAULT() bool {
	// e.g. "renault" and "foucault", well known to americans, but not "fault"
	if (m.current > 3) &&
		(m.stringAt(-3, "RAULT", "NAULT", "BAULT", "SAULT", "GAULT", "CAULT") ||
			m.stringAt(-4, "REAULT", "RIAULT", "NEAULT", "BEAULT")) &&
		!(m.rootOrInflections(m.inWord, "ASSAULT") ||
			m.stringAt(-8, "SOMERSAULT") ||
			m.stringAt(-9, "SUMMERSAULT")) {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-EUIL-", always found in a french word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeFrenchEUIL() bool {
	// e.g. "auteuil"
	if m.stringAt(-3, "EUIL") && (m.current == m.last) {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-OULX", always found in a french word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeFrenchOULX() bool {
	// e.g. "proulx"
	if m.stringAt(-2, "OULX") && ((m.current + 1) == m.last) {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encodes contexts where 'L' is not pronounced in "-LM-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentLInLM() bool {
	if m.stringAt(0, "LM", "LN") {
		// e.g. "lincoln", "holmes", "psalm", "salmon"
		if (m.stringAt(-2, "COLN", "CALM", "BALM", "MALM", "PALM") ||
			(m.stringAt(-1, "OLM") && ((m.current + 1) == m.last)) ||
			m.stringAt(-3, "PSALM", "QUALM") ||
			m.stringAt(-2, "SALMON", "HOLMES") ||
			m.stringAt(-1, "ALMOND") ||
			((m.current == 1) && m.stringAt(-1, "ALMS"))) &&
			(!m.stringAt(2, "A") &&
				!m.stringAt(-2, "BALMO") &&
				!m.stringAt(-2, "PALMER", "PALMOR", "BALMER") &&
				!m.stringAt(-3, "THALM")) {
			m.current++
			return true
		} else {
			m.metaphAdd("L")
			m.current++
			return true
		}
	}

	return false
}

/**
 * Encodes contexts where '-L-' is silent in 'LK', 'LV'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentLInLKLV() bool {
	if (m.stringAt(-2, "WALK", "YOLK", "FOLK", "HALF", "TALK", "CALF", "BALK", "CALK") ||
		(m.stringAt(-2, "POLK") &&
			!m.stringAt(-2, "POLKA", "WALKO")) ||
		(m.stringAt(-2, "HALV") &&
			!m.stringAt(-2, "HALVA", "HALVO")) ||
		(m.stringAt(-3, "CAULK", "CHALK", "BAULK", "FAULK") &&
			!m.stringAt(-4, "SCHALK")) ||
		(m.stringAt(-2, "SALVE", "CALVE") ||
			m.stringAt(-2, "SOLDER")) &&
			// exceptions to above cases where 'L' is usually pronounced
			!m.stringAt(-2, "SALVER", "CALVER")) &&
		!m.stringAt(-5, "GONSALVES", "GONCALVES") &&
		!m.stringAt(-2, "BALKAN", "TALKAL") &&
		!m.stringAt(-3, "PAULK", "CHALF") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode 'L' in contexts of "-OULD-" where it is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentLInOULD() bool {
	//'would', 'could'
	if m.stringAt(-3, "WOULD", "COULD") ||
		(m.stringAt(-4, "SHOULD") &&
			!m.stringAt(-4, "SHOULDER")) {
		m.metaphAddExactApprox("D", "T")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-ILLA-" and "-ILLE-" in spanish and french
 * contexts were americans know to pronounce it as a 'Y'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeLLAsVowelSpecialCases() bool {
	if m.stringAt(-5, "TORTILLA") ||
		m.stringAt(-8, "RATATOUILLE") ||
		// e.g. 'guillermo', "veillard"
		(m.stringAtStart("GUILL", "VEILL", "GAILL") &&
			// 'guillotine' usually has '-ll-' pronounced as 'L' in english
			!(m.stringAt(-3, "GUILLOT", "GUILLOR", "GUILLEN") ||
				m.stringEqual("GUILL"))) ||
		// e.g. "brouillard", "gremillion"
		m.stringAtStart("BROUILL", "GREMILL") ||
		m.stringAtStart("ROBILL") ||
		// e.g. 'mireille'
		(m.stringAt(-2, "EILLE") &&
			((m.current + 2) == m.last) &&
			// exception "reveille" usually pronounced as 're-vil-lee'
			!m.stringAt(-5, "REVEILLE")) {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode other spanish cases where "-LL-" is pronounced as 'Y'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeLLAsVowel() bool {
	//spanish e.g. "cabrillo", "gallegos" but also "gorilla", "ballerina" -
	// give both pronounciations since an american might pronounce "cabrillo"
	// in the spanish or the american fashion.
	if (((m.current + 3) == m.length) &&
		m.stringAt(-1, "ILLO", "ILLA", "ALLE")) ||
		(((m.stringAtEnd("AS", "OS") ||
			m.stringAtEnd("A", "O")) &&
			m.stringAt(-1, "AL", "IL")) &&
			!m.stringAt(-1, "ALLA")) ||
		m.stringAtStart("VILLE", "VILLA") ||
		m.stringAtStart("GALLARDO", "VALLADAR", "MAGALLAN", "CAVALLAR", "BALLASTE") ||
		m.stringAtStart("LLA") {
		m.metaphAdd("L", "")
		m.current += 2
		return true
	}
	return false
}

/**
 * Call routines to encode "-LL-", in proper order
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeLLAsVowelCases() bool {
	if m.charAt(m.current+1) == 'L' {
		if m.encodeLLAsVowelSpecialCases() {
			return true
		} else if m.encodeLLAsVowel() {
			return true
		}
		m.current += 2

	} else {
		m.current++
	}

	return false
}

/**
 * Encode vowel-encoding cases where "-LE-" is pronounced "-EL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeVowelLETransposition(save_current int) bool {
	// transposition of vowel sound and L occurs in many words,
	// e.g. "bristle", "dazzle", "goggle" => KAKAL
	if m.encodeVowels && (save_current > 1) &&
		!m.isVowel(save_current-1) &&
		(m.charAt(save_current+1) == 'E') &&
		(m.charAt(save_current-1) != 'L') &&
		(m.charAt(save_current-1) != 'R') &&
		// lots of exceptions to this:
		!m.isVowel(save_current+2) &&
		!m.stringAtStart("ECCLESI", "COMPLEC", "COMPLEJ", "ROBLEDO") &&
		!m.stringAtStart("MCCLE", "MCLEL") &&
		!m.stringAtStart("EMBLEM", "KADLEC") &&
		!(((save_current + 2) == m.last) && m.stringAt(save_current, "LET")) &&
		!m.stringAt(save_current, "LETTING") &&
		!m.stringAt(save_current, "LETELY", "LETTER", "LETION", "LETIAN", "LETING", "LETORY") &&
		!m.stringAt(save_current, "LETUS", "LETIV") &&
		!m.stringAt(save_current, "LESS", "LESQ", "LECT", "LEDG", "LETE", "LETH", "LETS", "LETT") &&
		!m.stringAt(save_current, "LEG", "LER", "LEX") &&
		// e.g. "complement" !=> KAMPALMENT
		!(m.stringAt(save_current, "LEMENT") &&
			!(m.stringAt(-5, "BATTLE", "TANGLE", "PUZZLE", "RABBLE", "BABBLE") ||
				m.stringAt(-4, "TABLE"))) &&
		!(((save_current + 2) == m.last) && m.stringAt((save_current-2), "OCLES", "ACLES", "AKLES")) &&
		!m.stringAt((save_current-3), "LISLE", "AISLE") &&
		!m.stringAtStart("ISLE") &&
		!m.stringAtStart("ROBLES") &&
		!m.stringAt((save_current-4), "PROBLEM", "RESPLEN") &&
		!m.stringAt((save_current-3), "REPLEN") &&
		!m.stringAt((save_current-2), "SPLE") &&
		(m.charAt(save_current-1) != 'H') &&
		(m.charAt(save_current-1) != 'W') {
		m.metaphAdd("AL")
		m.flag_AL_inversion = true

		// eat redundant 'L'
		if m.charAt(save_current+2) == 'L' {
			m.current = save_current + 3
		}
		return true
	}

	return false
}

/**
 * Encode special vowel-encoding cases where 'E' is not
 * silent at the end of a word as is the usual case
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeVowelPreserveVowelAfterL(save_current int) bool {
	// an example of where the vowel would NOT need to be preserved
	// would be, say, "hustled", where there is no vowel pronounced
	// between the 'l' and the 'd'
	if m.encodeVowels &&
		!m.isVowel(save_current-1) &&
		(m.charAt(save_current+1) == 'E') &&
		(save_current > 1) &&
		((save_current + 1) != m.last) &&
		!(m.stringAt((save_current+1), "ES", "ED") &&
			((save_current + 2) == m.last)) &&
		!m.stringAt((save_current-1), "RLEST") {
		m.metaphAdd("LA")
		m.current = m.skipVowels(m.current + 1)
		return true
	}

	return false
}

/**
 * Call routines to encode "-LE-", in proper order
 *
 * @param save_current index of actual current letter
 *
 */
func (m *Metaphone3) encodeLECases(save_current int) {
	if m.encodeVowelLETransposition(save_current) {
		return
	} else {
		if m.encodeVowelPreserveVowelAfterL(save_current) {
			return
		} else {
			m.metaphAdd("L")
		}
	}
}

/**
 * Encode 'L'
 *
 * Includes special vowel transposition
 * encoding, where 'LE' => AL
 *
 */
func (m *Metaphone3) encodeL() {
	// logic below needs to know this
	// after 'm.current' variable changed
	var save_current = m.current

	m.interpolateVowelWhenConsLAtEnd()

	if m.encodeLELYToL() ||
		m.encodeCOLONEL() ||
		m.encodeFrenchAULT() ||
		m.encodeFrenchEUIL() ||
		m.encodeFrenchOULX() ||
		m.encodeSilentLInLM() ||
		m.encodeSilentLInLKLV() ||
		m.encodeSilentLInOULD() {
		return
	}

	if m.encodeLLAsVowelCases() {
		return
	}

	m.encodeLECases(save_current)
}

/**
 * Encode cases where 'M' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentMAtBeginning() bool {
	//skip these when at start of word
	if (m.current == 0) && m.stringAt(0, "MN") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode special cases "Mr." and "Mrs."
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeMRAndMRS() bool {
	if (m.current == 0) && m.stringAt(0, "MR") {
		// exceptions for "mr." and "mrs."
		if (m.length == 2) && m.stringAt(0, "MR") {
			if m.encodeVowels {
				m.metaphAdd("MASTAR")
			} else {
				m.metaphAdd("MSTR")
			}
			m.current += 2
			return true
		} else if (m.length == 3) && m.stringAt(0, "MRS") {
			if m.encodeVowels {
				m.metaphAdd("MASAS")
			} else {
				m.metaphAdd("MSS")
			}
			m.current += 3
			return true
		}
	}

	return false
}

/**
 * Encode "Mac-" and "Mc-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeMAC() bool {
	// should only find irish and
	// scottish names e.g. 'macintosh'
	if (m.current == 0) &&
		(m.stringAtStart("MACIVER", "MACEWEN") ||
			m.stringAtStart("MACELROY", "MACILROY") ||
			m.stringAtStart("MACINTOSH") ||
			m.stringAtStart("MC")) {
		if m.encodeVowels {
			m.metaphAdd("MAK")
		} else {
			m.metaphAdd("MK")
		}

		if m.stringAtStart("MC") {
			if m.stringAt(2, "K", "G", "Q") &&
				// watch out for e.g. "McGeorge"
				!m.stringAt(2, "GEOR") {
				m.current += 3
			} else {
				m.current += 2
			}
		} else {
			m.current += 3
		}

		return true
	}

	return false
}

/**
 * Encode silent 'M' in context of "-MPT-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeMPT() bool {
	if m.stringAt(-2, "COMPTROL") ||
		m.stringAt(-4, "ACCOMPT") {
		m.metaphAdd("N")
		m.current += 2
		return true
	}

	return false
}

/**
 * Test if 'B' is silent in these contexts
 *
 * @return true if 'B' is silent in this context
 *
 */
func (m *Metaphone3) testSilentMB1() bool {
	// e.g. "LAMB", "COMB", "LIMB", "DUMB", "BOMB"
	// Handle combining roots first
	if ((m.current == 3) &&
		m.stringAt(-3, "THUMB")) ||
		((m.current == 2) &&
			m.stringAt(-2, "DUMB", "BOMB", "DAMN", "LAMB", "NUMB", "TOMB")) {
		return true
	}

	return false
}

/**
 * Test if 'B' is pronounced in this context
 *
 * @return true if 'B' is pronounced in this context
 *
 */
func (m *Metaphone3) testPronouncedMB() bool {
	if m.stringAt(-2, "NUMBER") ||
		(m.stringAt(2, "A") &&
			!m.stringAt(-2, "DUMBASS")) ||
		m.stringAt(2, "O") ||
		m.stringAt(-2, "LAMBEN", "LAMBER", "LAMBET", "TOMBIG", "LAMBRE") {
		return true
	}

	return false
}

/**
 * Test whether "-B-" is silent in these contexts
 *
 * @return true if 'B' is silent in this context
 *
 */
func (m *Metaphone3) testSilentMB2() bool {
	// 'M' is the current letter
	if (m.charAt(m.current+1) == 'B') && (m.current > 1) &&
		(((m.current + 1) == m.last) ||
			// other situations where "-MB-" is at end of root
			// but not at end of word. The tests are for standard
			// noun suffixes.
			// e.g. "climbing" => KLMNK
			m.stringAt(2, "ING", "ABL") ||
			m.stringAt(2, "LIKE") ||
			((m.charAt(m.current+2) == 'S') && ((m.current + 2) == m.last)) ||
			m.stringAt(-5, "BUNCOMB") ||
			// e.g. "bomber",
			(m.stringAt(2, "ED", "ER") &&
				((m.current + 3) == m.last) &&
				(m.stringAtStart("CLIMB", "PLUMB") ||
					// e.g. "beachcomber"
					!m.stringAt(-1, "IMBER", "AMBER", "EMBER", "UMBER")) &&
				// exceptions
				!m.stringAt(-2, "CUMBER", "SOMBER"))) {
		return true
	}

	return false
}

/**
 * Test if 'B' is pronounced in these "-MB-" contexts
 *
 * @return true if "-B-" is pronounced in these contexts
 *
 */
func (m *Metaphone3) testPronouncedMB2() bool {
	// e.g. "bombastic", "umbrage", "flamboyant"
	if m.stringAt(-1, "OMBAS", "OMBAD", "UMBRA") ||
		m.stringAt(-3, "FLAM") {
		return true
	}

	return false
}

/**
 * Tests for contexts where "-N-" is silent when after "-M-"
 *
 * @return true if "-N-" is silent in these contexts
 *
 */
func (m *Metaphone3) testMN() bool {
	return (m.charAt(m.current+1) == 'N') &&
		(((m.current + 1) == m.last) ||
			// or at the end of a word but followed by suffixes
			(m.stringAt(2, "ING", "EST") && ((m.current + 4) == m.last)) ||
			((m.charAt(m.current+2) == 'S') && ((m.current + 2) == m.last)) ||
			(m.stringAt(2, "LY", "ER", "ED") &&
				((m.current + 3) == m.last)) ||
			m.stringAt(-2, "DAMNEDEST") ||
			m.stringAt(-5, "GODDAMNIT"))
}

/**
 * Call routines to encode "-MB-", in proper order
 *
 */
func (m *Metaphone3) encodeMB() {
	if m.testSilentMB1() {
		if m.testPronouncedMB() {
			m.current++
		} else {
			m.current += 2
		}
	} else if m.testSilentMB2() {
		if m.testPronouncedMB2() {
			m.current++
		} else {
			m.current += 2
		}
	} else if m.testMN() {
		m.current += 2
	} else {
		// eat redundant 'M'
		if m.charAt(m.current+1) == 'M' {
			m.current += 2
		} else {
			m.current++
		}
	}
}

/**
 * Encode "-M-"
 *
 */
func (m *Metaphone3) encodeM() {
	if m.encodeSilentMAtBeginning() ||
		m.encodeMRAndMRS() ||
		m.encodeMAC() ||
		m.encodeMPT() {
		return
	}

	// Silent 'B' should really be handled
	// under 'B", not here under 'M'!
	m.encodeMB()

	m.metaphAdd("M")
}

/**
 * Encode "-NCE-" and "-NSE-"
 * "entrance" is pronounced exactly the same as "entrants"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeNCE() bool {
	//'acceptance', 'accountancy'
	if m.stringAt(1, "C", "S") &&
		m.stringAt(2, "E", "Y", "I") &&
		(((m.current + 2) == m.last) ||
			((m.current+3) == m.last) &&
				(m.charAt(m.current+3) == 'S')) {
		m.metaphAdd("NTS")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-N-"
 *
 */
func (m *Metaphone3) encodeN() {
	if m.encodeNCE() {
		return
	}

	// eat redundant 'N'
	if m.charAt(m.current+1) == 'N' {
		m.current += 2
	} else {
		m.current++
	}

	if !m.stringAt(-3, "MONSIEUR") &&
		// e.g. "aloneness",
		!m.stringAt(-3, "NENESS") {
		m.metaphAdd("N")
	}
}

/**
 * Encode cases where "-P-" is silent at the start of a word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentPAtBeginning() bool {
	//skip these when at start of word
	if (m.current == 0) &&
		m.stringAt(0, "PN", "PF", "PS", "PT") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode cases where "-P-" is silent before "-T-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodePT() bool {
	// 'pterodactyl', 'receipt', 'asymptote'
	if m.charAt(m.current+1) == 'T' {
		if ((m.current == 0) && m.stringAt(0, "PTERO")) ||
			m.stringAt(-5, "RECEIPT") ||
			m.stringAt(-4, "ASYMPTOT") {
			m.metaphAdd("T")
			m.current += 2
			return true
		}
	}
	return false
}

/**
 * Encode "-PH-", usually as F, with exceptions for
 * cases where it is silent, or where the 'P' and 'T'
 * are pronounced seperately because they belong to
 * two different words in a combining form
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodePH() bool {
	if m.charAt(m.current+1) == 'H' {
		// 'PH' silent in these contexts
		if m.stringAt(0, "PHTHALEIN") ||
			((m.current == 0) && m.stringAt(0, "PHTH")) ||
			m.stringAt(-3, "APOPHTHEGM") {
			m.metaphAdd("0")
			m.current += 4
			// combining forms
			//'sheepherd', 'upheaval', 'cupholder'
		} else if (m.current > 0) &&
			(m.stringAt(2, "EAD", "OLE", "ELD", "ILL", "OLD", "EAP", "ERD",
				"ARD", "ANG", "ORN", "EAV", "ART") ||
				m.stringAt(2, "OUSE") ||
				(m.stringAt(2, "AM") && !m.stringAt(-1, "LPHAM")) ||
				m.stringAt(2, "AMMER", "AZARD", "UGGER") ||
				m.stringAt(2, "OLSTER")) &&
			!m.stringAt(-3, "LYMPH", "NYMPH") {
			m.metaphAdd("P")
			m.advanceCounter(3, 2)
		} else {
			m.metaphAdd("F")
			m.current += 2
		}
		return true
	}

	return false
}

/**
 * Encode "-PPH-". I don't know why the greek poet's
 * name is transliterated this way...
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodePPH() bool {
	// 'sappho'
	if (m.charAt(m.current+1) == 'P') &&
		((m.current + 2) < m.length) && (m.charAt(m.current+2) == 'H') {
		m.metaphAdd("F")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode "-CORPS-" where "-PS-" not pronounced
 * since the cognate is here from the french
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeRPS() bool {
	//'-corps-', 'corpsman'
	if m.stringAt(-3, "CORPS") &&
		!m.stringAt(-3, "CORPSE") {
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-COUP-" where "-P-" is not pronounced
 * since the word is from the french
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeCOUP() bool {
	//'coup'
	if (m.current == m.last) &&
		m.stringAt(-3, "COUP") &&
		!m.stringAt(-5, "RECOUP") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode 'P' in non-initial contexts of "-PNEUM-"
 * where is also silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodePNEUM() bool {
	//'-pneum-'
	if m.stringAt(1, "NEUM") {
		m.metaphAdd("N")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode special case "-PSYCH-" where two encodings need to be
 * accounted for in one syllable, one for the 'PS' and one for
 * the 'CH'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodePSYCH() bool {
	//'-psych-'
	if m.stringAt(1, "SYCH") {
		if m.encodeVowels {
			m.metaphAdd("SAK")
		} else {
			m.metaphAdd("SK")
		}

		m.current += 5
		return true
	}

	return false
}

/**
 * Encode 'P' in context of "-PSALM-", where it has
 * become silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodePSALM() bool {
	//'-psalm-'
	if m.stringAt(1, "SALM") {
		// go ahead and encode entire word
		if m.encodeVowels {
			m.metaphAdd("SAM")
		} else {
			m.metaphAdd("SM")
		}

		m.current += 5
		return true
	}

	return false
}

/**
 * Eat redundant 'B' or 'P'
 *
 */
func (m *Metaphone3) encodePB() {
	// e.g. "campbell", "raspberry"
	// eat redundant 'P' or 'B'
	if m.stringAt(1, "P", "B") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode "-P-"
 *
 */
func (m *Metaphone3) encodeP() {
	if m.encodeSilentPAtBeginning() ||
		m.encodePT() ||
		m.encodePH() ||
		m.encodePPH() ||
		m.encodeRPS() ||
		m.encodeCOUP() ||
		m.encodePNEUM() ||
		m.encodePSYCH() ||
		m.encodePSALM() {
		return
	}

	m.encodePB()

	m.metaphAdd("P")
}

/**
 * Encode "-Q-"
 *
 */
func (m *Metaphone3) encodeQ() {
	// current pinyin
	if m.stringAt(0, "QIN") {
		m.metaphAdd("X")
		m.current++
		return
	}

	// eat redundant 'Q'
	if m.charAt(m.current+1) == 'Q' {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAdd("K")
}

/**
 * Encode "-RZ-" according
 * to american and polish pronunciations
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeRZ() bool {
	if m.stringAt(-2, "GARZ", "KURZ", "MARZ", "MERZ", "HERZ", "PERZ", "WARZ") ||
		m.stringAt(0, "RZANO", "RZOLA") ||
		m.stringAt(-1, "ARZA", "ARZN") {
		return false
	}

	// 'yastrzemski' usually has 'z' silent in
	// united states, but should get 'X' in poland
	if m.stringAt(-4, "YASTRZEMSKI") {
		m.metaphAdd("R", "X")
		m.current += 2
		return true
	}
	// 'BRZEZINSKI' gets two pronunciations
	// in the united states, neither of which
	// are authentically polish
	if m.stringAt(-1, "BRZEZINSKI") {
		m.metaphAdd("RS", "RJ")
		// skip over 2nd 'Z'
		m.current += 4
		return true
		// 'z' in 'rz after voiceless consonant gets 'X'
		// in alternate polish style pronunciation
	} else if m.stringAt(-1, "TRZ", "PRZ", "KRZ") ||
		(m.stringAt(0, "RZ") &&
			(m.isVowel(m.current-1) || (m.current == 0))) {
		m.metaphAdd("RS", "X")
		m.current += 2
		return true
		// 'z' in 'rz after voiceled consonant, vowel, or at
		// beginning gets 'J' in alternate polish style pronunciation
	} else if m.stringAt(-1, "BRZ", "DRZ", "GRZ") {
		m.metaphAdd("RS", "J")
		m.current += 2
		return true
	}

	return false
}

/**
 * Test whether 'R' is silent in this context
 *
 * @return true if 'R' is silent in this context
 *
 */
func (m *Metaphone3) testSilentR() bool {
	// test cases where 'R' is silent, either because the
	// word is from the french or because it is no longer pronounced.
	// e.g. "rogier", "monsieur", "surburban"
	return ((m.current == m.last) &&
		// reliably french word ending
		m.stringAt(-2, "IER") &&
		// e.g. "metier"
		(m.stringAt(-5, "MET", "VIV", "LUC") ||
			// e.g. "cartier", "bustier"
			m.stringAt(-6, "CART", "DOSS", "FOUR", "OLIV", "BUST", "DAUM", "ATEL",
				"SONN", "CORM", "MERC", "PELT", "POIR", "BERN", "FORT", "GREN",
				"SAUC", "GAGN", "GAUT", "GRAN", "FORC", "MESS", "LUSS", "MEUN",
				"POTH", "HOLL", "CHEN") ||
			// e.g. "croupier"
			m.stringAt(-7, "CROUP", "TORCH", "CLOUT", "FOURN", "GAUTH", "TROTT",
				"DEROS", "CHART") ||
			// e.g. "chevalier"
			m.stringAt(-8, "CHEVAL", "LAVOIS", "PELLET", "SOMMEL", "TREPAN", "LETELL", "COLOMB") ||
			m.stringAt(-9, "CHARCUT") ||
			m.stringAt(-10, "CHARPENT"))) ||
		m.stringAt(-2, "SURBURB", "WORSTED") ||
		m.stringAt(-2, "WORCESTER") ||
		m.stringAt(-7, "MONSIEUR") ||
		m.stringAt(-6, "POITIERS")
}

/**
 * Encode '-re-" as 'AR' in contexts
 * where this is the correct pronunciation
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeVowelRETransposition() bool {
	// -re inversion is just like
	// -le inversion
	// e.g. "fibre" => FABAR or "centre" => SANTAR
	if (m.encodeVowels) &&
		(m.charAt(m.current+1) == 'E') &&
		(m.length > 3) &&
		!m.stringAtStart("OUTRE", "LIBRE", "ANDRE") &&
		!m.stringEqual("FRED", "TRES") &&
		!m.stringAt(-2, "LDRED", "LFRED", "NDRED", "NFRED", "NDRES", "TRES", "IFRED") &&
		!m.isVowel(m.current-1) &&
		(((m.current + 1) == m.last) ||
			(((m.current + 2) == m.last) &&
				m.stringAt(2, "D", "S"))) {
		m.metaphAdd("AR")
		return true
	}

	return false
}

/**
 * Encode "-R-"
 *
 */
func (m *Metaphone3) encodeR() {
	if m.encodeRZ() {
		return
	}

	if !m.testSilentR() {
		if !m.encodeVowelRETransposition() {
			m.metaphAdd("R")
		}
	}

	// eat redundant 'R'; also skip 'S' as well as 'R' in "poitiers"
	if (m.charAt(m.current+1) == 'R') || m.stringAt(-6, "POITIERS") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Test for names derived from the swedish,
 * dutch, or slavic that should get an alternate
 * pronunciation of 'SV' to match the native
 * version
 *
 * @return true if swedish, dutch, or slavic derived name
 */
func (m *Metaphone3) namesBeginningWithSWThatGetAltSV() bool {
	if m.stringAtStart("SWANSON", "SWENSON", "SWINSON", "SWENSEN",
		"SWOBODA") ||
		m.stringAtStart("SWIDERSKI", "SWARTHOUT") ||
		m.stringAtStart("SWEARENGIN") {
		return true
	}

	return false
}

/**
 * Test for names derived from the german
 * that should get an alternate pronunciation
 * of 'XV' to match the german version spelled
 * "schw-"
 *
 * @return true if german derived name
 */
func (m *Metaphone3) namesBeginningWithSWThatGetAltXV() bool {
	if m.stringAtStart("SWART") ||
		m.stringAtStart("SWARTZ", "SWARTS", "SWIGER") ||
		m.stringAtStart("SWITZER", "SWANGER", "SWIGERT",
			"SWIGART", "SWIHART") ||
		m.stringAtStart("SWEITZER", "SWATZELL", "SWINDLER") ||
		m.stringAtStart("SWINEHART") ||
		m.stringAtStart("SWEARINGEN") {
		return true
	}

	return false
}

/**
 * Encode a couple of contexts where scandinavian, slavic
 * or german names should get an alternate, native
 * pronunciation of 'SV' or 'XV'
 *
 * @return true if handled
 *
 */
func (m *Metaphone3) encodeSpecialSW() bool {
	if m.current == 0 {
		//
		if m.namesBeginningWithSWThatGetAltSV() {
			m.metaphAdd("S", "SV")
			m.current += 2
			return true
		}

		//
		if m.namesBeginningWithSWThatGetAltXV() {
			m.metaphAdd("S", "XV")
			m.current += 2
			return true
		}
	}

	return false
}

/**
 * Encode "-SKJ-" as X ("sh"), since americans pronounce
 * the name Dag Hammerskjold as "hammer-shold"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSKJ() bool {
	// scandinavian
	if m.stringAt(0, "SKJO", "SKJU") &&
		m.isVowel(m.current+3) {
		m.metaphAdd("X")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode initial swedish "SJ-" as X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSJ() bool {
	if m.stringAtStart("SJ") {
		m.metaphAdd("X")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode final 'S' in words from the french, where they
 * are not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentFrenchSFinal() bool {
	// "louis" is an exception because it gets two pronuncuations
	if m.stringAtStart("LOUIS") && (m.current == m.last) {
		m.metaphAdd("S", "")
		m.current++
		return true
	}

	// french words familiar to americans where final s is silent
	if (m.current == m.last) &&
		(m.stringAtStart("YVES") ||
			(m.stringAtStart("HORS") && (m.current == 3)) ||
			m.stringAt(-4, "CAMUS", "YPRES") ||
			m.stringAt(-5, "MESNES", "DEBRIS", "BLANCS", "INGRES", "CANNES") ||
			m.stringAt(-6, "CHABLIS", "APROPOS", "JACQUES", "ELYSEES", "OEUVRES",
				"GEORGES", "DESPRES") ||
			m.stringAtStart("ARKANSAS", "FRANCAIS", "CRUDITES", "BRUYERES") ||
			m.stringAtStart("DESCARTES", "DESCHUTES", "DESCHAMPS", "DESROCHES", "DESCHENES") ||
			m.stringAtStart("RENDEZVOUS") ||
			m.stringAtStart("CONTRETEMPS", "DESLAURIERS")) ||
		((m.current == m.last) &&
			m.stringAt(-2, "AI", "OI", "UI") &&
			!m.stringAtStart("LOIS", "LUIS")) {
		m.current++
		return true
	}

	return false
}

/**
 * Encode non-final 'S' in words from the french where they
 * are not pronounced.
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentFrenchSInternal() bool {
	// french words familiar to americans where internal s is silent
	if m.stringAt(-2, "DESCARTES") ||
		m.stringAt(-2, "DESCHAM", "DESPRES", "DESROCH", "DESROSI", "DESJARD", "DESMARA",
			"DESCHEN", "DESHOTE", "DESLAUR") ||
		m.stringAt(-2, "MESNES") ||
		m.stringAt(-5, "DUQUESNE", "DUCHESNE") ||
		m.stringAt(-7, "BEAUCHESNE") ||
		m.stringAt(-3, "FRESNEL") ||
		m.stringAt(-3, "GROSVENOR") ||
		m.stringAt(-4, "LOUISVILLE") ||
		m.stringAt(-7, "ILLINOISAN") {
		m.current++
		return true
	}
	return false
}

/**
 * Encode silent 'S' in context of "-ISL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeISL() bool {
	//special cases 'island', 'isle', 'carlisle', 'carlysle'
	if (m.stringAt(-2, "LISL", "LYSL", "AISL") &&
		!m.stringAt(-3, "PAISLEY", "BAISLEY", "ALISLAM", "ALISLAH", "ALISLAA")) ||
		((m.current == 1) &&
			((m.stringAt(-1, "ISLE") ||
				m.stringAt(-1, "ISLAN")) &&
				!m.stringAt(-1, "ISLEY", "ISLER"))) {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-STL-" in contexts where the 'T' is silent. Also
 * encode "-USCLE-" in contexts where the 'C' is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSTL() bool {
	//'hustle', 'bustle', 'whistle'
	if (m.stringAt(0, "STLE", "STLI") &&
		!m.stringAt(2, "LESS", "LIKE", "LINE")) ||
		m.stringAt(-3, "THISTLY", "BRISTLY", "GRISTLY") ||
		// e.g. "corpuscle"
		m.stringAt(-1, "USCLE") {
		// KRISTEN, KRYSTLE, CRYSTLE, KRISTLE all pronounce the 't'
		// also, exceptions where "-LING" is a nominalizing suffix
		if m.stringAtStart("KRISTEN", "KRYSTLE", "CRYSTLE", "KRISTLE") ||
			m.stringAtStart("CHRISTENSEN", "CHRISTENSON") ||
			m.stringAt(-3, "FIRSTLING") ||
			m.stringAt(-2, "NESTLING", "WESTLING") {
			m.metaphAdd("ST")
			m.current += 2
		} else {
			if m.encodeVowels &&
				(m.charAt(m.current+3) == 'E') &&
				(m.charAt(m.current+4) != 'R') &&
				!m.stringAt(3, "ETTE", "ETTA") &&
				!m.stringAt(3, "EY") {
				m.metaphAdd("SAL")
				m.flag_AL_inversion = true
			} else {
				m.metaphAdd("SL")
			}
			m.current += 3
		}
		return true
	}

	return false
}

/**
 * Encode "christmas". Americans always pronounce this as "krissmuss"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeChristmas() bool {
	//'christmas'
	if m.stringAt(-4, "CHRISTMA") {
		m.metaphAdd("SM")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode "-STHM-" in contexts where the 'TH'
 * is silent.
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSTHM() bool {
	//'asthma', 'isthmus'
	if m.stringAt(0, "STHM") {
		m.metaphAdd("SM")
		m.current += 4
		return true
	}

	return false
}

/**
 * Encode "-ISTEN-" and "-STNT-" in contexts
 * where the 'T' is silent
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeISTEN() bool {
	// 't' is silent in verb, pronounced in name
	if m.stringAtStart("CHRISTEN") {
		// the word itself
		if m.rootOrInflections(m.inWord, "CHRISTEN") ||
			m.stringAtStart("CHRISTENDOM") {
			m.metaphAdd("S", "ST")
		} else {
			// e.g. 'christenson', 'christene'
			m.metaphAdd("ST")
		}
		m.current += 2
		return true
	}

	//e.g. 'glisten', 'listen'
	if m.stringAt(-2, "LISTEN", "RISTEN", "HASTEN", "FASTEN", "MUSTNT") ||
		m.stringAt(-3, "MOISTEN") {
		m.metaphAdd("S")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode special case "sugar"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSugar() bool {
	//special case 'sugar-'
	if m.stringAt(0, "SUGAR") {
		m.metaphAdd("X")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-SH-" as X ("sh"), except in cases
 * where the 'S' and 'H' belong to different combining
 * roots and are therefore pronounced seperately
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSH() bool {
	if m.stringAt(0, "SH") {
		// exception
		if m.stringAt(-2, "CASHMERE") {
			m.metaphAdd("J")
			m.current += 2
			return true
		}

		//combining forms, e.g. 'clotheshorse', 'woodshole'
		if (m.current > 0) &&
			// e.g. "mishap"
			((m.stringAt(1, "HAP") && ((m.current + 3) == m.last)) ||
				// e.g. "hartsheim", "clothshorse"
				m.stringAt(1, "HEIM", "HOEK", "HOLM", "HOLZ", "HOOD", "HEAD", "HEID",
					"HAAR", "HORS", "HOLE", "HUND", "HELM", "HAWK", "HILL") ||
				// e.g. "dishonor"
				m.stringAt(1, "HEART", "HATCH", "HOUSE", "HOUND", "HONOR") ||
				// e.g. "mishear"
				(m.stringAt(2, "EAR") && ((m.current + 4) == m.last)) ||
				// e.g. "hartshorn"
				(m.stringAt(2, "ORN") && !m.stringAt(-2, "UNSHORN")) ||
				// e.g. "newshour" but not "bashour", "manshour"
				(m.stringAt(1, "HOUR") &&
					!(m.stringAtStart("BASHOUR") || m.stringAtStart("MANSHOUR") || m.stringAtStart("ASHOUR"))) ||
				// e.g. "dishonest", "grasshopper"
				m.stringAt(2, "ARMON", "ONEST", "ALLOW", "OLDER", "OPPER", "EIMER", "ANDLE", "ONOUR") ||
				// e.g. "dishabille", "transhumance"
				m.stringAt(2, "ABILLE", "UMANCE", "ABITUA")) {
			if !m.stringAt(-1, "S") {
				m.metaphAdd("S")
			}
		} else {
			m.metaphAdd("X")
		}

		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-SCH-" in cases where the 'S' is pronounced
 * seperately from the "CH", in words from the dutch, italian,
 * and greek where it can be pronounced SK, and german words
 * where it is pronounced X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSCH() bool {
	// these words were combining forms many centuries ago
	if m.stringAt(1, "CH") {
		if (m.current > 0) &&
			// e.g. "mischief", "escheat"
			(m.stringAt(3, "IEF", "EAT") ||
				// e.g. "mischance"
				m.stringAt(3, "ANCE", "ARGE") ||
				// e.g. "eschew"
				m.stringAtStart("ESCHEW")) {
			m.metaphAdd("S")
			m.current++
			return true
		}

		//Schlesinger's rule
		//dutch, danish, italian, greek origin, e.g. "school", "schooner", "schiavone", "schiz-"
		if (m.stringAt(3, "OO", "ER", "EN", "UY", "ED", "EM", "IA", "IZ", "IS", "OL") &&
			!m.stringAt(0, "SCHOLT", "SCHISL", "SCHERR")) ||
			m.stringAt(3, "ISZ") ||
			(m.stringAt(-1, "ESCHAT", "ASCHIN", "ASCHAL", "ISCHAE", "ISCHIA") &&
				!m.stringAt(-2, "FASCHING")) ||
			(m.stringAt(-1, "ESCHI") && ((m.current + 3) == m.last)) ||
			(m.charAt(m.current+3) == 'Y') {
			// e.g. "schermerhorn", "schenker", "schistose"
			if m.stringAt(3, "ER", "EN", "IS") &&
				(((m.current + 4) == m.last) ||
					m.stringAt(3, "ENK", "ENB", "IST")) {
				m.metaphAdd("X", "SK")
			} else {
				m.metaphAdd("SK")
			}
			m.current += 3
			return true
		} else {
			// Fix for smith and schmidt not returning same code:
			// next two lines from metaphone.go at line 621: code for SCH
			if m.current == 0 && !m.isVowel(3) && (m.charAt(3) != 'W') {
				m.metaphAdd("X", "S")
			} else {
				m.metaphAdd("X")
			}
			m.current += 3
			return true
		}
	}

	return false
}

/**
 * Encode "-SUR<E,A,Y>-" to J, unless it is at the beginning,
 * or preceeded by 'N', 'K', or "NO"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSUR() bool {
	// 'erasure', 'usury'
	if m.stringAt(1, "URE", "URA", "URY") {
		//'sure', 'ensure'
		if (m.current == 0) ||
			m.stringAt(-1, "N", "K") ||
			m.stringAt(-2, "NO") {
			m.metaphAdd("X")
		} else {
			m.metaphAdd("J")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-SU<O,A>-" to X ("sh") unless it is preceeded by
 * an 'R', in which case it is encoded to S, or it is
 * preceeded by a vowel, in which case it is encoded to J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSU() bool {
	//'sensuous', 'consensual'
	if m.stringAt(1, "UO", "UA") && (m.current != 0) {
		// exceptions e.g. "persuade"
		if m.stringAt(-1, "RSUA") {
			m.metaphAdd("S")
			// exceptions e.g. "casual"
		} else if m.isVowel(m.current - 1) {
			m.metaphAdd("J", "S")
		} else {
			m.metaphAdd("X", "S")
		}

		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encodes "-SSIO-" in contexts where it is pronounced
 * either J or X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSSIO() bool {
	if m.stringAt(1, "SION") {
		//"abcission"
		if m.stringAt(-2, "CI") {
			m.metaphAdd("J")
			//'mission'
		} else {
			if m.isVowel(m.current - 1) {
				m.metaphAdd("X")
			}
		}

		m.advanceCounter(4, 2)
		return true
	}

	return false
}

/**
 * Encode "-SS-" in contexts where it is pronounced X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSS() bool {
	// e.g. "russian", "pressure"
	if m.stringAt(-1, "USSIA", "ESSUR", "ISSUR", "ISSUE") ||
		// e.g. "hessian", "assurance"
		m.stringAt(-1, "ESSIAN", "ASSURE", "ASSURA", "ISSUAB", "ISSUAN", "ASSIUS") {
		m.metaphAdd("X")
		m.advanceCounter(3, 2)
		return true
	}

	return false
}

/**
 * Encodes "-SIA-" in contexts where it is pronounced
 * as X ("sh"), J, or S
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSIA() bool {
	// e.g. "controversial", also "fuchsia", "ch" is silent
	if m.stringAt(-2, "CHSIA") ||
		m.stringAt(-1, "RSIAL") {
		m.metaphAdd("X")
		m.advanceCounter(3, 1)
		return true
	}

	// names generally get 'X' where terms, e.g. "aphasia" get 'J'
	if (m.stringAtStart("ALESIA", "ALYSIA", "ALISIA", "STASIA") &&
		(m.current == 3) &&
		!m.stringAtStart("ANASTASIA")) ||
		m.stringAt(-5, "DIONYSIAN") ||
		m.stringAt(-5, "THERESIA") {
		m.metaphAdd("X", "S")
		m.advanceCounter(3, 1)
		return true
	}

	if (m.stringAt(0, "SIA") && ((m.current + 2) == m.last)) ||
		(m.stringAt(0, "SIAN") && ((m.current + 3) == m.last)) ||
		m.stringAt(-5, "AMBROSIAL") {
		if (m.isVowel(m.current-1) || m.stringAt(-1, "R")) &&
			// exclude compounds based on names, or french or greek words
			!(m.stringAtStart("JAMES", "NICOS", "PEGAS", "PEPYS") ||
				m.stringAtStart("HOBBES", "HOLMES", "JAQUES", "KEYNES") ||
				m.stringAtStart("MALTHUS", "HOMOOUS") ||
				m.stringAtStart("MAGLEMOS", "HOMOIOUS") ||
				m.stringAtStart("LEVALLOIS", "TARDENOIS") ||
				m.stringAt(-4, "ALGES")) {
			m.metaphAdd("J")
		} else {
			m.metaphAdd("S")
		}

		m.advanceCounter(2, 1)
		return true
	}
	return false
}

/**
 * Encodes "-SIO-" in contexts where it is pronounced
 * as J or X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSIO() bool {
	// special case, irish name
	if m.stringAtStart("SIOBHAN") {
		m.metaphAdd("X")
		m.advanceCounter(3, 1)
		return true
	}

	if m.stringAt(1, "ION") {
		// e.g. "vision", "version"
		if m.isVowel(m.current-1) || m.stringAt(-2, "ER", "UR") {
			m.metaphAdd("J")
		} else {
			// e.g. "declension"
			m.metaphAdd("X")
		}

		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode cases where "-S-" might well be from a german name
 * and add encoding of german pronounciation in alternate m.metaph
 * so that it can be found in a genealogical search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeAnglicisations() bool {
	//german & anglicisations, e.g. 'smith' match 'schmidt', 'snider' match 'schneider'
	//also, -sz- in slavic language altho in hungarian it is pronounced 's'
	if ((m.current == 0) &&
		m.stringAt(1, "M", "N", "L")) ||
		m.stringAt(1, "Z") {
		m.metaphAdd("S", "X")

		// eat redundant 'Z'
		if m.stringAt(1, "Z") {
			m.current += 2
		} else {
			m.current++
		}

		return true
	}

	return false
}

/**
 * Encode "-SC<vowel>-" in contexts where it is silent,
 * or pronounced as X ("sh"), S, or SK
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSC() bool {
	if m.stringAt(0, "SC") {
		// exception 'viscount'
		if m.stringAt(-2, "VISCOUNT") {
			m.current += 1
			return true
		}

		// encode "-SC<front vowel>-"
		if m.stringAt(2, "I", "E", "Y") {
			// e.g. "conscious"
			if m.stringAt(2, "IOUS") ||
				// e.g. "prosciutto"
				m.stringAt(2, "IUT") ||
				m.stringAt(-4, "OMNISCIEN") ||
				// e.g. "conscious"
				m.stringAt(-3, "CONSCIEN", "CRESCEND", "CONSCION") ||
				m.stringAt(-2, "FASCIS") {
				m.metaphAdd("X")
			} else if m.stringAt(0, "SCEPTIC", "SCEPSIS") ||
				m.stringAt(0, "SCIVV", "SCIRO") ||
				// commonly pronounced this way in u.s.
				m.stringAt(0, "SCIPIO") ||
				m.stringAt(-2, "PISCITELLI") {
				m.metaphAdd("SK")
			} else {
				m.metaphAdd("S")
			}
			m.current += 2
			return true
		}

		m.metaphAdd("SK")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-S<EA,UI,IER>-" in contexts where it is pronounced
 * as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSEASUISIER() bool {
	// "nausea" by itself has => NJ as a more likely encoding. Other forms
	// using "nause-" (see m.encodeSEA()) have X or S as more familiar pronounciations
	if (m.stringAt(-3, "NAUSEA") && ((m.current + 2) == m.last)) ||
		// e.g. "casuistry", "frasier", "hoosier"
		m.stringAt(-2, "CASUI") ||
		(m.stringAt(-1, "OSIER", "ASIER") &&
			!(m.stringAtStart("EASIER") ||
				m.stringAtStart("OSIER") ||
				m.stringAt(-2, "ROSIER", "MOSIER"))) {
		m.metaphAdd("J", "X")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode cases where "-SE-" is pronounced as X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSEA() bool {
	if (m.stringAtStart("SEAN") && ((m.current + 3) == m.last)) ||
		(m.stringAt(-3, "NAUSEO") &&
			!m.stringAt(-3, "NAUSEAT")) {
		m.metaphAdd("X")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-S-"
 *
 */
func (m *Metaphone3) encodeS() {
	if m.encodeSKJ() ||
		m.encodeSpecialSW() ||
		m.encodeSJ() ||
		m.encodeSilentFrenchSFinal() ||
		m.encodeSilentFrenchSInternal() ||
		m.encodeISL() ||
		m.encodeSTL() ||
		m.encodeChristmas() ||
		m.encodeSTHM() ||
		m.encodeISTEN() ||
		m.encodeSugar() ||
		m.encodeSH() ||
		m.encodeSCH() ||
		m.encodeSUR() ||
		m.encodeSU() ||
		m.encodeSSIO() ||
		m.encodeSS() ||
		m.encodeSIA() ||
		m.encodeSIO() ||
		m.encodeAnglicisations() ||
		m.encodeSC() ||
		m.encodeSEASUISIER() ||
		m.encodeSEA() {
		return
	}

	m.metaphAdd("S")

	if m.stringAt(1, "S", "Z") &&
		!m.stringAt(1, "SH") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode some exceptions for initial 'T'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTInitial() bool {
	if m.current == 0 {
		// americans usually pronounce "tzar" as "zar"
		if m.stringAt(1, "SAR", "ZAR") {
			m.current++
			return true
		}

		// old 'École française d'Extrême-Orient' chinese pinyin where 'ts-' => 'X'
		if ((m.length == 3) && m.stringAt(1, "SO", "SA", "SU")) ||
			((m.length == 4) && m.stringAt(1, "SAO", "SAI")) ||
			((m.length == 5) && m.stringAt(1, "SING", "SANG")) {
			m.metaphAdd("X")
			m.advanceCounter(3, 2)
			return true
		}

		// "TS<vowel>-" at start can be pronounced both with and without 'T'
		if m.stringAt(1, "S") && m.isVowel(m.current+2) {
			m.metaphAdd("TS", "S")
			m.advanceCounter(3, 2)
			return true
		}

		// e.g. "Tjaarda"
		if m.stringAt(1, "J") {
			m.metaphAdd("X")
			m.advanceCounter(3, 2)
			return true
		}

		// cases where initial "TH-" is pronounced as T and not 0 ("th")
		if (m.stringAt(1, "HU") && (m.length == 3)) ||
			m.stringAt(1, "HAI", "HUY", "HAO") ||
			m.stringAt(1, "HYME", "HYMY", "HANH") ||
			m.stringAt(1, "HERES") {
			m.metaphAdd("T")
			m.advanceCounter(3, 2)
			return true
		}
	}

	return false
}

/**
 * Encode "-TCH-", reliably X ("sh", or in this case, "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTCH() bool {
	if m.stringAt(1, "CH") {
		m.metaphAdd("X")
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode the many cases where americans are aware that a certain word is
 * french and know to not pronounce the 'T'
 *
 * @return true if encoding handled in this routine, false if not
 * TOUCHET CHABOT BENOIT
 */
func (m *Metaphone3) encodeSilentFrenchT() bool {
	// french silent T familiar to americans
	if ((m.current == m.last) && m.stringAt(-4, "MONET", "GENET", "CHAUT")) ||
		m.stringAt(-2, "POTPOURRI") ||
		m.stringAt(-3, "BOATSWAIN") ||
		m.stringAt(-3, "MORTGAGE") ||
		(m.stringAt(-4, "BERET", "BIDET", "FILET", "DEBUT", "DEPOT", "PINOT", "TAROT") ||
			m.stringAt(-5, "BALLET", "BUFFET", "CACHET", "CHALET", "ESPRIT", "RAGOUT", "GOULET",
				"CHABOT", "BENOIT") ||
			m.stringAt(-6, "GOURMET", "BOUQUET", "CROCHET", "CROQUET", "PARFAIT", "PINCHOT",
				"CABARET", "PARQUET", "RAPPORT", "TOUCHET", "COURBET", "DIDEROT") ||
			m.stringAt(-7, "ENTREPOT", "CABERNET", "DUBONNET", "MASSENET", "MUSCADET", "RICOCHET", "ESCARGOT") ||
			m.stringAt(-8, "SOBRIQUET", "CABRIOLET", "CASSOULET", "OUBRIQUET", "CAMEMBERT")) &&
			!m.stringAt(1, "AN", "RY", "IC", "OM", "IN") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-TU<N,L,A,O>-" in cases where it is pronounced
 * X ("sh", or in this case, "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTUNTULTUATUO() bool {
	// e.g. "fortune", "fortunate"
	if m.stringAt(-3, "FORTUN") ||
		// e.g. "capitulate"
		(m.stringAt(0, "TUL") &&
			(m.isVowel(m.current-1) && m.isVowel(m.current+3))) ||
		// e.g. "obituary", "barbituate"
		m.stringAt(-2, "BITUA", "BITUE") ||
		// e.g. "actual"
		((m.current > 1) && m.stringAt(0, "TUA", "TUO")) {
		m.metaphAdd("X", "T")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-T<vowel>-" forms where 'T' is pronounced as X
 * ("sh", or in this case "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTUETEUTEOUTULTIE() bool {
	// 'constituent', 'pasteur'
	if m.stringAt(1, "UENT") ||
		m.stringAt(-4, "RIGHTEOUS") ||
		m.stringAt(-3, "STATUTE") ||
		m.stringAt(-3, "AMATEUR") ||
		// e.g. "blastula", "pasteur"
		(m.stringAt(-1, "NTULE", "NTULA", "STULE", "STULA", "STEUR")) ||
		// e.g. "statue"
		(((m.current + 2) == m.last) && m.stringAt(0, "TUE")) ||
		// e.g. "constituency"
		m.stringAt(0, "TUENC") ||
		// e.g. "statutory"
		m.stringAt(-3, "STATUTOR") ||
		// e.g. "patience"
		(((m.current + 5) == m.last) && m.stringAt(0, "TIENCE")) {
		m.metaphAdd("X", "T")
		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-TU-" forms in suffixes where it is usually
 * pronounced as X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTURTIUSuffixes() bool {
	// 'adventure', 'musculature'
	if (m.current > 0) && m.stringAt(1, "URE", "URA", "URI", "URY", "URO", "IUS") {
		// exceptions e.g. 'tessitura', mostly from romance languages
		if (m.stringAt(1, "URA", "URO") &&
			//&& !m.stringAt(1, "URIA")
			((m.current+3) == m.last)) &&
			!m.stringAt(-3, "VENTURA") ||
			// e.g. "kachaturian", "hematuria"
			m.stringAt(1, "URIA") {
			m.metaphAdd("T")
		} else {
			m.metaphAdd("X", "T")
		}

		m.advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-TI<O,A,U>-" as X ("sh"), except
 * in cases where it is part of a combining form,
 * or as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTI() bool {
	// '-tio-', '-tia-', '-tiu-'
	// except combining forms where T already pronounced e.g 'rooseveltian'
	if (m.stringAt(1, "IO") && !m.stringAt(-1, "ETIOL")) ||
		m.stringAt(1, "IAL") ||
		m.stringAt(-1, "RTIUM", "ATIUM") ||
		((m.stringAt(1, "IAN") && (m.current > 0)) &&
			!(m.stringAt(-4, "FAUSTIAN") ||
				m.stringAt(-5, "PROUSTIAN") ||
				m.stringAt(-2, "TATIANA") ||
				(m.stringAt(-3, "KANTIAN", "GENTIAN") ||
					m.stringAt(-8, "ROOSEVELTIAN"))) ||
			(((m.current + 2) == m.last) &&
				m.stringAt(0, "TIA") &&
				// exceptions to above rules where the pronounciation is usually X
				!(m.stringAt(-3, "HESTIA", "MASTIA") ||
					m.stringAt(-2, "OSTIA") ||
					m.stringAtStart("TIA") ||
					m.stringAt(-5, "IZVESTIA"))) ||
			m.stringAt(1, "IATE", "IATI", "IABL", "IATO", "IARY") ||
			m.stringAt(-5, "CHRISTIAN")) {
		if ((m.current == 2) && m.stringAtStart("ANTI")) ||
			m.stringAtStart("PATIO", "PITIA", "DUTIA") {
			m.metaphAdd("T")
		} else if m.stringAt(-4, "EQUATION") {
			m.metaphAdd("J")
		} else {
			if m.stringAt(0, "TION") {
				m.metaphAdd("X")
			} else if m.stringAtStart("KATIA", "LATIA") {
				m.metaphAdd("T", "X")
			} else {
				m.metaphAdd("X", "T")
			}
		}

		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-TIENT-" where "TI" is pronounced X ("sh")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTIENT() bool {
	// e.g. 'patient'
	if m.stringAt(1, "IENT") {
		m.metaphAdd("X", "T")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-TSCH-" as X ("ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTSCH() bool {
	//'deutsch'
	if m.stringAt(0, "TSCH") &&
		// combining forms in german where the 'T' is pronounced seperately
		!m.stringAt(-3, "WELT", "KLAT", "FEST") {
		// pronounced the same as "ch" in "chit" => X
		m.metaphAdd("X")
		m.current += 4
		return true
	}

	return false
}

/**
 * Encode "-TZSCH-" as X ("ch")
 *
 * "Neitzsche is peachy"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTZSCH() bool {
	//'neitzsche'
	if m.stringAt(0, "TZSCH") {
		m.metaphAdd("X")
		m.current += 5
		return true
	}

	return false
}

/**
 * Encodes cases where the 'H' in "-TH-" is the beginning of
 * another word in a combining form, special cases where it is
 * usually pronounced as 'T', and a special case where it has
 * become pronounced as X ("sh", in this case "ch")
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTHPronouncedSeparately() bool {
	//'adulthood', 'bithead', 'apartheid'
	if ((m.current > 0) &&
		m.stringAt(1, "HOOD", "HEAD", "HEID", "HAND", "HILL", "HOLD",
			"HAWK", "HEAP", "HERD", "HOLE", "HOOK", "HUNT",
			"HUMO", "HAUS", "HOFF", "HARD") &&
		!m.stringAt(-3, "SOUTH", "NORTH")) ||
		m.stringAt(1, "HOUSE", "HEART", "HASTE", "HYPNO", "HEQUE") ||
		// watch out for greek root "-thallic"
		(m.stringAt(1, "HALL") &&
			((m.current + 4) == m.last) &&
			!m.stringAt(-3, "SOUTH", "NORTH")) ||
		(m.stringAt(1, "HAM") &&
			((m.current + 3) == m.last) &&
			!(m.stringAtStart("GOTHAM", "WITHAM", "LATHAM") ||
				m.stringAtStart("BENTHAM", "WALTHAM", "WORTHAM") ||
				m.stringAtStart("GRANTHAM"))) ||
		(m.stringAt(1, "HATCH") &&
			!((m.current == 0) || m.stringAt(-2, "UNTHATCH"))) ||
		m.stringAt(-3, "WARTHOG") ||
		// and some special cases where "-TH-" is usually pronounced 'T'
		m.stringAt(-2, "ESTHER") ||
		m.stringAt(-3, "GOETHE") ||
		m.stringAt(-2, "NATHALIE") {
		// special case
		if m.stringAt(-3, "POSTHUM") {
			m.metaphAdd("X")
		} else {
			m.metaphAdd("T")
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode the "-TTH-" in "matthew", eating the redundant 'T'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTTH() bool {
	// 'matthew' vs. 'outthink'
	if m.stringAt(0, "TTH") {
		if m.stringAt(-2, "MATTH") {
			m.metaphAdd("0")
		} else {
			m.metaphAdd("T0")
		}
		m.current += 3
		return true
	}

	return false
}

/**
 * Encode "-TH-". 0 (zero) is used in Metaphone to encode this sound
 * when it is pronounced as a dipthong, either voiced or unvoiced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeTH() bool {
	if m.stringAt(0, "TH") {
		//'-clothes-'
		if m.stringAt(-3, "CLOTHES") {
			// vowel already encoded so skip right to S
			m.current += 3
			return true
		}

		//special case "thomas", "thames", "beethoven" or germanic words
		if m.stringAt(2, "OMAS", "OMPS", "OMPK", "OMSO", "OMSE",
			"AMES", "OVEN", "OFEN", "ILDA", "ILDE") ||
			m.stringEqual("THOM", "THOMS") ||
			m.stringAtStart("VAN ", "VON ") ||
			m.stringAtStart("SCH") {
			m.metaphAdd("T")

		} else {
			// give an 'etymological' 2nd
			// encoding for "smith"
			if m.stringAtStart("SM") {
				m.metaphAdd("0", "T")
			} else {
				m.metaphAdd("0")
			}
		}

		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-T-"
 *
 */
func (m *Metaphone3) encodeT() {
	if m.encodeTInitial() ||
		m.encodeTCH() ||
		m.encodeSilentFrenchT() ||
		m.encodeTUNTULTUATUO() ||
		m.encodeTUETEUTEOUTULTIE() ||
		m.encodeTURTIUSuffixes() ||
		m.encodeTI() ||
		m.encodeTIENT() ||
		m.encodeTSCH() ||
		m.encodeTZSCH() ||
		m.encodeTHPronouncedSeparately() ||
		m.encodeTTH() ||
		m.encodeTH() {
		return
	}

	// eat redundant 'T' or 'D'
	if m.stringAt(1, "T", "D") {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAdd("T")
}

/**
 * Encode "-V-"
 *
 */
func (m *Metaphone3) encodeV() {
	// eat redundant 'V'
	if m.charAt(m.current+1) == 'V' {
		m.current += 2
	} else {
		m.current++
	}

	m.metaphAddExactApprox("V", "F")
}

/**
 * Encode cases where 'W' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeSilentWAtBeginning() bool {
	//skip these when at start of word
	if (m.current == 0) &&
		m.stringAt(0, "WR") {
		m.current += 1
		return true
	}

	return false
}

/**
 * Encode polish patronymic suffix, mapping
 * alternate spellings to the same encoding,
 * and including easern european pronounciation
 * to the american so that both forms can
 * be found in a genealogy search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeWITZWICZ() bool {
	//polish e.g. 'filipowicz'
	if ((m.current + 3) == m.last) && m.stringAt(0, "WICZ", "WITZ") {
		if m.encodeVowels {
			if (len(m.primary) > 0) &&
				m.charAt(len(m.primary)-1) == 'A' {
				m.metaphAdd("TS", "FAX")
			} else {
				m.metaphAdd("ATS", "FAX")
			}
		} else {
			m.metaphAdd("TS", "FX")
		}
		m.current += 4
		return true
	}

	return false
}

/**
 * Encode "-WR-" as R ('W' always effectively silent)
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeWR() bool {
	//can also be in middle of word
	if m.stringAt(0, "WR") {
		m.metaphAdd("R")
		m.current += 2
		return true
	}

	return false
}

/**
 * Test whether the word in question
 * is a name of germanic or slavic origin, for
 * the purpose of determining whether to add an
 * alternate encoding of 'V'
 *
 * @return true if germanic or slavic name
 */
func (m *Metaphone3) germanicOrSlavicNameBeginningWithW() bool {
	return m.stringAtStart("WEE", "WIX", "WAX",
		"WOLF", "WEIS", "WAHL", "WALZ", "WEIL", "WERT",
		"WINE", "WILK", "WALT", "WOLL", "WADA", "WULF",
		"WEHR", "WURM", "WYSE", "WENZ", "WIRT", "WOLK",
		"WEIN", "WYSS", "WASS", "WANN", "WINT", "WINK",
		"WILE", "WIKE", "WIER", "WELK", "WISE",
		"WIRTH", "WIESE", "WITTE", "WENTZ", "WOLFF", "WENDT",
		"WERTZ", "WILKE", "WALTZ", "WEISE", "WOOLF", "WERTH",
		"WEESE", "WURTH", "WINES", "WARGO", "WIMER", "WISER",
		"WAGER", "WILLE", "WILDS", "WAGAR", "WERTS", "WITTY",
		"WIENS", "WIEBE", "WIRTZ", "WYMER", "WULFF", "WIBLE",
		"WINER", "WIEST", "WALKO", "WALLA", "WEBRE", "WEYER",
		"WYBLE", "WOMAC", "WILTZ", "WURST", "WOLAK", "WELKE",
		"WEDEL", "WEIST", "WYGAN", "WUEST", "WEISZ", "WALCK",
		"WEITZ", "WYDRA", "WANDA", "WILMA", "WEBER",
		"WETZEL", "WEINER", "WENZEL", "WESTER", "WALLEN", "WENGER",
		"WALLIN", "WEILER", "WIMMER", "WEIMER", "WYRICK", "WEGNER",
		"WINNER", "WESSEL", "WILKIE", "WEIGEL", "WOJCIK", "WENDEL",
		"WITTER", "WIENER", "WEISER", "WEXLER", "WACKER", "WISNER",
		"WITMER", "WINKLE", "WELTER", "WIDMER", "WITTEN", "WINDLE",
		"WASHER", "WOLTER", "WILKEY", "WIDNER", "WARMAN", "WEYANT",
		"WEIBEL", "WANNER", "WILKEN", "WILTSE", "WARNKE", "WALSER",
		"WEIKEL", "WESNER", "WITZEL", "WROBEL", "WAGNON", "WINANS",
		"WENNER", "WOLKEN", "WILNER", "WYSONG", "WYCOFF", "WUNDER",
		"WINKEL", "WIDMAN", "WELSCH", "WEHNER", "WEIGLE", "WETTER",
		"WUNSCH", "WHITTY", "WAXMAN", "WILKER", "WILHAM", "WITTIG",
		"WITMAN", "WESTRA", "WEHRLE", "WASSER", "WILLER", "WEGMAN",
		"WARFEL", "WYNTER", "WERNER", "WAGNER", "WISSER",
		"WISEMAN", "WINKLER", "WILHELM", "WELLMAN", "WAMPLER", "WACHTER",
		"WALTHER", "WYCKOFF", "WEIDNER", "WOZNIAK", "WEILAND", "WILFONG",
		"WIEGAND", "WILCHER", "WIELAND", "WILDMAN", "WALDMAN", "WORTMAN",
		"WYSOCKI", "WEIDMAN", "WITTMAN", "WIDENER", "WOLFSON", "WENDELL",
		"WEITZEL", "WILLMAN", "WALDRUP", "WALTMAN", "WALCZAK", "WEIGAND",
		"WESSELS", "WIDEMAN", "WOLTERS", "WIREMAN", "WILHOIT", "WEGENER",
		"WOTRING", "WINGERT", "WIESNER", "WAYMIRE", "WHETZEL", "WENTZEL",
		"WINEGAR", "WESTMAN", "WYNKOOP", "WALLICK", "WURSTER", "WINBUSH",
		"WILBERT", "WALLACH", "WYNKOOP", "WALLICK", "WURSTER", "WINBUSH",
		"WILBERT", "WALLACH", "WEISSER", "WEISNER", "WINDERS", "WILLMON",
		"WILLEMS", "WIERSMA", "WACHTEL", "WARNICK", "WEIDLER", "WALTRIP",
		"WHETSEL", "WHELESS", "WELCHER", "WALBORN", "WILLSEY", "WEINMAN",
		"WAGAMAN", "WOMMACK", "WINGLER", "WINKLES", "WIEDMAN", "WHITNER",
		"WOLFRAM", "WARLICK", "WEEDMAN", "WHISMAN", "WINLAND", "WEESNER",
		"WARTHEN", "WETZLER", "WENDLER", "WALLNER", "WOLBERT", "WITTMER",
		"WISHART", "WILLIAM",
		"WESTPHAL", "WICKLUND", "WEISSMAN", "WESTLUND", "WOLFGANG", "WILLHITE",
		"WEISBERG", "WALRAVEN", "WOLFGRAM", "WILHOITE", "WECHSLER", "WENDLING",
		"WESTBERG", "WENDLAND", "WININGER", "WHISNANT", "WESTRICK", "WESTLING",
		"WESTBURY", "WEITZMAN", "WEHMEYER", "WEINMANN", "WISNESKI", "WHELCHEL",
		"WEISHAAR", "WAGGENER", "WALDROUP", "WESTHOFF", "WIEDEMAN", "WASINGER",
		"WINBORNE",
		"WHISENANT", "WEINSTEIN", "WESTERMAN", "WASSERMAN", "WITKOWSKI", "WEINTRAUB",
		"WINKELMAN", "WINKFIELD", "WANAMAKER", "WIECZOREK", "WIECHMANN", "WOJTOWICZ",
		"WALKOWIAK", "WEINSTOCK", "WILLEFORD", "WARKENTIN", "WEISINGER", "WINKLEMAN",
		"WILHEMINA",
		"WISNIEWSKI", "WUNDERLICH", "WHISENHUNT", "WEINBERGER", "WROBLEWSKI",
		"WAGUESPACK", "WEISGERBER", "WESTERVELT", "WESTERLUND", "WASILEWSKI",
		"WILDERMUTH", "WESTENDORF", "WESOLOWSKI", "WEINGARTEN", "WINEBARGER",
		"WESTERBERG", "WANNAMAKER", "WEISSINGER",
		"WALDSCHMIDT", "WEINGARTNER", "WINEBRENNER",
		"WOLFENBARGER",
		"WOJCIECHOWSKI")
}

/**
 * Encode "W-", adding central and eastern european
 * pronounciations so that both forms can be found
 * in a genealogy search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeInitialWVowel() bool {
	if (m.current == 0) && m.isVowel(m.current+1) {
		//Witter should match Vitter
		if m.germanicOrSlavicNameBeginningWithW() {
			if m.encodeVowels {
				m.metaphAddExactApprox("A", "VA", "A", "FA")
			} else {
				m.metaphAddExactApprox("A", "V", "A", "F")
			}
		} else {
			m.metaphAdd("A")
		}

		// don't encode vowels twice
		m.current = m.skipVowels(m.current + 1)
		return true
	}

	return false
}

/**
 * Encode "-WH-" either as H, or close enough to 'U' to be
 * considered a vowel
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeWH() bool {
	if m.stringAt(0, "WH") {
		// cases where it is pronounced as H
		// e.g. 'who', 'whole'
		if (m.charAt(m.current+2) == 'O') &&
			// exclude cases where it is pronounced like a vowel
			!(m.stringAt(2, "OOSH") ||
				m.stringAt(2, "OOP", "OMP", "ORL", "ORT") ||
				m.stringAt(2, "OA", "OP")) {
			m.metaphAdd("H")
			m.advanceCounter(3, 2)
			return true
		} else {
			// combining forms, e.g. 'hollowhearted', 'rawhide'
			if m.stringAt(2, "IDE", "ARD", "EAD", "AWK", "ERD",
				"OOK", "AND", "OLE", "OOD") ||
				m.stringAt(2, "EART", "OUSE", "OUND") ||
				m.stringAt(2, "AMMER") {
				m.metaphAdd("H")
				m.current += 2
				return true
			} else if m.current == 0 {
				m.metaphAdd("A")
				// don't encode vowels twice
				m.current = m.skipVowels(m.current + 2)
				return true
			}
		}
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-W-" when in eastern european names, adding
 * the eastern european pronounciation to the american so
 * that both forms can be found in a genealogy search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeEasternEuropeanW() bool {
	//Arnow should match Arnoff
	if ((m.current == m.last) && m.isVowel(m.current-1)) ||
		m.stringAt(-1, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
		(m.stringAt(0, "WICKI", "WACKI") && ((m.current + 4) == m.last)) ||
		m.stringAt(0, "WIAK") && ((m.current+3) == m.last) ||
		m.stringAtStart("SCH") {
		m.metaphAddExactApprox("", "V", "", "F")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-W-"
 *
 */
func (m *Metaphone3) encodeW() {
	if m.encodeSilentWAtBeginning() ||
		m.encodeWITZWICZ() ||
		m.encodeWR() ||
		m.encodeInitialWVowel() ||
		m.encodeWH() ||
		m.encodeEasternEuropeanW() {
		return
	}

	// e.g. 'zimbabwe'
	if m.encodeVowels &&
		m.stringAt(0, "WE") &&
		((m.current + 1) == m.last) {
		m.metaphAdd("A")
	}

	//else skip it
	m.current++

}

/**
 * Encode initial X where it is usually pronounced as S
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeInitialX() bool {
	// current chinese pinyin spelling
	if m.stringAtStart("XIA", "XIO", "XIE") ||
		m.stringAtStart("XU") {
		m.metaphAdd("X")
		m.current++
		return true
	}

	// else
	if m.current == 0 {
		m.metaphAdd("S")
		m.current++
		return true
	}

	return false
}

/**
 * Encode X when from greek roots where it is usually pronounced as S
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGreekX() bool {
	// 'xylophone', xylem', 'xanthoma', 'xeno-'
	if m.stringAt(1, "YLO", "YLE", "ENO") ||
		m.stringAt(1, "ANTH") {
		m.metaphAdd("S")
		m.current++
		return true
	}

	return false
}

/**
 * Encode special cases, "LUXUR-", "Texeira"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeXSpecialCases() bool {
	// 'luxury'
	if m.stringAt(-2, "LUXUR") {
		m.metaphAddExactApprox("GJ", "KJ")
		m.current++
		return true
	}

	// 'texeira' portuguese/galician name
	if m.stringAtStart("TEXEIRA") ||
		m.stringAtStart("TEIXEIRA") {
		m.metaphAdd("X")
		m.current++
		return true
	}

	return false
}

/**
 * Encode special case where americans know the
 * proper mexican indian pronounciation of this name
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeXToH() bool {
	// TODO: look for other mexican indian words
	// where 'X' is usually pronounced this way
	if m.stringAt(-2, "OAXACA") ||
		m.stringAt(-3, "QUIXOTE") {
		m.metaphAdd("H")
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-X-" in vowel contexts where it is usually
 * pronounced KX ("ksh")
 * account also for BBC pronounciation of => KS
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeXVowel() bool {
	// e.g. "sexual", "connexion" (british), "noxious"
	if m.stringAt(1, "UAL", "ION", "IOU") {
		m.metaphAdd("KX", "KS")
		m.advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode cases of "-X", encoding as silent when part
 * of a french word where it is not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeFrenchXFinal() bool {
	//french e.g. "breaux", "paix"
	if !((m.current == m.last) &&
		(m.stringAt(-3, "IAU", "EAU", "IEU") ||
			m.stringAt(-2, "AI", "AU", "OU", "OI", "EU"))) {
		m.metaphAdd("KS")
	}

	return false
}

/**
 * Encode "-X-"
 *
 */
func (m *Metaphone3) encodeX() {
	if m.encodeInitialX() ||
		m.encodeGreekX() ||
		m.encodeXSpecialCases() ||
		m.encodeXToH() ||
		m.encodeXVowel() ||
		m.encodeFrenchXFinal() {
		return
	}

	// eat redundant 'X' or other redundant cases
	if m.stringAt(1, "X", "Z", "S") ||
		// e.g. "excite", "exceed"
		m.stringAt(1, "CI", "CE") {
		m.current += 2
	} else {
		m.current++
	}
}

/**
 * Encode cases of "-ZZ-" where it is obviously part
 * of an italian word where "-ZZ-" is pronounced as TS
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeZZ() bool {
	// "abruzzi", 'pizza'
	if (m.charAt(m.current+1) == 'Z') &&
		((m.stringAt(2, "I", "O", "A") &&
			((m.current + 2) == m.last)) ||
			m.stringAt(-2, "MOZZARELL", "PIZZICATO", "PUZZONLAN")) {
		m.metaphAdd("TS", "S")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode special cases where "-Z-" is pronounced as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeZUZIERZS() bool {
	if ((m.current == 1) && m.stringAt(-1, "AZUR")) ||
		(m.stringAt(0, "ZIER") &&
			!m.stringAt(-2, "VIZIER")) ||
		m.stringAt(0, "ZSA") {
		m.metaphAdd("J", "S")

		if m.stringAt(0, "ZSA") {
			m.current += 2
		} else {
			m.current++
		}
		return true
	}

	return false
}

/**
 * Encode cases where americans recognize "-EZ" as part
 * of a french word where Z not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeFrenchEZ() bool {
	if ((m.current == 3) && m.stringAt(-3, "CHEZ")) ||
		m.stringAt(-5, "RENDEZ") {
		m.current++
		return true
	}

	return false
}

/**
 * Encode cases where "-Z-" is in a german word
 * where Z => TS in german
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeGermanZ() bool {
	if ((m.current == 2) && ((m.current + 1) == m.last) && m.stringAt(-2, "NAZI")) ||
		m.stringAt(-2, "NAZIFY", "MOZART") ||
		m.stringAt(-3, "HOLZ", "HERZ", "MERZ", "FITZ") ||
		(m.stringAt(-3, "GANZ") && !m.isVowel(m.current+1)) ||
		m.stringAt(-4, "STOLZ", "PRINZ") ||
		m.stringAt(-4, "VENEZIA") ||
		m.stringAt(-3, "HERZOG") ||
		// german words beginning with "sch-" but not schlimazel, schmooze
		(strings.Contains(string(m.inWord), "SCH") && !(m.stringAtEnd("IZE", "OZE", "ZEL"))) ||
		((m.current > 0) && m.stringAt(0, "ZEIT")) ||
		m.stringAt(-3, "WEIZ") {
		if (m.current > 0) && m.charAt(m.current-1) == 'T' {
			m.metaphAdd("S")
		} else {
			m.metaphAdd("TS")
		}
		m.current++
		return true
	}

	return false
}

/**
 * Encode "-ZH-" as J
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func (m *Metaphone3) encodeZH() bool {
	//chinese pinyin e.g. 'zhao', also english "phonetic spelling"
	if m.charAt(m.current+1) == 'H' {
		m.metaphAdd("J")
		m.current += 2
		return true
	}

	return false
}

/**
 * Encode "-Z-"
 *
 */
func (m *Metaphone3) encodeZ() {
	if m.encodeZZ() ||
		m.encodeZUZIERZS() ||
		m.encodeFrenchEZ() ||
		m.encodeGermanZ() {
		return
	}

	if m.encodeZH() {
		return
	} else {
		m.metaphAdd("S")
	}

	// eat redundant 'Z'
	if m.charAt(m.current+1) == 'Z' {
		m.current += 2
	} else {
		m.current++
	}
}
