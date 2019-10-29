//solves https://leetcode.com/problems/find-and-replace-pattern/submissions/

package main

import (
	"fmt"
	"sync"
)

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	fmt.Println(findAndReplacePattern(words, "abb"))
}

func findAndReplacePattern(words []string, pattern string) []string {
	//spawn goroutines to analyse each word, sending matches to a channel
	ch := make(chan string, len(words))
	wg := new(sync.WaitGroup)
	wg.Add(len(words))
	for i := range words {
		go sendIfMatchesPattern(words[i], pattern, wg, ch)
	}
	wg.Wait()
	close(ch)
	//consume channel with matched words
	var strings []string
	for s := range ch {
		strings = append(strings, s)
	}
	return strings
}

func sendIfMatchesPattern(word string, pattern string, wg *sync.WaitGroup, matches chan string) {
	defer wg.Done()
	//represent bijection with two maps
	wordToPattern := make(map[byte]byte)
	patternToWord := make(map[byte]byte)
	//verify each rune from word respects bijection
	for i := range word {
		if ok := respectsMapping(wordToPattern, word[i], pattern[i]); !ok {
			return
		}
		if ok := respectsMapping(patternToWord, pattern[i], word[i]); !ok {
			return
		}
	}
	matches <- word
}

func respectsMapping(m map[byte]byte, key byte, val byte) bool {
	if last, seen := m[key]; seen {
		return last == val
	}
	m[key] = val
	return true
}
