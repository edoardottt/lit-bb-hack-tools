package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/edoardottt/golazy"
)

var (
	ErrDomainFormat = errors.New("domain formatted in a bad way")
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")

	flag.Parse()

	if *helpPtr {
		help()
	}

	set := make(map[string]bool)

	input := ScanTargets()
	for _, elem := range input {
		rootHost, err := GetRootHost(elem)
		if err == nil && rootHost != "" {
			_, exists := set[rootHost]
			if !exists {
				set[rootHost] = true
			}
		}
	}

	for k := range set {
		fmt.Println(k)
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique root hosts.
	$> cat urls | earh`

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

// GetRootHost takes as input a string and
// tries to parse it as url, if it's a
// well formatted url this function returns
// the second level domain.
func GetRootHost(input string) (string, error) {
	if !HasProtocol(input) {
		input = "http://" + input
	}

	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	// divide host and port, then split by dot
	parts := strings.Split(strings.Split(u.Host, ":")[0], ".")
	// return the last two parts
	if len(parts) > 1 {
		return parts[len(parts)-2] + "." + parts[len(parts)-1], nil
	}

	return "", fmt.Errorf("%w", ErrDomainFormat)
}

// HasProtocol takes as input a string and
// checks if it has a protocol ( like in a
// URI/URL).
func HasProtocol(input string) bool {
	res := strings.Index(input, "://")
	return res >= 0
}
