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
		path := ExtractPath(elem)
		if path != "" {
			result = append(result, path)
		}
	}
	for _, elem := range golazy.RemoveDuplicateValues(result) {
		fmt.Println(elem)
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique urls without queries.
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
	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return golazy.RemoveDuplicateValues(result)
}

// ExtractPath.
func ExtractPath(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	if u.Host == "" {
		return ""
	}
	return u.Scheme + "://" + u.Host + u.Path
}
