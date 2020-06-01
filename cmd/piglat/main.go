package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ortymid/piglat/pkg/piglat"
)

func main() {
	fmt.Println("Piglat 0.1.0\nEnter English text and press Enter to see its translation to Pig Latin.")
	for {
		fmt.Print(">>> ")

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Cannot read your text. Reason:", err, "\nTry again...")
			continue
		}

		writer := piglat.NewWriter(os.Stdout)
		_, err = writer.WriteString(text)
		if err != nil {
			fmt.Println("Cannot convert your text. Reason:", err, "\nTry again...")
			continue
		}
	}
}
