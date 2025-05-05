package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func dotSepar(words []string, count int) {
	for _, word := range words {
		fmt.Print(word)
		for range count {
			fmt.Print(".")
		}
	}
}

func commaSepar(words []string, count int) {
	for _, word := range words {
		fmt.Print(word)
		for range count {
			fmt.Print(",")
		}
	}
}

func parseSeparator(input string) (sep string, count int) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	var digits strings.Builder
	var separator strings.Builder
	hasDigits := false
	hasSeparator := false

	for _, r := range input {
		if unicode.IsDigit(r) {
			digits.WriteRune(r)
			hasDigits = true
		} else if !unicode.IsSpace(r) {
			separator.WriteRune(r)
			hasSeparator = true
		}
	}

	if hasDigits {
		if num, err := strconv.Atoi(digits.String()); err == nil && num > 0 {
			count = num
		}
	}
	if hasSeparator {
		sep = string([]rune(separator.String())[0])
	}

	return
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var wordCount int
	words := make([]string, 0)
	var separatorInput string

	fmt.Println("Введите количество слов:")
	fmt.Fscan(in, &wordCount)
	in.ReadString('\n')

	fmt.Println("Вводите слова:")
	for range wordCount {
		var word string
		fmt.Fscan(in, &word)
		words = append(words, word)
	}

	fmt.Println("Введите сепаратор и количество повторений (например, '.3' или '2,' или ':'):")

	in.ReadString('\n')
	separatorInput, _ = in.ReadString('\n')
	separatorInput = strings.TrimSpace(separatorInput)

	separator, sepCount := parseSeparator(separatorInput)

	if separator == "." {
		dotSepar(words, sepCount)
	} else if separator == "," {
		commaSepar(words, sepCount)
	} else {
		fmt.Fprint(out, "Введены некорректные данные")
	}
}
