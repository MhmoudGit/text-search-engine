package main

import (
	"strings"
	"unicode"

	snowball "github.com/kljensen/snowball/english"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func LowerCaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func StopWordFilter(tokens []string) []string {
	var stopWords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {}, "in": {},
		"are": {}, "is": {}, "to": {}, "of": {}, "that": {}, "the": {},
	}
	r := make([]string, len(tokens))
	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func StemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowball.Stem(token, false)
	}
	return r
}

func Analyze(text string) []string {
	tokens := Tokenize(text)
	tokens = LowerCaseFilter(tokens)
	tokens = StopWordFilter(tokens)
	tokens = StemmerFilter(tokens)
	return tokens
}