package main

import (
	"testing"
)

func TestConvertWord(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"Empty slice", "", ""},
		{"Consonant", "n", "n"},
		{"Consonant", "pig", "igpay"},
		{"Consonant", "latin", "atinlay"},
		{"Consonant", "me", "emay"},
		{"Uppercase consonant", "Pig", "igpay"},
		{"Consonant cluster", "smile", "ilesmay"},
		{"Consonant cluster", "string", "ingstray"},
		{"Vowel", "i", "iay"},
		{"Vowel", "are", "areay"},
		{"Vowel", "eat", "eatay"},
		{"Starts with whitespace", " whitespace", " whitespace"},
		{"Starts with number", "1number", "1number"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ConvertWord(test.in)
			if err != nil {
				t.Errorf("uexpected error: %s", err)
			}
			if got != test.want {
				t.Errorf("%s (len: %v) != %s (len: %v)", got, len(got), test.want, len(test.want))
			}
		})
	}
}

func TestConvertSentence(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"Empty slice", "", ""},
		{"Phrase", "pig latin", "igpay atinlay"},
		{"Phrase", "rock n roll", "ockray n ollray"},
		{"Punctuation", "hi!", "ihay!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ConvertSentence(test.in)
			if err != nil {
				t.Errorf("uexpected error: %s", err)
			}
			if got != test.want {
				t.Errorf("%s (len: %v) != %s (len: %v)", got, len(got), test.want, len(test.want))
			}
		})
	}
}

// for _, test := range tests {
// 	t.Run(test.name, func(t *testing.T) {
// 		buf := new(bytes.Buffer)
// 		w := NewWriter(buf)
// 		w.Write(test.in))
// 		got := buf.Bytes()
// 		if !bytes.Equal(got, test.want)) {
// 			t.Errorf("%s != %s", got, test.want)
// 		}
// 	})
// }
