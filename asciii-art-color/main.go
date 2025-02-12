package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// colors maps color names to their ANSI escape codes
var colors = map[string]string{
	"black":             "\033[30m",
	"red":               "\033[31m",
	"green":             "\033[32m",
	"yellow":            "\033[33m",
	"blue":              "\033[34m",
	"magenta":           "\033[35m",
	"cyan":              "\033[36m",
	"white":             "\033[37m",
	"gray":              "\033[90m",
	"light_red":         "\033[91m",
	"light_green":       "\033[92m",
	"light_yellow":      "\033[93m",
	"light_blue":        "\033[94m",
	"light_magenta":     "\033[95m",
	"light_cyan":        "\033[96m",
	"light_white":       "\033[97m",
	"#ff0000":           "\033[31m",
	"rgb(255, 0, 0)":    "\033[31m",
	"hsl(0, 100%, 50%)": "\033[31m",
	"orange":            "\033[38;5;208m",
	"violet":            "\033[38;5;165m",
	"pink":              "\033[38;5;218m",
	"sky_blue":          "\033[38;5;45m",
	"olive_green":       "\033[38;5;100m",
	"light_olive":       "\033[38;5;142m",
	"turquoise":         "\033[38;5;80m",
	"lavender":          "\033[38;5;183m",
	"lime_green":        "\033[38;5;148m",
	"salmon":            "\033[38;5;173m",
}

// Ascii function converts text to ASCII art using specified color and letters
func Ascii(word, templ, color, letters string) {
	// Check if word contains invalid characters
	for i := 0; i < len(word); i++ {
		if (word[i] < 32 || word[i] > 127) && word[i] != 10 {
			fmt.Println("Incorrect input.")
			return
		}
	}

	// Read the ASCII art template file
	template, err := ioutil.ReadFile(templ + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Split the template into individual characters
	splitted := strings.Split(string(template), "\n\n")
	if len(splitted) != 95 {
		fmt.Println("Incorrect template:", len(splitted), "instead of 95")
		return
	}

	lines := strings.Split(word, "\n")
	res := ""

	// Process each line of the input text
	for _, line := range lines {
		if line == "" && res != "" {
			res += string('\n')
			continue
		}

		// Process each row of the ASCII art template
		for row := 0; row < 8; row++ {
			for d := 0; d < len(line); d++ {
				ch := line[d]
				if ch == letters[0] {
					if strings.HasPrefix(line[d:], letters) {
						// Apply color to specific letters
						for i := d; i < d+len(letters); i++ {
							temp := fmt.Sprintf("%s%s\033[0m", color, strings.Split(splitted[line[i]-32], "\n")[row])
							res += temp
						}
						d += len(letters) - 1
					} else {
						res += strings.Split(splitted[ch-32], "\n")[row]
					}
				} else {
					res += strings.Split(splitted[ch-32], "\n")[row]
				}
			}
			res += string('\n')
		}
	}

	// Print the resulting ASCII art
	fmt.Print(res)
}

func main() {
	// Parse command-line flags
	colorPtr := flag.String("color", "", "Specify a color for the ASCII art")
	flag.Parse()

	// Check if the color flag is specified
	if len(os.Args) != 1 {
		col := os.Args[1]
		if len(col) < 8 || col[:8] != "--color=" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored>", "something")
			return
		}
	}

	// Check the number of arguments
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored>", "something")
		return
	}

	// Check if the color is specified
	if *colorPtr == "" {
		fmt.Println("Usage: go run . --color=<color> <text>")
		return
	}

	// Check if the specified color is available
	color, ok := colors[*colorPtr]
	if !ok {
		fmt.Println("This color is not available:", *colorPtr)
		return
	}

	arg := flag.Args()
	// Retrieve the text to convert to ASCII art
	word := arg[len(arg)-1]
	letter := arg[0]

	word = strings.ReplaceAll(word, "\\n", "\n")

	if letter == "" {
		letter = word
	}

	// Call the Ascii function with flag values
	Ascii(word, "standard", color, letter)
}
