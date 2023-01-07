/**
 * @Author: Noaghzil
 * @Date:   2023-01-07 22:17:03
 * @Last Modified by:   Noaghzil
 * 内存多，耗时少
 * @Last Modified time: 2023-01-07 22:19:11
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseWordFromSentence(sentence string) []string {
	sep := " "
	words := strings.Split(sentence, sep)
	return words
}

func areNumbersAscending(sentence string) bool {
	var result bool = true
	wordList := parseWordFromSentence(sentence)
	var minNum int64 = 0
	for _, value := range wordList {
		if it64, err := strconv.ParseInt(value, 10, 64); err == nil {
			if minNum < it64 {
				minNum = it64
			} else {
				result = false
			}
		}
	}
	fmt.Println("result: ", result)
	return result
}

func main() {
	var sentence_list []string = []string{
		"1 box has 3 blue 4 red 6 green and 12 yellow marbles",
		"hello world 5 x 5",
		"sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
		"4 5 11 26",
	}
	for _, sentence := range sentence_list {
		fmt.Println(sentence)
		areNumbersAscending(sentence)
	}
}
