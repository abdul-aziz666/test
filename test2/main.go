package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "aku suka makan nasi"
	fmt.Println("\nresult : ", searchWords(kalimat, palindrome(kalimat)))
}

func palindrome(kalimat string) (result string) {
	kalimat1 := strings.ReplaceAll(kalimat, " ", "")
	kalimat1 = strings.ToLower(kalimat1)
	var ok bool
	var val string
	var index int

	for i := (len(kalimat1) - 1); i >= 0; i-- {
		if strings.Contains(kalimat1, val) {
			ok = true
			index++
		} else {
			ok = false
			i = i + index
			index = 0
		}
		if !ok {
			val = ""
		}
		val += fmt.Sprintf("%c", kalimat1[i])
	}
	return val
}
func searchWords(word string, contain string) string {
	resultList := []string{}
	wordList := strings.Split(word, " ")
	for _, v := range wordList {
		if strings.Contains(contain, v) {
			resultList = append(resultList, v)
		}
	}
	return strings.Join(resultList, " ")
}
