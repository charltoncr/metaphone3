// metaphone3.go - ported to the Go programming language by Ron Charlton
// on 2023-01-05 from the original Java code at
// https://github.com/OpenRefine/OpenRefine/blob/master/main/src/com/google/refine/clustering/binning/Metaphone3.java
//
// $Id: metaphone3.go,v 2.12 2023-01-17 16:31:36-05 ron Exp $
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

// Encode returns a primary and alternate encoding for word.  It honors
// the values set by SetEncodeVowels and SetEncodeExact, as well as by
// SetMaxLength.  The encodings will be the same for words that sound similar.
func (m *Metaphone3) Encode(word string) (metaph, metaph2 string) {
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

	// MetaphAdd adds a string to primary and secondary.  Call it with 1 or 2
	// arguments.  The first argument is appended to primary (and to
	// secondary if a second argument is not provided).  Any second
	// argument is appended to secondary.  But don't append an 'A' next to
	// another 'A'.
	MetaphAdd := func(s ...string) {
		if len(s) < 1 || len(s) > 2 {
			panic("MetaphAdd requires one or two arguments")
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
	MetaphAddExactApprox := func(s ...string) {
		if len(s) != 2 && len(s) != 4 {
			panic("MetaphAddExactApprox requires 2 or 4 arguments")
		}
		var mainExact, altExact, main, alt string
		if len(s) == 2 {
			mainExact = s[0]
			main = s[1]
			if m_encodeExact {
				MetaphAdd(mainExact)
			} else {
				MetaphAdd(main)
			}
		} else {
			mainExact = s[0]
			altExact = s[1]
			main = s[2]
			alt = s[3]
			if m_encodeExact {
				MetaphAdd(mainExact, altExact)
			} else {
				MetaphAdd(main, alt)
			}
		}
	}

	/**
	 * Subscript safe .charAt()
	 *
	 * @param at index of character to access
	 * @return null if index out of bounds, .charAt() otherwise
	 */
	CharAt := func(at int) rune {
		// check substring bounds
		if at < 0 || at >= m_length {
			return 0
		}

		return m_inWord[at]
	}

	// StringAtPos determines if any of a list of string arguments appear
	// in m_inWord at start.
	StringAtPos := func(start int, s ...string) bool {
		if len(s) > 0 {
			if start >= 0 {
			forOuterLoop:
				for _, str := range s {
					if (start + len(str)) <= m_length {
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
		}
		return false
	}

	// StringAt determines if any of a list of string arguments appear
	// in m_inWord at m_current+index.
	StringAt := func(index int, s ...string) bool {
		start := m_current + index
		return StringAtPos(start, s...)
	}

	// StringAtStart determines if any of a list of string arguments appear
	// in m_inWord at its beginning.
	StringAtStart := func(s ...string) bool {
		return StringAtPos(0, s...)
	}

	// StringAtEnd determines if any of a list of string arguments appear
	// in m_inWord at its end.
	StringAtEnd := func(s ...string) bool {
		if len(s) > 0 {
			start := m_length - len(s[0])
			return StringAtPos(start, s...)
		} else {
			return false
		}
	}

	/**
	 * Test for close front vowels
	 *
	 * @return true if close front vowel
	 */
	Front_Vowel := func(at int) bool {
		c := CharAt(at)
		return c == 'E' || c == 'I' || c == 'Y'
	}

	/**
	 * Detect names or words that begin with spellings
	 * typical of german or slavic words, for the purpose
	 * of choosing alternate pronunciations correctly
	 *
	 */
	SlavoGermanic := func() bool {
		return StringAtStart("SCH") ||
			StringAtStart("SW") ||
			(CharAt(0) == 'J') ||
			(CharAt(0) == 'W')
	}

	/**
	 * Tests if character is a vowel
	 *
	 * @param at rune to be tested in input word or integer location of same.
	 * @return true if character is a vowel, false if not
	 *
	 */
	IsVowel := func(at int) bool {
		return strings.ContainsRune("AEIOUYÀÁÂÃÄÅÆÈÉÊËÌÍÎÏÒÓÔÕÖØÙÚÛÜÝ", CharAt(at))
	}

	/**
	 * Skips over vowels in a string. Has exceptions for skipping consonants that
	 * will not be encoded.
	 *
	 * @param at position, in string to be encoded, of character to start skipping from
	 *
	 * @return position of next consonant in string to be encoded
	 */
	SkipVowels := func(at int) int {
		if at < 0 {
			return 0
		}
		if at >= m_length {
			return m_length
		}
		for IsVowel(at) || (CharAt(at) == 'W') {
			if StringAtPos(at, "WICZ", "WITZ", "WIAK") ||
				StringAtPos(at-1, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
				(StringAtPos(at, "WICKI", "WACKI") && ((at + 4) == m_last)) {
				break
			}
			at++
			if ((CharAt(at-1) == 'W') && (CharAt(at) == 'H')) &&
				!(StringAtPos(at, "HOP") ||
					StringAtPos(at, "HIDE", "HARD", "HEAD", "HAWK", "HERD", "HOOK", "HAND", "HOLE") ||
					StringAtPos(at, "HEART", "HOUSE", "HOUND") ||
					StringAtPos(at, "HAMMER")) {
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
	AdvanceCounter := func(ifNotEncodeVowels, ifEncodeVowels int) {
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
	RootOrInflections := func(InWord []rune, root string) bool {
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
	O_Silent := func() bool {
		// if "iron" at beginning or end of word and not "irony"
		if (CharAt(m_current) == 'O') && StringAt(-2, "IRON") {
			if (StringAtStart("IRON") ||
				(StringAt(-2, "IRON") &&
					(m_last == (m_current + 1)))) &&
				!StringAt(-2, "IRONIC") {
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
	E_Silent_Suffix := func(at int) bool {
		return (m_current == (at - 1)) &&
			(m_length > (at + 1)) &&
			(IsVowel((at + 1)) ||
				(StringAtPos(at, "ST", "SL") &&
					(m_length > (at + 2))))
	}

	/**
	 * Detect endings that will
	 * cause the 'e' to be pronounced
	 *
	 */
	E_Pronouncing_Suffix := func(at int) bool {
		// e.g. 'bridgewood' - the other vowels will get eaten
		// up so we need to put one in here
		if (m_length == (at + 4)) && StringAtPos(at, "WOOD") {
			return true
		}

		// same as above
		if (m_length == (at + 5)) && StringAtPos(at, "WATER", "WORTH") {
			return true
		}

		// e.g. 'bridgette'
		if (m_length == (at + 3)) && StringAtPos(at, "TTE", "LIA", "NOW", "ROS", "RAS") {
			return true
		}

		// e.g. 'olena'
		if (m_length == (at + 2)) && StringAtPos(at, "TA", "TT", "NA", "NO", "NE",
			"RS", "RE", "LA", "AU", "RO", "RA") {
			return true
		}

		// e.g. 'bridget'
		if (m_length == (at + 1)) && StringAtPos(at, "T", "R") {
			return true
		}

		return false
	}

	/**
	 * Detect internal silent 'E's e.g. "roseman",
	 * "firestone"
	 *
	 */
	Silent_Internal_E := func() bool {
		// 'olesen' but not 'olen'	RAKE BLAKE
		return (StringAtStart("OLE") &&
			E_Silent_Suffix(3) && !E_Pronouncing_Suffix(3)) ||
			(StringAtStart("BARE", "FIRE", "FORE", "GATE", "HAGE", "HAVE",
				"HAZE", "HOLE", "CAPE", "HUSE", "LACE", "LINE",
				"LIVE", "LOVE", "MORE", "MOSE", "MORE", "NICE",
				"RAKE", "ROBE", "ROSE", "SISE", "SIZE", "WARE",
				"WAKE", "WISE", "WINE") &&
				E_Silent_Suffix(4) && !E_Pronouncing_Suffix(4)) ||
			(StringAtStart("BLAKE", "BRAKE", "BRINE", "CARLE", "CLEVE", "DUNNE",
				"HEDGE", "HOUSE", "JEFFE", "LUNCE", "STOKE", "STONE",
				"THORE", "WEDGE", "WHITE") &&
				E_Silent_Suffix(5) && !E_Pronouncing_Suffix(5)) ||
			(StringAtStart("BRIDGE", "CHEESE") &&
				E_Silent_Suffix(6) && !E_Pronouncing_Suffix(6)) ||
			StringAt(-5, "CHARLES")
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
	E_Pronounced_At_End := func() bool {
		return (m_current == m_last) &&
			(StringAt(-6, "STROPHE") ||
				// if a vowel is before the 'E', vowel eater will have eaten it.
				//otherwise, consonant + 'E' will need 'E' pronounced
				(m_length == 2) ||
				((m_length == 3) && !IsVowel(0)) ||
				// these german name endings can be relied on to have the 'e' pronounced
				(StringAtEnd("BKE", "DKE", "FKE", "KKE", "LKE",
					"NKE", "MKE", "PKE", "TKE", "VKE", "ZKE") &&
					!StringAtStart("FINKE", "FUNKE") &&
					!StringAtStart("FRANKE")) ||
				StringAtEnd("SCHKE") ||
				(StringAtStart("ACME", "NIKE", "CAFE", "RENE", "LUPE", "JOSE", "ESME") && (m_length == 4)) ||
				(StringAtStart("LETHE", "CADRE", "TILDE", "SIGNE", "POSSE", "LATTE", "ANIME", "DOLCE", "CROCE",
					"ADOBE", "OUTRE", "JESSE", "JAIME", "JAFFE", "BENGE", "RUNGE",
					"CHILE", "DESME", "CONDE", "URIBE", "LIBRE", "ANDRE") && (m_length == 5)) ||
				(StringAtStart("HECATE", "PSYCHE", "DAPHNE", "PENSKE", "CLICHE", "RECIPE",
					"TAMALE", "SESAME", "SIMILE", "FINALE", "KARATE", "RENATE", "SHANTE",
					"OBERLE", "COYOTE", "KRESGE", "STONGE", "STANGE", "SWAYZE", "FUENTE",
					"SALOME", "URRIBE") && (m_length == 6)) ||
				(StringAtStart("ECHIDNE", "ARIADNE", "MEINEKE", "PORSCHE", "ANEMONE", "EPITOME",
					"SYNCOPE", "SOUFFLE", "ATTACHE", "MACHETE", "KARAOKE", "BUKKAKE",
					"VICENTE", "ELLERBE", "VERSACE") && (m_length == 7)) ||
				(StringAtStart("PENELOPE", "CALLIOPE", "CHIPOTLE", "ANTIGONE", "KAMIKAZE", "EURIDICE",
					"YOSEMITE", "FERRANTE") && (m_length == 8)) ||
				(StringAtStart("HYPERBOLE", "GUACAMOLE", "XANTHIPPE") && (m_length == 9)) ||
				(StringAtStart("SYNECDOCHE") && (m_length == 10)))
	}

	/**
	 * Encodes "-UE".
	 *
	 * @return true if encoding handled in this routine, false if not
	 */
	Skip_Silent_UE := func() bool {
		// always silent except for cases listed below
		if (StringAt(-1, "QUE", "GUE") &&
			!StringAtStart("BARBEQUE", "PALENQUE", "APPLIQUE") &&
			// '-que' cases usually french but missing the acute accent
			!StringAtStart("RISQUE") &&
			!StringAt(-3, "ARGUE", "SEGUE") &&
			!StringAtStart("PIROGUE", "ENRIQUE") &&
			!StringAtStart("COMMUNIQUE")) &&
			(m_current > 1) &&
			(((m_current + 1) == m_last) ||
				StringAtStart("JACQUES")) {
			m_current = SkipVowels(m_current)
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
	E_Silent := func() bool {
		if E_Pronounced_At_End() {
			return false
		}

		// 'e' silent when last letter, altho
		return (m_current == m_last) ||
			// also silent if before plural 's'
			// or past tense or participle 'd', e.g.
			// 'grapes' and 'banished' => PNXT
			(StringAtEnd("S", "D") &&
				(m_current > 1) &&
				((m_current + 1) == m_last) &&
				// and not e.g. "nested", "rises", or "pieces" => RASAS
				!(StringAt(-1, "TED", "SES", "CES") ||
					StringAtStart("ANTIPODES", "ANOPHELES") ||
					StringAtStart("MOHAMMED", "MUHAMMED", "MOUHAMED") ||
					StringAtStart("MOHAMED") ||
					StringAtStart("NORRED", "MEDVED", "MERCED", "ALLRED", "KHALED", "RASHED", "MASJED") ||
					StringAtStart("JARED", "AHMED", "HAMED", "JAVED") ||
					StringAtStart("ABED", "IMED"))) ||
			// e.g.  'wholeness', 'boneless', 'barely'
			(StringAt(+1, "NESS", "LESS") && ((m_current + 4) == m_last)) ||
			(StringAt(+1, "LY") && ((m_current + 2) == m_last) &&
				!StringAtStart("CICELY"))
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
	E_Pronounced_Exceptions := func() bool {
		// greek names e.g. "herakles" or hispanic names e.g. "robles", where 'e' is pronounced, other exceptions
		return (((m_current + 1) == m_last) &&
			(StringAt(-3, "OCLES", "ACLES", "AKLES") ||
				StringAtStart("INES") ||
				StringAtStart("LOPES", "ESTES", "GOMES", "NUNES", "ALVES", "ICKES",
					"INNES", "PERES", "WAGES", "NEVES", "BENES", "DONES") ||
				StringAtStart("CORTES", "CHAVES", "VALDES", "ROBLES", "TORRES", "FLORES", "BORGES",
					"NIEVES", "MONTES", "SOARES", "VALLES", "GEDDES", "ANDRES", "VIAJES",
					"CALLES", "FONTES", "HERMES", "ACEVES", "BATRES", "MATHES") ||
				StringAtStart("DELORES", "MORALES", "DOLORES", "ANGELES", "ROSALES", "MIRELES", "LINARES",
					"PERALES", "PAREDES", "BRIONES", "SANCHES", "CAZARES", "REVELES", "ESTEVES",
					"ALVARES", "MATTHES", "SOLARES", "CASARES", "CACERES", "STURGES", "RAMIRES",
					"FUNCHES", "BENITES", "FUENTES", "PUENTES", "TABARES", "HENTGES", "VALORES") ||
				StringAtStart("GONZALES", "MERCEDES", "FAGUNDES", "JOHANNES", "GONSALES", "BERMUDES",
					"CESPEDES", "BETANCES", "TERRONES", "DIOGENES", "CORRALES", "CABRALES",
					"MARTINES", "GRAJALES") ||
				StringAtStart("CERVANTES", "FERNANDES", "GONCALVES", "BENEVIDES", "CIFUENTES", "SIFUENTES",
					"SERVANTES", "HERNANDES", "BENAVIDES") ||
				StringAtStart("ARCHIMEDES", "CARRIZALES", "MAGALLANES"))) ||
			StringAt(-2, "FRED", "DGES", "DRED", "GNES") ||
			StringAt(-5, "PROBLEM", "RESPLEN") ||
			StringAt(-4, "REPLEN") ||
			StringAt(-3, "SPLE")
	}

	/**
	 * Encodes cases where non-initial 'e' is pronounced, taking
	 * care to detect unusual cases from the greek.
	 *
	 * Only executed if non initial vowel encoding is turned on
	 *
	 *
	 */
	Encode_E_Pronounced := func() {
		// special cases with two pronunciations
		// 'agape' 'lame' 'resume'
		if (StringAtStart("LAME", "SAKE", "PATE") && (m_length == 4)) ||
			(StringAtStart("AGAPE") && (m_length == 5)) ||
			((m_current == 5) && StringAtStart("RESUME")) {
			MetaphAdd("", "A")
			return
		}

		// special case "inge" => 'INGA', 'INJ'
		if StringAtStart("INGE") && (m_length == 4) {
			MetaphAdd("A", "")
			return
		}

		// special cases with two pronunciations
		// special handling due to the difference in
		// the pronunciation of the '-D'
		if (m_current == 5) && StringAtStart("BLESSED", "LEARNED") {
			MetaphAddExactApprox("D", "AD", "T", "AT")
			m_current += 2
			return
		}

		// encode all vowels and diphthongs to the same value
		if (!E_Silent() && !flag_AL_inversion && !Silent_Internal_E()) ||
			E_Pronounced_Exceptions() {
			MetaphAdd("A")
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
	Encode_Silent_B := func() bool {
		//'debt', 'doubt', 'subtle'
		if StringAt(-2, "DEBT") ||
			StringAt(-2, "SUBTL") ||
			StringAt(-2, "SUBTIL") ||
			StringAt(-3, "DOUBT") {
			MetaphAdd("T")
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
	Encode_Vowels := func() {
		if m_current == 0 {
			// all init vowels map to 'A'
			// as of Double Metaphone
			MetaphAdd("A")
		} else if m_encodeVowels {
			if CharAt(m_current) != 'E' {
				if Skip_Silent_UE() {
					return
				}

				if O_Silent() {
					m_current++
					return
				}

				// encode all vowels and
				// diphthongs to the same value
				MetaphAdd("A")
			} else {
				Encode_E_Pronounced()
			}
		}

		if !(!IsVowel(m_current-2) && StringAt(-1, "LEWA", "LEWO", "LEWI")) {
			m_current = SkipVowels(m_current)
		} else {
			m_current++
		}
	}

	/**
	 * Encodes 'B'
	 *
	 *
	 */
	Encode_B := func() {
		if Encode_Silent_B() {
			return
		}

		// "-mb", e.g", "dumb", already skipped over under
		// 'M', altho it should really be handled here...
		MetaphAddExactApprox("B", "P")

		if (CharAt(m_current+1) == 'B') ||
			((CharAt(m_current+1) == 'P') &&
				((m_current+1 < m_last) && (CharAt(m_current+2) != 'H'))) {
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
	Encode_CH_To_H := func() bool {
		// hebrew => 'H', e.g. 'channukah', 'chabad'
		if ((m_current == 0) &&
			(StringAt(+2, "AIM", "ETH", "ELM") ||
				StringAt(+2, "ASID", "AZAN") ||
				StringAt(+2, "UPPAH", "UTZPA", "ALLAH", "ALUTZ", "AMETZ") ||
				StringAt(+2, "ESHVAN", "ADARIM", "ANUKAH") ||
				StringAt(+2, "ALLLOTH", "ANNUKAH", "AROSETH"))) ||
			// and an irish name with the same encoding
			StringAt(-3, "CLACHAN") {
			MetaphAdd("H")
			AdvanceCounter(3, 2)
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
	Encode_Silent_C_At_Beginning := func() bool {
		//skip these when at start of word
		if (m_current == 0) && StringAt(0, "CT", "CN") {
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
	Encode_CA_To_S := func() bool {
		// Special case: 'caesar'.
		// Also, where cedilla not used, as in "linguica" => LNKS
		if ((m_current == 0) && StringAt(0, "CAES", "CAEC", "CAEM")) ||
			StringAtStart("FRANCAIS", "FRANCAIX", "LINGUICA") ||
			StringAtStart("FACADE") ||
			StringAtStart("GONCALVES", "PROVENCAL") {
			MetaphAdd("S")
			AdvanceCounter(2, 1)
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
	Encode_CO_To_S := func() bool {
		// e.g. 'coelecanth' => SLKN0
		if (StringAt(0, "COEL") &&
			(IsVowel(m_current+4) || ((m_current + 3) == m_last))) ||
			StringAt(0, "COENA", "COENO") ||
			StringAtStart("FRANCOIS", "MELANCON") ||
			StringAtStart("GARCON") {
			MetaphAdd("S")
			AdvanceCounter(3, 1)
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
	Encode_CHAE := func() bool {
		// e.g. 'michael'
		if (m_current > 0) && StringAt(+2, "AE") {
			if StringAtStart("RACHAEL") {
				MetaphAdd("X")
			} else if !StringAt(-1, "C", "K", "G", "Q") {
				MetaphAdd("K")
			}

			AdvanceCounter(4, 2)
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
	Encode_Silent_CH := func() bool {
		// '-ch-' not pronounced
		if StringAt(-2, "FUCHSIA") ||
			StringAt(-2, "YACHT") ||
			StringAtStart("STRACHAN") ||
			StringAtStart("CRICHTON") ||
			(StringAt(-3, "DRACHM")) &&
				!StringAt(-3, "DRACHMA") {
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
	Encode_CH_To_X := func() bool {
		// e.g. 'approach', 'beach'
		if (StringAt(-2, "OACH", "EACH", "EECH", "OUCH", "OOCH", "MUCH", "SUCH") &&
			!StringAt(-3, "JOACH")) ||
			// e.g. 'dacha', 'macho'
			(((m_current + 2) == m_last) && StringAt(-1, "ACHA", "ACHO")) ||
			(StringAt(0, "CHOT", "CHOD", "CHAT") && ((m_current + 3) == m_last)) ||
			((StringAt(-1, "OCHE") && ((m_current + 2) == m_last)) &&
				!StringAt(-2, "DOCHE")) ||
			StringAt(-4, "ATTACH", "DETACH", "KOVACH") ||
			StringAt(-5, "SPINACH") ||
			StringAtStart("MACHAU") ||
			StringAt(-4, "PARACHUT") ||
			StringAt(-5, "MASSACHU") ||
			(StringAt(-3, "THACH") && !StringAt(-1, "ACHE")) ||
			StringAt(-2, "VACHON") {
			MetaphAdd("X")
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
	Encode_English_CH_To_K := func() bool {
		//'ache', 'echo', alternate spelling of 'michael'
		if ((m_current == 1) && RootOrInflections(m_inWord, "ACHE")) ||
			(((m_current > 3) && RootOrInflections(m_inWord[m_current-1:], "ACHE")) &&
				(StringAtStart("EAR") ||
					StringAtStart("HEAD", "BACK") ||
					StringAtStart("HEART", "BELLY", "TOOTH"))) ||
			StringAt(-1, "ECHO") ||
			StringAt(-2, "MICHEAL") ||
			StringAt(-4, "JERICHO") ||
			StringAt(-5, "LEPRECH") {
			MetaphAdd("K", "X")
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
	Encode_Germanic_CH_To_K := func() bool {
		// various germanic
		// "<consonant><vowel>CH-" implies a german word where 'ch' => K
		if ((m_current > 1) &&
			!IsVowel(m_current-2) &&
			StringAt(-1, "ACH") &&
			!StringAt(-2, "MACHADO", "MACHUCA", "LACHANC", "LACHAPE", "KACHATU") &&
			!StringAt(-3, "KHACHAT") &&
			((CharAt(m_current+2) != 'I') &&
				((CharAt(m_current+2) != 'E') ||
					StringAt(-2, "BACHER", "MACHER", "MACHEN", "LACHER"))) ||
			// e.g. 'brecht', 'fuchs'
			(StringAt(+2, "T", "S") &&
				!(StringAtStart("WHICHSOEVER") || StringAtStart("LUNCHTIME"))) ||
			// e.g. 'andromache'
			StringAtStart("SCHR") ||
			((m_current > 2) && StringAt(-2, "MACHE")) ||
			((m_current == 2) && StringAt(-2, "ZACH")) ||
			StringAt(-4, "SCHACH") ||
			StringAt(-1, "ACHEN") ||
			StringAt(-3, "SPICH", "ZURCH", "BUECH") ||
			(StringAt(-3, "KIRCH", "JOACH", "BLECH", "MALCH") &&
				// "kirch" and "blech" both get 'X'
				!(StringAt(-3, "KIRCHNER") || ((m_current + 1) == m_last))) ||
			(((m_current + 1) == m_last) && StringAt(-2, "NICH", "LICH", "BACH")) ||
			(((m_current + 1) == m_last) &&
				StringAt(-3, "URICH", "BRICH", "ERICH", "DRICH", "NRICH") &&
				!StringAt(-5, "ALDRICH") &&
				!StringAt(-6, "GOODRICH") &&
				!StringAt(-7, "GINGERICH"))) ||
			(((m_current + 1) == m_last) && StringAt(-4, "ULRICH", "LFRICH", "LLRICH",
				"EMRICH", "ZURICH", "EYRICH")) ||
			// e.g., 'wachtler', 'wechsler', but not 'tichner'
			((StringAt(-1, "A", "O", "U", "E") || (m_current == 0)) &&
				StringAt(+2, "L", "R", "N", "M", "B", "H", "F", "V", "W", " ")) {
			// "CHR/L-" e.g. 'chris' do not get
			// alt pronunciation of 'X'
			if StringAt(+2, "R", "L") || SlavoGermanic() {
				MetaphAdd("K")
			} else {
				MetaphAdd("K", "X")
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
	Encode_ARCH := func() bool {
		if StringAt(-2, "ARCH") {
			// "-ARCH-" has many combining forms where "-CH-" => K because of its
			// derivation from the greek
			if ((IsVowel(m_current+2) && StringAt(-2, "ARCHA", "ARCHI", "ARCHO", "ARCHU", "ARCHY")) ||
				StringAt(-2, "ARCHEA", "ARCHEG", "ARCHEO", "ARCHET", "ARCHEL", "ARCHES", "ARCHEP",
					"ARCHEM", "ARCHEN") ||
				(StringAt(-2, "ARCH") && ((m_current + 1) == m_last)) ||
				StringAtStart("MENARCH")) &&
				(!RootOrInflections(m_inWord, "ARCH") &&
					!StringAt(-4, "SEARCH", "POARCH") &&
					!StringAtStart("ARCHENEMY", "ARCHIBALD", "ARCHULETA", "ARCHAMBAU") &&
					!StringAtStart("ARCHER", "ARCHIE") &&
					!((((StringAt(-3, "LARCH", "MARCH", "PARCH") ||
						StringAt(-4, "STARCH")) &&
						!(StringAtStart("EPARCH") ||
							StringAtStart("NOMARCH") ||
							StringAtStart("EXILARCH", "HIPPARCH", "MARCHESE") ||
							StringAtStart("ARISTARCH") ||
							StringAtStart("MARCHETTI"))) ||
						RootOrInflections(m_inWord, "STARCH")) &&
						(!StringAt(-2, "ARCHU", "ARCHY") ||
							StringAtStart("STARCHY")))) {
				MetaphAdd("K", "X")
			} else {
				MetaphAdd("X")
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
	Encode_Greek_CH_Initial := func() bool {
		// greek roots e.g. 'chemistry', 'chorus', ch at beginning of root
		if (StringAt(0, "CHAMOM", "CHARAC", "CHARIS", "CHARTO", "CHARTU", "CHARYB", "CHRIST", "CHEMIC", "CHILIA") ||
			(StringAt(0, "CHEMI", "CHEMO", "CHEMU", "CHEMY", "CHOND", "CHONA", "CHONI", "CHOIR", "CHASM",
				"CHARO", "CHROM", "CHROI", "CHAMA", "CHALC", "CHALD", "CHAET", "CHIRO", "CHILO", "CHELA", "CHOUS",
				"CHEIL", "CHEIR", "CHEIM", "CHITI", "CHEOP") &&
				!(StringAt(0, "CHEMIN") || StringAt(-2, "ANCHONDO"))) ||
			(StringAt(0, "CHISM", "CHELI") &&
				// exclude spanish "machismo"
				!(StringAtStart("MACHISMO") ||
					// exclude some french words
					StringAtStart("REVANCHISM") ||
					StringAtStart("RICHELIEU") ||
					(StringAtStart("CHISM") && (m_length == 5)) ||
					StringAtStart("MICHEL"))) ||
			// include e.g. "chorus", "chyme", "chaos"
			(StringAt(0, "CHOR", "CHOL", "CHYM", "CHYL", "CHLO", "CHOS", "CHUS", "CHOE") &&
				!StringAtStart("CHOLLO", "CHOLLA", "CHORIZ")) ||
			// "chaos" => K but not "chao"
			(StringAt(0, "CHAO") && ((m_current + 3) != m_last)) ||
			// e.g. "abranchiate"
			(StringAt(0, "CHIA") && !(StringAtStart("APPALACHIA") || StringAtStart("CHIAPAS"))) ||
			// e.g. "chimera"
			StringAt(0, "CHIMERA", "CHIMAER", "CHIMERI") ||
			// e.g. "chameleon"
			((m_current == 0) && StringAt(0, "CHAME", "CHELO", "CHITO")) ||
			// e.g. "spirochete"
			((((m_current + 4) == m_last) || ((m_current + 5) == m_last)) && StringAt(-1, "OCHETE"))) &&
			// more exceptions where "-CH-" => X e.g. "chortle", "crocheter"
			!((StringAtStart("CHORE", "CHOLO", "CHOLA") && (m_length == 5)) ||
				StringAt(0, "CHORT", "CHOSE") ||
				StringAt(-3, "CROCHET") ||
				StringAtStart("CHEMISE", "CHARISE", "CHARISS", "CHAROLE")) {
			// "CHR/L-" e.g. 'christ', 'chlorine' do not get
			// alt pronunciation of 'X'
			if StringAt(+2, "R", "L") {
				MetaphAdd("K")
			} else {
				MetaphAdd("K", "X")
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
	Encode_Greek_CH_Non_Initial := func() bool {
		//greek & other roots e.g. 'tachometer', 'orchid', ch in middle or end of root
		if StringAt(-2, "ORCHID", "NICHOL", "MECHAN", "LICHEN", "MACHIC", "PACHEL", "RACHIF", "RACHID",
			"RACHIS", "RACHIC", "MICHAL") ||
			StringAt(-3, "MELCH", "GLOCH", "TRACH", "TROCH", "BRACH", "SYNCH", "PSYCH",
				"STICH", "PULCH", "EPOCH") ||
			(StringAt(-3, "TRICH") && !StringAt(-5, "OSTRICH")) ||
			(StringAt(-2, "TYCH", "TOCH", "BUCH", "MOCH", "CICH", "DICH", "NUCH", "EICH", "LOCH",
				"DOCH", "ZECH", "WYCH") &&
				!(StringAt(-4, "INDOCHINA") || StringAt(-2, "BUCHON"))) ||
			StringAt(-2, "LYCHN", "TACHO", "ORCHO", "ORCHI", "LICHO") ||
			(StringAt(-1, "OCHER", "ECHIN", "ECHID") && ((m_current == 1) || (m_current == 2))) ||
			StringAt(-4, "BRONCH", "STOICH", "STRYCH", "TELECH", "PLANCH", "CATECH", "MANICH", "MALACH",
				"BIANCH", "DIDACH") ||
			(StringAt(-1, "ICHA", "ICHN") && (m_current == 1)) ||
			StringAt(-2, "ORCHESTR") ||
			StringAt(-4, "BRANCHIO", "BRANCHIF") ||
			(StringAt(-1, "ACHAB", "ACHAD", "ACHAN", "ACHAZ") &&
				!StringAt(-2, "MACHADO", "LACHANC")) ||
			StringAt(-1, "ACHISH", "ACHILL", "ACHAIA", "ACHENE") ||
			StringAt(-1, "ACHAIAN", "ACHATES", "ACHIRAL", "ACHERON") ||
			StringAt(-1, "ACHILLEA", "ACHIMAAS", "ACHILARY", "ACHELOUS", "ACHENIAL", "ACHERNAR") ||
			StringAt(-1, "ACHALASIA", "ACHILLEAN", "ACHIMENES") ||
			StringAt(-1, "ACHIMELECH", "ACHITOPHEL") ||
			// e.g. 'inchoate'
			(((m_current - 2) == 0) && (StringAt(-2, "INCHOA") ||
				// e.g. 'ischemia'
				StringAtStart("ISCH"))) ||
			// e.g. 'ablimelech', 'antioch', 'pentateuch'
			(((m_current + 1) == m_last) && StringAt(-1, "A", "O", "U", "E") &&
				!(StringAtStart("DEBAUCH") ||
					StringAt(-2, "MUCH", "SUCH", "KOCH") ||
					StringAt(-5, "OODRICH", "ALDRICH"))) {
			MetaphAdd("K", "X")
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
	Encode_CH := func() bool {
		if StringAt(0, "CH") {
			if Encode_CHAE() ||
				Encode_CH_To_H() ||
				Encode_Silent_CH() ||
				Encode_ARCH() ||
				// Encode_CH_To_X() should be
				// called before the germanic
				// and greek encoding functions
				Encode_CH_To_X() ||
				Encode_English_CH_To_K() ||
				Encode_Germanic_CH_To_K() ||
				Encode_Greek_CH_Initial() ||
				Encode_Greek_CH_Non_Initial() {
				return true
			}

			if m_current > 0 {
				if StringAtStart("MC") && (m_current == 1) {
					//e.g., "McHugh"
					MetaphAdd("K")
				} else {
					MetaphAdd("X", "K")
				}
			} else {
				MetaphAdd("X")
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
	Encode_CCIA := func() bool {
		//e.g., 'focaccia'
		if StringAt(+1, "CIA") {
			MetaphAdd("X", "S")
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
	Encode_CC := func() bool {
		//double 'C', but not if e.g. 'McClellan'
		if StringAt(0, "CC") && !((m_current == 1) && (CharAt(0) == 'M')) {
			// exception
			if StringAt(-3, "FLACCID") {
				MetaphAdd("S")
				AdvanceCounter(3, 2)
				return true
			}

			//'bacci', 'bertucci', other italian
			if (((m_current + 2) == m_last) && StringAt(+2, "I")) ||
				StringAt(+2, "IO") ||
				(((m_current + 4) == m_last) && StringAt(+2, "INO", "INI")) {
				MetaphAdd("X")
				AdvanceCounter(3, 2)
				return true
			}

			//'accident', 'accede' 'succeed'
			if StringAt(+2, "I", "E", "Y") &&
				//except 'bellocchio','bacchus', 'soccer' get K
				!((CharAt(m_current+2) == 'H') ||
					StringAt(-2, "SOCCER")) {
				MetaphAdd("KS")
				AdvanceCounter(3, 2)
				return true

			} else {
				//Pierce's rule
				MetaphAdd("K")
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
	Encode_CK_CG_CQ := func() bool {
		if StringAt(0, "CK", "CG", "CQ") {
			// eastern european spelling e.g. 'gorecki' == 'goresky'
			if StringAt(0, "CKI", "CKY") &&
				((m_current + 2) == m_last) &&
				(m_length > 6) {
				MetaphAdd("K", "SK")
			} else {
				MetaphAdd("K")
			}
			m_current += 2

			if StringAt(0, "K", "G", "Q") {
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
	Encode_British_Silent_CE := func() bool {
		// english place names like e.g.'gloucester' pronounced glo-ster
		return (StringAt(+1, "ESTER") && ((m_current + 5) == m_last)) ||
			StringAt(+1, "ESTERSHIRE")
	}

	/**
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_CE := func() bool {
		// 'ocean', 'commercial', 'provincial', 'cello', 'fettucini', 'medici'
		if (StringAt(+1, "EAN") && IsVowel(m_current-1)) ||
			// e.g. 'rosacea'
			(StringAt(-1, "ACEA") &&
				((m_current + 2) == m_last) &&
				!StringAtStart("PANACEA")) ||
			// e.g. 'botticelli', 'concerto'
			StringAt(+1, "ELLI", "ERTO", "EORL") ||
			// some italian names familiar to americans
			(StringAt(-3, "CROCE") && ((m_current + 1) == m_last)) ||
			StringAt(-3, "DOLCE") ||
			// e.g. 'cello'
			(StringAt(+1, "ELLO") &&
				((m_current + 4) == m_last)) {
			MetaphAdd("X", "S")
			return true
		}

		return false
	}

	/**
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_CI := func() bool {
		// with consonant before C
		// e.g. 'fettucini', but exception for the americanized pronunciation of 'mancini'
		if ((StringAt(+1, "INI") && !StringAtStart("MANCINI")) && ((m_current + 3) == m_last)) ||
			// e.g. 'medici'
			(StringAt(-1, "ICI") && ((m_current + 1) == m_last)) ||
			// e.g. "commercial', 'provincial', 'cistercian'
			StringAt(-1, "RCIAL", "NCIAL", "RCIAN", "UCIUS") ||
			// special cases
			StringAt(-3, "MARCIA") ||
			StringAt(-2, "ANCIENT") {
			MetaphAdd("X", "S")
			return true
		}

		// with vowel before C (or at beginning?)
		if ((StringAt(0, "CIO", "CIE", "CIA") &&
			IsVowel(m_current-1)) ||
			// e.g. "ciao"
			StringAt(+1, "IAO")) &&
			!StringAt(-4, "COERCION") {
			if (StringAt(0, "CIAN", "CIAL", "CIAO", "CIES", "CIOL", "CION") ||
				// exception - "glacier" => 'X' but "spacier" = > 'S'
				StringAt(-3, "GLACIER") ||
				StringAt(0, "CIENT", "CIENC", "CIOUS", "CIATE", "CIATI", "CIATO", "CIABL", "CIARY") ||
				(((m_current + 2) == m_last) && StringAt(0, "CIA", "CIO")) ||
				(((m_current + 3) == m_last) && StringAt(0, "CIAS", "CIOS"))) &&
				// exceptions
				!(StringAt(-4, "ASSOCIATION") ||
					StringAtStart("OCIE") ||
					// exceptions mostly because these names are usually from
					// the spanish rather than the italian in america
					StringAt(-2, "LUCIO") ||
					StringAt(-2, "MACIAS") ||
					StringAt(-3, "GRACIE", "GRACIA") ||
					StringAt(-2, "LUCIANO") ||
					StringAt(-3, "MARCIANO") ||
					StringAt(-4, "PALACIO") ||
					StringAt(-4, "FELICIANO") ||
					StringAt(-5, "MAURICIO") ||
					StringAt(-7, "ENCARNACION") ||
					StringAt(-4, "POLICIES") ||
					StringAt(-2, "HACIENDA") ||
					StringAt(-6, "ANDALUCIA") ||
					StringAt(-2, "SOCIO", "SOCIE")) {
				MetaphAdd("X", "S")
			} else {
				MetaphAdd("S", "X")
			}

			return true
		}

		// exception
		if StringAt(-4, "COERCION") {
			MetaphAdd("J")
			return true
		}

		return false
	}

	/**
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Latinate_Suffixes := func() bool {
		if StringAt(+1, "EOUS", "IOUS") {
			MetaphAdd("X", "S")
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
	Encode_C_Front_Vowel := func() bool {
		if StringAt(0, "CI", "CE", "CY") {
			if Encode_British_Silent_CE() ||
				Encode_CE() ||
				Encode_CI() ||
				Encode_Latinate_Suffixes() {
				AdvanceCounter(2, 1)
				return true
			}

			MetaphAdd("S")
			AdvanceCounter(2, 1)
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
	Encode_Silent_C := func() bool {
		if StringAt(+1, "T", "S") {
			if StringAtStart("CONNECTICUT") ||
				StringAtStart("INDICT", "TUCSON") {
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
	Encode_CZ := func() bool {
		if StringAt(+1, "Z") &&
			!StringAt(-1, "ECZEMA") {
			if StringAt(0, "CZAR") {
				MetaphAdd("S")
			} else {
				// otherwise most likely a czech word...
				MetaphAdd("X")
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
	Encode_CS := func() bool {
		// give an 'etymological' 2nd
		// encoding for "kovacs" so
		// that it matches "kovach"
		if StringAtStart("KOVACS") {
			MetaphAdd("KS", "X")
			m_current += 2
			return true
		}

		if StringAt(-1, "ACS") &&
			((m_current + 1) == m_last) &&
			!StringAt(-4, "ISAACS") {
			MetaphAdd("X")
			m_current += 2
			return true
		}

		return false
	}

	/**
	 * Encodes 'C'
	 *
	 */
	Encode_C := func() {
		if Encode_Silent_C_At_Beginning() ||
			Encode_CA_To_S() ||
			Encode_CO_To_S() ||
			Encode_CH() ||
			Encode_CCIA() ||
			Encode_CC() ||
			Encode_CK_CG_CQ() ||
			Encode_C_Front_Vowel() ||
			Encode_Silent_C() ||
			Encode_CZ() ||
			Encode_CS() {
			return
		}

		//else
		if !StringAt(-1, "C", "K", "G", "Q") {
			MetaphAdd("K")
		}

		//name sent in 'mac caffrey', 'mac gregor
		if StringAt(+1, " C", " Q", " G") {
			m_current += 2
		} else {
			if StringAt(+1, "C", "K", "Q") &&
				!StringAt(+1, "CE", "CI") {
				m_current += 2
				// account for combinations such as Ro-ckc-liffe
				if StringAt(0, "C", "K", "Q") &&
					!StringAt(+1, "CE", "CI") {
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
	Encode_DG := func() bool {
		if StringAt(0, "DG") {
			// excludes exceptions e.g. 'edgar',
			// or cases where 'g' is first letter of combining form
			// e.g. 'handgun', 'waldglas'
			if StringAt(+2, "A", "O") ||
				// e.g. "midgut"
				StringAt(+1, "GUN", "GUT") ||
				// e.g. "handgrip"
				StringAt(+1, "GEAR", "GLAS", "GRIP", "GREN", "GILL", "GRAF") ||
				// e.g. "mudgard"
				StringAt(+1, "GUARD", "GUILT", "GRAVE", "GRASS") ||
				// e.g. "woodgrouse"
				StringAt(+1, "GROUSE") {
				MetaphAddExactApprox("DG", "TK")
			} else {
				//e.g. "edge", "abridgment"
				MetaphAdd("J")
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
	Encode_DJ := func() bool {
		// e.g. "adjacent"
		if StringAt(0, "DJ") {
			MetaphAdd("J")
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
	Encode_DT_DD := func() bool {
		// eat redundant 'T' or 'D'
		if StringAt(0, "DT", "DD") {
			if StringAt(0, "DTH") {
				MetaphAddExactApprox("D0", "T0")
				m_current += 3
			} else {
				if m_encodeExact {
					// devoice it
					if StringAt(0, "DT") {
						MetaphAdd("T")
					} else {
						MetaphAdd("D")
					}
				} else {
					MetaphAdd("T")
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
	Encode_D_To_J := func() bool {
		// e.g. "module", "adulate"
		if (StringAt(0, "DUL") &&
			(IsVowel(m_current-1) && IsVowel(m_current+3))) ||
			// e.g. "soldier", "grandeur", "procedure"
			(((m_current + 3) == m_last) &&
				StringAt(-1, "LDIER", "NDEUR", "EDURE", "RDURE")) ||
			StringAt(-3, "CORDIAL") ||
			// e.g.  "pendulum", "education"
			StringAt(-1, "NDULA", "NDULU", "EDUCA") ||
			// e.g. "individual", "individual", "residuum"
			StringAt(-1, "ADUA", "IDUA", "IDUU") {
			MetaphAddExactApprox("J", "D", "J", "T")
			AdvanceCounter(2, 1)
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
	Encode_DOUS := func() bool {
		// e.g. "assiduous", "arduous"
		if StringAt(+1, "UOUS") {
			MetaphAddExactApprox("J", "D", "J", "T")
			AdvanceCounter(4, 1)
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
	Encode_Silent_D := func() bool {
		// silent 'D' e.g. 'wednesday', 'handsome'
		if StringAt(-2, "WEDNESDAY") ||
			StringAt(-3, "HANDKER", "HANDSOM", "WINDSOR") ||
			// french silent D at end in words or names familiar to americans
			StringAt(-5, "PERNOD", "ARTAUD", "RENAUD") ||
			StringAt(-6, "RIMBAUD", "MICHAUD", "BICHAUD") {
			m_current++
			return true
		}

		return false
	}

	/**
	 * Encode "-D-"
	 *
	 */
	Encode_D := func() {
		if Encode_DG() ||
			Encode_DJ() ||
			Encode_DT_DD() ||
			Encode_D_To_J() ||
			Encode_DOUS() ||
			Encode_Silent_D() {
			return
		}

		if m_encodeExact {
			// "final de-voicing" in this case
			// e.g. 'missed' == 'mist'
			if (m_current == m_last) &&
				StringAt(-3, "SSED") {
				MetaphAdd("T")
			} else {
				MetaphAdd("D")
			}
		} else {
			MetaphAdd("T")
		}
		m_current++
	}

	/**
	 * Encode "-F-"
	 *
	 */
	Encode_F := func() {
		// Encode cases where "-FT-" => "T" is usually silent
		// e.g. 'often', 'soften'
		// This should really be covered under "T"!
		if StringAt(-1, "OFTEN") {
			MetaphAdd("F", "FT")
			m_current += 2
			return
		}

		// eat redundant 'F'
		if CharAt(m_current+1) == 'F' {
			m_current += 2
		} else {
			m_current++
		}

		MetaphAdd("F")

	}

	/**
	 * Encode cases where 'G' is silent at beginning of word
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Silent_G_At_Beginning := func() bool {
		//skip these when at start of word
		if (m_current == 0) &&
			StringAt(0, "GN") {
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
	Encode_GG := func() bool {
		if CharAt(m_current+1) == 'G' {
			// italian e.g, 'loggia', 'caraveggio', also 'suggest' and 'exaggerate'
			if StringAt(-1, "AGGIA", "OGGIA", "AGGIO", "EGGIO", "EGGIA", "IGGIO") ||
				// 'ruggiero' but not 'snuggies'
				(StringAt(-1, "UGGIE") && !(((m_current + 3) == m_last) || ((m_current + 4) == m_last))) ||
				(((m_current + 2) == m_last) && StringAt(-1, "AGGI", "OGGI")) ||
				StringAt(-2, "SUGGES", "XAGGER", "REGGIE") {
				// expection where "-GG-" => KJ
				if StringAt(-2, "SUGGEST") {
					MetaphAddExactApprox("G", "K")
				}

				MetaphAdd("J")
				AdvanceCounter(3, 2)
			} else {
				MetaphAddExactApprox("G", "K")
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
	Encode_GK := func() bool {
		// 'gingko'
		if CharAt(m_current+1) == 'K' {
			MetaphAdd("K")
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
	Encode_GH_After_Consonant := func() bool {
		// e.g. 'burgher', 'bingham'
		if (m_current > 0) &&
			!IsVowel(m_current-1) &&
			// not e.g. 'greenhalgh'
			!(StringAt(-3, "HALGH") &&
				((m_current + 1) == m_last)) {
			MetaphAddExactApprox("G", "K")
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
	Encode_Initial_GH := func() bool {
		if m_current < 3 {
			// e.g. "ghislane", "ghiradelli"
			if m_current == 0 {
				if CharAt(m_current+2) == 'I' {
					MetaphAdd("J")
				} else {
					MetaphAddExactApprox("G", "K")
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
	Encode_GH_To_J := func() bool {
		// e.g., 'greenhalgh', 'dunkenhalgh', english names
		if StringAt(-2, "ALGH") && ((m_current + 1) == m_last) {
			MetaphAdd("J", "")
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
	Encode_GH_To_H := func() bool {
		// special cases
		// e.g., 'donoghue', 'donaghy'
		if (StringAt(-4, "DONO", "DONA") && IsVowel(m_current+2)) ||
			StringAt(-5, "CALLAGHAN") {
			MetaphAdd("H")
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
	Encode_UGHT := func() bool {
		//e.g. "ought", "aught", "daughter", "slaughter"
		if StringAt(-1, "UGHT") {
			if (StringAt(-3, "LAUGH") &&
				!(StringAt(-4, "SLAUGHT") ||
					StringAt(-3, "LAUGHTO"))) ||
				StringAt(-4, "DRAUGH") {
				MetaphAdd("FT")
			} else {
				MetaphAdd("T")
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
	Encode_GH_H_Part_Of_Other_Word := func() bool {
		// if the 'H' is the beginning of another word or syllable
		if StringAt(+1, "HOUS", "HEAD", "HOLE", "HORN", "HARN") {
			MetaphAddExactApprox("G", "K")
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
	Encode_Silent_GH := func() bool {
		//Parker's rule (with some further refinements) - e.g., 'hugh'
		if ((((m_current > 1) && StringAt(-2, "B", "H", "D", "G", "L")) ||
			//e.g., 'bough'
			((m_current > 2) &&
				StringAt(-3, "B", "H", "D", "K", "W", "N", "P", "V") &&
				!StringAtStart("ENOUGH")) ||
			//e.g., 'broughton'
			((m_current > 3) && StringAt(-4, "B", "H")) ||
			//'plough', 'slaugh'
			((m_current > 3) && StringAt(-4, "PL", "SL")) ||
			((m_current > 0) &&
				// 'sigh', 'light'
				((CharAt(m_current-1) == 'I') ||
					StringAtStart("PUGH") ||
					// e.g. 'MCDONAGH', 'MURTAGH', 'CREAGH'
					(StringAt(-1, "AGH") &&
						((m_current + 1) == m_last)) ||
					StringAt(-4, "GERAGH", "DRAUGH") ||
					(StringAt(-3, "GAUGH", "GEOGH", "MAUGH") &&
						!StringAtStart("MCGAUGHEY")) ||
					// exceptions to 'tough', 'rough', 'lough'
					(StringAt(-2, "OUGH") &&
						(m_current > 3) &&
						!StringAt(-4, "CCOUGH", "ENOUGH", "TROUGH", "CLOUGH"))))) &&
			// suffixes starting w/ vowel where "-GH-" is usually silent
			(StringAt(-3, "VAUGH", "FEIGH", "LEIGH") ||
				StringAt(-2, "HIGH", "TIGH") ||
				((m_current + 1) == m_last) ||
				(StringAt(+2, "IE", "EY", "ES", "ER", "ED", "TY") &&
					((m_current + 3) == m_last) &&
					!StringAt(-5, "GALLAGHER")) ||
				(StringAt(+2, "Y") && ((m_current + 2) == m_last)) ||
				(StringAt(+2, "ING", "OUT") && ((m_current + 4) == m_last)) ||
				(StringAt(+2, "ERTY") && ((m_current + 5) == m_last)) ||
				(!IsVowel(m_current+2) ||
					StringAt(-3, "GAUGH", "GEOGH", "MAUGH") ||
					StringAt(-4, "BROUGHAM")))) &&
			// exceptions where '-g-' pronounced
			!(StringAtStart("BALOGH", "SABAGH") ||
				StringAt(-2, "BAGHDAD") ||
				StringAt(-3, "WHIGH") ||
				StringAt(-5, "SABBAGH", "AKHLAGH")) {
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
	Encode_GH_Special_Cases := func() bool {
		var handled = false

		// special case: 'hiccough' == 'hiccup'
		if StringAt(-6, "HICCOUGH") {
			MetaphAdd("P")
			handled = true
		} else if StringAtStart("LOUGH") {
			// special case: 'lough' alt spelling for scots 'loch'
			MetaphAdd("K")
			handled = true
		} else if StringAtStart("BALOGH") {
			// hungarian
			MetaphAddExactApprox("G", "", "K", "")
			handled = true
		} else if StringAt(-3, "LAUGHLIN", "COUGHLAN", "LOUGHLIN") {
			// "maclaughlin"
			MetaphAdd("K", "F")
			handled = true
		} else if StringAt(-3, "GOUGH") ||
			StringAt(-7, "COLCLOUGH") {
			MetaphAdd("", "F")
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
	Encode_GH_To_F := func() bool {
		// the cases covered here would fall under
		// the GH_To_F rule below otherwise
		if Encode_GH_Special_Cases() {
			return true
		} else {
			//e.g., 'laugh', 'cough', 'rough', 'tough'
			if (m_current > 2) &&
				(CharAt(m_current-1) == 'U') &&
				IsVowel(m_current-2) &&
				StringAt(-3, "C", "G", "L", "R", "T", "N", "S") &&
				!StringAt(-4, "BREUGHEL", "FLAUGHER") {
				MetaphAdd("F")
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
	Encode_GH := func() bool {
		if CharAt(m_current+1) == 'H' {
			if Encode_GH_After_Consonant() ||
				Encode_Initial_GH() ||
				Encode_GH_To_J() ||
				Encode_GH_To_H() ||
				Encode_UGHT() ||
				Encode_GH_H_Part_Of_Other_Word() ||
				Encode_Silent_GH() ||
				Encode_GH_To_F() {
				return true
			}

			MetaphAddExactApprox("G", "K")
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
	Encode_Silent_G := func() bool {
		// e.g. "phlegm", "apothegm", "voigt"
		if (((m_current + 1) == m_last) &&
			(StringAt(-1, "EGM", "IGM", "AGM") ||
				StringAt(0, "GT"))) ||
			(StringAtStart("HUGES") && (m_length == 5)) {
			m_current++
			return true
		}

		// vietnamese names e.g. "Nguyen" but not "Ng"
		if StringAtStart("NG") && (m_current != m_last) {
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
	Encode_GN := func() bool {
		if CharAt(m_current+1) == 'N' {
			// 'align' 'sign', 'resign' but not 'resignation'
			// also 'impugn', 'impugnable', but not 'repugnant'
			if ((m_current > 1) &&
				((StringAt(-1, "I", "U", "E") ||
					StringAt(-3, "LORGNETTE") ||
					StringAt(-2, "LAGNIAPPE") ||
					StringAt(-2, "COGNAC") ||
					StringAt(-3, "CHAGNON") ||
					StringAt(-5, "COMPAGNIE") ||
					StringAt(-4, "BOLOGN")) &&
					// Exceptions: following are cases where 'G' is pronounced
					// in "assign" 'g' is silent, but not in "assignation"
					!(StringAt(+2, "ATION") ||
						StringAt(+2, "ATOR") ||
						StringAt(+2, "ATE", "ITY") ||
						// exception to exceptions, not pronounced:
						(StringAt(+2, "AN", "AC", "IA", "UM") &&
							!(StringAt(-3, "POIGNANT") ||
								StringAt(-2, "COGNAC"))) ||
						StringAtStart("SPIGNER", "STEGNER") ||
						(StringAtStart("SIGNE") && (m_length == 5)) ||
						StringAt(-2, "LIGNI", "LIGNO", "REGNA", "DIGNI", "WEGNE",
							"TIGNE", "RIGNE", "REGNE", "TIGNO") ||
						StringAt(-2, "SIGNAL", "SIGNIF", "SIGNAT") ||
						StringAt(-1, "IGNIT")) &&
					!StringAt(-2, "SIGNET", "LIGNEO"))) ||
				//not e.g. 'cagney', 'magna'
				(((m_current + 2) == m_last) &&
					StringAt(0, "GNE", "GNA") &&
					!StringAt(-2, "SIGNA", "MAGNA", "SIGNE")) {
				MetaphAddExactApprox("N", "GN", "N", "KN")
			} else {
				MetaphAddExactApprox("GN", "KN")
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
	Encode_GL := func() bool {
		//'tagliaro', 'puglia' BUT add K in alternative
		// since americans sometimes do this
		if StringAt(+1, "LIA", "LIO", "LIE") &&
			IsVowel(m_current-1) {
			MetaphAddExactApprox("L", "GL", "L", "KL")
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
	Initial_G_Soft := func() bool {
		if ((StringAt(+1, "EL", "EM", "EN", "EO", "ER", "ES", "IA", "IN", "IO", "IP", "IU", "YM", "YN", "YP", "YR", "EE") ||
			StringAt(+1, "IRA", "IRO")) &&
			// except for smaller set of cases where => K, e.g. "gerber"
			!(StringAt(+1, "ELD", "ELT", "ERT", "INZ", "ERH", "ITE", "ERD", "ERL", "ERN",
				"INT", "EES", "EEK", "ELB", "EER") ||
				StringAt(+1, "ERSH", "ERST", "INSB", "INGR", "EROW", "ERKE", "EREN") ||
				StringAt(+1, "ELLER", "ERDIE", "ERBER", "ESUND", "ESNER", "INGKO", "INKGO",
					"IPPER", "ESELL", "IPSON", "EEZER", "ERSON", "ELMAN") ||
				StringAt(+1, "ESTALT", "ESTAPO", "INGHAM", "ERRITY", "ERRISH", "ESSNER", "ENGLER") ||
				StringAt(+1, "YNAECOL", "YNECOLO", "ENTHNER", "ERAGHTY") ||
				StringAt(+1, "INGERICH", "EOGHEGAN"))) ||
			(IsVowel(m_current+1) &&
				(StringAt(+1, "EE ", "EEW") ||
					(StringAt(+1, "IGI", "IRA", "IBE", "AOL", "IDE", "IGL") &&
						!StringAt(+1, "IDEON")) ||
					StringAt(+1, "ILES", "INGI", "ISEL") ||
					(StringAt(+1, "INGER") && !StringAt(+1, "INGERICH")) ||
					StringAt(+1, "IBBER", "IBBET", "IBLET", "IBRAN", "IGOLO", "IRARD", "IGANT") ||
					StringAt(+1, "IRAFFE", "EEWHIZ") ||
					StringAt(+1, "ILLETTE", "IBRALTA"))) {
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
	Encode_Initial_G_Front_Vowel := func() bool {
		// 'g' followed by vowel at beginning
		if (m_current == 0) && Front_Vowel(m_current+1) {
			// special case "gila" as in "gila monster"
			if StringAt(+1, "ILA") && (m_length == 4) {
				MetaphAdd("H")
			} else if Initial_G_Soft() {
				MetaphAddExactApprox("J", "G", "J", "K")
			} else {
				// only code alternate 'J' if front vowel
				if (m_inWord[m_current+1] == 'E') || (m_inWord[m_current+1] == 'I') {
					MetaphAddExactApprox("G", "J", "K", "J")
				} else {
					MetaphAddExactApprox("G", "K")
				}
			}

			AdvanceCounter(2, 1)
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
	Encode_NGER := func() bool {
		if (m_current > 1) &&
			StringAt(-1, "NGER") {
			// default 'G' => J  such as 'ranger', 'stranger', 'manger', 'messenger', 'orangery', 'granger'
			// 'boulanger', 'challenger', 'danger', 'changer', 'harbinger', 'lounger', 'ginger', 'passenger'
			// except for these the following
			if !(RootOrInflections(m_inWord, "ANGER") ||
				RootOrInflections(m_inWord, "LINGER") ||
				RootOrInflections(m_inWord, "MALINGER") ||
				RootOrInflections(m_inWord, "FINGER") ||
				(StringAt(-3, "HUNG", "FING", "BUNG", "WING", "RING", "DING", "ZENG",
					"ZING", "JUNG", "LONG", "PING", "CONG", "MONG", "BANG",
					"GANG", "HANG", "LANG", "SANG", "SING", "WANG", "ZANG") &&
					// exceptions to above where 'G' => J
					!(StringAt(-6, "BOULANG", "SLESING", "KISSING", "DERRING") ||
						StringAt(-8, "SCHLESING") ||
						StringAt(-5, "SALING", "BELANG") ||
						StringAt(-6, "BARRING") ||
						StringAt(-6, "PHALANGER") ||
						StringAt(-4, "CHANG"))) ||
				StringAt(-4, "STING", "YOUNG") ||
				StringAt(-5, "STRONG") ||
				StringAtStart("UNG", "ENG", "ING") ||
				StringAt(0, "GERICH") ||
				StringAtStart("SENGER") ||
				StringAt(-3, "WENGER", "MUNGER", "SONGER", "KINGER") ||
				StringAt(-4, "FLINGER", "SLINGER", "STANGER", "STENGER", "KLINGER", "CLINGER") ||
				StringAt(-5, "SPRINGER", "SPRENGER") ||
				StringAt(-3, "LINGERF") ||
				StringAt(-2, "ANGERLY", "ANGERBO", "INGERSO")) {
				MetaphAddExactApprox("J", "G", "J", "K")
			} else {
				MetaphAddExactApprox("G", "J", "K", "J")
			}

			AdvanceCounter(2, 1)
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
	Encode_GER := func() bool {
		if (m_current > 0) && StringAt(+1, "ER") {
			// Exceptions to 'GE' where 'G' => K
			// e.g. "JAGER", "TIGER", "LIGER", "LAGER", "LUGER", "AUGER", "EAGER", "HAGER", "SAGER"
			if (((m_current == 2) && IsVowel(m_current-1) && !IsVowel(m_current-2) &&
				!(StringAt(-2, "PAGER", "WAGER", "NIGER", "ROGER", "LEGER", "CAGER")) ||
				StringAt(-2, "AUGER", "EAGER", "INGER", "YAGER")) ||
				StringAt(-3, "SEEGER", "JAEGER", "GEIGER", "KRUGER", "SAUGER", "BURGER",
					"MEAGER", "MARGER", "RIEGER", "YAEGER", "STEGER", "PRAGER", "SWIGER",
					"YERGER", "TORGER", "FERGER", "HILGER", "ZEIGER", "YARGER",
					"COWGER", "CREGER", "KROGER", "KREGER", "GRAGER", "STIGER", "BERGER") ||
				// 'berger' but not 'bergerac'
				(StringAt(-3, "BERGER") && ((m_current + 2) == m_last)) ||
				StringAt(-4, "KREIGER", "KRUEGER", "METZGER", "KRIEGER", "KROEGER", "STEIGER",
					"DRAEGER", "BUERGER", "BOERGER", "FIBIGER") ||
				// e.g. 'harshbarger', 'winebarger'
				(StringAt(-3, "BARGER") && (m_current > 4)) ||
				// e.g. 'weisgerber'
				(StringAt(0, "GERBER") && (m_current > 0)) ||
				StringAt(-5, "SCHWAGER", "LYBARGER", "SPRENGER", "GALLAGER", "WILLIGER") ||
				StringAtStart("HARGER") ||
				(StringAtStart("AGER", "EGER") && (m_length == 4)) ||
				StringAt(-1, "YGERNE") ||
				StringAt(-6, "SCHWEIGER")) &&
				!(StringAt(-5, "BELLIGEREN") ||
					StringAtStart("MARGERY") ||
					StringAt(-3, "BERGERAC")) {
				if SlavoGermanic() {
					MetaphAddExactApprox("G", "K")
				} else {
					MetaphAddExactApprox("G", "J", "K", "J")
				}
			} else {
				MetaphAddExactApprox("J", "G", "J", "K")
			}

			AdvanceCounter(2, 1)
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
	Encode_GEL := func() bool {
		// more likely to be "-GEL-" => JL
		if StringAt(+1, "EL") && (m_current > 0) {
			// except for
			// "BAGEL", "HEGEL", "HUGEL", "KUGEL", "NAGEL", "VOGEL", "FOGEL", "PAGEL"
			if ((m_length == 5) &&
				IsVowel(m_current-1) &&
				!IsVowel(m_current-2) &&
				!StringAt(-2, "NIGEL", "RIGEL")) ||
				// or the following as combining forms
				StringAt(-2, "ENGEL", "HEGEL", "NAGEL", "VOGEL") ||
				StringAt(-3, "MANGEL", "WEIGEL", "FLUGEL", "RANGEL", "HAUGEN", "RIEGEL", "VOEGEL") ||
				StringAt(-4, "SPEIGEL", "STEIGEL", "WRANGEL", "SPIEGEL") ||
				StringAt(-4, "DANEGELD") {
				if SlavoGermanic() {
					MetaphAddExactApprox("G", "K")
				} else {
					MetaphAddExactApprox("G", "J", "K", "J")
				}
			} else {
				MetaphAddExactApprox("J", "G", "J", "K")
			}

			AdvanceCounter(2, 1)
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
	Hard_GE_At_End := func() bool {
		return StringAtStart("RENEGE", "STONGE", "STANGE", "PRANGE", "KRESGE") ||
			StringAtStart("BYRGE", "BIRGE", "BERGE", "HAUGE") ||
			StringAtStart("HAGE") ||
			StringAtStart("LANGE", "SYNGE", "BENGE", "RUNGE", "HELGE") ||
			StringAtStart("INGE", "LAGE")
	}

	/**
	 * Detect words where "-ge-" or "-gi-" get a 'hard' 'g'
	 * even though this is usually a 'soft' 'g' context
	 *
	 * @return true if 'hard' 'g' detected
	 *
	 */
	Internal_Hard_G_Other := func() bool {
		return (StringAt(0, "GETH", "GEAR", "GEIS", "GIRL", "GIVI", "GIVE", "GIFT",
			"GIRD", "GIRT", "GILV", "GILD", "GELD") &&
			!StringAt(-3, "GINGIV")) ||
			// "gish" but not "largish"
			(StringAt(+1, "ISH") && (m_current > 0) && !StringAtStart("LARG")) ||
			(StringAt(-2, "MAGED", "MEGID") && !((m_current + 2) == m_last)) ||
			StringAt(0, "GEZ") ||
			StringAtStart("WEGE", "HAGE") ||
			(StringAt(-2, "ONGEST", "UNGEST") &&
				((m_current + 3) == m_last) &&
				!StringAt(-3, "CONGEST")) ||
			StringAtStart("VOEGE", "BERGE", "HELGE") ||
			(StringAtStart("ENGE", "BOGY") && (m_length == 4)) ||
			StringAt(0, "GIBBON") ||
			StringAtStart("CORREGIDOR") ||
			StringAtStart("INGEBORG") ||
			(StringAt(0, "GILL") &&
				(((m_current + 3) == m_last) || ((m_current + 4) == m_last)) &&
				!StringAtStart("STURGILL"))
	}

	/**
	 * Detect words where "-gy-", "-gie-", "-gee-",
	 * or "-gio-" get a 'hard' 'g' even though this is
	 * usually a 'soft' 'g' context
	 *
	 * @return true if 'hard' 'g' detected
	 *
	 */
	Internal_Hard_G_Open_Syllable := func() bool {
		return StringAt(+1, "EYE") ||
			StringAt(-2, "FOGY", "POGY", "YOGI") ||
			StringAt(-2, "MAGEE", "MCGEE", "HAGIO") ||
			StringAt(-1, "RGEY", "OGEY") ||
			StringAt(-3, "HOAGY", "STOGY", "PORGY") ||
			StringAt(-5, "CARNEGIE") ||
			(StringAt(-1, "OGEY", "OGIE") && ((m_current + 2) == m_last))
	}

	/**
	 * Detect a number of contexts, mostly german names, that
	 * take a 'hard' 'g'.
	 *
	 * @return true if 'hard' 'g' detected, false if not
	 *
	 */
	Internal_Hard_GEN_GIN_GET_GIT := func() bool {
		return (StringAt(-3, "FORGET", "TARGET", "MARGIT", "MARGET", "TURGEN",
			"BERGEN", "MORGEN", "JORGEN", "HAUGEN", "JERGEN",
			"JURGEN", "LINGEN", "BORGEN", "LANGEN", "KLAGEN", "STIGER", "BERGER") &&
			!StringAt(0, "GENETIC", "GENESIS") &&
			!StringAt(-4, "PLANGENT")) ||
			(StringAt(-3, "BERGIN", "FEAGIN", "DURGIN") && ((m_current + 2) == m_last)) ||
			(StringAt(-2, "ENGEN") && !StringAt(+3, "DER", "ETI", "ESI")) ||
			StringAt(-4, "JUERGEN") ||
			StringAtStart("NAGIN", "MAGIN", "HAGIN") ||
			(StringAtStart("ENGIN", "DEGEN", "LAGEN", "MAGEN", "NAGIN") && (m_length == 5)) ||
			(StringAt(-2, "BEGET", "BEGIN", "HAGEN", "FAGIN",
				"BOGEN", "WIGIN", "NTGEN", "EIGEN",
				"WEGEN", "WAGEN") &&
				!StringAt(-5, "OSPHAGEN"))
	}

	/**
	 * Detect a number of contexts of '-ng-' that will
	 * take a 'hard' 'g' despite being followed by a
	 * front vowel.
	 *
	 * @return true if 'hard' 'g' detected, false if not
	 *
	 */
	Internal_Hard_NG := func() bool {
		return (StringAt(-3, "DANG", "FANG", "SING") &&
			// exception to exception
			!StringAt(-5, "DISINGEN")) ||
			StringAtStart("INGEB", "ENGEB") ||
			(StringAt(-3, "RING", "WING", "HANG", "LONG") &&
				!(StringAt(-4, "CRING", "FRING", "ORANG", "TWING", "CHANG", "PHANG") ||
					StringAt(-5, "SYRING") ||
					StringAt(-3, "RINGENC", "RINGENT", "LONGITU", "LONGEVI") ||
					// e.g. 'longino', 'mastrangelo'
					(StringAt(0, "GELO", "GINO") && ((m_current + 3) == m_last)))) ||
			(StringAt(-1, "NGY") &&
				// exceptions to exception
				!(StringAt(-3, "RANGY", "MANGY", "MINGY") ||
					StringAt(-4, "SPONGY", "STINGY")))
	}

	/**
	 * Exceptions to default encoding to 'J':
	 * encode "-G-" to 'G' in "-g<frontvowel>-" words
	 * where we are not at "-GE" at the end of the word
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Internal_Hard_G := func() bool {
		// if not "-GE" at end
		return !(((m_current + 1) == m_last) && (CharAt(m_current+1) == 'E')) &&
			(Internal_Hard_NG() ||
				Internal_Hard_GEN_GIN_GET_GIT() ||
				Internal_Hard_G_Open_Syllable() ||
				Internal_Hard_G_Other())
	}
	/**
	 * Encode "-G-" followed by a vowel when non-initial leter.
	 * Default for this is a 'J' sound, so check exceptions where
	 * it is pronounced 'G'
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Non_Initial_G_Front_Vowel := func() bool {
		// -gy-, gi-, ge-
		if StringAt(+1, "E", "I", "Y") {
			// '-ge' at end
			// almost always 'j 'sound
			if StringAt(0, "GE") && (m_current == (m_last - 1)) {
				if Hard_GE_At_End() {
					if SlavoGermanic() {
						MetaphAddExactApprox("G", "K")
					} else {
						MetaphAddExactApprox("G", "J", "K", "J")
					}
				} else {
					MetaphAdd("J")
				}
			} else {
				if Internal_Hard_G() {
					// don't encode KG or KK if e.g. "mcgill"
					if !((m_current == 2) && StringAtStart("MC")) ||
						((m_current == 3) && StringAtStart("MAC")) {
						if SlavoGermanic() {
							MetaphAddExactApprox("G", "K")
						} else {
							MetaphAddExactApprox("G", "J", "K", "J")
						}
					}
				} else {
					MetaphAddExactApprox("J", "G", "J", "K")
				}
			}

			AdvanceCounter(2, 1)
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
	Encode_GA_To_J := func() bool {
		// 'margary', 'margarine'
		if (StringAt(-3, "MARGARY", "MARGARI") &&
			// but not in spanish forms such as "margatita"
			!StringAt(-3, "MARGARIT")) ||
			StringAtStart("GAOL") ||
			StringAt(-2, "ALGAE") {
			MetaphAddExactApprox("J", "G", "J", "K")
			AdvanceCounter(2, 1)
			return true
		}

		return false
	}

	/**
	 * Encode "-G-"
	 *
	 */
	Encode_G := func() {
		if Encode_Silent_G_At_Beginning() ||
			Encode_GG() ||
			Encode_GK() ||
			Encode_GH() ||
			Encode_Silent_G() ||
			Encode_GN() ||
			Encode_GL() ||
			Encode_Initial_G_Front_Vowel() ||
			Encode_NGER() ||
			Encode_GER() ||
			Encode_GEL() ||
			Encode_Non_Initial_G_Front_Vowel() ||
			Encode_GA_To_J() {
			return
		}

		if !StringAt(-1, "C", "K", "G", "Q") {
			MetaphAddExactApprox("G", "K")
		}

		m_current++
	}

	/**
	 * Encode cases where initial 'H' is not pronounced (in American)
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Initial_Silent_H := func() bool {
		//'hour', 'herb', 'heir', 'honor'
		if StringAt(+1, "OUR", "ERB", "EIR") ||
			StringAt(+1, "ONOR") ||
			StringAt(+1, "ONOUR", "ONEST") {
			// british pronounce H in this word
			// americans give it 'H' for the name,
			// no 'H' for the plant
			if (m_current == 0) && StringAt(0, "HERB") {
				if m_encodeVowels {
					MetaphAdd("HA", "A")
				} else {
					MetaphAdd("H", "A")
				}
			} else if (m_current == 0) || m_encodeVowels {
				MetaphAdd("A")
			}

			m_current++
			// don't encode vowels twice
			m_current = SkipVowels(m_current)
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
	Encode_Initial_HS := func() bool {
		// old chinese pinyin transliteration
		// e.g., 'HSIAO'
		if (m_current == 0) && StringAtStart("HS") {
			MetaphAdd("X")
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
	Encode_Initial_HU_HW := func() bool {
		// spanish spellings and chinese pinyin transliteration
		if StringAtStart("HUA", "HUE", "HWA") {
			if !StringAt(0, "HUEY") {
				MetaphAdd("A")

				if !m_encodeVowels {
					m_current += 3
				} else {
					m_current++
					// don't encode vowels twice
					for IsVowel(m_current) || (CharAt(m_current) == 'W') {
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
	Encode_Non_Initial_Silent_H := func() bool {
		//exceptions - 'h' not pronounced
		// "PROHIB" BUT NOT "PROHIBIT"
		if StringAt(-2, "NIHIL", "VEHEM", "LOHEN", "NEHEM",
			"MAHON", "MAHAN", "COHEN", "GAHAN") ||
			StringAt(-3, "GRAHAM", "PROHIB", "FRAHER",
				"TOOHEY", "TOUHEY") ||
			StringAt(-3, "TOUHY") ||
			StringAtStart("CHIHUAHUA") {
			if !m_encodeVowels {
				m_current += 2
			} else {
				m_current++
				// don't encode vowels twice
				m_current = SkipVowels(m_current)
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
	Encode_H_Pronounced := func() bool {
		if (((m_current == 0) ||
			IsVowel(m_current-1) ||
			((m_current > 0) &&
				(CharAt(m_current-1) == 'W'))) &&
			IsVowel(m_current+1)) ||
			// e.g. 'alWahhab'
			((CharAt(m_current+1) == 'H') && IsVowel(m_current+2)) {
			MetaphAdd("H")
			AdvanceCounter(2, 1)
			return true
		}

		return false
	}

	/**
	 * Encode 'H'
	 *
	 *
	 */
	Encode_H := func() {
		if Encode_Initial_Silent_H() ||
			Encode_Initial_HS() ||
			Encode_Initial_HU_HW() ||
			Encode_Non_Initial_Silent_H() {
			return
		}

		//only keep if first & before vowel or btw. 2 vowels
		if !Encode_H_Pronounced() {
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
	Encode_Spanish_J := func() bool {
		//obvious spanish, e.g. "jose", "san jacinto"
		if (StringAt(+1, "UAN", "ACI", "ALI", "EFE", "ICA", "IME", "OAQ", "UAR") &&
			!StringAt(0, "JIMERSON", "JIMERSEN")) ||
			(StringAt(+1, "OSE") && ((m_current + 3) == m_last)) ||
			StringAt(+1, "EREZ", "UNTA", "AIME", "AVIE", "AVIA") ||
			StringAt(+1, "IMINEZ", "ARAMIL") ||
			(((m_current + 2) == m_last) && StringAt(-2, "MEJIA")) ||
			StringAt(-2, "TEJED", "TEJAD", "LUJAN", "FAJAR", "BEJAR", "BOJOR", "CAJIG",
				"DEJAS", "DUJAR", "DUJAN", "MIJAR", "MEJOR", "NAJAR",
				"NOJOS", "RAJED", "RIJAL", "REJON", "TEJAN", "UIJAN") ||
			StringAt(-3, "ALEJANDR", "GUAJARDO", "TRUJILLO") ||
			(StringAt(-2, "RAJAS") && (m_current > 2)) ||
			(StringAt(-2, "MEJIA") && !StringAt(-2, "MEJIAN")) ||
			StringAt(-1, "OJEDA") ||
			StringAt(-3, "LEIJA", "MINJA") ||
			StringAt(-3, "VIAJES", "GRAJAL") ||
			StringAt(0, "JAUREGUI") ||
			StringAt(-4, "HINOJOSA") ||
			StringAtStart("SAN ") ||
			(((m_current + 1) == m_last) &&
				(CharAt(m_current+1) == 'O') &&
				// exceptions
				!(StringAtStart("TOJO") ||
					StringAtStart("BANJO") ||
					StringAtStart("MARYJO"))) {
			// americans pronounce "juan" as 'wan'
			// and "marijuana" and "tijuana" also
			// do not get the 'H' as in spanish, so
			// just treat it like a vowel in these cases
			if !(StringAt(0, "JUAN") || StringAt(0, "JOAQ")) {
				MetaphAdd("H")
			} else {
				if m_current == 0 {
					MetaphAdd("A")
				}
			}
			AdvanceCounter(2, 1)
			return true
		}

		// Jorge gets 2nd HARHA. also JULIO, JESUS
		if StringAt(+1, "ORGE", "ULIO", "ESUS") &&
			!StringAtStart("JORGEN") {
			// get both consonants for "jorge"
			if ((m_current + 4) == m_last) && StringAt(+1, "ORGE") {
				if m_encodeVowels {
					MetaphAdd("JARJ", "HARHA")
				} else {
					MetaphAdd("JRJ", "HRH")
				}
				AdvanceCounter(5, 5)
				return true
			}

			MetaphAdd("J", "H")
			AdvanceCounter(2, 1)
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
	Encode_German_J := func() bool {
		if StringAt(+1, "AH") ||
			(StringAt(+1, "OHANN") && ((m_current + 5) == m_last)) ||
			(StringAt(+1, "UNG") && !StringAt(+1, "UNGL")) ||
			StringAt(+1, "UGO") {
			MetaphAdd("A")
			AdvanceCounter(2, 1)
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
	Encode_Spanish_OJ_UJ := func() bool {
		if StringAt(+1, "OJOBA", "UJUY ") {
			if m_encodeVowels {
				MetaphAdd("HAH")
			} else {
				MetaphAdd("HH")
			}

			AdvanceCounter(4, 3)
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
	Encode_Spanish_J_2 := func() bool {
		// spanish forms e.g. "brujo", "badajoz"
		if (((m_current - 2) == 0) &&
			StringAt(-2, "BOJA", "BAJA", "BEJA", "BOJO", "MOJA", "MOJI", "MEJI")) ||
			(((m_current - 3) == 0) &&
				StringAt(-3, "FRIJO", "BRUJO", "BRUJA", "GRAJE", "GRIJA", "LEIJA", "QUIJA")) ||
			(((m_current + 3) == m_last) &&
				StringAt(-1, "AJARA")) ||
			(((m_current + 2) == m_last) &&
				StringAt(-1, "AJOS", "EJOS", "OJAS", "OJOS", "UJON", "AJOZ", "AJAL", "UJAR", "EJON", "EJAN")) ||
			(((m_current + 1) == m_last) &&
				(StringAt(-1, "OJA", "EJA") && !StringAtStart("DEJA"))) {
			MetaphAdd("H")
			AdvanceCounter(2, 1)
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
	Encode_J_As_Vowel := func() bool {
		if StringAt(0, "JEWSK") {
			MetaphAdd("J", "")
			return true
		}

		// e.g. "stijl", "sejm" - dutch, scandanavian, and eastern european spellings
		return (StringAt(+1, "L", "T", "K", "S", "N", "M") &&
			// except words from hindi and arabic
			!StringAt(+2, "A")) ||
			StringAtStart("HALLELUJA", "LJUBLJANA") ||
			StringAtStart("LJUB", "BJOR") ||
			StringAtStart("HAJEK") ||
			StringAtStart("WOJ") ||
			// e.g. 'fjord'
			StringAtStart("FJ") ||
			// e.g. 'rekjavik', 'blagojevic'
			StringAt(0, "JAVIK", "JEVIC") ||
			(((m_current + 1) == m_last) && StringAtStart("SONJA", "TANJA", "TONJA"))
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
	Names_Beginning_With_J_That_Get_Alt_Y := func() bool {
		if StringAtStart("JAN", "JON", "JAN", "JIN", "JEN") ||
			StringAtStart("JUHL", "JULY", "JOEL", "JOHN", "JOSH",
				"JUDE", "JUNE", "JONI", "JULI", "JENA",
				"JUNG", "JINA", "JANA", "JENI", "JOEL",
				"JANN", "JONA", "JENE", "JULE", "JANI",
				"JONG", "JOHN", "JEAN", "JUNG", "JONE",
				"JARA", "JUST", "JOST", "JAHN", "JACO",
				"JANG", "JUDE", "JONE") ||
			StringAtStart("JOANN", "JANEY", "JANAE", "JOANA", "JUTTA",
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
			StringAtStart("JORDAN", "JORDON", "JOSEPH", "JOSHUA", "JOSIAH",
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
			StringAtStart("JAKOB") ||
			StringAtStart("JOHNSON", "JOHNNIE", "JASMINE", "JEANNIE", "JOHANNA",
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
			StringAtStart("JOSEFINA", "JEANNINE", "JULIANNE", "JULIANNA", "JONATHAN",
				"JONATHON", "JEANETTE", "JANNETTE", "JEANETTA", "JOHNETTA",
				"JENNEFER", "JULIENNE", "JOSPHINE", "JEANELLE", "JOHNETTE",
				"JULIEANN", "JOSEFINE", "JULIETTA", "JOHNSTON", "JACOBSON",
				"JACOBSEN", "JOHANSEN", "JOHANSON", "JAWORSKI", "JENNETTE",
				"JELLISON", "JOHANNES", "JASINSKI", "JUERGENS", "JARNAGIN",
				"JEREMIAH", "JEPPESEN", "JARNIGAN", "JANOUSEK") ||
			StringAtStart("JOHNATHAN", "JOHNATHON", "JORGENSEN", "JEANMARIE", "JOSEPHINA",
				"JEANNETTE", "JOSEPHINE", "JEANNETTA", "JORGENSON", "JANKOWSKI",
				"JOHNSTONE", "JABLONSKI", "JOSEPHSON", "JOHANNSEN", "JURGENSEN",
				"JIMMERSON", "JOHANSSON") ||
			StringAtStart("JAKUBOWSKI") {
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
	Encode_J_To_J := func() bool {
		if IsVowel(m_current + 1) {
			if (m_current == 0) &&
				Names_Beginning_With_J_That_Get_Alt_Y() {
				// 'Y' is a vowel so encode
				// is as 'A'
				if m_encodeVowels {
					MetaphAdd("JA", "A")
				} else {
					MetaphAdd("J", "A")
				}
			} else {
				if m_encodeVowels {
					MetaphAdd("JA")
				} else {
					MetaphAdd("J")
				}
			}

			m_current++
			m_current = SkipVowels(m_current)
			return false
		} else {
			MetaphAdd("J")
			m_current++
			return true
		}

		//		return false
	}

	/**
	 * Call routines to encode 'J', in proper order
	 *
	 */
	Encode_Other_J := func() {
		if m_current == 0 {
			if Encode_German_J() {
				return
			} else {
				if Encode_J_To_J() {
					return
				}
			}
		} else {
			if Encode_Spanish_J_2() {
				return
			} else if !Encode_J_As_Vowel() {
				MetaphAdd("J")
			}

			//it could happen! e.g. "hajj"
			// eat redundant 'J'
			if CharAt(m_current+1) == 'J' {
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
	Encode_J := func() {
		if Encode_Spanish_J() || Encode_Spanish_OJ_UJ() {
			return
		}

		Encode_Other_J()
	}

	/**
	 * Encode cases where 'K' is not pronounced
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Silent_K := func() bool {
		//skip this except for special cases
		if (m_current == 0) && StringAt(0, "KN") {
			if !(StringAt(+2, "ESSET", "IEVEL") || StringAt(+2, "ISH")) {
				m_current += 1
				return true
			}
		}

		// e.g. "know", "knit", "knob"
		if (StringAt(+1, "NOW", "NIT", "NOT", "NOB") &&
			// exception, "slipknot" => SLPNT but "banknote" => PNKNT
			!StringAtStart("BANKNOTE")) ||
			StringAt(+1, "NOCK", "NUCK", "NIFE", "NACK") ||
			StringAt(+1, "NIGHT") {
			// N already encoded before
			// e.g. "penknife"
			if (m_current > 0) && CharAt(m_current-1) == 'N' {
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
	Encode_K := func() {
		if !Encode_Silent_K() {
			MetaphAdd("K")

			// eat redundant 'K's and 'Q's
			if (CharAt(m_current+1) == 'K') ||
				(CharAt(m_current+1) == 'Q') {
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
	Interpolate_Vowel_When_Cons_L_At_End := func() {
		if m_encodeVowels == true {
			// e.g. "ertl", "vogl"
			if (m_current == m_last) &&
				StringAt(-1, "D", "G", "T") {
				MetaphAdd("A")
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
	Encode_LELY_To_L := func() bool {
		// e.g. "agilely", "docilely"
		if StringAt(-1, "ILELY") &&
			((m_current + 3) == m_last) {
			MetaphAdd("L")
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
	Encode_COLONEL := func() bool {
		if StringAt(-2, "COLONEL") {
			MetaphAdd("R")
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
	Encode_French_AULT := func() bool {
		// e.g. "renault" and "foucault", well known to americans, but not "fault"
		if (m_current > 3) &&
			(StringAt(-3, "RAULT", "NAULT", "BAULT", "SAULT", "GAULT", "CAULT") ||
				StringAt(-4, "REAULT", "RIAULT", "NEAULT", "BEAULT")) &&
			!(RootOrInflections(m_inWord, "ASSAULT") ||
				StringAt(-8, "SOMERSAULT") ||
				StringAt(-9, "SUMMERSAULT")) {
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
	Encode_French_EUIL := func() bool {
		// e.g. "auteuil"
		if StringAt(-3, "EUIL") && (m_current == m_last) {
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
	Encode_French_OULX := func() bool {
		// e.g. "proulx"
		if StringAt(-2, "OULX") && ((m_current + 1) == m_last) {
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
	Encode_Silent_L_In_LM := func() bool {
		if StringAt(0, "LM", "LN") {
			// e.g. "lincoln", "holmes", "psalm", "salmon"
			if (StringAt(-2, "COLN", "CALM", "BALM", "MALM", "PALM") ||
				(StringAt(-1, "OLM") && ((m_current + 1) == m_last)) ||
				StringAt(-3, "PSALM", "QUALM") ||
				StringAt(-2, "SALMON", "HOLMES") ||
				StringAt(-1, "ALMOND") ||
				((m_current == 1) && StringAt(-1, "ALMS"))) &&
				(!StringAt(+2, "A") &&
					!StringAt(-2, "BALMO") &&
					!StringAt(-2, "PALMER", "PALMOR", "BALMER") &&
					!StringAt(-3, "THALM")) {
				m_current++
				return true
			} else {
				MetaphAdd("L")
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
	Encode_Silent_L_In_LK_LV := func() bool {
		if (StringAt(-2, "WALK", "YOLK", "FOLK", "HALF", "TALK", "CALF", "BALK", "CALK") ||
			(StringAt(-2, "POLK") &&
				!StringAt(-2, "POLKA", "WALKO")) ||
			(StringAt(-2, "HALV") &&
				!StringAt(-2, "HALVA", "HALVO")) ||
			(StringAt(-3, "CAULK", "CHALK", "BAULK", "FAULK") &&
				!StringAt(-4, "SCHALK")) ||
			(StringAt(-2, "SALVE", "CALVE") ||
				StringAt(-2, "SOLDER")) &&
				// exceptions to above cases where 'L' is usually pronounced
				!StringAt(-2, "SALVER", "CALVER")) &&
			!StringAt(-5, "GONSALVES", "GONCALVES") &&
			!StringAt(-2, "BALKAN", "TALKAL") &&
			!StringAt(-3, "PAULK", "CHALF") {
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
	Encode_Silent_L_In_OULD := func() bool {
		//'would', 'could'
		if StringAt(-3, "WOULD", "COULD") ||
			(StringAt(-4, "SHOULD") &&
				!StringAt(-4, "SHOULDER")) {
			MetaphAddExactApprox("D", "T")
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
	Encode_LL_As_Vowel_Special_Cases := func() bool {
		if StringAt(-5, "TORTILLA") ||
			StringAt(-8, "RATATOUILLE") ||
			// e.g. 'guillermo', "veillard"
			(StringAtStart("GUILL", "VEILL", "GAILL") &&
				// 'guillotine' usually has '-ll-' pronounced as 'L' in english
				!(StringAt(-3, "GUILLOT", "GUILLOR", "GUILLEN") ||
					(StringAtStart("GUILL") && (m_length == 5)))) ||
			// e.g. "brouillard", "gremillion"
			StringAtStart("BROUILL", "GREMILL") ||
			StringAtStart("ROBILL") ||
			// e.g. 'mireille'
			(StringAt(-2, "EILLE") &&
				((m_current + 2) == m_last) &&
				// exception "reveille" usually pronounced as 're-vil-lee'
				!StringAt(-5, "REVEILLE")) {
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
	Encode_LL_As_Vowel := func() bool {
		//spanish e.g. "cabrillo", "gallegos" but also "gorilla", "ballerina" -
		// give both pronounciations since an american might pronounce "cabrillo"
		// in the spanish or the american fashion.
		if (((m_current + 3) == m_length) &&
			StringAt(-1, "ILLO", "ILLA", "ALLE")) ||
			(((StringAtEnd("AS", "OS") ||
				StringAtEnd("A", "O")) &&
				StringAt(-1, "AL", "IL")) &&
				!StringAt(-1, "ALLA")) ||
			StringAtStart("VILLE", "VILLA") ||
			StringAtStart("GALLARDO", "VALLADAR", "MAGALLAN", "CAVALLAR", "BALLASTE") ||
			StringAtStart("LLA") {
			MetaphAdd("L", "")
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
	Encode_LL_As_Vowel_Cases := func() bool {
		if CharAt(m_current+1) == 'L' {
			if Encode_LL_As_Vowel_Special_Cases() {
				return true
			} else if Encode_LL_As_Vowel() {
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
	Encode_Vowel_LE_Transposition := func(save_current int) bool {
		// transposition of vowel sound and L occurs in many words,
		// e.g. "bristle", "dazzle", "goggle" => KAKAL
		if m_encodeVowels && (save_current > 1) &&
			!IsVowel(save_current-1) &&
			(CharAt(save_current+1) == 'E') &&
			(CharAt(save_current-1) != 'L') &&
			(CharAt(save_current-1) != 'R') &&
			// lots of exceptions to this:
			!IsVowel(save_current+2) &&
			!StringAtStart("ECCLESI", "COMPLEC", "COMPLEJ", "ROBLEDO") &&
			!StringAtStart("MCCLE", "MCLEL") &&
			!StringAtStart("EMBLEM", "KADLEC") &&
			!(((save_current + 2) == m_last) && StringAt(save_current, "LET")) &&
			!StringAt(save_current, "LETTING") &&
			!StringAt(save_current, "LETELY", "LETTER", "LETION", "LETIAN", "LETING", "LETORY") &&
			!StringAt(save_current, "LETUS", "LETIV") &&
			!StringAt(save_current, "LESS", "LESQ", "LECT", "LEDG", "LETE", "LETH", "LETS", "LETT") &&
			!StringAt(save_current, "LEG", "LER", "LEX") &&
			// e.g. "complement" !=> KAMPALMENT
			!(StringAt(save_current, "LEMENT") &&
				!(StringAt(-5, "BATTLE", "TANGLE", "PUZZLE", "RABBLE", "BABBLE") ||
					StringAt(-4, "TABLE"))) &&
			!(((save_current + 2) == m_last) && StringAt((save_current-2), "OCLES", "ACLES", "AKLES")) &&
			!StringAt((save_current-3), "LISLE", "AISLE") &&
			!StringAtStart("ISLE") &&
			!StringAtStart("ROBLES") &&
			!StringAt((save_current-4), "PROBLEM", "RESPLEN") &&
			!StringAt((save_current-3), "REPLEN") &&
			!StringAt((save_current-2), "SPLE") &&
			(CharAt(save_current-1) != 'H') &&
			(CharAt(save_current-1) != 'W') {
			MetaphAdd("AL")
			flag_AL_inversion = true

			// eat redundant 'L'
			if CharAt(save_current+2) == 'L' {
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
	Encode_Vowel_Preserve_Vowel_After_L := func(save_current int) bool {
		// an example of where the vowel would NOT need to be preserved
		// would be, say, "hustled", where there is no vowel pronounced
		// between the 'l' and the 'd'
		if m_encodeVowels &&
			!IsVowel(save_current-1) &&
			(CharAt(save_current+1) == 'E') &&
			(save_current > 1) &&
			((save_current + 1) != m_last) &&
			!(StringAt((save_current+1), "ES", "ED") &&
				((save_current + 2) == m_last)) &&
			!StringAt((save_current-1), "RLEST") {
			MetaphAdd("LA")
			m_current = SkipVowels(m_current)
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
	Encode_LE_Cases := func(save_current int) {
		if Encode_Vowel_LE_Transposition(save_current) {
			return
		} else {
			if Encode_Vowel_Preserve_Vowel_After_L(save_current) {
				return
			} else {
				MetaphAdd("L")
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
	Encode_L := func() {
		// logic below needs to know this
		// after 'm_current' variable changed
		var save_current = m_current

		Interpolate_Vowel_When_Cons_L_At_End()

		if Encode_LELY_To_L() ||
			Encode_COLONEL() ||
			Encode_French_AULT() ||
			Encode_French_EUIL() ||
			Encode_French_OULX() ||
			Encode_Silent_L_In_LM() ||
			Encode_Silent_L_In_LK_LV() ||
			Encode_Silent_L_In_OULD() {
			return
		}

		if Encode_LL_As_Vowel_Cases() {
			return
		}

		Encode_LE_Cases(save_current)
	}

	/**
	 * Encode cases where 'M' is silent at beginning of word
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Silent_M_At_Beginning := func() bool {
		//skip these when at start of word
		if (m_current == 0) && StringAt(0, "MN") {
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
	Encode_MR_And_MRS := func() bool {
		if (m_current == 0) && StringAt(0, "MR") {
			// exceptions for "mr." and "mrs."
			if (m_length == 2) && StringAt(0, "MR") {
				if m_encodeVowels {
					MetaphAdd("MASTAR")
				} else {
					MetaphAdd("MSTR")
				}
				m_current += 2
				return true
			} else if (m_length == 3) && StringAt(0, "MRS") {
				if m_encodeVowels {
					MetaphAdd("MASAS")
				} else {
					MetaphAdd("MSS")
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
	Encode_MAC := func() bool {
		// should only find irish and
		// scottish names e.g. 'macintosh'
		if (m_current == 0) &&
			(StringAtStart("MACIVER", "MACEWEN") ||
				StringAtStart("MACELROY", "MACILROY") ||
				StringAtStart("MACINTOSH") ||
				StringAtStart("MC")) {
			if m_encodeVowels {
				MetaphAdd("MAK")
			} else {
				MetaphAdd("MK")
			}

			if StringAtStart("MC") {
				if StringAt(+2, "K", "G", "Q") &&
					// watch out for e.g. "McGeorge"
					!StringAt(+2, "GEOR") {
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
	Encode_MPT := func() bool {
		if StringAt(-2, "COMPTROL") ||
			StringAt(-4, "ACCOMPT") {
			MetaphAdd("N")
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
	Test_Silent_MB_1 := func() bool {
		// e.g. "LAMB", "COMB", "LIMB", "DUMB", "BOMB"
		// Handle combining roots first
		if ((m_current == 3) &&
			StringAt(-3, "THUMB")) ||
			((m_current == 2) &&
				StringAt(-2, "DUMB", "BOMB", "DAMN", "LAMB", "NUMB", "TOMB")) {
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
	Test_Pronounced_MB := func() bool {
		if StringAt(-2, "NUMBER") ||
			(StringAt(+2, "A") &&
				!StringAt(-2, "DUMBASS")) ||
			StringAt(+2, "O") ||
			StringAt(-2, "LAMBEN", "LAMBER", "LAMBET", "TOMBIG", "LAMBRE") {
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
	Test_Silent_MB_2 := func() bool {
		// 'M' is the current letter
		if (CharAt(m_current+1) == 'B') && (m_current > 1) &&
			(((m_current + 1) == m_last) ||
				// other situations where "-MB-" is at end of root
				// but not at end of word. The tests are for standard
				// noun suffixes.
				// e.g. "climbing" => KLMNK
				StringAt(+2, "ING", "ABL") ||
				StringAt(+2, "LIKE") ||
				((CharAt(m_current+2) == 'S') && ((m_current + 2) == m_last)) ||
				StringAt(-5, "BUNCOMB") ||
				// e.g. "bomber",
				(StringAt(+2, "ED", "ER") &&
					((m_current + 3) == m_last) &&
					(StringAtStart("CLIMB", "PLUMB") ||
						// e.g. "beachcomber"
						!StringAt(-1, "IMBER", "AMBER", "EMBER", "UMBER")) &&
					// exceptions
					!StringAt(-2, "CUMBER", "SOMBER"))) {
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
	Test_Pronounced_MB_2 := func() bool {
		// e.g. "bombastic", "umbrage", "flamboyant"
		if StringAt(-1, "OMBAS", "OMBAD", "UMBRA") ||
			StringAt(-3, "FLAM") {
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
	Test_MN := func() bool {
		return (CharAt(m_current+1) == 'N') &&
			(((m_current + 1) == m_last) ||
				// or at the end of a word but followed by suffixes
				(StringAt(+2, "ING", "EST") && ((m_current + 4) == m_last)) ||
				((CharAt(m_current+2) == 'S') && ((m_current + 2) == m_last)) ||
				(StringAt(+2, "LY", "ER", "ED") &&
					((m_current + 3) == m_last)) ||
				StringAt(-2, "DAMNEDEST") ||
				StringAt(-5, "GODDAMNIT"))
	}

	/**
	 * Call routines to encode "-MB-", in proper order
	 *
	 */
	Encode_MB := func() {
		if Test_Silent_MB_1() {
			if Test_Pronounced_MB() {
				m_current++
			} else {
				m_current += 2
			}
		} else if Test_Silent_MB_2() {
			if Test_Pronounced_MB_2() {
				m_current++
			} else {
				m_current += 2
			}
		} else if Test_MN() {
			m_current += 2
		} else {
			// eat redundant 'M'
			if CharAt(m_current+1) == 'M' {
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
	Encode_M := func() {
		if Encode_Silent_M_At_Beginning() ||
			Encode_MR_And_MRS() ||
			Encode_MAC() ||
			Encode_MPT() {
			return
		}

		// Silent 'B' should really be handled
		// under 'B", not here under 'M'!
		Encode_MB()

		MetaphAdd("M")
	}

	/**
	 * Encode "-NCE-" and "-NSE-"
	 * "entrance" is pronounced exactly the same as "entrants"
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_NCE := func() bool {
		//'acceptance', 'accountancy'
		if StringAt(+1, "C", "S") &&
			StringAt(+2, "E", "Y", "I") &&
			(((m_current + 2) == m_last) ||
				((m_current+3) == m_last) &&
					(CharAt(m_current+3) == 'S')) {
			MetaphAdd("NTS")
			m_current += 2
			return true
		}

		return false
	}

	/**
	 * Encode "-N-"
	 *
	 */
	Encode_N := func() {
		if Encode_NCE() {
			return
		}

		// eat redundant 'N'
		if CharAt(m_current+1) == 'N' {
			m_current += 2
		} else {
			m_current++
		}

		if !StringAt(-3, "MONSIEUR") &&
			// e.g. "aloneness",
			!StringAt(-3, "NENESS") {
			MetaphAdd("N")
		}
	}

	/**
	 * Encode cases where "-P-" is silent at the start of a word
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Silent_P_At_Beginning := func() bool {
		//skip these when at start of word
		if (m_current == 0) &&
			StringAt(0, "PN", "PF", "PS", "PT") {
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
	Encode_PT := func() bool {
		// 'pterodactyl', 'receipt', 'asymptote'
		if CharAt(m_current+1) == 'T' {
			if ((m_current == 0) && StringAt(0, "PTERO")) ||
				StringAt(-5, "RECEIPT") ||
				StringAt(-4, "ASYMPTOT") {
				MetaphAdd("T")
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
	Encode_PH := func() bool {
		if CharAt(m_current+1) == 'H' {
			// 'PH' silent in these contexts
			if StringAt(0, "PHTHALEIN") ||
				((m_current == 0) && StringAt(0, "PHTH")) ||
				StringAt(-3, "APOPHTHEGM") {
				MetaphAdd("0")
				m_current += 4
				// combining forms
				//'sheepherd', 'upheaval', 'cupholder'
			} else if (m_current > 0) &&
				(StringAt(+2, "EAD", "OLE", "ELD", "ILL", "OLD", "EAP", "ERD",
					"ARD", "ANG", "ORN", "EAV", "ART") ||
					StringAt(+2, "OUSE") ||
					(StringAt(+2, "AM") && !StringAt(-1, "LPHAM")) ||
					StringAt(+2, "AMMER", "AZARD", "UGGER") ||
					StringAt(+2, "OLSTER")) &&
				!StringAt(-3, "LYMPH", "NYMPH") {
				MetaphAdd("P")
				AdvanceCounter(3, 2)
			} else {
				MetaphAdd("F")
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
	Encode_PPH := func() bool {
		// 'sappho'
		if (CharAt(m_current+1) == 'P') &&
			((m_current + 2) < m_length) && (CharAt(m_current+2) == 'H') {
			MetaphAdd("F")
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
	Encode_RPS := func() bool {
		//'-corps-', 'corpsman'
		if StringAt(-3, "CORPS") &&
			!StringAt(-3, "CORPSE") {
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
	Encode_COUP := func() bool {
		//'coup'
		if (m_current == m_last) &&
			StringAt(-3, "COUP") &&
			!StringAt(-5, "RECOUP") {
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
	Encode_PNEUM := func() bool {
		//'-pneum-'
		if StringAt(+1, "NEUM") {
			MetaphAdd("N")
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
	Encode_PSYCH := func() bool {
		//'-psych-'
		if StringAt(+1, "SYCH") {
			if m_encodeVowels {
				MetaphAdd("SAK")
			} else {
				MetaphAdd("SK")
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
	Encode_PSALM := func() bool {
		//'-psalm-'
		if StringAt(+1, "SALM") {
			// go ahead and encode entire word
			if m_encodeVowels {
				MetaphAdd("SAM")
			} else {
				MetaphAdd("SM")
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
	Encode_PB := func() {
		// e.g. "campbell", "raspberry"
		// eat redundant 'P' or 'B'
		if StringAt(+1, "P", "B") {
			m_current += 2
		} else {
			m_current++
		}
	}

	/**
	 * Encode "-P-"
	 *
	 */
	Encode_P := func() {
		if Encode_Silent_P_At_Beginning() ||
			Encode_PT() ||
			Encode_PH() ||
			Encode_PPH() ||
			Encode_RPS() ||
			Encode_COUP() ||
			Encode_PNEUM() ||
			Encode_PSYCH() ||
			Encode_PSALM() {
			return
		}

		Encode_PB()

		MetaphAdd("P")
	}

	/**
	 * Encode "-Q-"
	 *
	 */
	Encode_Q := func() {
		// current pinyin
		if StringAt(0, "QIN") {
			MetaphAdd("X")
			m_current++
			return
		}

		// eat redundant 'Q'
		if CharAt(m_current+1) == 'Q' {
			m_current += 2
		} else {
			m_current++
		}

		MetaphAdd("K")
	}

	/**
	 * Encode "-RZ-" according
	 * to american and polish pronunciations
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_RZ := func() bool {
		if StringAt(-2, "GARZ", "KURZ", "MARZ", "MERZ", "HERZ", "PERZ", "WARZ") ||
			StringAt(0, "RZANO", "RZOLA") ||
			StringAt(-1, "ARZA", "ARZN") {
			return false
		}

		// 'yastrzemski' usually has 'z' silent in
		// united states, but should get 'X' in poland
		if StringAt(-4, "YASTRZEMSKI") {
			MetaphAdd("R", "X")
			m_current += 2
			return true
		}
		// 'BRZEZINSKI' gets two pronunciations
		// in the united states, neither of which
		// are authentically polish
		if StringAt(-1, "BRZEZINSKI") {
			MetaphAdd("RS", "RJ")
			// skip over 2nd 'Z'
			m_current += 4
			return true
			// 'z' in 'rz after voiceless consonant gets 'X'
			// in alternate polish style pronunciation
		} else if StringAt(-1, "TRZ", "PRZ", "KRZ") ||
			(StringAt(0, "RZ") &&
				(IsVowel(m_current-1) || (m_current == 0))) {
			MetaphAdd("RS", "X")
			m_current += 2
			return true
			// 'z' in 'rz after voiceled consonant, vowel, or at
			// beginning gets 'J' in alternate polish style pronunciation
		} else if StringAt(-1, "BRZ", "DRZ", "GRZ") {
			MetaphAdd("RS", "J")
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
	Test_Silent_R := func() bool {
		// test cases where 'R' is silent, either because the
		// word is from the french or because it is no longer pronounced.
		// e.g. "rogier", "monsieur", "surburban"
		return ((m_current == m_last) &&
			// reliably french word ending
			StringAt(-2, "IER") &&
			// e.g. "metier"
			(StringAt(-5, "MET", "VIV", "LUC") ||
				// e.g. "cartier", "bustier"
				StringAt(-6, "CART", "DOSS", "FOUR", "OLIV", "BUST", "DAUM", "ATEL",
					"SONN", "CORM", "MERC", "PELT", "POIR", "BERN", "FORT", "GREN",
					"SAUC", "GAGN", "GAUT", "GRAN", "FORC", "MESS", "LUSS", "MEUN",
					"POTH", "HOLL", "CHEN") ||
				// e.g. "croupier"
				StringAt(-7, "CROUP", "TORCH", "CLOUT", "FOURN", "GAUTH", "TROTT",
					"DEROS", "CHART") ||
				// e.g. "chevalier"
				StringAt(-8, "CHEVAL", "LAVOIS", "PELLET", "SOMMEL", "TREPAN", "LETELL", "COLOMB") ||
				StringAt(-9, "CHARCUT") ||
				StringAt(-10, "CHARPENT"))) ||
			StringAt(-2, "SURBURB", "WORSTED") ||
			StringAt(-2, "WORCESTER") ||
			StringAt(-7, "MONSIEUR") ||
			StringAt(-6, "POITIERS")
	}

	/**
	 * Encode '-re-" as 'AR' in contexts
	 * where this is the correct pronunciation
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Vowel_RE_Transposition := func() bool {
		// -re inversion is just like
		// -le inversion
		// e.g. "fibre" => FABAR or "centre" => SANTAR
		if (m_encodeVowels) &&
			(CharAt(m_current+1) == 'E') &&
			(m_length > 3) &&
			!StringAtStart("OUTRE", "LIBRE", "ANDRE") &&
			!(StringAtStart("FRED", "TRES") && (m_length == 4)) &&
			!StringAt(-2, "LDRED", "LFRED", "NDRED", "NFRED", "NDRES", "TRES", "IFRED") &&
			!IsVowel(m_current-1) &&
			(((m_current + 1) == m_last) ||
				(((m_current + 2) == m_last) &&
					StringAt(+2, "D", "S"))) {
			MetaphAdd("AR")
			return true
		}

		return false
	}

	/**
	 * Encode "-R-"
	 *
	 */
	Encode_R := func() {
		if Encode_RZ() {
			return
		}

		if !Test_Silent_R() {
			if !Encode_Vowel_RE_Transposition() {
				MetaphAdd("R")
			}
		}

		// eat redundant 'R'; also skip 'S' as well as 'R' in "poitiers"
		if (CharAt(m_current+1) == 'R') || StringAt(-6, "POITIERS") {
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
	Names_Beginning_With_SW_That_Get_Alt_SV := func() bool {
		if StringAtStart("SWANSON", "SWENSON", "SWINSON", "SWENSEN",
			"SWOBODA") ||
			StringAtStart("SWIDERSKI", "SWARTHOUT") ||
			StringAtStart("SWEARENGIN") {
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
	Names_Beginning_With_SW_That_Get_Alt_XV := func() bool {
		if StringAtStart("SWART") ||
			StringAtStart("SWARTZ", "SWARTS", "SWIGER") ||
			StringAtStart("SWITZER", "SWANGER", "SWIGERT",
				"SWIGART", "SWIHART") ||
			StringAtStart("SWEITZER", "SWATZELL", "SWINDLER") ||
			StringAtStart("SWINEHART") ||
			StringAtStart("SWEARINGEN") {
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
	Encode_Special_SW := func() bool {
		if m_current == 0 {
			//
			if Names_Beginning_With_SW_That_Get_Alt_SV() {
				MetaphAdd("S", "SV")
				m_current += 2
				return true
			}

			//
			if Names_Beginning_With_SW_That_Get_Alt_XV() {
				MetaphAdd("S", "XV")
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
	Encode_SKJ := func() bool {
		// scandinavian
		if StringAt(0, "SKJO", "SKJU") &&
			IsVowel(m_current+3) {
			MetaphAdd("X")
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
	Encode_SJ := func() bool {
		if StringAtStart("SJ") {
			MetaphAdd("X")
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
	Encode_Silent_French_S_Final := func() bool {
		// "louis" is an exception because it gets two pronuncuations
		if StringAtStart("LOUIS") && (m_current == m_last) {
			MetaphAdd("S", "")
			m_current++
			return true
		}

		// french words familiar to americans where final s is silent
		if (m_current == m_last) &&
			(StringAtStart("YVES") ||
				(StringAtStart("HORS") && (m_current == 3)) ||
				StringAt(-4, "CAMUS", "YPRES") ||
				StringAt(-5, "MESNES", "DEBRIS", "BLANCS", "INGRES", "CANNES") ||
				StringAt(-6, "CHABLIS", "APROPOS", "JACQUES", "ELYSEES", "OEUVRES",
					"GEORGES", "DESPRES") ||
				StringAtStart("ARKANSAS", "FRANCAIS", "CRUDITES", "BRUYERES") ||
				StringAtStart("DESCARTES", "DESCHUTES", "DESCHAMPS", "DESROCHES", "DESCHENES") ||
				StringAtStart("RENDEZVOUS") ||
				StringAtStart("CONTRETEMPS", "DESLAURIERS")) ||
			((m_current == m_last) &&
				StringAt(-2, "AI", "OI", "UI") &&
				!StringAtStart("LOIS", "LUIS")) {
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
	Encode_Silent_French_S_Internal := func() bool {
		// french words familiar to americans where internal s is silent
		if StringAt(-2, "DESCARTES") ||
			StringAt(-2, "DESCHAM", "DESPRES", "DESROCH", "DESROSI", "DESJARD", "DESMARA",
				"DESCHEN", "DESHOTE", "DESLAUR") ||
			StringAt(-2, "MESNES") ||
			StringAt(-5, "DUQUESNE", "DUCHESNE") ||
			StringAt(-7, "BEAUCHESNE") ||
			StringAt(-3, "FRESNEL") ||
			StringAt(-3, "GROSVENOR") ||
			StringAt(-4, "LOUISVILLE") ||
			StringAt(-7, "ILLINOISAN") {
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
	Encode_ISL := func() bool {
		//special cases 'island', 'isle', 'carlisle', 'carlysle'
		if (StringAt(-2, "LISL", "LYSL", "AISL") &&
			!StringAt(-3, "PAISLEY", "BAISLEY", "ALISLAM", "ALISLAH", "ALISLAA")) ||
			((m_current == 1) &&
				((StringAt(-1, "ISLE") ||
					StringAt(-1, "ISLAN")) &&
					!StringAt(-1, "ISLEY", "ISLER"))) {
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
	Encode_STL := func() bool {
		//'hustle', 'bustle', 'whistle'
		if (StringAt(0, "STLE", "STLI") &&
			!StringAt(+2, "LESS", "LIKE", "LINE")) ||
			StringAt(-3, "THISTLY", "BRISTLY", "GRISTLY") ||
			// e.g. "corpuscle"
			StringAt(-1, "USCLE") {
			// KRISTEN, KRYSTLE, CRYSTLE, KRISTLE all pronounce the 't'
			// also, exceptions where "-LING" is a nominalizing suffix
			if StringAtStart("KRISTEN", "KRYSTLE", "CRYSTLE", "KRISTLE") ||
				StringAtStart("CHRISTENSEN", "CHRISTENSON") ||
				StringAt(-3, "FIRSTLING") ||
				StringAt(-2, "NESTLING", "WESTLING") {
				MetaphAdd("ST")
				m_current += 2
			} else {
				if m_encodeVowels &&
					(CharAt(m_current+3) == 'E') &&
					(CharAt(m_current+4) != 'R') &&
					!StringAt(+3, "ETTE", "ETTA") &&
					!StringAt(+3, "EY") {
					MetaphAdd("SAL")
					flag_AL_inversion = true
				} else {
					MetaphAdd("SL")
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
	Encode_Christmas := func() bool {
		//'christmas'
		if StringAt(-4, "CHRISTMA") {
			MetaphAdd("SM")
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
	Encode_STHM := func() bool {
		//'asthma', 'isthmus'
		if StringAt(0, "STHM") {
			MetaphAdd("SM")
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
	Encode_ISTEN := func() bool {
		// 't' is silent in verb, pronounced in name
		if StringAtStart("CHRISTEN") {
			// the word itself
			if RootOrInflections(m_inWord, "CHRISTEN") ||
				StringAtStart("CHRISTENDOM") {
				MetaphAdd("S", "ST")
			} else {
				// e.g. 'christenson', 'christene'
				MetaphAdd("ST")
			}
			m_current += 2
			return true
		}

		//e.g. 'glisten', 'listen'
		if StringAt(-2, "LISTEN", "RISTEN", "HASTEN", "FASTEN", "MUSTNT") ||
			StringAt(-3, "MOISTEN") {
			MetaphAdd("S")
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
	Encode_Sugar := func() bool {
		//special case 'sugar-'
		if StringAt(0, "SUGAR") {
			MetaphAdd("X")
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
	Encode_SH := func() bool {
		if StringAt(0, "SH") {
			// exception
			if StringAt(-2, "CASHMERE") {
				MetaphAdd("J")
				m_current += 2
				return true
			}

			//combining forms, e.g. 'clotheshorse', 'woodshole'
			if (m_current > 0) &&
				// e.g. "mishap"
				((StringAt(+1, "HAP") && ((m_current + 3) == m_last)) ||
					// e.g. "hartsheim", "clothshorse"
					StringAt(+1, "HEIM", "HOEK", "HOLM", "HOLZ", "HOOD", "HEAD", "HEID",
						"HAAR", "HORS", "HOLE", "HUND", "HELM", "HAWK", "HILL") ||
					// e.g. "dishonor"
					StringAt(+1, "HEART", "HATCH", "HOUSE", "HOUND", "HONOR") ||
					// e.g. "mishear"
					(StringAt(+2, "EAR") && ((m_current + 4) == m_last)) ||
					// e.g. "hartshorn"
					(StringAt(+2, "ORN") && !StringAt(-2, "UNSHORN")) ||
					// e.g. "newshour" but not "bashour", "manshour"
					(StringAt(+1, "HOUR") &&
						!(StringAtStart("BASHOUR") || StringAtStart("MANSHOUR") || StringAtStart("ASHOUR"))) ||
					// e.g. "dishonest", "grasshopper"
					StringAt(+2, "ARMON", "ONEST", "ALLOW", "OLDER", "OPPER", "EIMER", "ANDLE", "ONOUR") ||
					// e.g. "dishabille", "transhumance"
					StringAt(+2, "ABILLE", "UMANCE", "ABITUA")) {
				if !StringAt(-1, "S") {
					MetaphAdd("S")
				}
			} else {
				MetaphAdd("X")
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
	Encode_SCH := func() bool {
		// these words were combining forms many centuries ago
		if StringAt(+1, "CH") {
			if (m_current > 0) &&
				// e.g. "mischief", "escheat"
				(StringAt(+3, "IEF", "EAT") ||
					// e.g. "mischance"
					StringAt(+3, "ANCE", "ARGE") ||
					// e.g. "eschew"
					StringAtStart("ESCHEW")) {
				MetaphAdd("S")
				m_current++
				return true
			}

			//Schlesinger's rule
			//dutch, danish, italian, greek origin, e.g. "school", "schooner", "schiavone", "schiz-"
			if (StringAt(+3, "OO", "ER", "EN", "UY", "ED", "EM", "IA", "IZ", "IS", "OL") &&
				!StringAt(0, "SCHOLT", "SCHISL", "SCHERR")) ||
				StringAt(+3, "ISZ") ||
				(StringAt(-1, "ESCHAT", "ASCHIN", "ASCHAL", "ISCHAE", "ISCHIA") &&
					!StringAt(-2, "FASCHING")) ||
				(StringAt(-1, "ESCHI") && ((m_current + 3) == m_last)) ||
				(CharAt(m_current+3) == 'Y') {
				// e.g. "schermerhorn", "schenker", "schistose"
				if StringAt(+3, "ER", "EN", "IS") &&
					(((m_current + 4) == m_last) ||
						StringAt(+3, "ENK", "ENB", "IST")) {
					MetaphAdd("X", "SK")
				} else {
					MetaphAdd("SK")
				}
				m_current += 3
				return true
			} else {
				// Fix for smith and schmidt not returning same code:
				// next two lines from metaphone.go at line 621: code for SCH
				if m_current == 0 && !IsVowel(3) && (CharAt(3) != 'W') {
					MetaphAdd("X", "S")
				} else {
					MetaphAdd("X")
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
	Encode_SUR := func() bool {
		// 'erasure', 'usury'
		if StringAt(+1, "URE", "URA", "URY") {
			//'sure', 'ensure'
			if (m_current == 0) ||
				StringAt(-1, "N", "K") ||
				StringAt(-2, "NO") {
				MetaphAdd("X")
			} else {
				MetaphAdd("J")
			}

			AdvanceCounter(2, 1)
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
	Encode_SU := func() bool {
		//'sensuous', 'consensual'
		if StringAt(+1, "UO", "UA") && (m_current != 0) {
			// exceptions e.g. "persuade"
			if StringAt(-1, "RSUA") {
				MetaphAdd("S")
				// exceptions e.g. "casual"
			} else if IsVowel(m_current - 1) {
				MetaphAdd("J", "S")
			} else {
				MetaphAdd("X", "S")
			}

			AdvanceCounter(3, 1)
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
	Encode_SSIO := func() bool {
		if StringAt(+1, "SION") {
			//"abcission"
			if StringAt(-2, "CI") {
				MetaphAdd("J")
				//'mission'
			} else {
				if IsVowel(m_current - 1) {
					MetaphAdd("X")
				}
			}

			AdvanceCounter(4, 2)
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
	Encode_SS := func() bool {
		// e.g. "russian", "pressure"
		if StringAt(-1, "USSIA", "ESSUR", "ISSUR", "ISSUE") ||
			// e.g. "hessian", "assurance"
			StringAt(-1, "ESSIAN", "ASSURE", "ASSURA", "ISSUAB", "ISSUAN", "ASSIUS") {
			MetaphAdd("X")
			AdvanceCounter(3, 2)
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
	Encode_SIA := func() bool {
		// e.g. "controversial", also "fuchsia", "ch" is silent
		if StringAt(-2, "CHSIA") ||
			StringAt(-1, "RSIAL") {
			MetaphAdd("X")
			AdvanceCounter(3, 1)
			return true
		}

		// names generally get 'X' where terms, e.g. "aphasia" get 'J'
		if (StringAtStart("ALESIA", "ALYSIA", "ALISIA", "STASIA") &&
			(m_current == 3) &&
			!StringAtStart("ANASTASIA")) ||
			StringAt(-5, "DIONYSIAN") ||
			StringAt(-5, "THERESIA") {
			MetaphAdd("X", "S")
			AdvanceCounter(3, 1)
			return true
		}

		if (StringAt(0, "SIA") && ((m_current + 2) == m_last)) ||
			(StringAt(0, "SIAN") && ((m_current + 3) == m_last)) ||
			StringAt(-5, "AMBROSIAL") {
			if (IsVowel(m_current-1) || StringAt(-1, "R")) &&
				// exclude compounds based on names, or french or greek words
				!(StringAtStart("JAMES", "NICOS", "PEGAS", "PEPYS") ||
					StringAtStart("HOBBES", "HOLMES", "JAQUES", "KEYNES") ||
					StringAtStart("MALTHUS", "HOMOOUS") ||
					StringAtStart("MAGLEMOS", "HOMOIOUS") ||
					StringAtStart("LEVALLOIS", "TARDENOIS") ||
					StringAt(-4, "ALGES")) {
				MetaphAdd("J")
			} else {
				MetaphAdd("S")
			}

			AdvanceCounter(2, 1)
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
	Encode_SIO := func() bool {
		// special case, irish name
		if StringAtStart("SIOBHAN") {
			MetaphAdd("X")
			AdvanceCounter(3, 1)
			return true
		}

		if StringAt(+1, "ION") {
			// e.g. "vision", "version"
			if IsVowel(m_current-1) || StringAt(-2, "ER", "UR") {
				MetaphAdd("J")
			} else {
				// e.g. "declension"
				MetaphAdd("X")
			}

			AdvanceCounter(3, 1)
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
	Encode_Anglicisations := func() bool {
		//german & anglicisations, e.g. 'smith' match 'schmidt', 'snider' match 'schneider'
		//also, -sz- in slavic language altho in hungarian it is pronounced 's'
		if ((m_current == 0) &&
			StringAt(+1, "M", "N", "L")) ||
			StringAt(+1, "Z") {
			MetaphAdd("S", "X")

			// eat redundant 'Z'
			if StringAt(+1, "Z") {
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
	Encode_SC := func() bool {
		if StringAt(0, "SC") {
			// exception 'viscount'
			if StringAt(-2, "VISCOUNT") {
				m_current += 1
				return true
			}

			// encode "-SC<front vowel>-"
			if StringAt(+2, "I", "E", "Y") {
				// e.g. "conscious"
				if StringAt(+2, "IOUS") ||
					// e.g. "prosciutto"
					StringAt(+2, "IUT") ||
					StringAt(-4, "OMNISCIEN") ||
					// e.g. "conscious"
					StringAt(-3, "CONSCIEN", "CRESCEND", "CONSCION") ||
					StringAt(-2, "FASCIS") {
					MetaphAdd("X")
				} else if StringAt(0, "SCEPTIC", "SCEPSIS") ||
					StringAt(0, "SCIVV", "SCIRO") ||
					// commonly pronounced this way in u.s.
					StringAt(0, "SCIPIO") ||
					StringAt(-2, "PISCITELLI") {
					MetaphAdd("SK")
				} else {
					MetaphAdd("S")
				}
				m_current += 2
				return true
			}

			MetaphAdd("SK")
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
	Encode_SEA_SUI_SIER := func() bool {
		// "nausea" by itself has => NJ as a more likely encoding. Other forms
		// using "nause-" (see Encode_SEA()) have X or S as more familiar pronounciations
		if (StringAt(-3, "NAUSEA") && ((m_current + 2) == m_last)) ||
			// e.g. "casuistry", "frasier", "hoosier"
			StringAt(-2, "CASUI") ||
			(StringAt(-1, "OSIER", "ASIER") &&
				!(StringAtStart("EASIER") ||
					StringAtStart("OSIER") ||
					StringAt(-2, "ROSIER", "MOSIER"))) {
			MetaphAdd("J", "X")
			AdvanceCounter(3, 1)
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
	Encode_SEA := func() bool {
		if (StringAtStart("SEAN") && ((m_current + 3) == m_last)) ||
			(StringAt(-3, "NAUSEO") &&
				!StringAt(-3, "NAUSEAT")) {
			MetaphAdd("X")
			AdvanceCounter(3, 1)
			return true
		}

		return false
	}

	/**
	 * Encode "-S-"
	 *
	 */
	Encode_S := func() {
		if Encode_SKJ() ||
			Encode_Special_SW() ||
			Encode_SJ() ||
			Encode_Silent_French_S_Final() ||
			Encode_Silent_French_S_Internal() ||
			Encode_ISL() ||
			Encode_STL() ||
			Encode_Christmas() ||
			Encode_STHM() ||
			Encode_ISTEN() ||
			Encode_Sugar() ||
			Encode_SH() ||
			Encode_SCH() ||
			Encode_SUR() ||
			Encode_SU() ||
			Encode_SSIO() ||
			Encode_SS() ||
			Encode_SIA() ||
			Encode_SIO() ||
			Encode_Anglicisations() ||
			Encode_SC() ||
			Encode_SEA_SUI_SIER() ||
			Encode_SEA() {
			return
		}

		MetaphAdd("S")

		if StringAt(+1, "S", "Z") &&
			!StringAt(+1, "SH") {
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
	Encode_T_Initial := func() bool {
		if m_current == 0 {
			// americans usually pronounce "tzar" as "zar"
			if StringAt(+1, "SAR", "ZAR") {
				m_current++
				return true
			}

			// old 'École française d'Extrême-Orient' chinese pinyin where 'ts-' => 'X'
			if ((m_length == 3) && StringAt(+1, "SO", "SA", "SU")) ||
				((m_length == 4) && StringAt(+1, "SAO", "SAI")) ||
				((m_length == 5) && StringAt(+1, "SING", "SANG")) {
				MetaphAdd("X")
				AdvanceCounter(3, 2)
				return true
			}

			// "TS<vowel>-" at start can be pronounced both with and without 'T'
			if StringAt(+1, "S") && IsVowel(m_current+2) {
				MetaphAdd("TS", "S")
				AdvanceCounter(3, 2)
				return true
			}

			// e.g. "Tjaarda"
			if StringAt(+1, "J") {
				MetaphAdd("X")
				AdvanceCounter(3, 2)
				return true
			}

			// cases where initial "TH-" is pronounced as T and not 0 ("th")
			if (StringAt(+1, "HU") && (m_length == 3)) ||
				StringAt(+1, "HAI", "HUY", "HAO") ||
				StringAt(+1, "HYME", "HYMY", "HANH") ||
				StringAt(+1, "HERES") {
				MetaphAdd("T")
				AdvanceCounter(3, 2)
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
	Encode_TCH := func() bool {
		if StringAt(+1, "CH") {
			MetaphAdd("X")
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
	Encode_Silent_French_T := func() bool {
		// french silent T familiar to americans
		if ((m_current == m_last) && StringAt(-4, "MONET", "GENET", "CHAUT")) ||
			StringAt(-2, "POTPOURRI") ||
			StringAt(-3, "BOATSWAIN") ||
			StringAt(-3, "MORTGAGE") ||
			(StringAt(-4, "BERET", "BIDET", "FILET", "DEBUT", "DEPOT", "PINOT", "TAROT") ||
				StringAt(-5, "BALLET", "BUFFET", "CACHET", "CHALET", "ESPRIT", "RAGOUT", "GOULET",
					"CHABOT", "BENOIT") ||
				StringAt(-6, "GOURMET", "BOUQUET", "CROCHET", "CROQUET", "PARFAIT", "PINCHOT",
					"CABARET", "PARQUET", "RAPPORT", "TOUCHET", "COURBET", "DIDEROT") ||
				StringAt(-7, "ENTREPOT", "CABERNET", "DUBONNET", "MASSENET", "MUSCADET", "RICOCHET", "ESCARGOT") ||
				StringAt(-8, "SOBRIQUET", "CABRIOLET", "CASSOULET", "OUBRIQUET", "CAMEMBERT")) &&
				!StringAt(+1, "AN", "RY", "IC", "OM", "IN") {
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
	Encode_TUN_TUL_TUA_TUO := func() bool {
		// e.g. "fortune", "fortunate"
		if StringAt(-3, "FORTUN") ||
			// e.g. "capitulate"
			(StringAt(0, "TUL") &&
				(IsVowel(m_current-1) && IsVowel(m_current+3))) ||
			// e.g. "obituary", "barbituate"
			StringAt(-2, "BITUA", "BITUE") ||
			// e.g. "actual"
			((m_current > 1) && StringAt(0, "TUA", "TUO")) {
			MetaphAdd("X", "T")
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
	Encode_TUE_TEU_TEOU_TUL_TIE := func() bool {
		// 'constituent', 'pasteur'
		if StringAt(+1, "UENT") ||
			StringAt(-4, "RIGHTEOUS") ||
			StringAt(-3, "STATUTE") ||
			StringAt(-3, "AMATEUR") ||
			// e.g. "blastula", "pasteur"
			(StringAt(-1, "NTULE", "NTULA", "STULE", "STULA", "STEUR")) ||
			// e.g. "statue"
			(((m_current + 2) == m_last) && StringAt(0, "TUE")) ||
			// e.g. "constituency"
			StringAt(0, "TUENC") ||
			// e.g. "statutory"
			StringAt(-3, "STATUTOR") ||
			// e.g. "patience"
			(((m_current + 5) == m_last) && StringAt(0, "TIENCE")) {
			MetaphAdd("X", "T")
			AdvanceCounter(2, 1)
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
	Encode_TUR_TIU_Suffixes := func() bool {
		// 'adventure', 'musculature'
		if (m_current > 0) && StringAt(+1, "URE", "URA", "URI", "URY", "URO", "IUS") {
			// exceptions e.g. 'tessitura', mostly from romance languages
			if (StringAt(+1, "URA", "URO") &&
				//&& !StringAt(+1, "URIA")
				((m_current+3) == m_last)) &&
				!StringAt(-3, "VENTURA") ||
				// e.g. "kachaturian", "hematuria"
				StringAt(+1, "URIA") {
				MetaphAdd("T")
			} else {
				MetaphAdd("X", "T")
			}

			AdvanceCounter(2, 1)
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
	Encode_TI := func() bool {
		// '-tio-', '-tia-', '-tiu-'
		// except combining forms where T already pronounced e.g 'rooseveltian'
		if (StringAt(+1, "IO") && !StringAt(-1, "ETIOL")) ||
			StringAt(+1, "IAL") ||
			StringAt(-1, "RTIUM", "ATIUM") ||
			((StringAt(+1, "IAN") && (m_current > 0)) &&
				!(StringAt(-4, "FAUSTIAN") ||
					StringAt(-5, "PROUSTIAN") ||
					StringAt(-2, "TATIANA") ||
					(StringAt(-3, "KANTIAN", "GENTIAN") ||
						StringAt(-8, "ROOSEVELTIAN"))) ||
				(((m_current + 2) == m_last) &&
					StringAt(0, "TIA") &&
					// exceptions to above rules where the pronounciation is usually X
					!(StringAt(-3, "HESTIA", "MASTIA") ||
						StringAt(-2, "OSTIA") ||
						StringAtStart("TIA") ||
						StringAt(-5, "IZVESTIA"))) ||
				StringAt(+1, "IATE", "IATI", "IABL", "IATO", "IARY") ||
				StringAt(-5, "CHRISTIAN")) {
			if ((m_current == 2) && StringAtStart("ANTI")) ||
				StringAtStart("PATIO", "PITIA", "DUTIA") {
				MetaphAdd("T")
			} else if StringAt(-4, "EQUATION") {
				MetaphAdd("J")
			} else {
				if StringAt(0, "TION") {
					MetaphAdd("X")
				} else if StringAtStart("KATIA", "LATIA") {
					MetaphAdd("T", "X")
				} else {
					MetaphAdd("X", "T")
				}
			}

			AdvanceCounter(3, 1)
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
	Encode_TIENT := func() bool {
		// e.g. 'patient'
		if StringAt(+1, "IENT") {
			MetaphAdd("X", "T")
			AdvanceCounter(3, 1)
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
	Encode_TSCH := func() bool {
		//'deutsch'
		if StringAt(0, "TSCH") &&
			// combining forms in german where the 'T' is pronounced seperately
			!StringAt(-3, "WELT", "KLAT", "FEST") {
			// pronounced the same as "ch" in "chit" => X
			MetaphAdd("X")
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
	Encode_TZSCH := func() bool {
		//'neitzsche'
		if StringAt(0, "TZSCH") {
			MetaphAdd("X")
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
	Encode_TH_Pronounced_Separately := func() bool {
		//'adulthood', 'bithead', 'apartheid'
		if ((m_current > 0) &&
			StringAt(+1, "HOOD", "HEAD", "HEID", "HAND", "HILL", "HOLD",
				"HAWK", "HEAP", "HERD", "HOLE", "HOOK", "HUNT",
				"HUMO", "HAUS", "HOFF", "HARD") &&
			!StringAt(-3, "SOUTH", "NORTH")) ||
			StringAt(+1, "HOUSE", "HEART", "HASTE", "HYPNO", "HEQUE") ||
			// watch out for greek root "-thallic"
			(StringAt(+1, "HALL") &&
				((m_current + 4) == m_last) &&
				!StringAt(-3, "SOUTH", "NORTH")) ||
			(StringAt(+1, "HAM") &&
				((m_current + 3) == m_last) &&
				!(StringAtStart("GOTHAM", "WITHAM", "LATHAM") ||
					StringAtStart("BENTHAM", "WALTHAM", "WORTHAM") ||
					StringAtStart("GRANTHAM"))) ||
			(StringAt(+1, "HATCH") &&
				!((m_current == 0) || StringAt(-2, "UNTHATCH"))) ||
			StringAt(-3, "WARTHOG") ||
			// and some special cases where "-TH-" is usually pronounced 'T'
			StringAt(-2, "ESTHER") ||
			StringAt(-3, "GOETHE") ||
			StringAt(-2, "NATHALIE") {
			// special case
			if StringAt(-3, "POSTHUM") {
				MetaphAdd("X")
			} else {
				MetaphAdd("T")
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
	Encode_TTH := func() bool {
		// 'matthew' vs. 'outthink'
		if StringAt(0, "TTH") {
			if StringAt(-2, "MATTH") {
				MetaphAdd("0")
			} else {
				MetaphAdd("T0")
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
	Encode_TH := func() bool {
		if StringAt(0, "TH") {
			//'-clothes-'
			if StringAt(-3, "CLOTHES") {
				// vowel already encoded so skip right to S
				m_current += 3
				return true
			}

			//special case "thomas", "thames", "beethoven" or germanic words
			if StringAt(+2, "OMAS", "OMPS", "OMPK", "OMSO", "OMSE",
				"AMES", "OVEN", "OFEN", "ILDA", "ILDE") ||
				(StringAtStart("THOM") && (m_length == 4)) ||
				(StringAtStart("THOMS") && (m_length == 5)) ||
				StringAtStart("VAN ", "VON ") ||
				StringAtStart("SCH") {
				MetaphAdd("T")

			} else {
				// give an 'etymological' 2nd
				// encoding for "smith"
				if StringAtStart("SM") {
					MetaphAdd("0", "T")
				} else {
					MetaphAdd("0")
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
	Encode_T := func() {
		if Encode_T_Initial() ||
			Encode_TCH() ||
			Encode_Silent_French_T() ||
			Encode_TUN_TUL_TUA_TUO() ||
			Encode_TUE_TEU_TEOU_TUL_TIE() ||
			Encode_TUR_TIU_Suffixes() ||
			Encode_TI() ||
			Encode_TIENT() ||
			Encode_TSCH() ||
			Encode_TZSCH() ||
			Encode_TH_Pronounced_Separately() ||
			Encode_TTH() ||
			Encode_TH() {
			return
		}

		// eat redundant 'T' or 'D'
		if StringAt(+1, "T", "D") {
			m_current += 2
		} else {
			m_current++
		}

		MetaphAdd("T")
	}

	/**
	 * Encode "-V-"
	 *
	 */
	Encode_V := func() {
		// eat redundant 'V'
		if CharAt(m_current+1) == 'V' {
			m_current += 2
		} else {
			m_current++
		}

		MetaphAddExactApprox("V", "F")
	}

	/**
	 * Encode cases where 'W' is silent at beginning of word
	 *
	 * @return true if encoding handled in this routine, false if not
	 *
	 */
	Encode_Silent_W_At_Beginning := func() bool {
		//skip these when at start of word
		if (m_current == 0) &&
			StringAt(0, "WR") {
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
	Encode_WITZ_WICZ := func() bool {
		//polish e.g. 'filipowicz'
		if ((m_current + 3) == m_last) && StringAt(0, "WICZ", "WITZ") {
			if m_encodeVowels {
				if (len(m_primary) > 0) &&
					CharAt(len(m_primary)-1) == 'A' {
					MetaphAdd("TS", "FAX")
				} else {
					MetaphAdd("ATS", "FAX")
				}
			} else {
				MetaphAdd("TS", "FX")
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
	Encode_WR := func() bool {
		//can also be in middle of word
		if StringAt(0, "WR") {
			MetaphAdd("R")
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
	Germanic_Or_Slavic_Name_Beginning_With_W := func() bool {
		if StringAtStart("WEE", "WIX", "WAX") ||
			StringAtStart("WOLF", "WEIS", "WAHL", "WALZ", "WEIL", "WERT",
				"WINE", "WILK", "WALT", "WOLL", "WADA", "WULF",
				"WEHR", "WURM", "WYSE", "WENZ", "WIRT", "WOLK",
				"WEIN", "WYSS", "WASS", "WANN", "WINT", "WINK",
				"WILE", "WIKE", "WIER", "WELK", "WISE") ||
			StringAtStart("WIRTH", "WIESE", "WITTE", "WENTZ", "WOLFF", "WENDT",
				"WERTZ", "WILKE", "WALTZ", "WEISE", "WOOLF", "WERTH",
				"WEESE", "WURTH", "WINES", "WARGO", "WIMER", "WISER",
				"WAGER", "WILLE", "WILDS", "WAGAR", "WERTS", "WITTY",
				"WIENS", "WIEBE", "WIRTZ", "WYMER", "WULFF", "WIBLE",
				"WINER", "WIEST", "WALKO", "WALLA", "WEBRE", "WEYER",
				"WYBLE", "WOMAC", "WILTZ", "WURST", "WOLAK", "WELKE",
				"WEDEL", "WEIST", "WYGAN", "WUEST", "WEISZ", "WALCK",
				"WEITZ", "WYDRA", "WANDA", "WILMA", "WEBER") ||
			StringAtStart("WETZEL", "WEINER", "WENZEL", "WESTER", "WALLEN", "WENGER",
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
			StringAtStart("WISEMAN", "WINKLER", "WILHELM", "WELLMAN", "WAMPLER", "WACHTER",
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
			StringAtStart("WESTPHAL", "WICKLUND", "WEISSMAN", "WESTLUND", "WOLFGANG", "WILLHITE",
				"WEISBERG", "WALRAVEN", "WOLFGRAM", "WILHOITE", "WECHSLER", "WENDLING",
				"WESTBERG", "WENDLAND", "WININGER", "WHISNANT", "WESTRICK", "WESTLING",
				"WESTBURY", "WEITZMAN", "WEHMEYER", "WEINMANN", "WISNESKI", "WHELCHEL",
				"WEISHAAR", "WAGGENER", "WALDROUP", "WESTHOFF", "WIEDEMAN", "WASINGER",
				"WINBORNE") ||
			StringAtStart("WHISENANT", "WEINSTEIN", "WESTERMAN", "WASSERMAN", "WITKOWSKI", "WEINTRAUB",
				"WINKELMAN", "WINKFIELD", "WANAMAKER", "WIECZOREK", "WIECHMANN", "WOJTOWICZ",
				"WALKOWIAK", "WEINSTOCK", "WILLEFORD", "WARKENTIN", "WEISINGER", "WINKLEMAN",
				"WILHEMINA") ||
			StringAtStart("WISNIEWSKI", "WUNDERLICH", "WHISENHUNT", "WEINBERGER", "WROBLEWSKI",
				"WAGUESPACK", "WEISGERBER", "WESTERVELT", "WESTERLUND", "WASILEWSKI",
				"WILDERMUTH", "WESTENDORF", "WESOLOWSKI", "WEINGARTEN", "WINEBARGER",
				"WESTERBERG", "WANNAMAKER", "WEISSINGER") ||
			StringAtStart("WALDSCHMIDT", "WEINGARTNER", "WINEBRENNER") ||
			StringAtStart("WOLFENBARGER") ||
			StringAtStart("WOJCIECHOWSKI") {
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
	Encode_Initial_W_Vowel := func() bool {
		if (m_current == 0) && IsVowel(m_current+1) {
			//Witter should match Vitter
			if Germanic_Or_Slavic_Name_Beginning_With_W() {
				if m_encodeVowels {
					MetaphAddExactApprox("A", "VA", "A", "FA")
				} else {
					MetaphAddExactApprox("A", "V", "A", "F")
				}
			} else {
				MetaphAdd("A")
			}

			m_current++
			// don't encode vowels twice
			m_current = SkipVowels(m_current)
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
	Encode_WH := func() bool {
		if StringAt(0, "WH") {
			// cases where it is pronounced as H
			// e.g. 'who', 'whole'
			if (CharAt(m_current+2) == 'O') &&
				// exclude cases where it is pronounced like a vowel
				!(StringAt(+2, "OOSH") ||
					StringAt(+2, "OOP", "OMP", "ORL", "ORT") ||
					StringAt(+2, "OA", "OP")) {
				MetaphAdd("H")
				AdvanceCounter(3, 2)
				return true
			} else {
				// combining forms, e.g. 'hollowhearted', 'rawhide'
				if StringAt(+2, "IDE", "ARD", "EAD", "AWK", "ERD",
					"OOK", "AND", "OLE", "OOD") ||
					StringAt(+2, "EART", "OUSE", "OUND") ||
					StringAt(+2, "AMMER") {
					MetaphAdd("H")
					m_current += 2
					return true
				} else if m_current == 0 {
					MetaphAdd("A")
					m_current += 2
					// don't encode vowels twice
					m_current = SkipVowels(m_current)
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
	Encode_Eastern_European_W := func() bool {
		//Arnow should match Arnoff
		if ((m_current == m_last) && IsVowel(m_current-1)) ||
			StringAt(-1, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
			(StringAt(0, "WICKI", "WACKI") && ((m_current + 4) == m_last)) ||
			StringAt(0, "WIAK") && ((m_current+3) == m_last) ||
			StringAtStart("SCH") {
			MetaphAddExactApprox("", "V", "", "F")
			m_current++
			return true
		}

		return false
	}

	/**
	 * Encode "-W-"
	 *
	 */
	Encode_W := func() {
		if Encode_Silent_W_At_Beginning() ||
			Encode_WITZ_WICZ() ||
			Encode_WR() ||
			Encode_Initial_W_Vowel() ||
			Encode_WH() ||
			Encode_Eastern_European_W() {
			return
		}

		// e.g. 'zimbabwe'
		if m_encodeVowels &&
			StringAt(0, "WE") &&
			((m_current + 1) == m_last) {
			MetaphAdd("A")
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
	Encode_Initial_X := func() bool {
		// current chinese pinyin spelling
		if StringAtStart("XIA", "XIO", "XIE") ||
			StringAtStart("XU") {
			MetaphAdd("X")
			m_current++
			return true
		}

		// else
		if m_current == 0 {
			MetaphAdd("S")
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
	Encode_Greek_X := func() bool {
		// 'xylophone', xylem', 'xanthoma', 'xeno-'
		if StringAt(+1, "YLO", "YLE", "ENO") ||
			StringAt(+1, "ANTH") {
			MetaphAdd("S")
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
	Encode_X_Special_Cases := func() bool {
		// 'luxury'
		if StringAt(-2, "LUXUR") {
			MetaphAddExactApprox("GJ", "KJ")
			m_current++
			return true
		}

		// 'texeira' portuguese/galician name
		if StringAtStart("TEXEIRA") ||
			StringAtStart("TEIXEIRA") {
			MetaphAdd("X")
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
	Encode_X_To_H := func() bool {
		// TODO: look for other mexican indian words
		// where 'X' is usually pronounced this way
		if StringAt(-2, "OAXACA") ||
			StringAt(-3, "QUIXOTE") {
			MetaphAdd("H")
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
	Encode_X_Vowel := func() bool {
		// e.g. "sexual", "connexion" (british), "noxious"
		if StringAt(+1, "UAL", "ION", "IOU") {
			MetaphAdd("KX", "KS")
			AdvanceCounter(3, 1)
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
	Encode_French_X_Final := func() bool {
		//french e.g. "breaux", "paix"
		if !((m_current == m_last) &&
			(StringAt(-3, "IAU", "EAU", "IEU") ||
				StringAt(-2, "AI", "AU", "OU", "OI", "EU"))) {
			MetaphAdd("KS")
		}

		return false
	}

	/**
	 * Encode "-X-"
	 *
	 */
	Encode_X := func() {
		if Encode_Initial_X() ||
			Encode_Greek_X() ||
			Encode_X_Special_Cases() ||
			Encode_X_To_H() ||
			Encode_X_Vowel() ||
			Encode_French_X_Final() {
			return
		}

		// eat redundant 'X' or other redundant cases
		if StringAt(+1, "X", "Z", "S") ||
			// e.g. "excite", "exceed"
			StringAt(+1, "CI", "CE") {
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
	Encode_ZZ := func() bool {
		// "abruzzi", 'pizza'
		if (CharAt(m_current+1) == 'Z') &&
			((StringAt(+2, "I", "O", "A") &&
				((m_current + 2) == m_last)) ||
				StringAt(-2, "MOZZARELL", "PIZZICATO", "PUZZONLAN")) {
			MetaphAdd("TS", "S")
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
	Encode_ZU_ZIER_ZS := func() bool {
		if ((m_current == 1) && StringAt(-1, "AZUR")) ||
			(StringAt(0, "ZIER") &&
				!StringAt(-2, "VIZIER")) ||
			StringAt(0, "ZSA") {
			MetaphAdd("J", "S")

			if StringAt(0, "ZSA") {
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
	Encode_French_EZ := func() bool {
		if ((m_current == 3) && StringAt(-3, "CHEZ")) ||
			StringAt(-5, "RENDEZ") {
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
	Encode_German_Z := func() bool {
		if ((m_current == 2) && ((m_current + 1) == m_last) && StringAt(-2, "NAZI")) ||
			StringAt(-2, "NAZIFY", "MOZART") ||
			StringAt(-3, "HOLZ", "HERZ", "MERZ", "FITZ") ||
			(StringAt(-3, "GANZ") && !IsVowel(m_current+1)) ||
			StringAt(-4, "STOLZ", "PRINZ") ||
			StringAt(-4, "VENEZIA") ||
			StringAt(-3, "HERZOG") ||
			// german words beginning with "sch-" but not schlimazel, schmooze
			(strings.Contains(string(m_inWord), "SCH") && !(StringAtEnd("IZE", "OZE", "ZEL"))) ||
			((m_current > 0) && StringAt(0, "ZEIT")) ||
			StringAt(-3, "WEIZ") {
			if (m_current > 0) && CharAt(m_current-1) == 'T' {
				MetaphAdd("S")
			} else {
				MetaphAdd("TS")
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
	Encode_ZH := func() bool {
		//chinese pinyin e.g. 'zhao', also english "phonetic spelling"
		if CharAt(m_current+1) == 'H' {
			MetaphAdd("J")
			m_current += 2
			return true
		}

		return false
	}

	/**
	 * Encode "-Z-"
	 *
	 */
	Encode_Z := func() {
		if Encode_ZZ() ||
			Encode_ZU_ZIER_ZS() ||
			Encode_French_EZ() ||
			Encode_German_Z() {
			return
		}

		if Encode_ZH() {
			return
		} else {
			MetaphAdd("S")
		}

		// eat redundant 'Z'
		if CharAt(m_current+1) == 'Z' {
			m_current += 2
		} else {
			m_current++
		}
	}

	/**
	 * Encodes input string to one or two key values according to Metaphone 3 rules.
	 *
	 */
	Encode := func() {
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
			switch CharAt(m_current) {
			case 'B':
				Encode_B()
			case 'ß':
			case 'Ç':
				MetaphAdd("S")
				m_current++
			case 'C':
				Encode_C()
			case 'D':
				Encode_D()
			case 'F':
				Encode_F()
			case 'G':
				Encode_G()
			case 'H':
				Encode_H()
			case 'J':
				Encode_J()
			case 'K':
				Encode_K()
			case 'L':
				Encode_L()
			case 'M':
				Encode_M()
			case 'N':
				Encode_N()
			case 'Ñ':
				MetaphAdd("N")
				m_current++
			case 'P':
				Encode_P()
			case 'Q':
				Encode_Q()
			case 'R':
				Encode_R()
			case 'S':
				Encode_S()
			case 'T':
				Encode_T()
			case 'Ð', 'Þ': // eth, thorn
				MetaphAdd("0")
				m_current++
			case 'V':
				Encode_V()
			case 'W':
				Encode_W()
			case 'X':
				Encode_X()
			case '':
				MetaphAdd("X")
				m_current++
			case '':
				MetaphAdd("S")
				m_current++
			case 'Z':
				Encode_Z()
			default:
				if IsVowel(m_current) {
					Encode_Vowels()
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

	m_encodeVowels = m.doVowels
	m_encodeExact = m.doExact
	m_metaphLength = m.maxlen
	m_inWord = []rune(strings.ToUpper(word))
	m_length = len(m_inWord)
	Encode()
	return string(m_primary), string(m_secondary)
}
