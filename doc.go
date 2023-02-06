// Package metaphone3 consists of a metaphone3 implementation in
// metaphone3.go and convenience functions in convenience.go and wordFreq.go.
// metaphone3.go can be used standalone; it does not require convenience.go
// or wordFreq.go to function.
//
// Metaphone3 is the relatively new (2010) and improved successor to Double
// Metaphone (1999) and Metaphone (1990).  It encodes similar sounding words
// to the same key, useful for finding matching words for misspellings with
// spelling checkers.  See the comments near the top of metaphone3.go for
// a detailed explanation of how Metaphone3 works.
//
// Typical use:
//
//	m := metaphone3.NewMetaphone3(4)
//	main, alternate := m.Encode("knewmoanya")
//
// Convenience functions simplify the use of metaphone3, eliminating the
// direct use of NewMetaphone3 and Encode.
//
// Convenience function NewMetaphMapFromFile returns a MetaphMap made
// from a file containing a word list, and a maximum length for the
// Metaphone3 return values.
//
// Convenience function MatchWord returns all words in the MetaphMap
// that sound like a given word or misspelling.  Letter case and non-alphabetic
// characters in word are ignored.  Typical use:
//
//		import "fmt"
//		import "metaphone3"
//		// ...
//		// File myWordlist.txt should contain a comprehensive word
//	 	// list, one word per line.  Errors are ignored here.
//		m, _ := metaphone3.NewMetaphMapFromFile("myWordlist.txt", 4)
//		matches := m.MatchWord("knewmoanya")
//		for _, word = range matches {
//			fmt.Println(word)
//		}
//
// Other convenience functions add more words to m above, and
// provide more control over how vowels and consonants are encoded, as well as
// working with word lists in string slices.
//
// metaphone3.go is based on a file that is copyright 2010 by Laurence Philips
// and is open source.
//
// You can buy a later metaphone 3 version with better proper name encoding
// and other corrections, and in a variety of programming languages (but not
// Go), online for $240 as of 2023-01-10.
package metaphone3
