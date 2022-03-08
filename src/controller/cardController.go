package controller

import (
	"math/rand"
	"strings"
	"time"
)

var cardNumber = [13]string{
	"A",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"J",
	"Q",
	"K",
}

var suits = [4]string{
	"♣",
	"♦",
	"♥",
	"♠",
}

func RemoveDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func DrawCards(count int, prev ...[]string) []string {
	result := []string{}
	var combineString strings.Builder
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		combineString.WriteString(cardNumber[rand.Intn(len(cardNumber))])
		combineString.WriteString(suits[rand.Intn(len(suits))])
		result = append(result, combineString.String())
		combineString.Reset()
	}
	if len(prev) > 0 {
		result = append(result, prev[0]...)
	}
	result = RemoveDuplicates(result)
	if len(result) != count {
		return append(result, DrawCards(count-len(result))...)
	}
	return result
}
