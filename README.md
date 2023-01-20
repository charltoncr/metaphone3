<!-- title: Metaphone3 Read Me -->
<!-- $Id: README.md,v 1.16 2023-01-20 09:07:28-05 ron Exp $ -->

# Metaphone3

Metaphone3 generates two keys for a word.  Two words sounds similar when
either non-empty key of one word matches either of the other word's keys.
Producing keys for the words in a word list and matching them with the keys
for an attempted spelling allows a spelling checker program to suggest correct
spellings.

Metaphone3 is the relatively new (2010) and improved successor to Double
Metaphone (1999) and Metaphone (1990).

- func NewMetaphone3(maxLen int) *Metaphone3
- func (m *Metaphone3) SetEncodeVowels(b bool)
- func (m *Metaphone3) SetEncodeExact(b bool)
- func (m *Metaphone3) SetMaxLength(max int)
- func (m *Metaphone3) Encode(word string) (key1, key2 string)

**NewMetaphone3** returns a new metaphone3 instance with its own maximum key
length. Both encodeVowels and encodeExact default to false.

**Encode** returns main and alternate keys for a word.  Keys will match for similar sounding words.

**SetEncodeVowels** determines whether vowels after the first character are encoded.

**SetEncodeExact** determines whether consonants are encoded more precisely.

**SetMaxLength** sets the the maximum allowed length for metaphone3 keys.

Example use:

```go
import "github.com/charltoncr/metaphone3"
// ...
meta := metaphone3.NewMetaphone3(4)
m, m2 := meta.Encode("knewmoanya")
n, n2 := meta.Encode("pneumonia")
if m == n || m == n2 || m2 == n || len(m2) > 0 && m2 == n2 {
    // match
}
// m is "NMN", as is n, so the two spellings match.
// The maximum allowed length for each of m, m2,
// n and n2 is 4 in this example.
```

# Metaphone3 Convenience Functions

The Metaphone3 convenience functions ease the use of Metaphone3.
Two function calls are sufficient to read all words in a file, create a
map of words that have the same metaphone return values, and find all words
in the map that match a given word/misspelling.  See the example below.

- func NewMetaphMap(wordlist []string, maxLen int) (*MetaphMap, error)
- func NewMetaphMapFromFile(fileName string, maxLen int) (
  metaph *MetaphMap, err error)
- func NewMetaphMapFromFileExact(fileName string, maxLen int, encodeVowels, encodeExact bool) (metaph *MetaphMap, err error)
- func (metaph *MetaphMap) AddWordsToMap(wordlist []string)
- func (metaph *MetaphMap) AddWordsFromFile(fileName string) error
- func (metaph *MetaphMap) MatchWord(word string) (output []string)
- func (metaph *MetaphMap) Len() int

**NewMetaphMap** returns a MetaphMap made from a wordlist and a maximum length
for the metaphone3 keys.

**NewMetaphMapFromFile** returns a MetaphMap made from a word list file and
a maximum length for the metaphone3 keys.

**NewMetaphMapFromFileExact** works like NewMetaphMapFromFile but allows
control of how vowels and consonants are encoded.

**AddWordsToMap** adds words from a word list to an existing MetaphMap.
This is useful for combining word lists, for example, a general word list and
a user's personal word list.

**AddWordsFromFile** adds words from a file to an existing MetaphMap.

**MatchWord** returns all words in a MetaphMap that sound like word. Case in
word is ignored, as are non-alphabetic characters.

**Len** returns the number of sounds-alike keys in the metaph map.

Example use:

```go
package main

import (
    "fmt"
    "github.com/charltoncr/metaphone3"
)
func main() {
    // The file specified by fileName should contain a comprehesive word
    // list with one word per line.  (Error check is omitted for brevity.)
    fileName := "spellCheckerWords.txt" // (can be a *.txt.gz file)
    metaphMap, _ := metaphone3.NewMetaphMapFromFile(fileName, 4)
    matches := metaphMap.MatchWord("knewmoanya")
    for _, word := range matches {
        fmt.Println(word)
    }
}
```

Ron Charlton
