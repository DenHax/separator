package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doSeparate(words []string, count int, separator string) {
	for _, word := range words {
		fmt.Print(word)
		for range count {
			fmt.Print(separator)
		}
	}
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var wordCount int
	words := make([]string, 0)

	fmt.Println("Введите количество слов:")
	fmt.Fscan(in, &wordCount)
	in.ReadString('\n')

	fmt.Println("Вводите слова:")
	for range wordCount {
		var word string
		fmt.Fscan(in, &word)
		words = append(words, word)
	}

	in.ReadString('\n')

	var separator string
	var sepCount int = 1
	var sepCountInput string

	fmt.Println("Введите сепаратор:")
	fmt.Fscan(in, &separator)
	in.ReadString('\n')
	fmt.Println("Введите количество повторений сепаратора (по умолчанию: 1):")

	sepCountInput, _ = in.ReadString('\n')
	sepCountInput = strings.TrimSpace(sepCountInput)

	if sepCountInput != "" {
		sepCount, _ = strconv.Atoi(sepCountInput)
	}

	doSeparate(words, sepCount, separator)
}
