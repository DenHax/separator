package main

import (
	"bufio"
	"fmt"
	"os"
)

func dotSepar(words []string) {
	for _, word := range words {
		fmt.Print(word)
		fmt.Print(".")
	}
}

func commaSepar(words []string) {
	for _, word := range words {
		fmt.Print(word)
		fmt.Print(",")
	}
}

func semiSepar(words []string) {
	for _, word := range words {
		fmt.Print(word)
		fmt.Print(":")
	}
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	words := make([]string, 0)
	var separator string

	fmt.Println("Введите количество слов:")
	fmt.Fscan(in, &n)
	in.ReadString('\n')

	fmt.Println("Вводите слова:")
	for range n {
		var word string
		fmt.Fscan(in, &word)
		words = append(words, word)
	}

	fmt.Println("Введите сепаратор:")
	fmt.Fscan(in, &separator)
	in.ReadString('\n')

	if separator == "." {
		dotSepar(words)
	} else if separator == "," {
		commaSepar(words)
	} else {
		fmt.Fprint(out, "Введены некорректные данные")
	}
	// TODO: semiSepar doesnt work right
	// else if separator == ";" {
	//   semiSepar(words)
	// }

}
