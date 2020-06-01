package piglat

import (
	"regexp"
	"strings"
)

var (
	consonants = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}
	vowels     = []rune{'a', 'e', 'i', 'o', 'u'}
	suffix     = "ay"
)

// ConvertWord converts word w to Pig Latin. The word should already be trimmed and
// start with proper English consonant or vowel, otherwise, it will be returned
// without any change. You may want to consider ConvertSentence instead.
// Uppercase letters become lowercase.
func ConvertWord(word string) (_ string, err error) {
	var builder strings.Builder

	if len(word) > 0 {
		lowerWord := strings.ToLower(word)
		first, _, err := strings.NewReader(lowerWord).ReadRune()
		if err != nil {
			return "", err
		}
		if isConsonant(first) {
			var vowelFound bool
			for i, r := range lowerWord {
				if isVowel(r) {
					_, err = builder.WriteString(lowerWord[i:]) // after prefix
					_, err = builder.WriteString(lowerWord[:i]) // prefix
					_, err = builder.WriteString(suffix)
					vowelFound = true
					break
				}
			}
			if !vowelFound { // no vowel found, just copy
				_, err = builder.WriteString(word)
			}
		} else if isVowel(first) {
			_, err = builder.WriteString(lowerWord)
			_, err = builder.WriteString(suffix)
		} else {
			_, err = builder.WriteString(word)
		}
	}

	return builder.String(), err
}

// ConvertSentence converts English words contained in the provided string to Pig Latin.
// See ConvertWord for the rules of conversion.
func ConvertSentence(sentence string) (_ string, err error) {
	var builder strings.Builder

	// leverage regexp to find all possible words
	reWord := regexp.MustCompile(`[A-Za-z]+`)
	words := reWord.FindAllStringIndex(sentence, -1)

	var prevWordEnd int
	for _, bounds := range words {
		wordStart, wordEnd := bounds[0], bounds[1]
		word := sentence[wordStart:wordEnd]

		var cWord string
		cWord, err = ConvertWord(word)

		_, err = builder.WriteString(sentence[prevWordEnd:wordStart]) // what's between words
		_, err = builder.WriteString(cWord)

		prevWordEnd = wordEnd
	}
	if prevWordEnd < len(sentence) {
		builder.WriteString(sentence[prevWordEnd:]) // the rest e.g. punctuation
	}

	return builder.String(), err
}

func isConsonant(b rune) bool {
	for _, c := range consonants {
		if b == c {
			return true
		}
	}
	return false
}

func isVowel(b rune) bool {
	for _, v := range vowels {
		if b == v {
			return true
		}
	}
	return false
}
