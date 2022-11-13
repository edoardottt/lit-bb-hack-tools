package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/edoardottt/golazy"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")

	flag.Parse()

	if *helpPtr {
		help()
	}

	input := ScanTargets()

	var result []string
	for _, elem := range input {
		result = append(result, ExtractParameters(elem)...)
	}

	for _, elem := range golazy.RemoveDuplicateValues(result) {
		fmt.Println(elem)
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique parameters.
	$> cat urls | eaparam`

	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
	os.Exit(0)
}

// ScanTargets return the array of elements
// taken as input on stdin.
func ScanTargets() []string {

	var result []string
	// accept domains on stdin.
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}

	return golazy.RemoveDuplicateValues(result)
}

// ExtractParameters.
func ExtractParameters(input string) []string {
	var result []string

	u, err := url.Parse(input)
	if err != nil {
		return []string{}
	}

	decodedValue, err := url.QueryUnescape(u.RawQuery)
	if err != nil {
		return []string{}
	}

	couples := strings.Split(decodedValue, "&")
	for _, pair := range couples {
		values := strings.Split(pair, "=")
		if values[0] != "" && !strings.Contains(values[0], ";") && !strings.Contains(values[0], "{") &&
			!strings.Contains(values[0], "}") && !strings.Contains(values[0], "$") &&
			!strings.Contains(values[0], " ") && !strings.Contains(values[0], "?") &&
			!strings.Contains(values[0], "/") && !strings.Contains(values[0], "@") &&
			!strings.Contains(values[0], "(") && !strings.Contains(values[0], ")") &&
			!strings.Contains(values[0], "\"") && !strings.Contains(values[0], ".") &&
			!strings.Contains(values[0], "*") {
			result = append(result, values[0])
		}
	}

	return result
}
