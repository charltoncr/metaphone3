// metaphone3.go - ported to the Go programming language by Ron Charlton
// on 2023-01-05 from the original Java code at
// https://github.com/OpenRefine/OpenRefine/blob/master/main/src/com/google/refine/clustering/binning/Metaphone3.java
//
// $Id: metaphone3.go,v 3.4 2023-01-18 11:34:48-05 ron Exp $
//
// This open source Go file is based on Metaphone3.java 2.1.3 that is
// copyright 2010 by Laurence Philips, and is also open source.
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

    1. Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
    2. Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
    3. Neither the name of Google Inc. nor the names of its
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

==============================================================================

Metaphone 3
VERSION 2.1.3

by Lawrence Philips

Metaphone 3 is designed to return an *approximate* phonetic key (and an alternate
approximate phonetic key when appropriate) that should be the same for English
words, and most names familiar in the United States, that are pronounced *similarly*.
The key value is *not* intended to be an *exact* phonetic, or even phonemic,
representation of the word. This is because a certain degree of 'fuzziness' has
proven to be useful in compensating for variations in pronunciation, as well as
misheard pronunciations. For example, although Americans are not usually aware of it,
the letter 's' is normally pronounced 'z' at the end of words such as "sounds".

The 'approximate' aspect of the encoding is implemented according to the following rules:

(1) All vowels are encoded to the same value - 'A'. If the parameter encodeVowels
is set to false, only *initial* vowels will be encoded at all. If encodeVowels is set
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
words found in american publications from the 89% for Double Metaphone, to over 98%.

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

Metaphone 3 is designed to return an <i>approximate</i> phonetic key (and an alternate
approximate phonetic key when appropriate) that should be the same for English
words, and most names familiar in the United States, that are pronounced "similarly".
The key value is <i>not</i> intended to be an exact phonetic, or even phonemic,
representation of the word. This is because a certain degree of 'fuzziness' has
proven to be useful in compensating for variations in pronunciation, as well as
misheard pronunciations. For example, although Americans are not usually aware of it,
the letter 's' is normally pronounced 'z' at the end of words such as "sounds".

The 'approximate' aspect of the encoding is implemented according to the following rules:

(1) All vowels are encoded to the same value - 'A'. If the parameter encodeVowels
is set to false, only *initial* vowels will be encoded at all. If encodeVowels is set
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
americans are not usually aware of it, "TH" is pronounced in a voiced (e.g. "that") as
well as an unvoiced (e.g. "theater") form, which are naturally mapped to the same encoding.)

In the "Exact" encoding, voiced/unvoiced pairs are <i>not</i> mapped to the same encoding, except
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
are different in the United Kingdom, for example "tube" -> "CHOOBE" -> XAP rather than american TAP.
*/
// End of Metaphone3.java copyright and header comments.

package metaphone3

import (
	"strings"
)

// Metaphone3 defines a type and return length for Metaphone3, as well as
// recording whether vowels after the first character are encoded and whether
// consonants are encoded more exactly.
type Metaphone3 struct {
	maxlen   int
	doVowels bool
	doExact  bool
}

// NewMetaphone3 returns a Metaphone3 instance with a maximum
// length for the Metaphone3 return values.
// Argument maxLen is 4 in the original Double Metaphone algorithm.
func NewMetaphone3(maxLen int) *Metaphone3 {
	return &Metaphone3{
		maxlen:   maxLen,
		doVowels: false,
		doExact:  false,
	}
}

// SetEncodeVowels determines whether or not vowels after the first character
// are encoded.
func (m *Metaphone3) SetEncodeVowels(b bool) {
	m.doVowels = b
}

// SetEncodeExact determines whether or not consonants are encoded more exactly.
func (m *Metaphone3) SetEncodeExact(b bool) {
	m.doExact = b
}

// SetMaxLength sets the maximum length for each of the return values from
// Encode.  If max is less than 1 it is set to 4.
func (m *Metaphone3) SetMaxLength(max int) {
	if max < 1 {
		max = 4
	}
	m.maxlen = max
}

var (
	/** Length of word sent in to be encoded, as
	* measured at beginning of encoding. */
	m_length int

	/** Length of encoded key string. */
	m_metaphLength int

	/** Flag whether or not to encode non-initial vowels. */
	m_encodeVowels bool

	/** Flag whether or not to encode consonants as exactly
	* as possible. */
	m_encodeExact bool

	/** Internal copy of word to be encoded, allocated separately
	* from string pointed to in incoming parameter. */
	m_inWord []rune

	/** Running copy of primary key. */
	m_primary []rune

	/** Running copy of secondary key. */
	m_secondary []rune

	/** Index of character in m_inWord currently being
	* encoded. */
	m_current int

	/** Index of last character in m_inWord. */
	m_last int

	/** Flag that an AL inversion has already been done. */
	flag_AL_inversion bool
)

// Encode returns main and alternate keys for word.  It honors
// the values set by SetEncodeVowels, SetEncodeExact and
// SetMaxLength.  The keys will match for words that sound similar.
// Either key can match either of the other word's keys for similar
// sounding words.
func (m *Metaphone3) Encode(word string) (metaph, metaph2 string) {
	m_encodeVowels = m.doVowels
	m_encodeExact = m.doExact
	m_metaphLength = m.maxlen
	m_inWord = []rune(strings.ToUpper(word))
	m_length = len(m_inWord)
	encode()
	return string(m_primary), string(m_secondary)
}

// metaphAdd adds a string to primary and secondary.  Call it with 1 or 2
// arguments.  The first argument is appended to primary (and to
// secondary if a second argument is not provided).  Any second
// argument is appended to secondary.  But don't append an 'A' next to
// another 'A'.
func metaphAdd(s ...string) {
	if len(s) < 1 || len(s) > 2 {
		panic("metaphAdd requires one or two arguments")
	}
	main := s[0]
	primaryLen := len(m_primary)
	secondaryLen := len(m_secondary)
	if !(main == "A" &&
		(primaryLen > 0) &&
		(m_primary[primaryLen-1] == 'A')) {
		m_primary = append(m_primary, []rune(main)...)
	}
	if len(s) == 1 {
		if !(main == "A" &&
			(secondaryLen > 0) &&
			(m_secondary[secondaryLen-1] == 'A')) {
			m_secondary = append(m_secondary, []rune(main)...)
		}
	} else {
		alt := s[1]
		if !(alt == "A" &&
			(secondaryLen > 0) &&
			(m_secondary[secondaryLen-1] == 'A')) {
			if len(alt) > 0 {
				m_secondary = append(m_secondary, []rune(alt)...)
			}
		}
	}
}

/**
 * Adds an encoding character to the encoded key value string - Exact/Approx version
 *
 * @param mainExact primary encoding character to be added to encoded key string if
 * m_encodeExact is set
 *
 * @param altExact alternative encoding character to be added to encoded alternative
 * key string if m_encodeExact is set
 *
 * @param main primary encoding character to be added to encoded key string
 *
 * @param alt alternative encoding character to be added to encoded alternative key string
 *
 */
func metaphAddExactApprox(s ...string) {
	if len(s) != 2 && len(s) != 4 {
		panic("metaphAddExactApprox requires 2 or 4 arguments")
	}
	var mainExact, altExact, main, alt string
	if len(s) == 2 {
		mainExact = s[0]
		main = s[1]
		if m_encodeExact {
			metaphAdd(mainExact)
		} else {
			metaphAdd(main)
		}
	} else {
		mainExact = s[0]
		altExact = s[1]
		main = s[2]
		alt = s[3]
		if m_encodeExact {
			metaphAdd(mainExact, altExact)
		} else {
			metaphAdd(main, alt)
		}
	}
}

/**
 * Subscript safe .charAt()
 *
 * @param at index of character to access
 * @return null if index out of bounds, .charAt() otherwise
 */
func charAt(at int) rune {
	// check substring bounds
	if at < 0 || at >= m_length {
		return 0
	}

	return m_inWord[at]
}

// stringAtPos determines if any of a list of string arguments appear
// in m_inWord at position pos.
func stringAtPos(pos int, s ...string) bool {
	if len(s) > 0 && pos >= 0 {
	forOuterLoop:
		for _, str := range s {
			if (pos + len(str)) <= m_length {
				j := pos
				for _, r := range str {
					if r != m_inWord[j] {
						continue forOuterLoop
					}
					j++
				}
				return true
			}
		}
	}
	return false
}

// stringAt determines if any of a list of string arguments appear
// in m_inWord at m_current+index.
func stringAt(index int, s ...string) bool {
	start := m_current + index
	return stringAtPos(start, s...)
}

// stringAtStart determines if any of a list of string arguments appear
// in m_inWord at its beginning.
func stringAtStart(s ...string) bool {
	return stringAtPos(0, s...)
}

// stringAtEnd determines if any of a list of string arguments appear
// in m_inWord at its end.
func stringAtEnd(s ...string) bool {
	if len(s) > 0 {
	forOuterLoop:
		for _, str := range s {
			start := m_length - len(str)
			if start >= 0 {
				j := start
				for _, r := range str {
					if r != m_inWord[j] {
						continue forOuterLoop
					}
					j++
				}
				return true
			}
		}
	}
	return false
}

/**
 * Test for close front vowels
 *
 * @return true if close front vowel
 */
func frontVowel(at int) bool {
	c := charAt(at)
	return c == 'E' || c == 'I' || c == 'Y'
}

/**
 * Detect names or words that begin with spellings
 * typical of german or slavic words, for the purpose
 * of choosing alternate pronunciations correctly
 *
 */
func slavoGermanic() bool {
	return stringAtStart("SCH") ||
		stringAtStart("SW") ||
		(charAt(0) == 'J') ||
		(charAt(0) == 'W')
}

/**
 * Tests if character is a vowel
 *
 * @param at rune to be tested in input word or integer location of same.
 * @return true if character is a vowel, false if not
 *
 */
func isVowel(at int) bool {
	return strings.ContainsRune("AEIOUYÀÁÂÃÄÅÆÈÉÊËÌÍÎÏÒÓÔÕÖØÙÚÛÜÝ", charAt(at))
}

/**
 * Skips over vowels in a string. Has exceptions for skipping consonants that
 * will not be encoded.
 *
 * @param at position, in string to be encoded, of character to start skipping from
 *
 * @return position of next consonant in string to be encoded
 */
func skipVowels(at int) int {
	if at < 0 {
		return 0
	}
	if at >= m_length {
		return m_length
	}
	for isVowel(at) || (charAt(at) == 'W') {
		if stringAtPos(at, "WICZ", "WITZ", "WIAK") ||
			stringAtPos(at-1, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
			(stringAtPos(at, "WICKI", "WACKI") && ((at + 4) == m_last)) {
			break
		}
		at++
		if ((charAt(at-1) == 'W') && (charAt(at) == 'H')) &&
			!(stringAtPos(at, "HOP") ||
				stringAtPos(at, "HIDE", "HARD", "HEAD", "HAWK", "HERD", "HOOK", "HAND", "HOLE") ||
				stringAtPos(at, "HEART", "HOUSE", "HOUND") ||
				stringAtPos(at, "HAMMER")) {
			at++
		}
		if at >= m_length {
			break
		}
	}
	return at
}

/**
 * Advanced counter m_current so that it indexes the next character to be encoded
 *
 * @param ifNotEncodeVowels number of characters to advance if not encoding internal vowels
 * @param ifEncodeVowels number of characters to advance if encoding internal vowels
 *
 */
func advanceCounter(ifNotEncodeVowels, ifEncodeVowels int) {
	if !m_encodeVowels {
		m_current += ifNotEncodeVowels
	} else {
		m_current += ifEncodeVowels
	}
}

/**
 * Tests whether the word is the root or a regular english inflection
 * of it, e.g. "ache", "achy", "aches", "ached", "aching", "achingly"
 * This is for cases where we want to match only the root and corresponding
 * inflected forms, and not completely different words which may have the
 * same substring in them.
 */
func rootOrInflections(InWord []rune, root string) bool {
	inWord := string(InWord)
	rootrune := []rune(root)
	len := len(rootrune)
	lastrune := rootrune[len-1]
	var test string

	test = root + "S"
	if inWord == root || inWord == test {
		return true
	}

	if lastrune != 'E' {
		test = root + "ES"
	}

	if inWord == test {
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

	test = root + "ING"
	if inWord == test {
		return true
	}

	test = root + "INGLY"
	if inWord == test {
		return true
	}

	test = root + "Y"
	return inWord == test
}

/**
 * Tests for cases where non-initial 'o' is not pronounced
 * Only executed if non initial vowel encoding is turned on
 *
 * @return true if encoded as silent - no addition to m_metaph key
 *
 */
func oSilent() bool {
	// if "iron" at beginning or end of word and not "irony"
	if (charAt(m_current) == 'O') && stringAt(-2, "IRON") {
		if (stringAtStart("IRON") ||
			(stringAt(-2, "IRON") &&
				(m_last == (m_current + 1)))) &&
			!stringAt(-2, "IRONIC") {
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
func eSilentSuffix(at int) bool {
	return (m_current == (at - 1)) &&
		(m_length > (at + 1)) &&
		(isVowel((at + 1)) ||
			(stringAtPos(at, "ST", "SL") &&
				(m_length > (at + 2))))
}

/**
 * Detect endings that will
 * cause the 'e' to be pronounced
 *
 */
func ePronouncingSuffix(at int) bool {
	// e.g. 'bridgewood' - the other vowels will get eaten
	// up so we need to put one in here
	if (m_length == (at + 4)) && stringAtPos(at, "WOOD") {
		return true
	}

	// same as above
	if (m_length == (at + 5)) && stringAtPos(at, "WATER", "WORTH") {
		return true
	}

	// e.g. 'bridgette'
	if (m_length == (at + 3)) && stringAtPos(at, "TTE", "LIA", "NOW", "ROS", "RAS") {
		return true
	}

	// e.g. 'olena'
	if (m_length == (at + 2)) && stringAtPos(at, "TA", "TT", "NA", "NO", "NE",
		"RS", "RE", "LA", "AU", "RO", "RA") {
		return true
	}

	// e.g. 'bridget'
	if (m_length == (at + 1)) && stringAtPos(at, "T", "R") {
		return true
	}

	return false
}

/**
 * Detect internal silent 'E's e.g. "roseman",
 * "firestone"
 *
 */
func silentInternalE() bool {
	// 'olesen' but not 'olen'	RAKE BLAKE
	return (stringAtStart("OLE") &&
		eSilentSuffix(3) && !ePronouncingSuffix(3)) ||
		(stringAtStart("BARE", "FIRE", "FORE", "GATE", "HAGE", "HAVE",
			"HAZE", "HOLE", "CAPE", "HUSE", "LACE", "LINE",
			"LIVE", "LOVE", "MORE", "MOSE", "MORE", "NICE",
			"RAKE", "ROBE", "ROSE", "SISE", "SIZE", "WARE",
			"WAKE", "WISE", "WINE") &&
			eSilentSuffix(4) && !ePronouncingSuffix(4)) ||
		(stringAtStart("BLAKE", "BRAKE", "BRINE", "CARLE", "CLEVE", "DUNNE",
			"HEDGE", "HOUSE", "JEFFE", "LUNCE", "STOKE", "STONE",
			"THORE", "WEDGE", "WHITE") &&
			eSilentSuffix(5) && !ePronouncingSuffix(5)) ||
		(stringAtStart("BRIDGE", "CHEESE") &&
			eSilentSuffix(6) && !ePronouncingSuffix(6)) ||
		stringAt(-5, "CHARLES")
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
func ePronouncedAtEnd() bool {
	return (m_current == m_last) &&
		(stringAt(-6, "STROPHE") ||
			// if a vowel is before the 'E', vowel eater will have eaten it.
			//otherwise, consonant + 'E' will need 'E' pronounced
			(m_length == 2) ||
			((m_length == 3) && !isVowel(0)) ||
			// these german name endings can be relied on to have the 'e' pronounced
			(stringAtEnd("BKE", "DKE", "FKE", "KKE", "LKE",
				"NKE", "MKE", "PKE", "TKE", "VKE", "ZKE") &&
				!stringAtStart("FINKE", "FUNKE") &&
				!stringAtStart("FRANKE")) ||
			stringAtEnd("SCHKE") ||
			(stringAtStart("ACME", "NIKE", "CAFE", "RENE", "LUPE", "JOSE", "ESME") && (m_length == 4)) ||
			(stringAtStart("LETHE", "CADRE", "TILDE", "SIGNE", "POSSE", "LATTE", "ANIME", "DOLCE", "CROCE",
				"ADOBE", "OUTRE", "JESSE", "JAIME", "JAFFE", "BENGE", "RUNGE",
				"CHILE", "DESME", "CONDE", "URIBE", "LIBRE", "ANDRE") && (m_length == 5)) ||
			(stringAtStart("HECATE", "PSYCHE", "DAPHNE", "PENSKE", "CLICHE", "RECIPE",
				"TAMALE", "SESAME", "SIMILE", "FINALE", "KARATE", "RENATE", "SHANTE",
				"OBERLE", "COYOTE", "KRESGE", "STONGE", "STANGE", "SWAYZE", "FUENTE",
				"SALOME", "URRIBE") && (m_length == 6)) ||
			(stringAtStart("ECHIDNE", "ARIADNE", "MEINEKE", "PORSCHE", "ANEMONE", "EPITOME",
				"SYNCOPE", "SOUFFLE", "ATTACHE", "MACHETE", "KARAOKE", "BUKKAKE",
				"VICENTE", "ELLERBE", "VERSACE") && (m_length == 7)) ||
			(stringAtStart("PENELOPE", "CALLIOPE", "CHIPOTLE", "ANTIGONE", "KAMIKAZE", "EURIDICE",
				"YOSEMITE", "FERRANTE") && (m_length == 8)) ||
			(stringAtStart("HYPERBOLE", "GUACAMOLE", "XANTHIPPE") && (m_length == 9)) ||
			(stringAtStart("SYNECDOCHE") && (m_length == 10)))
}

/**
 * Encodes "-UE".
 *
 * @return true if encoding handled in this routine, false if not
 */
func skipSilentUE() bool {
	// always silent except for cases listed below
	if (stringAt(-1, "QUE", "GUE") &&
		!stringAtStart("BARBEQUE", "PALENQUE", "APPLIQUE") &&
		// '-que' cases usually french but missing the acute accent
		!stringAtStart("RISQUE") &&
		!stringAt(-3, "ARGUE", "SEGUE") &&
		!stringAtStart("PIROGUE", "ENRIQUE") &&
		!stringAtStart("COMMUNIQUE")) &&
		(m_current > 1) &&
		(((m_current + 1) == m_last) ||
			stringAtStart("JACQUES")) {
		m_current = skipVowels(m_current)
		return true
	}

	return false
}

/**
 * Tests and encodes cases where non-initial 'e' is never pronounced
 * Only executed if non initial vowel encoding is turned on
 *
 * @return true if encoded as silent - no addition to m_metaph key
 *
 */
func eSilent() bool {
	if ePronouncedAtEnd() {
		return false
	}

	// 'e' silent when last letter, altho
	return (m_current == m_last) ||
		// also silent if before plural 's'
		// or past tense or participle 'd', e.g.
		// 'grapes' and 'banished' => PNXT
		(stringAtEnd("S", "D") &&
			(m_current > 1) &&
			((m_current + 1) == m_last) &&
			// and not e.g. "nested", "rises", or "pieces" => RASAS
			!(stringAt(-1, "TED", "SES", "CES") ||
				stringAtStart("ANTIPODES", "ANOPHELES") ||
				stringAtStart("MOHAMMED", "MUHAMMED", "MOUHAMED") ||
				stringAtStart("MOHAMED") ||
				stringAtStart("NORRED", "MEDVED", "MERCED", "ALLRED", "KHALED", "RASHED", "MASJED") ||
				stringAtStart("JARED", "AHMED", "HAMED", "JAVED") ||
				stringAtStart("ABED", "IMED"))) ||
		// e.g.  'wholeness', 'boneless', 'barely'
		(stringAt(+1, "NESS", "LESS") && ((m_current + 4) == m_last)) ||
		(stringAt(+1, "LY") && ((m_current + 2) == m_last) &&
			!stringAtStart("CICELY"))
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
func ePronouncedExceptions() bool {
	// greek names e.g. "herakles" or hispanic names e.g. "robles", where 'e' is pronounced, other exceptions
	return (((m_current + 1) == m_last) &&
		(stringAt(-3, "OCLES", "ACLES", "AKLES") ||
			stringAtStart("INES") ||
			stringAtStart("LOPES", "ESTES", "GOMES", "NUNES", "ALVES", "ICKES",
				"INNES", "PERES", "WAGES", "NEVES", "BENES", "DONES") ||
			stringAtStart("CORTES", "CHAVES", "VALDES", "ROBLES", "TORRES", "FLORES", "BORGES",
				"NIEVES", "MONTES", "SOARES", "VALLES", "GEDDES", "ANDRES", "VIAJES",
				"CALLES", "FONTES", "HERMES", "ACEVES", "BATRES", "MATHES") ||
			stringAtStart("DELORES", "MORALES", "DOLORES", "ANGELES", "ROSALES", "MIRELES", "LINARES",
				"PERALES", "PAREDES", "BRIONES", "SANCHES", "CAZARES", "REVELES", "ESTEVES",
				"ALVARES", "MATTHES", "SOLARES", "CASARES", "CACERES", "STURGES", "RAMIRES",
				"FUNCHES", "BENITES", "FUENTES", "PUENTES", "TABARES", "HENTGES", "VALORES") ||
			stringAtStart("GONZALES", "MERCEDES", "FAGUNDES", "JOHANNES", "GONSALES", "BERMUDES",
				"CESPEDES", "BETANCES", "TERRONES", "DIOGENES", "CORRALES", "CABRALES",
				"MARTINES", "GRAJALES") ||
			stringAtStart("CERVANTES", "FERNANDES", "GONCALVES", "BENEVIDES", "CIFUENTES", "SIFUENTES",
				"SERVANTES", "HERNANDES", "BENAVIDES") ||
			stringAtStart("ARCHIMEDES", "CARRIZALES", "MAGALLANES"))) ||
		stringAt(-2, "FRED", "DGES", "DRED", "GNES") ||
		stringAt(-5, "PROBLEM", "RESPLEN") ||
		stringAt(-4, "REPLEN") ||
		stringAt(-3, "SPLE")
}

/**
 * Encodes cases where non-initial 'e' is pronounced, taking
 * care to detect unusual cases from the greek.
 *
 * Only executed if non initial vowel encoding is turned on
 *
 *
 */
func encodeEPronounced() {
	// special cases with two pronunciations
	// 'agape' 'lame' 'resume'
	if (stringAtStart("LAME", "SAKE", "PATE") && (m_length == 4)) ||
		(stringAtStart("AGAPE") && (m_length == 5)) ||
		((m_current == 5) && stringAtStart("RESUME")) {
		metaphAdd("", "A")
		return
	}

	// special case "inge" => 'INGA', 'INJ'
	if stringAtStart("INGE") && (m_length == 4) {
		metaphAdd("A", "")
		return
	}

	// special cases with two pronunciations
	// special handling due to the difference in
	// the pronunciation of the '-D'
	if (m_current == 5) && stringAtStart("BLESSED", "LEARNED") {
		metaphAddExactApprox("D", "AD", "T", "AT")
		m_current += 2
		return
	}

	// encode all vowels and diphthongs to the same value
	if (!eSilent() && !flag_AL_inversion && !silentInternalE()) ||
		ePronouncedExceptions() {
		metaphAdd("A")
	}

	// now that we've visited the vowel in question
	flag_AL_inversion = false
}

/**
 * Encodes silent 'B' for cases not covered under "-mb-"
 *
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSilentB() bool {
	//'debt', 'doubt', 'subtle'
	if stringAt(-2, "DEBT") ||
		stringAt(-2, "SUBTL") ||
		stringAt(-2, "SUBTIL") ||
		stringAt(-3, "DOUBT") {
		metaphAdd("T")
		m_current += 2
		return true
	}

	return false
}

/**
 * Encodes all initial vowels to A.
 *
 * Encodes non-initial vowels to A if m_encodeVowels is true
 *
 *
 */
func encodeVowels() {
	if m_current == 0 {
		// all init vowels map to 'A'
		// as of Double Metaphone
		metaphAdd("A")
	} else if m_encodeVowels {
		if charAt(m_current) != 'E' {
			if skipSilentUE() {
				return
			}

			if oSilent() {
				m_current++
				return
			}

			// encode all vowels and
			// diphthongs to the same value
			metaphAdd("A")
		} else {
			encodeEPronounced()
		}
	}

	if !(!isVowel(m_current-2) && stringAt(-1, "LEWA", "LEWO", "LEWI")) {
		m_current = skipVowels(m_current)
	} else {
		m_current++
	}
}

/**
 * Encodes 'B'
 *
 *
 */
func encodeB() {
	if encodeSilentB() {
		return
	}

	// "-mb", e.g", "dumb", already skipped over under
	// 'M', altho it should really be handled here...
	metaphAddExactApprox("B", "P")

	if (charAt(m_current+1) == 'B') ||
		((charAt(m_current+1) == 'P') &&
			((m_current+1 < m_last) && (charAt(m_current+2) != 'H'))) {
		m_current += 2
	} else {
		m_current++
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
func encodeCHToH() bool {
	// hebrew => 'H', e.g. 'channukah', 'chabad'
	if ((m_current == 0) &&
		(stringAt(+2, "AIM", "ETH", "ELM") ||
			stringAt(+2, "ASID", "AZAN") ||
			stringAt(+2, "UPPAH", "UTZPA", "ALLAH", "ALUTZ", "AMETZ") ||
			stringAt(+2, "ESHVAN", "ADARIM", "ANUKAH") ||
			stringAt(+2, "ALLLOTH", "ANNUKAH", "AROSETH"))) ||
		// and an irish name with the same encoding
		stringAt(-3, "CLACHAN") {
		metaphAdd("H")
		advanceCounter(3, 2)
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
func encodeSilentCAtBeginning() bool {
	//skip these when at start of word
	if (m_current == 0) && stringAt(0, "CT", "CN") {
		m_current += 1
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
func encodeCAToS() bool {
	// Special case: 'caesar'.
	// Also, where cedilla not used, as in "linguica" => LNKS
	if ((m_current == 0) && stringAt(0, "CAES", "CAEC", "CAEM")) ||
		stringAtStart("FRANCAIS", "FRANCAIX", "LINGUICA") ||
		stringAtStart("FACADE") ||
		stringAtStart("GONCALVES", "PROVENCAL") {
		metaphAdd("S")
		advanceCounter(2, 1)
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
func encodeCOToS() bool {
	// e.g. 'coelecanth' => SLKN0
	if (stringAt(0, "COEL") &&
		(isVowel(m_current+4) || ((m_current + 3) == m_last))) ||
		stringAt(0, "COENA", "COENO") ||
		stringAtStart("FRANCOIS", "MELANCON") ||
		stringAtStart("GARCON") {
		metaphAdd("S")
		advanceCounter(3, 1)
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
func encodeCHAE() bool {
	// e.g. 'michael'
	if (m_current > 0) && stringAt(+2, "AE") {
		if stringAtStart("RACHAEL") {
			metaphAdd("X")
		} else if !stringAt(-1, "C", "K", "G", "Q") {
			metaphAdd("K")
		}

		advanceCounter(4, 2)
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
func encodeSilentCH() bool {
	// '-ch-' not pronounced
	if stringAt(-2, "FUCHSIA") ||
		stringAt(-2, "YACHT") ||
		stringAtStart("STRACHAN") ||
		stringAtStart("CRICHTON") ||
		(stringAt(-3, "DRACHM")) &&
			!stringAt(-3, "DRACHMA") {
		m_current += 2
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
func encodeCHToX() bool {
	// e.g. 'approach', 'beach'
	if (stringAt(-2, "OACH", "EACH", "EECH", "OUCH", "OOCH", "MUCH", "SUCH") &&
		!stringAt(-3, "JOACH")) ||
		// e.g. 'dacha', 'macho'
		(((m_current + 2) == m_last) && stringAt(-1, "ACHA", "ACHO")) ||
		(stringAt(0, "CHOT", "CHOD", "CHAT") && ((m_current + 3) == m_last)) ||
		((stringAt(-1, "OCHE") && ((m_current + 2) == m_last)) &&
			!stringAt(-2, "DOCHE")) ||
		stringAt(-4, "ATTACH", "DETACH", "KOVACH") ||
		stringAt(-5, "SPINACH") ||
		stringAtStart("MACHAU") ||
		stringAt(-4, "PARACHUT") ||
		stringAt(-5, "MASSACHU") ||
		(stringAt(-3, "THACH") && !stringAt(-1, "ACHE")) ||
		stringAt(-2, "VACHON") {
		metaphAdd("X")
		m_current += 2
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
func encodeEnglishCHToK() bool {
	//'ache', 'echo', alternate spelling of 'michael'
	if ((m_current == 1) && rootOrInflections(m_inWord, "ACHE")) ||
		(((m_current > 3) && rootOrInflections(m_inWord[m_current-1:], "ACHE")) &&
			(stringAtStart("EAR") ||
				stringAtStart("HEAD", "BACK") ||
				stringAtStart("HEART", "BELLY", "TOOTH"))) ||
		stringAt(-1, "ECHO") ||
		stringAt(-2, "MICHEAL") ||
		stringAt(-4, "JERICHO") ||
		stringAt(-5, "LEPRECH") {
		metaphAdd("K", "X")
		m_current += 2
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
func encodeGermanicCHToK() bool {
	// various germanic
	// "<consonant><vowel>CH-" implies a german word where 'ch' => K
	if ((m_current > 1) &&
		!isVowel(m_current-2) &&
		stringAt(-1, "ACH") &&
		!stringAt(-2, "MACHADO", "MACHUCA", "LACHANC", "LACHAPE", "KACHATU") &&
		!stringAt(-3, "KHACHAT") &&
		((charAt(m_current+2) != 'I') &&
			((charAt(m_current+2) != 'E') ||
				stringAt(-2, "BACHER", "MACHER", "MACHEN", "LACHER"))) ||
		// e.g. 'brecht', 'fuchs'
		(stringAt(+2, "T", "S") &&
			!(stringAtStart("WHICHSOEVER") || stringAtStart("LUNCHTIME"))) ||
		// e.g. 'andromache'
		stringAtStart("SCHR") ||
		((m_current > 2) && stringAt(-2, "MACHE")) ||
		((m_current == 2) && stringAt(-2, "ZACH")) ||
		stringAt(-4, "SCHACH") ||
		stringAt(-1, "ACHEN") ||
		stringAt(-3, "SPICH", "ZURCH", "BUECH") ||
		(stringAt(-3, "KIRCH", "JOACH", "BLECH", "MALCH") &&
			// "kirch" and "blech" both get 'X'
			!(stringAt(-3, "KIRCHNER") || ((m_current + 1) == m_last))) ||
		(((m_current + 1) == m_last) && stringAt(-2, "NICH", "LICH", "BACH")) ||
		(((m_current + 1) == m_last) &&
			stringAt(-3, "URICH", "BRICH", "ERICH", "DRICH", "NRICH") &&
			!stringAt(-5, "ALDRICH") &&
			!stringAt(-6, "GOODRICH") &&
			!stringAt(-7, "GINGERICH"))) ||
		(((m_current + 1) == m_last) && stringAt(-4, "ULRICH", "LFRICH", "LLRICH",
			"EMRICH", "ZURICH", "EYRICH")) ||
		// e.g., 'wachtler', 'wechsler', but not 'tichner'
		((stringAt(-1, "A", "O", "U", "E") || (m_current == 0)) &&
			stringAt(+2, "L", "R", "N", "M", "B", "H", "F", "V", "W", " ")) {
		// "CHR/L-" e.g. 'chris' do not get
		// alt pronunciation of 'X'
		if stringAt(+2, "R", "L") || slavoGermanic() {
			metaphAdd("K")
		} else {
			metaphAdd("K", "X")
		}
		m_current += 2
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
func encodeARCH() bool {
	if stringAt(-2, "ARCH") {
		// "-ARCH-" has many combining forms where "-CH-" => K because of its
		// derivation from the greek
		if ((isVowel(m_current+2) && stringAt(-2, "ARCHA", "ARCHI", "ARCHO", "ARCHU", "ARCHY")) ||
			stringAt(-2, "ARCHEA", "ARCHEG", "ARCHEO", "ARCHET", "ARCHEL", "ARCHES", "ARCHEP",
				"ARCHEM", "ARCHEN") ||
			(stringAt(-2, "ARCH") && ((m_current + 1) == m_last)) ||
			stringAtStart("MENARCH")) &&
			(!rootOrInflections(m_inWord, "ARCH") &&
				!stringAt(-4, "SEARCH", "POARCH") &&
				!stringAtStart("ARCHENEMY", "ARCHIBALD", "ARCHULETA", "ARCHAMBAU") &&
				!stringAtStart("ARCHER", "ARCHIE") &&
				!((((stringAt(-3, "LARCH", "MARCH", "PARCH") ||
					stringAt(-4, "STARCH")) &&
					!(stringAtStart("EPARCH") ||
						stringAtStart("NOMARCH") ||
						stringAtStart("EXILARCH", "HIPPARCH", "MARCHESE") ||
						stringAtStart("ARISTARCH") ||
						stringAtStart("MARCHETTI"))) ||
					rootOrInflections(m_inWord, "STARCH")) &&
					(!stringAt(-2, "ARCHU", "ARCHY") ||
						stringAtStart("STARCHY")))) {
			metaphAdd("K", "X")
		} else {
			metaphAdd("X")
		}
		m_current += 2
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
func encodeGreekCHInitial() bool {
	// greek roots e.g. 'chemistry', 'chorus', ch at beginning of root
	if (stringAt(0, "CHAMOM", "CHARAC", "CHARIS", "CHARTO", "CHARTU", "CHARYB", "CHRIST", "CHEMIC", "CHILIA") ||
		(stringAt(0, "CHEMI", "CHEMO", "CHEMU", "CHEMY", "CHOND", "CHONA", "CHONI", "CHOIR", "CHASM",
			"CHARO", "CHROM", "CHROI", "CHAMA", "CHALC", "CHALD", "CHAET", "CHIRO", "CHILO", "CHELA", "CHOUS",
			"CHEIL", "CHEIR", "CHEIM", "CHITI", "CHEOP") &&
			!(stringAt(0, "CHEMIN") || stringAt(-2, "ANCHONDO"))) ||
		(stringAt(0, "CHISM", "CHELI") &&
			// exclude spanish "machismo"
			!(stringAtStart("MACHISMO") ||
				// exclude some french words
				stringAtStart("REVANCHISM") ||
				stringAtStart("RICHELIEU") ||
				(stringAtStart("CHISM") && (m_length == 5)) ||
				stringAtStart("MICHEL"))) ||
		// include e.g. "chorus", "chyme", "chaos"
		(stringAt(0, "CHOR", "CHOL", "CHYM", "CHYL", "CHLO", "CHOS", "CHUS", "CHOE") &&
			!stringAtStart("CHOLLO", "CHOLLA", "CHORIZ")) ||
		// "chaos" => K but not "chao"
		(stringAt(0, "CHAO") && ((m_current + 3) != m_last)) ||
		// e.g. "abranchiate"
		(stringAt(0, "CHIA") && !(stringAtStart("APPALACHIA") || stringAtStart("CHIAPAS"))) ||
		// e.g. "chimera"
		stringAt(0, "CHIMERA", "CHIMAER", "CHIMERI") ||
		// e.g. "chameleon"
		((m_current == 0) && stringAt(0, "CHAME", "CHELO", "CHITO")) ||
		// e.g. "spirochete"
		((((m_current + 4) == m_last) || ((m_current + 5) == m_last)) && stringAt(-1, "OCHETE"))) &&
		// more exceptions where "-CH-" => X e.g. "chortle", "crocheter"
		!((stringAtStart("CHORE", "CHOLO", "CHOLA") && (m_length == 5)) ||
			stringAt(0, "CHORT", "CHOSE") ||
			stringAt(-3, "CROCHET") ||
			stringAtStart("CHEMISE", "CHARISE", "CHARISS", "CHAROLE")) {
		// "CHR/L-" e.g. 'christ', 'chlorine' do not get
		// alt pronunciation of 'X'
		if stringAt(+2, "R", "L") {
			metaphAdd("K")
		} else {
			metaphAdd("K", "X")
		}
		m_current += 2
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
func encodeGreekCHNonInitial() bool {
	//greek & other roots e.g. 'tachometer', 'orchid', ch in middle or end of root
	if stringAt(-2, "ORCHID", "NICHOL", "MECHAN", "LICHEN", "MACHIC", "PACHEL", "RACHIF", "RACHID",
		"RACHIS", "RACHIC", "MICHAL") ||
		stringAt(-3, "MELCH", "GLOCH", "TRACH", "TROCH", "BRACH", "SYNCH", "PSYCH",
			"STICH", "PULCH", "EPOCH") ||
		(stringAt(-3, "TRICH") && !stringAt(-5, "OSTRICH")) ||
		(stringAt(-2, "TYCH", "TOCH", "BUCH", "MOCH", "CICH", "DICH", "NUCH", "EICH", "LOCH",
			"DOCH", "ZECH", "WYCH") &&
			!(stringAt(-4, "INDOCHINA") || stringAt(-2, "BUCHON"))) ||
		stringAt(-2, "LYCHN", "TACHO", "ORCHO", "ORCHI", "LICHO") ||
		(stringAt(-1, "OCHER", "ECHIN", "ECHID") && ((m_current == 1) || (m_current == 2))) ||
		stringAt(-4, "BRONCH", "STOICH", "STRYCH", "TELECH", "PLANCH", "CATECH", "MANICH", "MALACH",
			"BIANCH", "DIDACH") ||
		(stringAt(-1, "ICHA", "ICHN") && (m_current == 1)) ||
		stringAt(-2, "ORCHESTR") ||
		stringAt(-4, "BRANCHIO", "BRANCHIF") ||
		(stringAt(-1, "ACHAB", "ACHAD", "ACHAN", "ACHAZ") &&
			!stringAt(-2, "MACHADO", "LACHANC")) ||
		stringAt(-1, "ACHISH", "ACHILL", "ACHAIA", "ACHENE") ||
		stringAt(-1, "ACHAIAN", "ACHATES", "ACHIRAL", "ACHERON") ||
		stringAt(-1, "ACHILLEA", "ACHIMAAS", "ACHILARY", "ACHELOUS", "ACHENIAL", "ACHERNAR") ||
		stringAt(-1, "ACHALASIA", "ACHILLEAN", "ACHIMENES") ||
		stringAt(-1, "ACHIMELECH", "ACHITOPHEL") ||
		// e.g. 'inchoate'
		(((m_current - 2) == 0) && (stringAt(-2, "INCHOA") ||
			// e.g. 'ischemia'
			stringAtStart("ISCH"))) ||
		// e.g. 'ablimelech', 'antioch', 'pentateuch'
		(((m_current + 1) == m_last) && stringAt(-1, "A", "O", "U", "E") &&
			!(stringAtStart("DEBAUCH") ||
				stringAt(-2, "MUCH", "SUCH", "KOCH") ||
				stringAt(-5, "OODRICH", "ALDRICH"))) {
		metaphAdd("K", "X")
		m_current += 2
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
func encodeCH() bool {
	if stringAt(0, "CH") {
		if encodeCHAE() ||
			encodeCHToH() ||
			encodeSilentCH() ||
			encodeARCH() ||
			// encodeCHToX() should be
			// called before the germanic
			// and greek encoding functions
			encodeCHToX() ||
			encodeEnglishCHToK() ||
			encodeGermanicCHToK() ||
			encodeGreekCHInitial() ||
			encodeGreekCHNonInitial() {
			return true
		}

		if m_current > 0 {
			if stringAtStart("MC") && (m_current == 1) {
				//e.g., "McHugh"
				metaphAdd("K")
			} else {
				metaphAdd("X", "K")
			}
		} else {
			metaphAdd("X")
		}
		m_current += 2
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
func encodeCCIA() bool {
	//e.g., 'focaccia'
	if stringAt(+1, "CIA") {
		metaphAdd("X", "S")
		m_current += 2
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
func encodeCC() bool {
	//double 'C', but not if e.g. 'McClellan'
	if stringAt(0, "CC") && !((m_current == 1) && (charAt(0) == 'M')) {
		// exception
		if stringAt(-3, "FLACCID") {
			metaphAdd("S")
			advanceCounter(3, 2)
			return true
		}

		//'bacci', 'bertucci', other italian
		if (((m_current + 2) == m_last) && stringAt(+2, "I")) ||
			stringAt(+2, "IO") ||
			(((m_current + 4) == m_last) && stringAt(+2, "INO", "INI")) {
			metaphAdd("X")
			advanceCounter(3, 2)
			return true
		}

		//'accident', 'accede' 'succeed'
		if stringAt(+2, "I", "E", "Y") &&
			//except 'bellocchio','bacchus', 'soccer' get K
			!((charAt(m_current+2) == 'H') ||
				stringAt(-2, "SOCCER")) {
			metaphAdd("KS")
			advanceCounter(3, 2)
			return true

		} else {
			//Pierce's rule
			metaphAdd("K")
			m_current += 2
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
func encodeCKCGCQ() bool {
	if stringAt(0, "CK", "CG", "CQ") {
		// eastern european spelling e.g. 'gorecki' == 'goresky'
		if stringAt(0, "CKI", "CKY") &&
			((m_current + 2) == m_last) &&
			(m_length > 6) {
			metaphAdd("K", "SK")
		} else {
			metaphAdd("K")
		}
		m_current += 2

		if stringAt(0, "K", "G", "Q") {
			m_current++
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
func encodeBritishSilentCE() bool {
	// english place names like e.g.'gloucester' pronounced glo-ster
	return (stringAt(+1, "ESTER") && ((m_current + 5) == m_last)) ||
		stringAt(+1, "ESTERSHIRE")
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeCE() bool {
	// 'ocean', 'commercial', 'provincial', 'cello', 'fettucini', 'medici'
	if (stringAt(+1, "EAN") && isVowel(m_current-1)) ||
		// e.g. 'rosacea'
		(stringAt(-1, "ACEA") &&
			((m_current + 2) == m_last) &&
			!stringAtStart("PANACEA")) ||
		// e.g. 'botticelli', 'concerto'
		stringAt(+1, "ELLI", "ERTO", "EORL") ||
		// some italian names familiar to americans
		(stringAt(-3, "CROCE") && ((m_current + 1) == m_last)) ||
		stringAt(-3, "DOLCE") ||
		// e.g. 'cello'
		(stringAt(+1, "ELLO") &&
			((m_current + 4) == m_last)) {
		metaphAdd("X", "S")
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeCI() bool {
	// with consonant before C
	// e.g. 'fettucini', but exception for the americanized pronunciation of 'mancini'
	if ((stringAt(+1, "INI") && !stringAtStart("MANCINI")) && ((m_current + 3) == m_last)) ||
		// e.g. 'medici'
		(stringAt(-1, "ICI") && ((m_current + 1) == m_last)) ||
		// e.g. "commercial', 'provincial', 'cistercian'
		stringAt(-1, "RCIAL", "NCIAL", "RCIAN", "UCIUS") ||
		// special cases
		stringAt(-3, "MARCIA") ||
		stringAt(-2, "ANCIENT") {
		metaphAdd("X", "S")
		return true
	}

	// with vowel before C (or at beginning?)
	if ((stringAt(0, "CIO", "CIE", "CIA") &&
		isVowel(m_current-1)) ||
		// e.g. "ciao"
		stringAt(+1, "IAO")) &&
		!stringAt(-4, "COERCION") {
		if (stringAt(0, "CIAN", "CIAL", "CIAO", "CIES", "CIOL", "CION") ||
			// exception - "glacier" => 'X' but "spacier" = > 'S'
			stringAt(-3, "GLACIER") ||
			stringAt(0, "CIENT", "CIENC", "CIOUS", "CIATE", "CIATI", "CIATO", "CIABL", "CIARY") ||
			(((m_current + 2) == m_last) && stringAt(0, "CIA", "CIO")) ||
			(((m_current + 3) == m_last) && stringAt(0, "CIAS", "CIOS"))) &&
			// exceptions
			!(stringAt(-4, "ASSOCIATION") ||
				stringAtStart("OCIE") ||
				// exceptions mostly because these names are usually from
				// the spanish rather than the italian in america
				stringAt(-2, "LUCIO") ||
				stringAt(-2, "MACIAS") ||
				stringAt(-3, "GRACIE", "GRACIA") ||
				stringAt(-2, "LUCIANO") ||
				stringAt(-3, "MARCIANO") ||
				stringAt(-4, "PALACIO") ||
				stringAt(-4, "FELICIANO") ||
				stringAt(-5, "MAURICIO") ||
				stringAt(-7, "ENCARNACION") ||
				stringAt(-4, "POLICIES") ||
				stringAt(-2, "HACIENDA") ||
				stringAt(-6, "ANDALUCIA") ||
				stringAt(-2, "SOCIO", "SOCIE")) {
			metaphAdd("X", "S")
		} else {
			metaphAdd("S", "X")
		}

		return true
	}

	// exception
	if stringAt(-4, "COERCION") {
		metaphAdd("J")
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeLatinateSuffixes() bool {
	if stringAt(+1, "EOUS", "IOUS") {
		metaphAdd("X", "S")
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
func encodeCFrontVowel() bool {
	if stringAt(0, "CI", "CE", "CY") {
		if encodeBritishSilentCE() ||
			encodeCE() ||
			encodeCI() ||
			encodeLatinateSuffixes() {
			advanceCounter(2, 1)
			return true
		}

		metaphAdd("S")
		advanceCounter(2, 1)
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
func encodeSilentC() bool {
	if stringAt(+1, "T", "S") {
		if stringAtStart("CONNECTICUT") ||
			stringAtStart("INDICT", "TUCSON") {
			m_current++
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
func encodeCZ() bool {
	if stringAt(+1, "Z") &&
		!stringAt(-1, "ECZEMA") {
		if stringAt(0, "CZAR") {
			metaphAdd("S")
		} else {
			// otherwise most likely a czech word...
			metaphAdd("X")
		}
		m_current += 2
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
func encodeCS() bool {
	// give an 'etymological' 2nd
	// encoding for "kovacs" so
	// that it matches "kovach"
	if stringAtStart("KOVACS") {
		metaphAdd("KS", "X")
		m_current += 2
		return true
	}

	if stringAt(-1, "ACS") &&
		((m_current + 1) == m_last) &&
		!stringAt(-4, "ISAACS") {
		metaphAdd("X")
		m_current += 2
		return true
	}

	return false
}

/**
 * Encodes 'C'
 *
 */
func encodeC() {
	if encodeSilentCAtBeginning() ||
		encodeCAToS() ||
		encodeCOToS() ||
		encodeCH() ||
		encodeCCIA() ||
		encodeCC() ||
		encodeCKCGCQ() ||
		encodeCFrontVowel() ||
		encodeSilentC() ||
		encodeCZ() ||
		encodeCS() {
		return
	}

	//else
	if !stringAt(-1, "C", "K", "G", "Q") {
		metaphAdd("K")
	}

	//name sent in 'mac caffrey', 'mac gregor
	if stringAt(+1, " C", " Q", " G") {
		m_current += 2
	} else {
		if stringAt(+1, "C", "K", "Q") &&
			!stringAt(+1, "CE", "CI") {
			m_current += 2
			// account for combinations such as Ro-ckc-liffe
			if stringAt(0, "C", "K", "Q") &&
				!stringAt(+1, "CE", "CI") {
				m_current++
			}
		} else {
			m_current++
		}
	}
}

/**
 * Encode "-DG-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeDG() bool {
	if stringAt(0, "DG") {
		// excludes exceptions e.g. 'edgar',
		// or cases where 'g' is first letter of combining form
		// e.g. 'handgun', 'waldglas'
		if stringAt(+2, "A", "O") ||
			// e.g. "midgut"
			stringAt(+1, "GUN", "GUT") ||
			// e.g. "handgrip"
			stringAt(+1, "GEAR", "GLAS", "GRIP", "GREN", "GILL", "GRAF") ||
			// e.g. "mudgard"
			stringAt(+1, "GUARD", "GUILT", "GRAVE", "GRASS") ||
			// e.g. "woodgrouse"
			stringAt(+1, "GROUSE") {
			metaphAddExactApprox("DG", "TK")
		} else {
			//e.g. "edge", "abridgment"
			metaphAdd("J")
		}
		m_current += 2
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
func encodeDJ() bool {
	// e.g. "adjacent"
	if stringAt(0, "DJ") {
		metaphAdd("J")
		m_current += 2
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
func encodeDTDD() bool {
	// eat redundant 'T' or 'D'
	if stringAt(0, "DT", "DD") {
		if stringAt(0, "DTH") {
			metaphAddExactApprox("D0", "T0")
			m_current += 3
		} else {
			if m_encodeExact {
				// devoice it
				if stringAt(0, "DT") {
					metaphAdd("T")
				} else {
					metaphAdd("D")
				}
			} else {
				metaphAdd("T")
			}
			m_current += 2
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
func encodeDToJ() bool {
	// e.g. "module", "adulate"
	if (stringAt(0, "DUL") &&
		(isVowel(m_current-1) && isVowel(m_current+3))) ||
		// e.g. "soldier", "grandeur", "procedure"
		(((m_current + 3) == m_last) &&
			stringAt(-1, "LDIER", "NDEUR", "EDURE", "RDURE")) ||
		stringAt(-3, "CORDIAL") ||
		// e.g.  "pendulum", "education"
		stringAt(-1, "NDULA", "NDULU", "EDUCA") ||
		// e.g. "individual", "individual", "residuum"
		stringAt(-1, "ADUA", "IDUA", "IDUU") {
		metaphAddExactApprox("J", "D", "J", "T")
		advanceCounter(2, 1)
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
func encodeDOUS() bool {
	// e.g. "assiduous", "arduous"
	if stringAt(+1, "UOUS") {
		metaphAddExactApprox("J", "D", "J", "T")
		advanceCounter(4, 1)
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
func encodeSilentD() bool {
	// silent 'D' e.g. 'wednesday', 'handsome'
	if stringAt(-2, "WEDNESDAY") ||
		stringAt(-3, "HANDKER", "HANDSOM", "WINDSOR") ||
		// french silent D at end in words or names familiar to americans
		stringAt(-5, "PERNOD", "ARTAUD", "RENAUD") ||
		stringAt(-6, "RIMBAUD", "MICHAUD", "BICHAUD") {
		m_current++
		return true
	}

	return false
}

/**
 * Encode "-D-"
 *
 */
func encodeD() {
	if encodeDG() ||
		encodeDJ() ||
		encodeDTDD() ||
		encodeDToJ() ||
		encodeDOUS() ||
		encodeSilentD() {
		return
	}

	if m_encodeExact {
		// "final de-voicing" in this case
		// e.g. 'missed' == 'mist'
		if (m_current == m_last) &&
			stringAt(-3, "SSED") {
			metaphAdd("T")
		} else {
			metaphAdd("D")
		}
	} else {
		metaphAdd("T")
	}
	m_current++
}

/**
 * Encode "-F-"
 *
 */
func encodeF() {
	// Encode cases where "-FT-" => "T" is usually silent
	// e.g. 'often', 'soften'
	// This should really be covered under "T"!
	if stringAt(-1, "OFTEN") {
		metaphAdd("F", "FT")
		m_current += 2
		return
	}

	// eat redundant 'F'
	if charAt(m_current+1) == 'F' {
		m_current += 2
	} else {
		m_current++
	}

	metaphAdd("F")

}

/**
 * Encode cases where 'G' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSilentGAtBeginning() bool {
	//skip these when at start of word
	if (m_current == 0) &&
		stringAt(0, "GN") {
		m_current += 1
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
func encodeGG() bool {
	if charAt(m_current+1) == 'G' {
		// italian e.g, 'loggia', 'caraveggio', also 'suggest' and 'exaggerate'
		if stringAt(-1, "AGGIA", "OGGIA", "AGGIO", "EGGIO", "EGGIA", "IGGIO") ||
			// 'ruggiero' but not 'snuggies'
			(stringAt(-1, "UGGIE") && !(((m_current + 3) == m_last) || ((m_current + 4) == m_last))) ||
			(((m_current + 2) == m_last) && stringAt(-1, "AGGI", "OGGI")) ||
			stringAt(-2, "SUGGES", "XAGGER", "REGGIE") {
			// expection where "-GG-" => KJ
			if stringAt(-2, "SUGGEST") {
				metaphAddExactApprox("G", "K")
			}

			metaphAdd("J")
			advanceCounter(3, 2)
		} else {
			metaphAddExactApprox("G", "K")
			m_current += 2
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
func encodeGK() bool {
	// 'gingko'
	if charAt(m_current+1) == 'K' {
		metaphAdd("K")
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeGHAfterConsonant() bool {
	// e.g. 'burgher', 'bingham'
	if (m_current > 0) &&
		!isVowel(m_current-1) &&
		// not e.g. 'greenhalgh'
		!(stringAt(-3, "HALGH") &&
			((m_current + 1) == m_last)) {
		metaphAddExactApprox("G", "K")
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeInitialGH() bool {
	if m_current < 3 {
		// e.g. "ghislane", "ghiradelli"
		if m_current == 0 {
			if charAt(m_current+2) == 'I' {
				metaphAdd("J")
			} else {
				metaphAddExactApprox("G", "K")
			}
			m_current += 2
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
func encodeGHToJ() bool {
	// e.g., 'greenhalgh', 'dunkenhalgh', english names
	if stringAt(-2, "ALGH") && ((m_current + 1) == m_last) {
		metaphAdd("J", "")
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeGHToH() bool {
	// special cases
	// e.g., 'donoghue', 'donaghy'
	if (stringAt(-4, "DONO", "DONA") && isVowel(m_current+2)) ||
		stringAt(-5, "CALLAGHAN") {
		metaphAdd("H")
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeUGHT() bool {
	//e.g. "ought", "aught", "daughter", "slaughter"
	if stringAt(-1, "UGHT") {
		if (stringAt(-3, "LAUGH") &&
			!(stringAt(-4, "SLAUGHT") ||
				stringAt(-3, "LAUGHTO"))) ||
			stringAt(-4, "DRAUGH") {
			metaphAdd("FT")
		} else {
			metaphAdd("T")
		}
		m_current += 3
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeGHHPartOfOtherWord() bool {
	// if the 'H' is the beginning of another word or syllable
	if stringAt(+1, "HOUS", "HEAD", "HOLE", "HORN", "HARN") {
		metaphAddExactApprox("G", "K")
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSilentGH() bool {
	//Parker's rule (with some further refinements) - e.g., 'hugh'
	if ((((m_current > 1) && stringAt(-2, "B", "H", "D", "G", "L")) ||
		//e.g., 'bough'
		((m_current > 2) &&
			stringAt(-3, "B", "H", "D", "K", "W", "N", "P", "V") &&
			!stringAtStart("ENOUGH")) ||
		//e.g., 'broughton'
		((m_current > 3) && stringAt(-4, "B", "H")) ||
		//'plough', 'slaugh'
		((m_current > 3) && stringAt(-4, "PL", "SL")) ||
		((m_current > 0) &&
			// 'sigh', 'light'
			((charAt(m_current-1) == 'I') ||
				stringAtStart("PUGH") ||
				// e.g. 'MCDONAGH', 'MURTAGH', 'CREAGH'
				(stringAt(-1, "AGH") &&
					((m_current + 1) == m_last)) ||
				stringAt(-4, "GERAGH", "DRAUGH") ||
				(stringAt(-3, "GAUGH", "GEOGH", "MAUGH") &&
					!stringAtStart("MCGAUGHEY")) ||
				// exceptions to 'tough', 'rough', 'lough'
				(stringAt(-2, "OUGH") &&
					(m_current > 3) &&
					!stringAt(-4, "CCOUGH", "ENOUGH", "TROUGH", "CLOUGH"))))) &&
		// suffixes starting w/ vowel where "-GH-" is usually silent
		(stringAt(-3, "VAUGH", "FEIGH", "LEIGH") ||
			stringAt(-2, "HIGH", "TIGH") ||
			((m_current + 1) == m_last) ||
			(stringAt(+2, "IE", "EY", "ES", "ER", "ED", "TY") &&
				((m_current + 3) == m_last) &&
				!stringAt(-5, "GALLAGHER")) ||
			(stringAt(+2, "Y") && ((m_current + 2) == m_last)) ||
			(stringAt(+2, "ING", "OUT") && ((m_current + 4) == m_last)) ||
			(stringAt(+2, "ERTY") && ((m_current + 5) == m_last)) ||
			(!isVowel(m_current+2) ||
				stringAt(-3, "GAUGH", "GEOGH", "MAUGH") ||
				stringAt(-4, "BROUGHAM")))) &&
		// exceptions where '-g-' pronounced
		!(stringAtStart("BALOGH", "SABAGH") ||
			stringAt(-2, "BAGHDAD") ||
			stringAt(-3, "WHIGH") ||
			stringAt(-5, "SABBAGH", "AKHLAGH")) {
		// silent - do nothing
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeGHSpecialCases() bool {
	var handled = false

	// special case: 'hiccough' == 'hiccup'
	if stringAt(-6, "HICCOUGH") {
		metaphAdd("P")
		handled = true
	} else if stringAtStart("LOUGH") {
		// special case: 'lough' alt spelling for scots 'loch'
		metaphAdd("K")
		handled = true
	} else if stringAtStart("BALOGH") {
		// hungarian
		metaphAddExactApprox("G", "", "K", "")
		handled = true
	} else if stringAt(-3, "LAUGHLIN", "COUGHLAN", "LOUGHLIN") {
		// "maclaughlin"
		metaphAdd("K", "F")
		handled = true
	} else if stringAt(-3, "GOUGH") ||
		stringAt(-7, "COLCLOUGH") {
		metaphAdd("", "F")
		handled = true
	}

	if handled {
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeGHToF() bool {
	// the cases covered here would fall under
	// the GH_To_F rule below otherwise
	if encodeGHSpecialCases() {
		return true
	} else {
		//e.g., 'laugh', 'cough', 'rough', 'tough'
		if (m_current > 2) &&
			(charAt(m_current-1) == 'U') &&
			isVowel(m_current-2) &&
			stringAt(-3, "C", "G", "L", "R", "T", "N", "S") &&
			!stringAt(-4, "BREUGHEL", "FLAUGHER") {
			metaphAdd("F")
			m_current += 2
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
func encodeGH() bool {
	if charAt(m_current+1) == 'H' {
		if encodeGHAfterConsonant() ||
			encodeInitialGH() ||
			encodeGHToJ() ||
			encodeGHToH() ||
			encodeUGHT() ||
			encodeGHHPartOfOtherWord() ||
			encodeSilentGH() ||
			encodeGHToF() {
			return true
		}

		metaphAddExactApprox("G", "K")
		m_current += 2
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
func encodeSilentG() bool {
	// e.g. "phlegm", "apothegm", "voigt"
	if (((m_current + 1) == m_last) &&
		(stringAt(-1, "EGM", "IGM", "AGM") ||
			stringAt(0, "GT"))) ||
		(stringAtStart("HUGES") && (m_length == 5)) {
		m_current++
		return true
	}

	// vietnamese names e.g. "Nguyen" but not "Ng"
	if stringAtStart("NG") && (m_current != m_last) {
		m_current++
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
func encodeGN() bool {
	if charAt(m_current+1) == 'N' {
		// 'align' 'sign', 'resign' but not 'resignation'
		// also 'impugn', 'impugnable', but not 'repugnant'
		if ((m_current > 1) &&
			((stringAt(-1, "I", "U", "E") ||
				stringAt(-3, "LORGNETTE") ||
				stringAt(-2, "LAGNIAPPE") ||
				stringAt(-2, "COGNAC") ||
				stringAt(-3, "CHAGNON") ||
				stringAt(-5, "COMPAGNIE") ||
				stringAt(-4, "BOLOGN")) &&
				// Exceptions: following are cases where 'G' is pronounced
				// in "assign" 'g' is silent, but not in "assignation"
				!(stringAt(+2, "ATION") ||
					stringAt(+2, "ATOR") ||
					stringAt(+2, "ATE", "ITY") ||
					// exception to exceptions, not pronounced:
					(stringAt(+2, "AN", "AC", "IA", "UM") &&
						!(stringAt(-3, "POIGNANT") ||
							stringAt(-2, "COGNAC"))) ||
					stringAtStart("SPIGNER", "STEGNER") ||
					(stringAtStart("SIGNE") && (m_length == 5)) ||
					stringAt(-2, "LIGNI", "LIGNO", "REGNA", "DIGNI", "WEGNE",
						"TIGNE", "RIGNE", "REGNE", "TIGNO") ||
					stringAt(-2, "SIGNAL", "SIGNIF", "SIGNAT") ||
					stringAt(-1, "IGNIT")) &&
				!stringAt(-2, "SIGNET", "LIGNEO"))) ||
			//not e.g. 'cagney', 'magna'
			(((m_current + 2) == m_last) &&
				stringAt(0, "GNE", "GNA") &&
				!stringAt(-2, "SIGNA", "MAGNA", "SIGNE")) {
			metaphAddExactApprox("N", "GN", "N", "KN")
		} else {
			metaphAddExactApprox("GN", "KN")
		}
		m_current += 2
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
func encodeGL() bool {
	//'tagliaro', 'puglia' BUT add K in alternative
	// since americans sometimes do this
	if stringAt(+1, "LIA", "LIO", "LIE") &&
		isVowel(m_current-1) {
		metaphAddExactApprox("L", "GL", "L", "KL")
		m_current += 2
		return true
	}

	return false
}

/**
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func initialGSoft() bool {
	if ((stringAt(+1, "EL", "EM", "EN", "EO", "ER", "ES", "IA", "IN", "IO", "IP", "IU", "YM", "YN", "YP", "YR", "EE") ||
		stringAt(+1, "IRA", "IRO")) &&
		// except for smaller set of cases where => K, e.g. "gerber"
		!(stringAt(+1, "ELD", "ELT", "ERT", "INZ", "ERH", "ITE", "ERD", "ERL", "ERN",
			"INT", "EES", "EEK", "ELB", "EER") ||
			stringAt(+1, "ERSH", "ERST", "INSB", "INGR", "EROW", "ERKE", "EREN") ||
			stringAt(+1, "ELLER", "ERDIE", "ERBER", "ESUND", "ESNER", "INGKO", "INKGO",
				"IPPER", "ESELL", "IPSON", "EEZER", "ERSON", "ELMAN") ||
			stringAt(+1, "ESTALT", "ESTAPO", "INGHAM", "ERRITY", "ERRISH", "ESSNER", "ENGLER") ||
			stringAt(+1, "YNAECOL", "YNECOLO", "ENTHNER", "ERAGHTY") ||
			stringAt(+1, "INGERICH", "EOGHEGAN"))) ||
		(isVowel(m_current+1) &&
			(stringAt(+1, "EE ", "EEW") ||
				(stringAt(+1, "IGI", "IRA", "IBE", "AOL", "IDE", "IGL") &&
					!stringAt(+1, "IDEON")) ||
				stringAt(+1, "ILES", "INGI", "ISEL") ||
				(stringAt(+1, "INGER") && !stringAt(+1, "INGERICH")) ||
				stringAt(+1, "IBBER", "IBBET", "IBLET", "IBRAN", "IGOLO", "IRARD", "IGANT") ||
				stringAt(+1, "IRAFFE", "EEWHIZ") ||
				stringAt(+1, "ILLETTE", "IBRALTA"))) {
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
func encodeInitialGFrontVowel() bool {
	// 'g' followed by vowel at beginning
	if (m_current == 0) && frontVowel(m_current+1) {
		// special case "gila" as in "gila monster"
		if stringAt(+1, "ILA") && (m_length == 4) {
			metaphAdd("H")
		} else if initialGSoft() {
			metaphAddExactApprox("J", "G", "J", "K")
		} else {
			// only code alternate 'J' if front vowel
			if (m_inWord[m_current+1] == 'E') || (m_inWord[m_current+1] == 'I') {
				metaphAddExactApprox("G", "J", "K", "J")
			} else {
				metaphAddExactApprox("G", "K")
			}
		}

		advanceCounter(2, 1)
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
func encodeNGER() bool {
	if (m_current > 1) &&
		stringAt(-1, "NGER") {
		// default 'G' => J  such as 'ranger', 'stranger', 'manger', 'messenger', 'orangery', 'granger'
		// 'boulanger', 'challenger', 'danger', 'changer', 'harbinger', 'lounger', 'ginger', 'passenger'
		// except for these the following
		if !(rootOrInflections(m_inWord, "ANGER") ||
			rootOrInflections(m_inWord, "LINGER") ||
			rootOrInflections(m_inWord, "MALINGER") ||
			rootOrInflections(m_inWord, "FINGER") ||
			(stringAt(-3, "HUNG", "FING", "BUNG", "WING", "RING", "DING", "ZENG",
				"ZING", "JUNG", "LONG", "PING", "CONG", "MONG", "BANG",
				"GANG", "HANG", "LANG", "SANG", "SING", "WANG", "ZANG") &&
				// exceptions to above where 'G' => J
				!(stringAt(-6, "BOULANG", "SLESING", "KISSING", "DERRING") ||
					stringAt(-8, "SCHLESING") ||
					stringAt(-5, "SALING", "BELANG") ||
					stringAt(-6, "BARRING") ||
					stringAt(-6, "PHALANGER") ||
					stringAt(-4, "CHANG"))) ||
			stringAt(-4, "STING", "YOUNG") ||
			stringAt(-5, "STRONG") ||
			stringAtStart("UNG", "ENG", "ING") ||
			stringAt(0, "GERICH") ||
			stringAtStart("SENGER") ||
			stringAt(-3, "WENGER", "MUNGER", "SONGER", "KINGER") ||
			stringAt(-4, "FLINGER", "SLINGER", "STANGER", "STENGER", "KLINGER", "CLINGER") ||
			stringAt(-5, "SPRINGER", "SPRENGER") ||
			stringAt(-3, "LINGERF") ||
			stringAt(-2, "ANGERLY", "ANGERBO", "INGERSO")) {
			metaphAddExactApprox("J", "G", "J", "K")
		} else {
			metaphAddExactApprox("G", "J", "K", "J")
		}

		advanceCounter(2, 1)
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
func encodeGER() bool {
	if (m_current > 0) && stringAt(+1, "ER") {
		// Exceptions to 'GE' where 'G' => K
		// e.g. "JAGER", "TIGER", "LIGER", "LAGER", "LUGER", "AUGER", "EAGER", "HAGER", "SAGER"
		if (((m_current == 2) && isVowel(m_current-1) && !isVowel(m_current-2) &&
			!(stringAt(-2, "PAGER", "WAGER", "NIGER", "ROGER", "LEGER", "CAGER")) ||
			stringAt(-2, "AUGER", "EAGER", "INGER", "YAGER")) ||
			stringAt(-3, "SEEGER", "JAEGER", "GEIGER", "KRUGER", "SAUGER", "BURGER",
				"MEAGER", "MARGER", "RIEGER", "YAEGER", "STEGER", "PRAGER", "SWIGER",
				"YERGER", "TORGER", "FERGER", "HILGER", "ZEIGER", "YARGER",
				"COWGER", "CREGER", "KROGER", "KREGER", "GRAGER", "STIGER", "BERGER") ||
			// 'berger' but not 'bergerac'
			(stringAt(-3, "BERGER") && ((m_current + 2) == m_last)) ||
			stringAt(-4, "KREIGER", "KRUEGER", "METZGER", "KRIEGER", "KROEGER", "STEIGER",
				"DRAEGER", "BUERGER", "BOERGER", "FIBIGER") ||
			// e.g. 'harshbarger', 'winebarger'
			(stringAt(-3, "BARGER") && (m_current > 4)) ||
			// e.g. 'weisgerber'
			(stringAt(0, "GERBER") && (m_current > 0)) ||
			stringAt(-5, "SCHWAGER", "LYBARGER", "SPRENGER", "GALLAGER", "WILLIGER") ||
			stringAtStart("HARGER") ||
			(stringAtStart("AGER", "EGER") && (m_length == 4)) ||
			stringAt(-1, "YGERNE") ||
			stringAt(-6, "SCHWEIGER")) &&
			!(stringAt(-5, "BELLIGEREN") ||
				stringAtStart("MARGERY") ||
				stringAt(-3, "BERGERAC")) {
			if slavoGermanic() {
				metaphAddExactApprox("G", "K")
			} else {
				metaphAddExactApprox("G", "J", "K", "J")
			}
		} else {
			metaphAddExactApprox("J", "G", "J", "K")
		}

		advanceCounter(2, 1)
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
func encodeGEL() bool {
	// more likely to be "-GEL-" => JL
	if stringAt(+1, "EL") && (m_current > 0) {
		// except for
		// "BAGEL", "HEGEL", "HUGEL", "KUGEL", "NAGEL", "VOGEL", "FOGEL", "PAGEL"
		if ((m_length == 5) &&
			isVowel(m_current-1) &&
			!isVowel(m_current-2) &&
			!stringAt(-2, "NIGEL", "RIGEL")) ||
			// or the following as combining forms
			stringAt(-2, "ENGEL", "HEGEL", "NAGEL", "VOGEL") ||
			stringAt(-3, "MANGEL", "WEIGEL", "FLUGEL", "RANGEL", "HAUGEN", "RIEGEL", "VOEGEL") ||
			stringAt(-4, "SPEIGEL", "STEIGEL", "WRANGEL", "SPIEGEL") ||
			stringAt(-4, "DANEGELD") {
			if slavoGermanic() {
				metaphAddExactApprox("G", "K")
			} else {
				metaphAddExactApprox("G", "J", "K", "J")
			}
		} else {
			metaphAddExactApprox("J", "G", "J", "K")
		}

		advanceCounter(2, 1)
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
func hardGEAtEnd() bool {
	return stringAtStart("RENEGE", "STONGE", "STANGE", "PRANGE", "KRESGE") ||
		stringAtStart("BYRGE", "BIRGE", "BERGE", "HAUGE") ||
		stringAtStart("HAGE") ||
		stringAtStart("LANGE", "SYNGE", "BENGE", "RUNGE", "HELGE") ||
		stringAtStart("INGE", "LAGE")
}

/**
 * Detect words where "-ge-" or "-gi-" get a 'hard' 'g'
 * even though this is usually a 'soft' 'g' context
 *
 * @return true if 'hard' 'g' detected
 *
 */
func internalHardGOther() bool {
	return (stringAt(0, "GETH", "GEAR", "GEIS", "GIRL", "GIVI", "GIVE", "GIFT",
		"GIRD", "GIRT", "GILV", "GILD", "GELD") &&
		!stringAt(-3, "GINGIV")) ||
		// "gish" but not "largish"
		(stringAt(+1, "ISH") && (m_current > 0) && !stringAtStart("LARG")) ||
		(stringAt(-2, "MAGED", "MEGID") && !((m_current + 2) == m_last)) ||
		stringAt(0, "GEZ") ||
		stringAtStart("WEGE", "HAGE") ||
		(stringAt(-2, "ONGEST", "UNGEST") &&
			((m_current + 3) == m_last) &&
			!stringAt(-3, "CONGEST")) ||
		stringAtStart("VOEGE", "BERGE", "HELGE") ||
		(stringAtStart("ENGE", "BOGY") && (m_length == 4)) ||
		stringAt(0, "GIBBON") ||
		stringAtStart("CORREGIDOR") ||
		stringAtStart("INGEBORG") ||
		(stringAt(0, "GILL") &&
			(((m_current + 3) == m_last) || ((m_current + 4) == m_last)) &&
			!stringAtStart("STURGILL"))
}

/**
 * Detect words where "-gy-", "-gie-", "-gee-",
 * or "-gio-" get a 'hard' 'g' even though this is
 * usually a 'soft' 'g' context
 *
 * @return true if 'hard' 'g' detected
 *
 */
func internalHardGOpenSyllable() bool {
	return stringAt(+1, "EYE") ||
		stringAt(-2, "FOGY", "POGY", "YOGI") ||
		stringAt(-2, "MAGEE", "MCGEE", "HAGIO") ||
		stringAt(-1, "RGEY", "OGEY") ||
		stringAt(-3, "HOAGY", "STOGY", "PORGY") ||
		stringAt(-5, "CARNEGIE") ||
		(stringAt(-1, "OGEY", "OGIE") && ((m_current + 2) == m_last))
}

/**
 * Detect a number of contexts, mostly german names, that
 * take a 'hard' 'g'.
 *
 * @return true if 'hard' 'g' detected, false if not
 *
 */
func internalHardGENGINGETGIT() bool {
	return (stringAt(-3, "FORGET", "TARGET", "MARGIT", "MARGET", "TURGEN",
		"BERGEN", "MORGEN", "JORGEN", "HAUGEN", "JERGEN",
		"JURGEN", "LINGEN", "BORGEN", "LANGEN", "KLAGEN", "STIGER", "BERGER") &&
		!stringAt(0, "GENETIC", "GENESIS") &&
		!stringAt(-4, "PLANGENT")) ||
		(stringAt(-3, "BERGIN", "FEAGIN", "DURGIN") && ((m_current + 2) == m_last)) ||
		(stringAt(-2, "ENGEN") && !stringAt(+3, "DER", "ETI", "ESI")) ||
		stringAt(-4, "JUERGEN") ||
		stringAtStart("NAGIN", "MAGIN", "HAGIN") ||
		(stringAtStart("ENGIN", "DEGEN", "LAGEN", "MAGEN", "NAGIN") && (m_length == 5)) ||
		(stringAt(-2, "BEGET", "BEGIN", "HAGEN", "FAGIN",
			"BOGEN", "WIGIN", "NTGEN", "EIGEN",
			"WEGEN", "WAGEN") &&
			!stringAt(-5, "OSPHAGEN"))
}

/**
 * Detect a number of contexts of '-ng-' that will
 * take a 'hard' 'g' despite being followed by a
 * front vowel.
 *
 * @return true if 'hard' 'g' detected, false if not
 *
 */
func internalHardNG() bool {
	return (stringAt(-3, "DANG", "FANG", "SING") &&
		// exception to exception
		!stringAt(-5, "DISINGEN")) ||
		stringAtStart("INGEB", "ENGEB") ||
		(stringAt(-3, "RING", "WING", "HANG", "LONG") &&
			!(stringAt(-4, "CRING", "FRING", "ORANG", "TWING", "CHANG", "PHANG") ||
				stringAt(-5, "SYRING") ||
				stringAt(-3, "RINGENC", "RINGENT", "LONGITU", "LONGEVI") ||
				// e.g. 'longino', 'mastrangelo'
				(stringAt(0, "GELO", "GINO") && ((m_current + 3) == m_last)))) ||
		(stringAt(-1, "NGY") &&
			// exceptions to exception
			!(stringAt(-3, "RANGY", "MANGY", "MINGY") ||
				stringAt(-4, "SPONGY", "STINGY")))
}

/**
 * Exceptions to default encoding to 'J':
 * encode "-G-" to 'G' in "-g<frontvowel>-" words
 * where we are not at "-GE" at the end of the word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func internalHardG() bool {
	// if not "-GE" at end
	return !(((m_current + 1) == m_last) && (charAt(m_current+1) == 'E')) &&
		(internalHardNG() ||
			internalHardGENGINGETGIT() ||
			internalHardGOpenSyllable() ||
			internalHardGOther())
}

/**
 * Encode "-G-" followed by a vowel when non-initial leter.
 * Default for this is a 'J' sound, so check exceptions where
 * it is pronounced 'G'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeNonInitialGFrontVowel() bool {
	// -gy-, gi-, ge-
	if stringAt(+1, "E", "I", "Y") {
		// '-ge' at end
		// almost always 'j 'sound
		if stringAt(0, "GE") && (m_current == (m_last - 1)) {
			if hardGEAtEnd() {
				if slavoGermanic() {
					metaphAddExactApprox("G", "K")
				} else {
					metaphAddExactApprox("G", "J", "K", "J")
				}
			} else {
				metaphAdd("J")
			}
		} else {
			if internalHardG() {
				// don't encode KG or KK if e.g. "mcgill"
				if !((m_current == 2) && stringAtStart("MC")) ||
					((m_current == 3) && stringAtStart("MAC")) {
					if slavoGermanic() {
						metaphAddExactApprox("G", "K")
					} else {
						metaphAddExactApprox("G", "J", "K", "J")
					}
				}
			} else {
				metaphAddExactApprox("J", "G", "J", "K")
			}
		}

		advanceCounter(2, 1)
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
func encodeGAToJ() bool {
	// 'margary', 'margarine'
	if (stringAt(-3, "MARGARY", "MARGARI") &&
		// but not in spanish forms such as "margatita"
		!stringAt(-3, "MARGARIT")) ||
		stringAtStart("GAOL") ||
		stringAt(-2, "ALGAE") {
		metaphAddExactApprox("J", "G", "J", "K")
		advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode "-G-"
 *
 */
func encodeG() {
	if encodeSilentGAtBeginning() ||
		encodeGG() ||
		encodeGK() ||
		encodeGH() ||
		encodeSilentG() ||
		encodeGN() ||
		encodeGL() ||
		encodeInitialGFrontVowel() ||
		encodeNGER() ||
		encodeGER() ||
		encodeGEL() ||
		encodeNonInitialGFrontVowel() ||
		encodeGAToJ() {
		return
	}

	if !stringAt(-1, "C", "K", "G", "Q") {
		metaphAddExactApprox("G", "K")
	}

	m_current++
}

/**
 * Encode cases where initial 'H' is not pronounced (in American)
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeInitialSilentH() bool {
	//'hour', 'herb', 'heir', 'honor'
	if stringAt(+1, "OUR", "ERB", "EIR") ||
		stringAt(+1, "ONOR") ||
		stringAt(+1, "ONOUR", "ONEST") {
		// british pronounce H in this word
		// americans give it 'H' for the name,
		// no 'H' for the plant
		if (m_current == 0) && stringAt(0, "HERB") {
			if m_encodeVowels {
				metaphAdd("HA", "A")
			} else {
				metaphAdd("H", "A")
			}
		} else if (m_current == 0) || m_encodeVowels {
			metaphAdd("A")
		}

		m_current++
		// don't encode vowels twice
		m_current = skipVowels(m_current)
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
func encodeInitialHS() bool {
	// old chinese pinyin transliteration
	// e.g., 'HSIAO'
	if (m_current == 0) && stringAtStart("HS") {
		metaphAdd("X")
		m_current += 2
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
func encodeInitialHUHW() bool {
	// spanish spellings and chinese pinyin transliteration
	if stringAtStart("HUA", "HUE", "HWA") {
		if !stringAt(0, "HUEY") {
			metaphAdd("A")

			if !m_encodeVowels {
				m_current += 3
			} else {
				m_current++
				// don't encode vowels twice
				for isVowel(m_current) || (charAt(m_current) == 'W') {
					m_current++
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
func encodeNonInitialSilentH() bool {
	//exceptions - 'h' not pronounced
	// "PROHIB" BUT NOT "PROHIBIT"
	if stringAt(-2, "NIHIL", "VEHEM", "LOHEN", "NEHEM",
		"MAHON", "MAHAN", "COHEN", "GAHAN") ||
		stringAt(-3, "GRAHAM", "PROHIB", "FRAHER",
			"TOOHEY", "TOUHEY") ||
		stringAt(-3, "TOUHY") ||
		stringAtStart("CHIHUAHUA") {
		if !m_encodeVowels {
			m_current += 2
		} else {
			m_current++
			// don't encode vowels twice
			m_current = skipVowels(m_current)
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
func encodeHPronounced() bool {
	if (((m_current == 0) ||
		isVowel(m_current-1) ||
		((m_current > 0) &&
			(charAt(m_current-1) == 'W'))) &&
		isVowel(m_current+1)) ||
		// e.g. 'alWahhab'
		((charAt(m_current+1) == 'H') && isVowel(m_current+2)) {
		metaphAdd("H")
		advanceCounter(2, 1)
		return true
	}

	return false
}

/**
 * Encode 'H'
 *
 *
 */
func encodeH() {
	if encodeInitialSilentH() ||
		encodeInitialHS() ||
		encodeInitialHUHW() ||
		encodeNonInitialSilentH() {
		return
	}

	//only keep if first & before vowel or btw. 2 vowels
	if !encodeHPronounced() {
		//also takes care of 'HH'
		m_current++
	}
}

/**
 * Encode cases where initial or medial "j" is in a spanish word or name
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSpanishJ() bool {
	//obvious spanish, e.g. "jose", "san jacinto"
	if (stringAt(+1, "UAN", "ACI", "ALI", "EFE", "ICA", "IME", "OAQ", "UAR") &&
		!stringAt(0, "JIMERSON", "JIMERSEN")) ||
		(stringAt(+1, "OSE") && ((m_current + 3) == m_last)) ||
		stringAt(+1, "EREZ", "UNTA", "AIME", "AVIE", "AVIA") ||
		stringAt(+1, "IMINEZ", "ARAMIL") ||
		(((m_current + 2) == m_last) && stringAt(-2, "MEJIA")) ||
		stringAt(-2, "TEJED", "TEJAD", "LUJAN", "FAJAR", "BEJAR", "BOJOR", "CAJIG",
			"DEJAS", "DUJAR", "DUJAN", "MIJAR", "MEJOR", "NAJAR",
			"NOJOS", "RAJED", "RIJAL", "REJON", "TEJAN", "UIJAN") ||
		stringAt(-3, "ALEJANDR", "GUAJARDO", "TRUJILLO") ||
		(stringAt(-2, "RAJAS") && (m_current > 2)) ||
		(stringAt(-2, "MEJIA") && !stringAt(-2, "MEJIAN")) ||
		stringAt(-1, "OJEDA") ||
		stringAt(-3, "LEIJA", "MINJA") ||
		stringAt(-3, "VIAJES", "GRAJAL") ||
		stringAt(0, "JAUREGUI") ||
		stringAt(-4, "HINOJOSA") ||
		stringAtStart("SAN ") ||
		(((m_current + 1) == m_last) &&
			(charAt(m_current+1) == 'O') &&
			// exceptions
			!(stringAtStart("TOJO") ||
				stringAtStart("BANJO") ||
				stringAtStart("MARYJO"))) {
		// americans pronounce "juan" as 'wan'
		// and "marijuana" and "tijuana" also
		// do not get the 'H' as in spanish, so
		// just treat it like a vowel in these cases
		if !(stringAt(0, "JUAN") || stringAt(0, "JOAQ")) {
			metaphAdd("H")
		} else {
			if m_current == 0 {
				metaphAdd("A")
			}
		}
		advanceCounter(2, 1)
		return true
	}

	// Jorge gets 2nd HARHA. also JULIO, JESUS
	if stringAt(+1, "ORGE", "ULIO", "ESUS") &&
		!stringAtStart("JORGEN") {
		// get both consonants for "jorge"
		if ((m_current + 4) == m_last) && stringAt(+1, "ORGE") {
			if m_encodeVowels {
				metaphAdd("JARJ", "HARHA")
			} else {
				metaphAdd("JRJ", "HRH")
			}
			advanceCounter(5, 5)
			return true
		}

		metaphAdd("J", "H")
		advanceCounter(2, 1)
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
func encodeGermanJ() bool {
	if stringAt(+1, "AH") ||
		(stringAt(+1, "OHANN") && ((m_current + 5) == m_last)) ||
		(stringAt(+1, "UNG") && !stringAt(+1, "UNGL")) ||
		stringAt(+1, "UGO") {
		metaphAdd("A")
		advanceCounter(2, 1)
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
func encodeSpanishOJUJ() bool {
	if stringAt(+1, "OJOBA", "UJUY ") {
		if m_encodeVowels {
			metaphAdd("HAH")
		} else {
			metaphAdd("HH")
		}

		advanceCounter(4, 3)
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
func encodeSpanishJ2() bool {
	// spanish forms e.g. "brujo", "badajoz"
	if (((m_current - 2) == 0) &&
		stringAt(-2, "BOJA", "BAJA", "BEJA", "BOJO", "MOJA", "MOJI", "MEJI")) ||
		(((m_current - 3) == 0) &&
			stringAt(-3, "FRIJO", "BRUJO", "BRUJA", "GRAJE", "GRIJA", "LEIJA", "QUIJA")) ||
		(((m_current + 3) == m_last) &&
			stringAt(-1, "AJARA")) ||
		(((m_current + 2) == m_last) &&
			stringAt(-1, "AJOS", "EJOS", "OJAS", "OJOS", "UJON", "AJOZ", "AJAL", "UJAR", "EJON", "EJAN")) ||
		(((m_current + 1) == m_last) &&
			(stringAt(-1, "OJA", "EJA") && !stringAtStart("DEJA"))) {
		metaphAdd("H")
		advanceCounter(2, 1)
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
func encodeJAsVowel() bool {
	if stringAt(0, "JEWSK") {
		metaphAdd("J", "")
		return true
	}

	// e.g. "stijl", "sejm" - dutch, scandanavian, and eastern european spellings
	return (stringAt(+1, "L", "T", "K", "S", "N", "M") &&
		// except words from hindi and arabic
		!stringAt(+2, "A")) ||
		stringAtStart("HALLELUJA", "LJUBLJANA") ||
		stringAtStart("LJUB", "BJOR") ||
		stringAtStart("HAJEK") ||
		stringAtStart("WOJ") ||
		// e.g. 'fjord'
		stringAtStart("FJ") ||
		// e.g. 'rekjavik', 'blagojevic'
		stringAt(0, "JAVIK", "JEVIC") ||
		(((m_current + 1) == m_last) && stringAtStart("SONJA", "TANJA", "TONJA"))
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
func namesBeginningWithJThatGetAltY() bool {
	if stringAtStart("JAN", "JON", "JAN", "JIN", "JEN") ||
		stringAtStart("JUHL", "JULY", "JOEL", "JOHN", "JOSH",
			"JUDE", "JUNE", "JONI", "JULI", "JENA",
			"JUNG", "JINA", "JANA", "JENI", "JOEL",
			"JANN", "JONA", "JENE", "JULE", "JANI",
			"JONG", "JOHN", "JEAN", "JUNG", "JONE",
			"JARA", "JUST", "JOST", "JAHN", "JACO",
			"JANG", "JUDE", "JONE") ||
		stringAtStart("JOANN", "JANEY", "JANAE", "JOANA", "JUTTA",
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
		stringAtStart("JORDAN", "JORDON", "JOSEPH", "JOSHUA", "JOSIAH",
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
		stringAtStart("JAKOB") ||
		stringAtStart("JOHNSON", "JOHNNIE", "JASMINE", "JEANNIE", "JOHANNA",
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
		stringAtStart("JOSEFINA", "JEANNINE", "JULIANNE", "JULIANNA", "JONATHAN",
			"JONATHON", "JEANETTE", "JANNETTE", "JEANETTA", "JOHNETTA",
			"JENNEFER", "JULIENNE", "JOSPHINE", "JEANELLE", "JOHNETTE",
			"JULIEANN", "JOSEFINE", "JULIETTA", "JOHNSTON", "JACOBSON",
			"JACOBSEN", "JOHANSEN", "JOHANSON", "JAWORSKI", "JENNETTE",
			"JELLISON", "JOHANNES", "JASINSKI", "JUERGENS", "JARNAGIN",
			"JEREMIAH", "JEPPESEN", "JARNIGAN", "JANOUSEK") ||
		stringAtStart("JOHNATHAN", "JOHNATHON", "JORGENSEN", "JEANMARIE", "JOSEPHINA",
			"JEANNETTE", "JOSEPHINE", "JEANNETTA", "JORGENSON", "JANKOWSKI",
			"JOHNSTONE", "JABLONSKI", "JOSEPHSON", "JOHANNSEN", "JURGENSEN",
			"JIMMERSON", "JOHANSSON") ||
		stringAtStart("JAKUBOWSKI") {
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
func encodeJToJ() bool {
	if isVowel(m_current + 1) {
		if (m_current == 0) &&
			namesBeginningWithJThatGetAltY() {
			// 'Y' is a vowel so encode
			// is as 'A'
			if m_encodeVowels {
				metaphAdd("JA", "A")
			} else {
				metaphAdd("J", "A")
			}
		} else {
			if m_encodeVowels {
				metaphAdd("JA")
			} else {
				metaphAdd("J")
			}
		}

		m_current++
		m_current = skipVowels(m_current)
		return false
	} else {
		metaphAdd("J")
		m_current++
		return true
	}

	//		return false
}

/**
 * Call routines to encode 'J', in proper order
 *
 */
func encodeOtherJ() {
	if m_current == 0 {
		if encodeGermanJ() {
			return
		} else {
			if encodeJToJ() {
				return
			}
		}
	} else {
		if encodeSpanishJ2() {
			return
		} else if !encodeJAsVowel() {
			metaphAdd("J")
		}

		//it could happen! e.g. "hajj"
		// eat redundant 'J'
		if charAt(m_current+1) == 'J' {
			m_current += 2
		} else {
			m_current++
		}
	}
}

/**
 * Encode 'J'
 *
 */
func encodeJ() {
	if encodeSpanishJ() || encodeSpanishOJUJ() {
		return
	}

	encodeOtherJ()
}

/**
 * Encode cases where 'K' is not pronounced
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSilentK() bool {
	//skip this except for special cases
	if (m_current == 0) && stringAt(0, "KN") {
		if !(stringAt(+2, "ESSET", "IEVEL") || stringAt(+2, "ISH")) {
			m_current += 1
			return true
		}
	}

	// e.g. "know", "knit", "knob"
	if (stringAt(+1, "NOW", "NIT", "NOT", "NOB") &&
		// exception, "slipknot" => SLPNT but "banknote" => PNKNT
		!stringAtStart("BANKNOTE")) ||
		stringAt(+1, "NOCK", "NUCK", "NIFE", "NACK") ||
		stringAt(+1, "NIGHT") {
		// N already encoded before
		// e.g. "penknife"
		if (m_current > 0) && charAt(m_current-1) == 'N' {
			m_current += 2
		} else {
			m_current++
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
func encodeK() {
	if !encodeSilentK() {
		metaphAdd("K")

		// eat redundant 'K's and 'Q's
		if (charAt(m_current+1) == 'K') ||
			(charAt(m_current+1) == 'Q') {
			m_current += 2
		} else {
			m_current++
		}
	}
}

/**
 * Cases where an L follows D, G, or T at the
 * end have a schwa pronounced before the L
 *
 */
func interpolateVowelWhenConsLAtEnd() {
	if m_encodeVowels == true {
		// e.g. "ertl", "vogl"
		if (m_current == m_last) &&
			stringAt(-1, "D", "G", "T") {
			metaphAdd("A")
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
func encodeLELYToL() bool {
	// e.g. "agilely", "docilely"
	if stringAt(-1, "ILELY") &&
		((m_current + 3) == m_last) {
		metaphAdd("L")
		m_current += 3
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
func encodeCOLONEL() bool {
	if stringAt(-2, "COLONEL") {
		metaphAdd("R")
		m_current += 2
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
func encodeFrenchAULT() bool {
	// e.g. "renault" and "foucault", well known to americans, but not "fault"
	if (m_current > 3) &&
		(stringAt(-3, "RAULT", "NAULT", "BAULT", "SAULT", "GAULT", "CAULT") ||
			stringAt(-4, "REAULT", "RIAULT", "NEAULT", "BEAULT")) &&
		!(rootOrInflections(m_inWord, "ASSAULT") ||
			stringAt(-8, "SOMERSAULT") ||
			stringAt(-9, "SUMMERSAULT")) {
		m_current += 2
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
func encodeFrenchEUIL() bool {
	// e.g. "auteuil"
	if stringAt(-3, "EUIL") && (m_current == m_last) {
		m_current++
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
func encodeFrenchOULX() bool {
	// e.g. "proulx"
	if stringAt(-2, "OULX") && ((m_current + 1) == m_last) {
		m_current += 2
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
func encodeSilentLInLM() bool {
	if stringAt(0, "LM", "LN") {
		// e.g. "lincoln", "holmes", "psalm", "salmon"
		if (stringAt(-2, "COLN", "CALM", "BALM", "MALM", "PALM") ||
			(stringAt(-1, "OLM") && ((m_current + 1) == m_last)) ||
			stringAt(-3, "PSALM", "QUALM") ||
			stringAt(-2, "SALMON", "HOLMES") ||
			stringAt(-1, "ALMOND") ||
			((m_current == 1) && stringAt(-1, "ALMS"))) &&
			(!stringAt(+2, "A") &&
				!stringAt(-2, "BALMO") &&
				!stringAt(-2, "PALMER", "PALMOR", "BALMER") &&
				!stringAt(-3, "THALM")) {
			m_current++
			return true
		} else {
			metaphAdd("L")
			m_current++
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
func encodeSilentLInLKLV() bool {
	if (stringAt(-2, "WALK", "YOLK", "FOLK", "HALF", "TALK", "CALF", "BALK", "CALK") ||
		(stringAt(-2, "POLK") &&
			!stringAt(-2, "POLKA", "WALKO")) ||
		(stringAt(-2, "HALV") &&
			!stringAt(-2, "HALVA", "HALVO")) ||
		(stringAt(-3, "CAULK", "CHALK", "BAULK", "FAULK") &&
			!stringAt(-4, "SCHALK")) ||
		(stringAt(-2, "SALVE", "CALVE") ||
			stringAt(-2, "SOLDER")) &&
			// exceptions to above cases where 'L' is usually pronounced
			!stringAt(-2, "SALVER", "CALVER")) &&
		!stringAt(-5, "GONSALVES", "GONCALVES") &&
		!stringAt(-2, "BALKAN", "TALKAL") &&
		!stringAt(-3, "PAULK", "CHALF") {
		m_current++
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
func encodeSilentLInOULD() bool {
	//'would', 'could'
	if stringAt(-3, "WOULD", "COULD") ||
		(stringAt(-4, "SHOULD") &&
			!stringAt(-4, "SHOULDER")) {
		metaphAddExactApprox("D", "T")
		m_current += 2
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
func encodeLLAsVowelSpecialCases() bool {
	if stringAt(-5, "TORTILLA") ||
		stringAt(-8, "RATATOUILLE") ||
		// e.g. 'guillermo', "veillard"
		(stringAtStart("GUILL", "VEILL", "GAILL") &&
			// 'guillotine' usually has '-ll-' pronounced as 'L' in english
			!(stringAt(-3, "GUILLOT", "GUILLOR", "GUILLEN") ||
				(stringAtStart("GUILL") && (m_length == 5)))) ||
		// e.g. "brouillard", "gremillion"
		stringAtStart("BROUILL", "GREMILL") ||
		stringAtStart("ROBILL") ||
		// e.g. 'mireille'
		(stringAt(-2, "EILLE") &&
			((m_current + 2) == m_last) &&
			// exception "reveille" usually pronounced as 're-vil-lee'
			!stringAt(-5, "REVEILLE")) {
		m_current += 2
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
func encodeLLAsVowel() bool {
	//spanish e.g. "cabrillo", "gallegos" but also "gorilla", "ballerina" -
	// give both pronounciations since an american might pronounce "cabrillo"
	// in the spanish or the american fashion.
	if (((m_current + 3) == m_length) &&
		stringAt(-1, "ILLO", "ILLA", "ALLE")) ||
		(((stringAtEnd("AS", "OS") ||
			stringAtEnd("A", "O")) &&
			stringAt(-1, "AL", "IL")) &&
			!stringAt(-1, "ALLA")) ||
		stringAtStart("VILLE", "VILLA") ||
		stringAtStart("GALLARDO", "VALLADAR", "MAGALLAN", "CAVALLAR", "BALLASTE") ||
		stringAtStart("LLA") {
		metaphAdd("L", "")
		m_current += 2
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
func encodeLLAsVowelCases() bool {
	if charAt(m_current+1) == 'L' {
		if encodeLLAsVowelSpecialCases() {
			return true
		} else if encodeLLAsVowel() {
			return true
		}
		m_current += 2

	} else {
		m_current++
	}

	return false
}

/**
 * Encode vowel-encoding cases where "-LE-" is pronounced "-EL-"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeVowelLETransposition(save_current int) bool {
	// transposition of vowel sound and L occurs in many words,
	// e.g. "bristle", "dazzle", "goggle" => KAKAL
	if m_encodeVowels && (save_current > 1) &&
		!isVowel(save_current-1) &&
		(charAt(save_current+1) == 'E') &&
		(charAt(save_current-1) != 'L') &&
		(charAt(save_current-1) != 'R') &&
		// lots of exceptions to this:
		!isVowel(save_current+2) &&
		!stringAtStart("ECCLESI", "COMPLEC", "COMPLEJ", "ROBLEDO") &&
		!stringAtStart("MCCLE", "MCLEL") &&
		!stringAtStart("EMBLEM", "KADLEC") &&
		!(((save_current + 2) == m_last) && stringAt(save_current, "LET")) &&
		!stringAt(save_current, "LETTING") &&
		!stringAt(save_current, "LETELY", "LETTER", "LETION", "LETIAN", "LETING", "LETORY") &&
		!stringAt(save_current, "LETUS", "LETIV") &&
		!stringAt(save_current, "LESS", "LESQ", "LECT", "LEDG", "LETE", "LETH", "LETS", "LETT") &&
		!stringAt(save_current, "LEG", "LER", "LEX") &&
		// e.g. "complement" !=> KAMPALMENT
		!(stringAt(save_current, "LEMENT") &&
			!(stringAt(-5, "BATTLE", "TANGLE", "PUZZLE", "RABBLE", "BABBLE") ||
				stringAt(-4, "TABLE"))) &&
		!(((save_current + 2) == m_last) && stringAt((save_current-2), "OCLES", "ACLES", "AKLES")) &&
		!stringAt((save_current-3), "LISLE", "AISLE") &&
		!stringAtStart("ISLE") &&
		!stringAtStart("ROBLES") &&
		!stringAt((save_current-4), "PROBLEM", "RESPLEN") &&
		!stringAt((save_current-3), "REPLEN") &&
		!stringAt((save_current-2), "SPLE") &&
		(charAt(save_current-1) != 'H') &&
		(charAt(save_current-1) != 'W') {
		metaphAdd("AL")
		flag_AL_inversion = true

		// eat redundant 'L'
		if charAt(save_current+2) == 'L' {
			m_current = save_current + 3
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
func encodeVowelPreserveVowelAfterL(save_current int) bool {
	// an example of where the vowel would NOT need to be preserved
	// would be, say, "hustled", where there is no vowel pronounced
	// between the 'l' and the 'd'
	if m_encodeVowels &&
		!isVowel(save_current-1) &&
		(charAt(save_current+1) == 'E') &&
		(save_current > 1) &&
		((save_current + 1) != m_last) &&
		!(stringAt((save_current+1), "ES", "ED") &&
			((save_current + 2) == m_last)) &&
		!stringAt((save_current-1), "RLEST") {
		metaphAdd("LA")
		m_current = skipVowels(m_current)
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
func encodeLECases(save_current int) {
	if encodeVowelLETransposition(save_current) {
		return
	} else {
		if encodeVowelPreserveVowelAfterL(save_current) {
			return
		} else {
			metaphAdd("L")
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
func encodeL() {
	// logic below needs to know this
	// after 'm_current' variable changed
	var save_current = m_current

	interpolateVowelWhenConsLAtEnd()

	if encodeLELYToL() ||
		encodeCOLONEL() ||
		encodeFrenchAULT() ||
		encodeFrenchEUIL() ||
		encodeFrenchOULX() ||
		encodeSilentLInLM() ||
		encodeSilentLInLKLV() ||
		encodeSilentLInOULD() {
		return
	}

	if encodeLLAsVowelCases() {
		return
	}

	encodeLECases(save_current)
}

/**
 * Encode cases where 'M' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSilentMAtBeginning() bool {
	//skip these when at start of word
	if (m_current == 0) && stringAt(0, "MN") {
		m_current += 1
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
func encodeMRAndMRS() bool {
	if (m_current == 0) && stringAt(0, "MR") {
		// exceptions for "mr." and "mrs."
		if (m_length == 2) && stringAt(0, "MR") {
			if m_encodeVowels {
				metaphAdd("MASTAR")
			} else {
				metaphAdd("MSTR")
			}
			m_current += 2
			return true
		} else if (m_length == 3) && stringAt(0, "MRS") {
			if m_encodeVowels {
				metaphAdd("MASAS")
			} else {
				metaphAdd("MSS")
			}
			m_current += 3
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
func encodeMAC() bool {
	// should only find irish and
	// scottish names e.g. 'macintosh'
	if (m_current == 0) &&
		(stringAtStart("MACIVER", "MACEWEN") ||
			stringAtStart("MACELROY", "MACILROY") ||
			stringAtStart("MACINTOSH") ||
			stringAtStart("MC")) {
		if m_encodeVowels {
			metaphAdd("MAK")
		} else {
			metaphAdd("MK")
		}

		if stringAtStart("MC") {
			if stringAt(+2, "K", "G", "Q") &&
				// watch out for e.g. "McGeorge"
				!stringAt(+2, "GEOR") {
				m_current += 3
			} else {
				m_current += 2
			}
		} else {
			m_current += 3
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
func encodeMPT() bool {
	if stringAt(-2, "COMPTROL") ||
		stringAt(-4, "ACCOMPT") {
		metaphAdd("N")
		m_current += 2
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
func testSilentMB1() bool {
	// e.g. "LAMB", "COMB", "LIMB", "DUMB", "BOMB"
	// Handle combining roots first
	if ((m_current == 3) &&
		stringAt(-3, "THUMB")) ||
		((m_current == 2) &&
			stringAt(-2, "DUMB", "BOMB", "DAMN", "LAMB", "NUMB", "TOMB")) {
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
func testPronouncedMB() bool {
	if stringAt(-2, "NUMBER") ||
		(stringAt(+2, "A") &&
			!stringAt(-2, "DUMBASS")) ||
		stringAt(+2, "O") ||
		stringAt(-2, "LAMBEN", "LAMBER", "LAMBET", "TOMBIG", "LAMBRE") {
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
func testSilentMB2() bool {
	// 'M' is the current letter
	if (charAt(m_current+1) == 'B') && (m_current > 1) &&
		(((m_current + 1) == m_last) ||
			// other situations where "-MB-" is at end of root
			// but not at end of word. The tests are for standard
			// noun suffixes.
			// e.g. "climbing" => KLMNK
			stringAt(+2, "ING", "ABL") ||
			stringAt(+2, "LIKE") ||
			((charAt(m_current+2) == 'S') && ((m_current + 2) == m_last)) ||
			stringAt(-5, "BUNCOMB") ||
			// e.g. "bomber",
			(stringAt(+2, "ED", "ER") &&
				((m_current + 3) == m_last) &&
				(stringAtStart("CLIMB", "PLUMB") ||
					// e.g. "beachcomber"
					!stringAt(-1, "IMBER", "AMBER", "EMBER", "UMBER")) &&
				// exceptions
				!stringAt(-2, "CUMBER", "SOMBER"))) {
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
func testPronouncedMB2() bool {
	// e.g. "bombastic", "umbrage", "flamboyant"
	if stringAt(-1, "OMBAS", "OMBAD", "UMBRA") ||
		stringAt(-3, "FLAM") {
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
func testMN() bool {
	return (charAt(m_current+1) == 'N') &&
		(((m_current + 1) == m_last) ||
			// or at the end of a word but followed by suffixes
			(stringAt(+2, "ING", "EST") && ((m_current + 4) == m_last)) ||
			((charAt(m_current+2) == 'S') && ((m_current + 2) == m_last)) ||
			(stringAt(+2, "LY", "ER", "ED") &&
				((m_current + 3) == m_last)) ||
			stringAt(-2, "DAMNEDEST") ||
			stringAt(-5, "GODDAMNIT"))
}

/**
 * Call routines to encode "-MB-", in proper order
 *
 */
func encodeMB() {
	if testSilentMB1() {
		if testPronouncedMB() {
			m_current++
		} else {
			m_current += 2
		}
	} else if testSilentMB2() {
		if testPronouncedMB2() {
			m_current++
		} else {
			m_current += 2
		}
	} else if testMN() {
		m_current += 2
	} else {
		// eat redundant 'M'
		if charAt(m_current+1) == 'M' {
			m_current += 2
		} else {
			m_current++
		}
	}
}

/**
 * Encode "-M-"
 *
 */
func encodeM() {
	if encodeSilentMAtBeginning() ||
		encodeMRAndMRS() ||
		encodeMAC() ||
		encodeMPT() {
		return
	}

	// Silent 'B' should really be handled
	// under 'B", not here under 'M'!
	encodeMB()

	metaphAdd("M")
}

/**
 * Encode "-NCE-" and "-NSE-"
 * "entrance" is pronounced exactly the same as "entrants"
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeNCE() bool {
	//'acceptance', 'accountancy'
	if stringAt(+1, "C", "S") &&
		stringAt(+2, "E", "Y", "I") &&
		(((m_current + 2) == m_last) ||
			((m_current+3) == m_last) &&
				(charAt(m_current+3) == 'S')) {
		metaphAdd("NTS")
		m_current += 2
		return true
	}

	return false
}

/**
 * Encode "-N-"
 *
 */
func encodeN() {
	if encodeNCE() {
		return
	}

	// eat redundant 'N'
	if charAt(m_current+1) == 'N' {
		m_current += 2
	} else {
		m_current++
	}

	if !stringAt(-3, "MONSIEUR") &&
		// e.g. "aloneness",
		!stringAt(-3, "NENESS") {
		metaphAdd("N")
	}
}

/**
 * Encode cases where "-P-" is silent at the start of a word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSilentPAtBeginning() bool {
	//skip these when at start of word
	if (m_current == 0) &&
		stringAt(0, "PN", "PF", "PS", "PT") {
		m_current += 1
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
func encodePT() bool {
	// 'pterodactyl', 'receipt', 'asymptote'
	if charAt(m_current+1) == 'T' {
		if ((m_current == 0) && stringAt(0, "PTERO")) ||
			stringAt(-5, "RECEIPT") ||
			stringAt(-4, "ASYMPTOT") {
			metaphAdd("T")
			m_current += 2
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
func encodePH() bool {
	if charAt(m_current+1) == 'H' {
		// 'PH' silent in these contexts
		if stringAt(0, "PHTHALEIN") ||
			((m_current == 0) && stringAt(0, "PHTH")) ||
			stringAt(-3, "APOPHTHEGM") {
			metaphAdd("0")
			m_current += 4
			// combining forms
			//'sheepherd', 'upheaval', 'cupholder'
		} else if (m_current > 0) &&
			(stringAt(+2, "EAD", "OLE", "ELD", "ILL", "OLD", "EAP", "ERD",
				"ARD", "ANG", "ORN", "EAV", "ART") ||
				stringAt(+2, "OUSE") ||
				(stringAt(+2, "AM") && !stringAt(-1, "LPHAM")) ||
				stringAt(+2, "AMMER", "AZARD", "UGGER") ||
				stringAt(+2, "OLSTER")) &&
			!stringAt(-3, "LYMPH", "NYMPH") {
			metaphAdd("P")
			advanceCounter(3, 2)
		} else {
			metaphAdd("F")
			m_current += 2
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
func encodePPH() bool {
	// 'sappho'
	if (charAt(m_current+1) == 'P') &&
		((m_current + 2) < m_length) && (charAt(m_current+2) == 'H') {
		metaphAdd("F")
		m_current += 3
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
func encodeRPS() bool {
	//'-corps-', 'corpsman'
	if stringAt(-3, "CORPS") &&
		!stringAt(-3, "CORPSE") {
		m_current += 2
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
func encodeCOUP() bool {
	//'coup'
	if (m_current == m_last) &&
		stringAt(-3, "COUP") &&
		!stringAt(-5, "RECOUP") {
		m_current++
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
func encodePNEUM() bool {
	//'-pneum-'
	if stringAt(+1, "NEUM") {
		metaphAdd("N")
		m_current += 2
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
func encodePSYCH() bool {
	//'-psych-'
	if stringAt(+1, "SYCH") {
		if m_encodeVowels {
			metaphAdd("SAK")
		} else {
			metaphAdd("SK")
		}

		m_current += 5
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
func encodePSALM() bool {
	//'-psalm-'
	if stringAt(+1, "SALM") {
		// go ahead and encode entire word
		if m_encodeVowels {
			metaphAdd("SAM")
		} else {
			metaphAdd("SM")
		}

		m_current += 5
		return true
	}

	return false
}

/**
 * Eat redundant 'B' or 'P'
 *
 */
func encodePB() {
	// e.g. "campbell", "raspberry"
	// eat redundant 'P' or 'B'
	if stringAt(+1, "P", "B") {
		m_current += 2
	} else {
		m_current++
	}
}

/**
 * Encode "-P-"
 *
 */
func encodeP() {
	if encodeSilentPAtBeginning() ||
		encodePT() ||
		encodePH() ||
		encodePPH() ||
		encodeRPS() ||
		encodeCOUP() ||
		encodePNEUM() ||
		encodePSYCH() ||
		encodePSALM() {
		return
	}

	encodePB()

	metaphAdd("P")
}

/**
 * Encode "-Q-"
 *
 */
func encodeQ() {
	// current pinyin
	if stringAt(0, "QIN") {
		metaphAdd("X")
		m_current++
		return
	}

	// eat redundant 'Q'
	if charAt(m_current+1) == 'Q' {
		m_current += 2
	} else {
		m_current++
	}

	metaphAdd("K")
}

/**
 * Encode "-RZ-" according
 * to american and polish pronunciations
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeRZ() bool {
	if stringAt(-2, "GARZ", "KURZ", "MARZ", "MERZ", "HERZ", "PERZ", "WARZ") ||
		stringAt(0, "RZANO", "RZOLA") ||
		stringAt(-1, "ARZA", "ARZN") {
		return false
	}

	// 'yastrzemski' usually has 'z' silent in
	// united states, but should get 'X' in poland
	if stringAt(-4, "YASTRZEMSKI") {
		metaphAdd("R", "X")
		m_current += 2
		return true
	}
	// 'BRZEZINSKI' gets two pronunciations
	// in the united states, neither of which
	// are authentically polish
	if stringAt(-1, "BRZEZINSKI") {
		metaphAdd("RS", "RJ")
		// skip over 2nd 'Z'
		m_current += 4
		return true
		// 'z' in 'rz after voiceless consonant gets 'X'
		// in alternate polish style pronunciation
	} else if stringAt(-1, "TRZ", "PRZ", "KRZ") ||
		(stringAt(0, "RZ") &&
			(isVowel(m_current-1) || (m_current == 0))) {
		metaphAdd("RS", "X")
		m_current += 2
		return true
		// 'z' in 'rz after voiceled consonant, vowel, or at
		// beginning gets 'J' in alternate polish style pronunciation
	} else if stringAt(-1, "BRZ", "DRZ", "GRZ") {
		metaphAdd("RS", "J")
		m_current += 2
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
func testSilentR() bool {
	// test cases where 'R' is silent, either because the
	// word is from the french or because it is no longer pronounced.
	// e.g. "rogier", "monsieur", "surburban"
	return ((m_current == m_last) &&
		// reliably french word ending
		stringAt(-2, "IER") &&
		// e.g. "metier"
		(stringAt(-5, "MET", "VIV", "LUC") ||
			// e.g. "cartier", "bustier"
			stringAt(-6, "CART", "DOSS", "FOUR", "OLIV", "BUST", "DAUM", "ATEL",
				"SONN", "CORM", "MERC", "PELT", "POIR", "BERN", "FORT", "GREN",
				"SAUC", "GAGN", "GAUT", "GRAN", "FORC", "MESS", "LUSS", "MEUN",
				"POTH", "HOLL", "CHEN") ||
			// e.g. "croupier"
			stringAt(-7, "CROUP", "TORCH", "CLOUT", "FOURN", "GAUTH", "TROTT",
				"DEROS", "CHART") ||
			// e.g. "chevalier"
			stringAt(-8, "CHEVAL", "LAVOIS", "PELLET", "SOMMEL", "TREPAN", "LETELL", "COLOMB") ||
			stringAt(-9, "CHARCUT") ||
			stringAt(-10, "CHARPENT"))) ||
		stringAt(-2, "SURBURB", "WORSTED") ||
		stringAt(-2, "WORCESTER") ||
		stringAt(-7, "MONSIEUR") ||
		stringAt(-6, "POITIERS")
}

/**
 * Encode '-re-" as 'AR' in contexts
 * where this is the correct pronunciation
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeVowelRETransposition() bool {
	// -re inversion is just like
	// -le inversion
	// e.g. "fibre" => FABAR or "centre" => SANTAR
	if (m_encodeVowels) &&
		(charAt(m_current+1) == 'E') &&
		(m_length > 3) &&
		!stringAtStart("OUTRE", "LIBRE", "ANDRE") &&
		!(stringAtStart("FRED", "TRES") && (m_length == 4)) &&
		!stringAt(-2, "LDRED", "LFRED", "NDRED", "NFRED", "NDRES", "TRES", "IFRED") &&
		!isVowel(m_current-1) &&
		(((m_current + 1) == m_last) ||
			(((m_current + 2) == m_last) &&
				stringAt(+2, "D", "S"))) {
		metaphAdd("AR")
		return true
	}

	return false
}

/**
 * Encode "-R-"
 *
 */
func encodeR() {
	if encodeRZ() {
		return
	}

	if !testSilentR() {
		if !encodeVowelRETransposition() {
			metaphAdd("R")
		}
	}

	// eat redundant 'R'; also skip 'S' as well as 'R' in "poitiers"
	if (charAt(m_current+1) == 'R') || stringAt(-6, "POITIERS") {
		m_current += 2
	} else {
		m_current++
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
func namesBeginningWithSWThatGetAltSV() bool {
	if stringAtStart("SWANSON", "SWENSON", "SWINSON", "SWENSEN",
		"SWOBODA") ||
		stringAtStart("SWIDERSKI", "SWARTHOUT") ||
		stringAtStart("SWEARENGIN") {
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
func namesBeginningWithSWThatGetAltXV() bool {
	if stringAtStart("SWART") ||
		stringAtStart("SWARTZ", "SWARTS", "SWIGER") ||
		stringAtStart("SWITZER", "SWANGER", "SWIGERT",
			"SWIGART", "SWIHART") ||
		stringAtStart("SWEITZER", "SWATZELL", "SWINDLER") ||
		stringAtStart("SWINEHART") ||
		stringAtStart("SWEARINGEN") {
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
func encodeSpecialSW() bool {
	if m_current == 0 {
		//
		if namesBeginningWithSWThatGetAltSV() {
			metaphAdd("S", "SV")
			m_current += 2
			return true
		}

		//
		if namesBeginningWithSWThatGetAltXV() {
			metaphAdd("S", "XV")
			m_current += 2
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
func encodeSKJ() bool {
	// scandinavian
	if stringAt(0, "SKJO", "SKJU") &&
		isVowel(m_current+3) {
		metaphAdd("X")
		m_current += 3
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
func encodeSJ() bool {
	if stringAtStart("SJ") {
		metaphAdd("X")
		m_current += 2
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
func encodeSilentFrenchSFinal() bool {
	// "louis" is an exception because it gets two pronuncuations
	if stringAtStart("LOUIS") && (m_current == m_last) {
		metaphAdd("S", "")
		m_current++
		return true
	}

	// french words familiar to americans where final s is silent
	if (m_current == m_last) &&
		(stringAtStart("YVES") ||
			(stringAtStart("HORS") && (m_current == 3)) ||
			stringAt(-4, "CAMUS", "YPRES") ||
			stringAt(-5, "MESNES", "DEBRIS", "BLANCS", "INGRES", "CANNES") ||
			stringAt(-6, "CHABLIS", "APROPOS", "JACQUES", "ELYSEES", "OEUVRES",
				"GEORGES", "DESPRES") ||
			stringAtStart("ARKANSAS", "FRANCAIS", "CRUDITES", "BRUYERES") ||
			stringAtStart("DESCARTES", "DESCHUTES", "DESCHAMPS", "DESROCHES", "DESCHENES") ||
			stringAtStart("RENDEZVOUS") ||
			stringAtStart("CONTRETEMPS", "DESLAURIERS")) ||
		((m_current == m_last) &&
			stringAt(-2, "AI", "OI", "UI") &&
			!stringAtStart("LOIS", "LUIS")) {
		m_current++
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
func encodeSilentFrenchSInternal() bool {
	// french words familiar to americans where internal s is silent
	if stringAt(-2, "DESCARTES") ||
		stringAt(-2, "DESCHAM", "DESPRES", "DESROCH", "DESROSI", "DESJARD", "DESMARA",
			"DESCHEN", "DESHOTE", "DESLAUR") ||
		stringAt(-2, "MESNES") ||
		stringAt(-5, "DUQUESNE", "DUCHESNE") ||
		stringAt(-7, "BEAUCHESNE") ||
		stringAt(-3, "FRESNEL") ||
		stringAt(-3, "GROSVENOR") ||
		stringAt(-4, "LOUISVILLE") ||
		stringAt(-7, "ILLINOISAN") {
		m_current++
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
func encodeISL() bool {
	//special cases 'island', 'isle', 'carlisle', 'carlysle'
	if (stringAt(-2, "LISL", "LYSL", "AISL") &&
		!stringAt(-3, "PAISLEY", "BAISLEY", "ALISLAM", "ALISLAH", "ALISLAA")) ||
		((m_current == 1) &&
			((stringAt(-1, "ISLE") ||
				stringAt(-1, "ISLAN")) &&
				!stringAt(-1, "ISLEY", "ISLER"))) {
		m_current++
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
func encodeSTL() bool {
	//'hustle', 'bustle', 'whistle'
	if (stringAt(0, "STLE", "STLI") &&
		!stringAt(+2, "LESS", "LIKE", "LINE")) ||
		stringAt(-3, "THISTLY", "BRISTLY", "GRISTLY") ||
		// e.g. "corpuscle"
		stringAt(-1, "USCLE") {
		// KRISTEN, KRYSTLE, CRYSTLE, KRISTLE all pronounce the 't'
		// also, exceptions where "-LING" is a nominalizing suffix
		if stringAtStart("KRISTEN", "KRYSTLE", "CRYSTLE", "KRISTLE") ||
			stringAtStart("CHRISTENSEN", "CHRISTENSON") ||
			stringAt(-3, "FIRSTLING") ||
			stringAt(-2, "NESTLING", "WESTLING") {
			metaphAdd("ST")
			m_current += 2
		} else {
			if m_encodeVowels &&
				(charAt(m_current+3) == 'E') &&
				(charAt(m_current+4) != 'R') &&
				!stringAt(+3, "ETTE", "ETTA") &&
				!stringAt(+3, "EY") {
				metaphAdd("SAL")
				flag_AL_inversion = true
			} else {
				metaphAdd("SL")
			}
			m_current += 3
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
func encodeChristmas() bool {
	//'christmas'
	if stringAt(-4, "CHRISTMA") {
		metaphAdd("SM")
		m_current += 3
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
func encodeSTHM() bool {
	//'asthma', 'isthmus'
	if stringAt(0, "STHM") {
		metaphAdd("SM")
		m_current += 4
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
func encodeISTEN() bool {
	// 't' is silent in verb, pronounced in name
	if stringAtStart("CHRISTEN") {
		// the word itself
		if rootOrInflections(m_inWord, "CHRISTEN") ||
			stringAtStart("CHRISTENDOM") {
			metaphAdd("S", "ST")
		} else {
			// e.g. 'christenson', 'christene'
			metaphAdd("ST")
		}
		m_current += 2
		return true
	}

	//e.g. 'glisten', 'listen'
	if stringAt(-2, "LISTEN", "RISTEN", "HASTEN", "FASTEN", "MUSTNT") ||
		stringAt(-3, "MOISTEN") {
		metaphAdd("S")
		m_current += 2
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
func encodeSugar() bool {
	//special case 'sugar-'
	if stringAt(0, "SUGAR") {
		metaphAdd("X")
		m_current++
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
func encodeSH() bool {
	if stringAt(0, "SH") {
		// exception
		if stringAt(-2, "CASHMERE") {
			metaphAdd("J")
			m_current += 2
			return true
		}

		//combining forms, e.g. 'clotheshorse', 'woodshole'
		if (m_current > 0) &&
			// e.g. "mishap"
			((stringAt(+1, "HAP") && ((m_current + 3) == m_last)) ||
				// e.g. "hartsheim", "clothshorse"
				stringAt(+1, "HEIM", "HOEK", "HOLM", "HOLZ", "HOOD", "HEAD", "HEID",
					"HAAR", "HORS", "HOLE", "HUND", "HELM", "HAWK", "HILL") ||
				// e.g. "dishonor"
				stringAt(+1, "HEART", "HATCH", "HOUSE", "HOUND", "HONOR") ||
				// e.g. "mishear"
				(stringAt(+2, "EAR") && ((m_current + 4) == m_last)) ||
				// e.g. "hartshorn"
				(stringAt(+2, "ORN") && !stringAt(-2, "UNSHORN")) ||
				// e.g. "newshour" but not "bashour", "manshour"
				(stringAt(+1, "HOUR") &&
					!(stringAtStart("BASHOUR") || stringAtStart("MANSHOUR") || stringAtStart("ASHOUR"))) ||
				// e.g. "dishonest", "grasshopper"
				stringAt(+2, "ARMON", "ONEST", "ALLOW", "OLDER", "OPPER", "EIMER", "ANDLE", "ONOUR") ||
				// e.g. "dishabille", "transhumance"
				stringAt(+2, "ABILLE", "UMANCE", "ABITUA")) {
			if !stringAt(-1, "S") {
				metaphAdd("S")
			}
		} else {
			metaphAdd("X")
		}

		m_current += 2
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
func encodeSCH() bool {
	// these words were combining forms many centuries ago
	if stringAt(+1, "CH") {
		if (m_current > 0) &&
			// e.g. "mischief", "escheat"
			(stringAt(+3, "IEF", "EAT") ||
				// e.g. "mischance"
				stringAt(+3, "ANCE", "ARGE") ||
				// e.g. "eschew"
				stringAtStart("ESCHEW")) {
			metaphAdd("S")
			m_current++
			return true
		}

		//Schlesinger's rule
		//dutch, danish, italian, greek origin, e.g. "school", "schooner", "schiavone", "schiz-"
		if (stringAt(+3, "OO", "ER", "EN", "UY", "ED", "EM", "IA", "IZ", "IS", "OL") &&
			!stringAt(0, "SCHOLT", "SCHISL", "SCHERR")) ||
			stringAt(+3, "ISZ") ||
			(stringAt(-1, "ESCHAT", "ASCHIN", "ASCHAL", "ISCHAE", "ISCHIA") &&
				!stringAt(-2, "FASCHING")) ||
			(stringAt(-1, "ESCHI") && ((m_current + 3) == m_last)) ||
			(charAt(m_current+3) == 'Y') {
			// e.g. "schermerhorn", "schenker", "schistose"
			if stringAt(+3, "ER", "EN", "IS") &&
				(((m_current + 4) == m_last) ||
					stringAt(+3, "ENK", "ENB", "IST")) {
				metaphAdd("X", "SK")
			} else {
				metaphAdd("SK")
			}
			m_current += 3
			return true
		} else {
			// Fix for smith and schmidt not returning same code:
			// next two lines from metaphone.go at line 621: code for SCH
			if m_current == 0 && !isVowel(3) && (charAt(3) != 'W') {
				metaphAdd("X", "S")
			} else {
				metaphAdd("X")
			}
			m_current += 3
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
func encodeSUR() bool {
	// 'erasure', 'usury'
	if stringAt(+1, "URE", "URA", "URY") {
		//'sure', 'ensure'
		if (m_current == 0) ||
			stringAt(-1, "N", "K") ||
			stringAt(-2, "NO") {
			metaphAdd("X")
		} else {
			metaphAdd("J")
		}

		advanceCounter(2, 1)
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
func encodeSU() bool {
	//'sensuous', 'consensual'
	if stringAt(+1, "UO", "UA") && (m_current != 0) {
		// exceptions e.g. "persuade"
		if stringAt(-1, "RSUA") {
			metaphAdd("S")
			// exceptions e.g. "casual"
		} else if isVowel(m_current - 1) {
			metaphAdd("J", "S")
		} else {
			metaphAdd("X", "S")
		}

		advanceCounter(3, 1)
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
func encodeSSIO() bool {
	if stringAt(+1, "SION") {
		//"abcission"
		if stringAt(-2, "CI") {
			metaphAdd("J")
			//'mission'
		} else {
			if isVowel(m_current - 1) {
				metaphAdd("X")
			}
		}

		advanceCounter(4, 2)
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
func encodeSS() bool {
	// e.g. "russian", "pressure"
	if stringAt(-1, "USSIA", "ESSUR", "ISSUR", "ISSUE") ||
		// e.g. "hessian", "assurance"
		stringAt(-1, "ESSIAN", "ASSURE", "ASSURA", "ISSUAB", "ISSUAN", "ASSIUS") {
		metaphAdd("X")
		advanceCounter(3, 2)
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
func encodeSIA() bool {
	// e.g. "controversial", also "fuchsia", "ch" is silent
	if stringAt(-2, "CHSIA") ||
		stringAt(-1, "RSIAL") {
		metaphAdd("X")
		advanceCounter(3, 1)
		return true
	}

	// names generally get 'X' where terms, e.g. "aphasia" get 'J'
	if (stringAtStart("ALESIA", "ALYSIA", "ALISIA", "STASIA") &&
		(m_current == 3) &&
		!stringAtStart("ANASTASIA")) ||
		stringAt(-5, "DIONYSIAN") ||
		stringAt(-5, "THERESIA") {
		metaphAdd("X", "S")
		advanceCounter(3, 1)
		return true
	}

	if (stringAt(0, "SIA") && ((m_current + 2) == m_last)) ||
		(stringAt(0, "SIAN") && ((m_current + 3) == m_last)) ||
		stringAt(-5, "AMBROSIAL") {
		if (isVowel(m_current-1) || stringAt(-1, "R")) &&
			// exclude compounds based on names, or french or greek words
			!(stringAtStart("JAMES", "NICOS", "PEGAS", "PEPYS") ||
				stringAtStart("HOBBES", "HOLMES", "JAQUES", "KEYNES") ||
				stringAtStart("MALTHUS", "HOMOOUS") ||
				stringAtStart("MAGLEMOS", "HOMOIOUS") ||
				stringAtStart("LEVALLOIS", "TARDENOIS") ||
				stringAt(-4, "ALGES")) {
			metaphAdd("J")
		} else {
			metaphAdd("S")
		}

		advanceCounter(2, 1)
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
func encodeSIO() bool {
	// special case, irish name
	if stringAtStart("SIOBHAN") {
		metaphAdd("X")
		advanceCounter(3, 1)
		return true
	}

	if stringAt(+1, "ION") {
		// e.g. "vision", "version"
		if isVowel(m_current-1) || stringAt(-2, "ER", "UR") {
			metaphAdd("J")
		} else {
			// e.g. "declension"
			metaphAdd("X")
		}

		advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode cases where "-S-" might well be from a german name
 * and add encoding of german pronounciation in alternate m_metaph
 * so that it can be found in a genealogical search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeAnglicisations() bool {
	//german & anglicisations, e.g. 'smith' match 'schmidt', 'snider' match 'schneider'
	//also, -sz- in slavic language altho in hungarian it is pronounced 's'
	if ((m_current == 0) &&
		stringAt(+1, "M", "N", "L")) ||
		stringAt(+1, "Z") {
		metaphAdd("S", "X")

		// eat redundant 'Z'
		if stringAt(+1, "Z") {
			m_current += 2
		} else {
			m_current++
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
func encodeSC() bool {
	if stringAt(0, "SC") {
		// exception 'viscount'
		if stringAt(-2, "VISCOUNT") {
			m_current += 1
			return true
		}

		// encode "-SC<front vowel>-"
		if stringAt(+2, "I", "E", "Y") {
			// e.g. "conscious"
			if stringAt(+2, "IOUS") ||
				// e.g. "prosciutto"
				stringAt(+2, "IUT") ||
				stringAt(-4, "OMNISCIEN") ||
				// e.g. "conscious"
				stringAt(-3, "CONSCIEN", "CRESCEND", "CONSCION") ||
				stringAt(-2, "FASCIS") {
				metaphAdd("X")
			} else if stringAt(0, "SCEPTIC", "SCEPSIS") ||
				stringAt(0, "SCIVV", "SCIRO") ||
				// commonly pronounced this way in u.s.
				stringAt(0, "SCIPIO") ||
				stringAt(-2, "PISCITELLI") {
				metaphAdd("SK")
			} else {
				metaphAdd("S")
			}
			m_current += 2
			return true
		}

		metaphAdd("SK")
		m_current += 2
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
func encodeSEASUISIER() bool {
	// "nausea" by itself has => NJ as a more likely encoding. Other forms
	// using "nause-" (see encodeSEA()) have X or S as more familiar pronounciations
	if (stringAt(-3, "NAUSEA") && ((m_current + 2) == m_last)) ||
		// e.g. "casuistry", "frasier", "hoosier"
		stringAt(-2, "CASUI") ||
		(stringAt(-1, "OSIER", "ASIER") &&
			!(stringAtStart("EASIER") ||
				stringAtStart("OSIER") ||
				stringAt(-2, "ROSIER", "MOSIER"))) {
		metaphAdd("J", "X")
		advanceCounter(3, 1)
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
func encodeSEA() bool {
	if (stringAtStart("SEAN") && ((m_current + 3) == m_last)) ||
		(stringAt(-3, "NAUSEO") &&
			!stringAt(-3, "NAUSEAT")) {
		metaphAdd("X")
		advanceCounter(3, 1)
		return true
	}

	return false
}

/**
 * Encode "-S-"
 *
 */
func encodeS() {
	if encodeSKJ() ||
		encodeSpecialSW() ||
		encodeSJ() ||
		encodeSilentFrenchSFinal() ||
		encodeSilentFrenchSInternal() ||
		encodeISL() ||
		encodeSTL() ||
		encodeChristmas() ||
		encodeSTHM() ||
		encodeISTEN() ||
		encodeSugar() ||
		encodeSH() ||
		encodeSCH() ||
		encodeSUR() ||
		encodeSU() ||
		encodeSSIO() ||
		encodeSS() ||
		encodeSIA() ||
		encodeSIO() ||
		encodeAnglicisations() ||
		encodeSC() ||
		encodeSEASUISIER() ||
		encodeSEA() {
		return
	}

	metaphAdd("S")

	if stringAt(+1, "S", "Z") &&
		!stringAt(+1, "SH") {
		m_current += 2
	} else {
		m_current++
	}
}

/**
 * Encode some exceptions for initial 'T'
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeTInitial() bool {
	if m_current == 0 {
		// americans usually pronounce "tzar" as "zar"
		if stringAt(+1, "SAR", "ZAR") {
			m_current++
			return true
		}

		// old 'École française d'Extrême-Orient' chinese pinyin where 'ts-' => 'X'
		if ((m_length == 3) && stringAt(+1, "SO", "SA", "SU")) ||
			((m_length == 4) && stringAt(+1, "SAO", "SAI")) ||
			((m_length == 5) && stringAt(+1, "SING", "SANG")) {
			metaphAdd("X")
			advanceCounter(3, 2)
			return true
		}

		// "TS<vowel>-" at start can be pronounced both with and without 'T'
		if stringAt(+1, "S") && isVowel(m_current+2) {
			metaphAdd("TS", "S")
			advanceCounter(3, 2)
			return true
		}

		// e.g. "Tjaarda"
		if stringAt(+1, "J") {
			metaphAdd("X")
			advanceCounter(3, 2)
			return true
		}

		// cases where initial "TH-" is pronounced as T and not 0 ("th")
		if (stringAt(+1, "HU") && (m_length == 3)) ||
			stringAt(+1, "HAI", "HUY", "HAO") ||
			stringAt(+1, "HYME", "HYMY", "HANH") ||
			stringAt(+1, "HERES") {
			metaphAdd("T")
			advanceCounter(3, 2)
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
func encodeTCH() bool {
	if stringAt(+1, "CH") {
		metaphAdd("X")
		m_current += 3
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
func encodeSilentFrenchT() bool {
	// french silent T familiar to americans
	if ((m_current == m_last) && stringAt(-4, "MONET", "GENET", "CHAUT")) ||
		stringAt(-2, "POTPOURRI") ||
		stringAt(-3, "BOATSWAIN") ||
		stringAt(-3, "MORTGAGE") ||
		(stringAt(-4, "BERET", "BIDET", "FILET", "DEBUT", "DEPOT", "PINOT", "TAROT") ||
			stringAt(-5, "BALLET", "BUFFET", "CACHET", "CHALET", "ESPRIT", "RAGOUT", "GOULET",
				"CHABOT", "BENOIT") ||
			stringAt(-6, "GOURMET", "BOUQUET", "CROCHET", "CROQUET", "PARFAIT", "PINCHOT",
				"CABARET", "PARQUET", "RAPPORT", "TOUCHET", "COURBET", "DIDEROT") ||
			stringAt(-7, "ENTREPOT", "CABERNET", "DUBONNET", "MASSENET", "MUSCADET", "RICOCHET", "ESCARGOT") ||
			stringAt(-8, "SOBRIQUET", "CABRIOLET", "CASSOULET", "OUBRIQUET", "CAMEMBERT")) &&
			!stringAt(+1, "AN", "RY", "IC", "OM", "IN") {
		m_current++
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
func encodeTUNTULTUATUO() bool {
	// e.g. "fortune", "fortunate"
	if stringAt(-3, "FORTUN") ||
		// e.g. "capitulate"
		(stringAt(0, "TUL") &&
			(isVowel(m_current-1) && isVowel(m_current+3))) ||
		// e.g. "obituary", "barbituate"
		stringAt(-2, "BITUA", "BITUE") ||
		// e.g. "actual"
		((m_current > 1) && stringAt(0, "TUA", "TUO")) {
		metaphAdd("X", "T")
		m_current++
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
func encodeTUETEUTEOUTULTIE() bool {
	// 'constituent', 'pasteur'
	if stringAt(+1, "UENT") ||
		stringAt(-4, "RIGHTEOUS") ||
		stringAt(-3, "STATUTE") ||
		stringAt(-3, "AMATEUR") ||
		// e.g. "blastula", "pasteur"
		(stringAt(-1, "NTULE", "NTULA", "STULE", "STULA", "STEUR")) ||
		// e.g. "statue"
		(((m_current + 2) == m_last) && stringAt(0, "TUE")) ||
		// e.g. "constituency"
		stringAt(0, "TUENC") ||
		// e.g. "statutory"
		stringAt(-3, "STATUTOR") ||
		// e.g. "patience"
		(((m_current + 5) == m_last) && stringAt(0, "TIENCE")) {
		metaphAdd("X", "T")
		advanceCounter(2, 1)
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
func encodeTURTIUSuffixes() bool {
	// 'adventure', 'musculature'
	if (m_current > 0) && stringAt(+1, "URE", "URA", "URI", "URY", "URO", "IUS") {
		// exceptions e.g. 'tessitura', mostly from romance languages
		if (stringAt(+1, "URA", "URO") &&
			//&& !stringAt(+1, "URIA")
			((m_current+3) == m_last)) &&
			!stringAt(-3, "VENTURA") ||
			// e.g. "kachaturian", "hematuria"
			stringAt(+1, "URIA") {
			metaphAdd("T")
		} else {
			metaphAdd("X", "T")
		}

		advanceCounter(2, 1)
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
func encodeTI() bool {
	// '-tio-', '-tia-', '-tiu-'
	// except combining forms where T already pronounced e.g 'rooseveltian'
	if (stringAt(+1, "IO") && !stringAt(-1, "ETIOL")) ||
		stringAt(+1, "IAL") ||
		stringAt(-1, "RTIUM", "ATIUM") ||
		((stringAt(+1, "IAN") && (m_current > 0)) &&
			!(stringAt(-4, "FAUSTIAN") ||
				stringAt(-5, "PROUSTIAN") ||
				stringAt(-2, "TATIANA") ||
				(stringAt(-3, "KANTIAN", "GENTIAN") ||
					stringAt(-8, "ROOSEVELTIAN"))) ||
			(((m_current + 2) == m_last) &&
				stringAt(0, "TIA") &&
				// exceptions to above rules where the pronounciation is usually X
				!(stringAt(-3, "HESTIA", "MASTIA") ||
					stringAt(-2, "OSTIA") ||
					stringAtStart("TIA") ||
					stringAt(-5, "IZVESTIA"))) ||
			stringAt(+1, "IATE", "IATI", "IABL", "IATO", "IARY") ||
			stringAt(-5, "CHRISTIAN")) {
		if ((m_current == 2) && stringAtStart("ANTI")) ||
			stringAtStart("PATIO", "PITIA", "DUTIA") {
			metaphAdd("T")
		} else if stringAt(-4, "EQUATION") {
			metaphAdd("J")
		} else {
			if stringAt(0, "TION") {
				metaphAdd("X")
			} else if stringAtStart("KATIA", "LATIA") {
				metaphAdd("T", "X")
			} else {
				metaphAdd("X", "T")
			}
		}

		advanceCounter(3, 1)
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
func encodeTIENT() bool {
	// e.g. 'patient'
	if stringAt(+1, "IENT") {
		metaphAdd("X", "T")
		advanceCounter(3, 1)
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
func encodeTSCH() bool {
	//'deutsch'
	if stringAt(0, "TSCH") &&
		// combining forms in german where the 'T' is pronounced seperately
		!stringAt(-3, "WELT", "KLAT", "FEST") {
		// pronounced the same as "ch" in "chit" => X
		metaphAdd("X")
		m_current += 4
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
func encodeTZSCH() bool {
	//'neitzsche'
	if stringAt(0, "TZSCH") {
		metaphAdd("X")
		m_current += 5
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
func encodeTHPronouncedSeparately() bool {
	//'adulthood', 'bithead', 'apartheid'
	if ((m_current > 0) &&
		stringAt(+1, "HOOD", "HEAD", "HEID", "HAND", "HILL", "HOLD",
			"HAWK", "HEAP", "HERD", "HOLE", "HOOK", "HUNT",
			"HUMO", "HAUS", "HOFF", "HARD") &&
		!stringAt(-3, "SOUTH", "NORTH")) ||
		stringAt(+1, "HOUSE", "HEART", "HASTE", "HYPNO", "HEQUE") ||
		// watch out for greek root "-thallic"
		(stringAt(+1, "HALL") &&
			((m_current + 4) == m_last) &&
			!stringAt(-3, "SOUTH", "NORTH")) ||
		(stringAt(+1, "HAM") &&
			((m_current + 3) == m_last) &&
			!(stringAtStart("GOTHAM", "WITHAM", "LATHAM") ||
				stringAtStart("BENTHAM", "WALTHAM", "WORTHAM") ||
				stringAtStart("GRANTHAM"))) ||
		(stringAt(+1, "HATCH") &&
			!((m_current == 0) || stringAt(-2, "UNTHATCH"))) ||
		stringAt(-3, "WARTHOG") ||
		// and some special cases where "-TH-" is usually pronounced 'T'
		stringAt(-2, "ESTHER") ||
		stringAt(-3, "GOETHE") ||
		stringAt(-2, "NATHALIE") {
		// special case
		if stringAt(-3, "POSTHUM") {
			metaphAdd("X")
		} else {
			metaphAdd("T")
		}
		m_current += 2
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
func encodeTTH() bool {
	// 'matthew' vs. 'outthink'
	if stringAt(0, "TTH") {
		if stringAt(-2, "MATTH") {
			metaphAdd("0")
		} else {
			metaphAdd("T0")
		}
		m_current += 3
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
func encodeTH() bool {
	if stringAt(0, "TH") {
		//'-clothes-'
		if stringAt(-3, "CLOTHES") {
			// vowel already encoded so skip right to S
			m_current += 3
			return true
		}

		//special case "thomas", "thames", "beethoven" or germanic words
		if stringAt(+2, "OMAS", "OMPS", "OMPK", "OMSO", "OMSE",
			"AMES", "OVEN", "OFEN", "ILDA", "ILDE") ||
			(stringAtStart("THOM") && (m_length == 4)) ||
			(stringAtStart("THOMS") && (m_length == 5)) ||
			stringAtStart("VAN ", "VON ") ||
			stringAtStart("SCH") {
			metaphAdd("T")

		} else {
			// give an 'etymological' 2nd
			// encoding for "smith"
			if stringAtStart("SM") {
				metaphAdd("0", "T")
			} else {
				metaphAdd("0")
			}
		}

		m_current += 2
		return true
	}

	return false
}

/**
 * Encode "-T-"
 *
 */
func encodeT() {
	if encodeTInitial() ||
		encodeTCH() ||
		encodeSilentFrenchT() ||
		encodeTUNTULTUATUO() ||
		encodeTUETEUTEOUTULTIE() ||
		encodeTURTIUSuffixes() ||
		encodeTI() ||
		encodeTIENT() ||
		encodeTSCH() ||
		encodeTZSCH() ||
		encodeTHPronouncedSeparately() ||
		encodeTTH() ||
		encodeTH() {
		return
	}

	// eat redundant 'T' or 'D'
	if stringAt(+1, "T", "D") {
		m_current += 2
	} else {
		m_current++
	}

	metaphAdd("T")
}

/**
 * Encode "-V-"
 *
 */
func encodeV() {
	// eat redundant 'V'
	if charAt(m_current+1) == 'V' {
		m_current += 2
	} else {
		m_current++
	}

	metaphAddExactApprox("V", "F")
}

/**
 * Encode cases where 'W' is silent at beginning of word
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeSilentWAtBeginning() bool {
	//skip these when at start of word
	if (m_current == 0) &&
		stringAt(0, "WR") {
		m_current += 1
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
func encodeWITZWICZ() bool {
	//polish e.g. 'filipowicz'
	if ((m_current + 3) == m_last) && stringAt(0, "WICZ", "WITZ") {
		if m_encodeVowels {
			if (len(m_primary) > 0) &&
				charAt(len(m_primary)-1) == 'A' {
				metaphAdd("TS", "FAX")
			} else {
				metaphAdd("ATS", "FAX")
			}
		} else {
			metaphAdd("TS", "FX")
		}
		m_current += 4
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
func encodeWR() bool {
	//can also be in middle of word
	if stringAt(0, "WR") {
		metaphAdd("R")
		m_current += 2
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
func germanicOrSlavicNameBeginningWithW() bool {
	if stringAtStart("WEE", "WIX", "WAX") ||
		stringAtStart("WOLF", "WEIS", "WAHL", "WALZ", "WEIL", "WERT",
			"WINE", "WILK", "WALT", "WOLL", "WADA", "WULF",
			"WEHR", "WURM", "WYSE", "WENZ", "WIRT", "WOLK",
			"WEIN", "WYSS", "WASS", "WANN", "WINT", "WINK",
			"WILE", "WIKE", "WIER", "WELK", "WISE") ||
		stringAtStart("WIRTH", "WIESE", "WITTE", "WENTZ", "WOLFF", "WENDT",
			"WERTZ", "WILKE", "WALTZ", "WEISE", "WOOLF", "WERTH",
			"WEESE", "WURTH", "WINES", "WARGO", "WIMER", "WISER",
			"WAGER", "WILLE", "WILDS", "WAGAR", "WERTS", "WITTY",
			"WIENS", "WIEBE", "WIRTZ", "WYMER", "WULFF", "WIBLE",
			"WINER", "WIEST", "WALKO", "WALLA", "WEBRE", "WEYER",
			"WYBLE", "WOMAC", "WILTZ", "WURST", "WOLAK", "WELKE",
			"WEDEL", "WEIST", "WYGAN", "WUEST", "WEISZ", "WALCK",
			"WEITZ", "WYDRA", "WANDA", "WILMA", "WEBER") ||
		stringAtStart("WETZEL", "WEINER", "WENZEL", "WESTER", "WALLEN", "WENGER",
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
			"WARFEL", "WYNTER", "WERNER", "WAGNER", "WISSER") ||
		stringAtStart("WISEMAN", "WINKLER", "WILHELM", "WELLMAN", "WAMPLER", "WACHTER",
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
			"WISHART", "WILLIAM") ||
		stringAtStart("WESTPHAL", "WICKLUND", "WEISSMAN", "WESTLUND", "WOLFGANG", "WILLHITE",
			"WEISBERG", "WALRAVEN", "WOLFGRAM", "WILHOITE", "WECHSLER", "WENDLING",
			"WESTBERG", "WENDLAND", "WININGER", "WHISNANT", "WESTRICK", "WESTLING",
			"WESTBURY", "WEITZMAN", "WEHMEYER", "WEINMANN", "WISNESKI", "WHELCHEL",
			"WEISHAAR", "WAGGENER", "WALDROUP", "WESTHOFF", "WIEDEMAN", "WASINGER",
			"WINBORNE") ||
		stringAtStart("WHISENANT", "WEINSTEIN", "WESTERMAN", "WASSERMAN", "WITKOWSKI", "WEINTRAUB",
			"WINKELMAN", "WINKFIELD", "WANAMAKER", "WIECZOREK", "WIECHMANN", "WOJTOWICZ",
			"WALKOWIAK", "WEINSTOCK", "WILLEFORD", "WARKENTIN", "WEISINGER", "WINKLEMAN",
			"WILHEMINA") ||
		stringAtStart("WISNIEWSKI", "WUNDERLICH", "WHISENHUNT", "WEINBERGER", "WROBLEWSKI",
			"WAGUESPACK", "WEISGERBER", "WESTERVELT", "WESTERLUND", "WASILEWSKI",
			"WILDERMUTH", "WESTENDORF", "WESOLOWSKI", "WEINGARTEN", "WINEBARGER",
			"WESTERBERG", "WANNAMAKER", "WEISSINGER") ||
		stringAtStart("WALDSCHMIDT", "WEINGARTNER", "WINEBRENNER") ||
		stringAtStart("WOLFENBARGER") ||
		stringAtStart("WOJCIECHOWSKI") {
		return true
	}

	return false
}

/**
 * Encode "W-", adding central and eastern european
 * pronounciations so that both forms can be found
 * in a genealogy search
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeInitialWVowel() bool {
	if (m_current == 0) && isVowel(m_current+1) {
		//Witter should match Vitter
		if germanicOrSlavicNameBeginningWithW() {
			if m_encodeVowels {
				metaphAddExactApprox("A", "VA", "A", "FA")
			} else {
				metaphAddExactApprox("A", "V", "A", "F")
			}
		} else {
			metaphAdd("A")
		}

		m_current++
		// don't encode vowels twice
		m_current = skipVowels(m_current)
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
func encodeWH() bool {
	if stringAt(0, "WH") {
		// cases where it is pronounced as H
		// e.g. 'who', 'whole'
		if (charAt(m_current+2) == 'O') &&
			// exclude cases where it is pronounced like a vowel
			!(stringAt(+2, "OOSH") ||
				stringAt(+2, "OOP", "OMP", "ORL", "ORT") ||
				stringAt(+2, "OA", "OP")) {
			metaphAdd("H")
			advanceCounter(3, 2)
			return true
		} else {
			// combining forms, e.g. 'hollowhearted', 'rawhide'
			if stringAt(+2, "IDE", "ARD", "EAD", "AWK", "ERD",
				"OOK", "AND", "OLE", "OOD") ||
				stringAt(+2, "EART", "OUSE", "OUND") ||
				stringAt(+2, "AMMER") {
				metaphAdd("H")
				m_current += 2
				return true
			} else if m_current == 0 {
				metaphAdd("A")
				m_current += 2
				// don't encode vowels twice
				m_current = skipVowels(m_current)
				return true
			}
		}
		m_current += 2
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
func encodeEasternEuropeanW() bool {
	//Arnow should match Arnoff
	if ((m_current == m_last) && isVowel(m_current-1)) ||
		stringAt(-1, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
		(stringAt(0, "WICKI", "WACKI") && ((m_current + 4) == m_last)) ||
		stringAt(0, "WIAK") && ((m_current+3) == m_last) ||
		stringAtStart("SCH") {
		metaphAddExactApprox("", "V", "", "F")
		m_current++
		return true
	}

	return false
}

/**
 * Encode "-W-"
 *
 */
func encodeW() {
	if encodeSilentWAtBeginning() ||
		encodeWITZWICZ() ||
		encodeWR() ||
		encodeInitialWVowel() ||
		encodeWH() ||
		encodeEasternEuropeanW() {
		return
	}

	// e.g. 'zimbabwe'
	if m_encodeVowels &&
		stringAt(0, "WE") &&
		((m_current + 1) == m_last) {
		metaphAdd("A")
	}

	//else skip it
	m_current++

}

/**
 * Encode initial X where it is usually pronounced as S
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeInitialX() bool {
	// current chinese pinyin spelling
	if stringAtStart("XIA", "XIO", "XIE") ||
		stringAtStart("XU") {
		metaphAdd("X")
		m_current++
		return true
	}

	// else
	if m_current == 0 {
		metaphAdd("S")
		m_current++
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
func encodeGreekX() bool {
	// 'xylophone', xylem', 'xanthoma', 'xeno-'
	if stringAt(+1, "YLO", "YLE", "ENO") ||
		stringAt(+1, "ANTH") {
		metaphAdd("S")
		m_current++
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
func encodeXSpecialCases() bool {
	// 'luxury'
	if stringAt(-2, "LUXUR") {
		metaphAddExactApprox("GJ", "KJ")
		m_current++
		return true
	}

	// 'texeira' portuguese/galician name
	if stringAtStart("TEXEIRA") ||
		stringAtStart("TEIXEIRA") {
		metaphAdd("X")
		m_current++
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
func encodeXToH() bool {
	// TODO: look for other mexican indian words
	// where 'X' is usually pronounced this way
	if stringAt(-2, "OAXACA") ||
		stringAt(-3, "QUIXOTE") {
		metaphAdd("H")
		m_current++
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
func encodeXVowel() bool {
	// e.g. "sexual", "connexion" (british), "noxious"
	if stringAt(+1, "UAL", "ION", "IOU") {
		metaphAdd("KX", "KS")
		advanceCounter(3, 1)
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
func encodeFrenchXFinal() bool {
	//french e.g. "breaux", "paix"
	if !((m_current == m_last) &&
		(stringAt(-3, "IAU", "EAU", "IEU") ||
			stringAt(-2, "AI", "AU", "OU", "OI", "EU"))) {
		metaphAdd("KS")
	}

	return false
}

/**
 * Encode "-X-"
 *
 */
func encodeX() {
	if encodeInitialX() ||
		encodeGreekX() ||
		encodeXSpecialCases() ||
		encodeXToH() ||
		encodeXVowel() ||
		encodeFrenchXFinal() {
		return
	}

	// eat redundant 'X' or other redundant cases
	if stringAt(+1, "X", "Z", "S") ||
		// e.g. "excite", "exceed"
		stringAt(+1, "CI", "CE") {
		m_current += 2
	} else {
		m_current++
	}
}

/**
 * Encode cases of "-ZZ-" where it is obviously part
 * of an italian word where "-ZZ-" is pronounced as TS
 *
 * @return true if encoding handled in this routine, false if not
 *
 */
func encodeZZ() bool {
	// "abruzzi", 'pizza'
	if (charAt(m_current+1) == 'Z') &&
		((stringAt(+2, "I", "O", "A") &&
			((m_current + 2) == m_last)) ||
			stringAt(-2, "MOZZARELL", "PIZZICATO", "PUZZONLAN")) {
		metaphAdd("TS", "S")
		m_current += 2
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
func encodeZUZIERZS() bool {
	if ((m_current == 1) && stringAt(-1, "AZUR")) ||
		(stringAt(0, "ZIER") &&
			!stringAt(-2, "VIZIER")) ||
		stringAt(0, "ZSA") {
		metaphAdd("J", "S")

		if stringAt(0, "ZSA") {
			m_current += 2
		} else {
			m_current++
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
func encodeFrenchEZ() bool {
	if ((m_current == 3) && stringAt(-3, "CHEZ")) ||
		stringAt(-5, "RENDEZ") {
		m_current++
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
func encodeGermanZ() bool {
	if ((m_current == 2) && ((m_current + 1) == m_last) && stringAt(-2, "NAZI")) ||
		stringAt(-2, "NAZIFY", "MOZART") ||
		stringAt(-3, "HOLZ", "HERZ", "MERZ", "FITZ") ||
		(stringAt(-3, "GANZ") && !isVowel(m_current+1)) ||
		stringAt(-4, "STOLZ", "PRINZ") ||
		stringAt(-4, "VENEZIA") ||
		stringAt(-3, "HERZOG") ||
		// german words beginning with "sch-" but not schlimazel, schmooze
		(strings.Contains(string(m_inWord), "SCH") && !(stringAtEnd("IZE", "OZE", "ZEL"))) ||
		((m_current > 0) && stringAt(0, "ZEIT")) ||
		stringAt(-3, "WEIZ") {
		if (m_current > 0) && charAt(m_current-1) == 'T' {
			metaphAdd("S")
		} else {
			metaphAdd("TS")
		}
		m_current++
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
func encodeZH() bool {
	//chinese pinyin e.g. 'zhao', also english "phonetic spelling"
	if charAt(m_current+1) == 'H' {
		metaphAdd("J")
		m_current += 2
		return true
	}

	return false
}

/**
 * Encode "-Z-"
 *
 */
func encodeZ() {
	if encodeZZ() ||
		encodeZUZIERZS() ||
		encodeFrenchEZ() ||
		encodeGermanZ() {
		return
	}

	if encodeZH() {
		return
	} else {
		metaphAdd("S")
	}

	// eat redundant 'Z'
	if charAt(m_current+1) == 'Z' {
		m_current += 2
	} else {
		m_current++
	}
}

/**
 * Encodes input string to one or two key values according to Metaphone 3 rules.
 *
 */
func encode() {
	flag_AL_inversion = false

	m_current = 0

	m_primary = []rune{}
	m_secondary = []rune{}

	if m_length < 1 {
		return
	}

	//zero based index
	m_last = m_length - 1

	///////////main loop//////////////////////////
	for len(m_primary) < m_metaphLength || len(m_secondary) < m_metaphLength {
		if m_current >= m_length {
			break
		}
		switch charAt(m_current) {
		case 'B':
			encodeB()
		case 'ß':
		case 'Ç':
			metaphAdd("S")
			m_current++
		case 'C':
			encodeC()
		case 'D':
			encodeD()
		case 'F':
			encodeF()
		case 'G':
			encodeG()
		case 'H':
			encodeH()
		case 'J':
			encodeJ()
		case 'K':
			encodeK()
		case 'L':
			encodeL()
		case 'M':
			encodeM()
		case 'N':
			encodeN()
		case 'Ñ':
			metaphAdd("N")
			m_current++
		case 'P':
			encodeP()
		case 'Q':
			encodeQ()
		case 'R':
			encodeR()
		case 'S':
			encodeS()
		case 'T':
			encodeT()
		case 'Ð', 'Þ': // eth, thorn
			metaphAdd("0")
			m_current++
		case 'V':
			encodeV()
		case 'W':
			encodeW()
		case 'X':
			encodeX()
		case '':
			metaphAdd("X")
			m_current++
		case '':
			metaphAdd("S")
			m_current++
		case 'Z':
			encodeZ()
		default:
			if isVowel(m_current) {
				encodeVowels()
			} else {
				m_current++
			}
		}
	}

	//only give back m_metaphLength number of chars in m_metaph
	if len(m_primary) > m_metaphLength {
		m_primary = m_primary[:m_metaphLength]
	}

	if len(m_secondary) > m_metaphLength {
		m_secondary = m_secondary[:m_metaphLength]
	}

	// it is possible for the two metaphs to be the same
	// after truncation. lose the second one if so
	if string(m_primary) == string(m_secondary) {
		m_secondary = []rune{}
	}
}
