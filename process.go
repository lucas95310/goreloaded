package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ProcessText(text string) string {
	words := strings.Fields(text)

	words = handleModifiers(words)

	text = strings.Join(words, " ")

	text = fixPunctuation(text)
	text = fixQuotes(text)
	text = fixArticles(text)

	return text
}

func handleModifiers(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		switch word {

		case "(hex)":
			if len(result) > 0 {
				n, _ := strconv.ParseInt(result[len(result)-1], 16, 64)
				result[len(result)-1] = fmt.Sprintf("%d", n)
			}

		case "(bin)":
			if len(result) > 0 {
				n, _ := strconv.ParseInt(result[len(result)-1], 2, 64)
				result[len(result)-1] = fmt.Sprintf("%d", n)
			}

		case "(up)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}

		case "(low)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}

		case "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = capitalize(result[len(result)-1])
			}

		default:
			matched := regexp.MustCompile(`^\((up|low|cap),\s*(\d+)\)$`).FindStringSubmatch(word)

			if matched != nil {
				action := matched[1]
				count, _ := strconv.Atoi(matched[2])

				start := len(result) - count
				if start < 0 {
					start = 0
				}

				for j := start; j < len(result); j++ {
					switch action {
					case "up":
						result[j] = strings.ToUpper(result[j])
					case "low":
						result[j] = strings.ToLower(result[j])
					case "cap":
						result[j] = capitalize(result[j])
					}
				}
			} else {
				result = append(result, word)
			}
		}
	}

	return result
}

func capitalize(word string) string {
	if word == "" {
		return word
	}

	runes := []rune(strings.ToLower(word))
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

func fixPunctuation(text string) string {
	re := regexp.MustCompile(`\s*([.,!?;:]+)\s*`)
	text = re.ReplaceAllString(text, "$1 ")

	text = strings.ReplaceAll(text, "... ", "... ")
	text = strings.TrimSpace(text)

	return text
}

func fixQuotes(text string) string {
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")
	return text
}

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		current := strings.ToLower(words[i])

		if current == "a" {
			next := strings.ToLower(words[i+1])

			if len(next) > 0 {
				first := next[0]

				if strings.ContainsRune("aeiouh", rune(first)) {
					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}
