/**
 * @Author: Noaghzil
 * @Date:   2023-01-07 21:00:54
 * @Last Modified by:   Noaghzil
 * 内存少，耗时多
 * @Last Modified time: 2023-01-07 22:17:33
 */

package main

import (
	"fmt"
	"unicode"
)

func areNumbersAscending(s string) bool {
	var result bool = true
	pre, i := 0, 0
	for i < len(s) {
		fmt.Println("one:", s[i], "i:", i)
		if unicode.IsDigit(rune(s[i])) {
			cur := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				cur = cur*10 + int(s[i]-'0')
				i++
			}
			if cur <= pre {
				result = false
			}
			pre = cur
		} else {
			i++
		}
	}
	fmt.Println("result: ", result)
	return true
}

func main() {
	var sentence_list []string = []string{
		// "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
		// "hello world 5 x 5",
		// "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
		"4 5 11 26",
	}
	for _, sentence := range sentence_list {
		fmt.Println(sentence)
		areNumbersAscending(sentence)
	}
}
