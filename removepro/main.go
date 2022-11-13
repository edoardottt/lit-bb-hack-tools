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
	var input []string
	var result []string
	if !ScanFlag() {
		input = ScanTargets()
		for _, elem := range input {
			result = append(result, RemoveProtocol(elem))
		}
	} else {
		for _, elem := range input {
			sub := RemovePort(RemoveProtocol(GetOnlySubs(elem)))
			if sub != "" {
				result = append(result, sub)
			}
		}
	}
	result = golazy.RemoveDuplicateValues(result)
	for _, elem := range result {
		fmt.Println(elem)
	}

}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique urls without protocols.  
With -subs you can output only domains without the queries.
	$> cat urls | removepro
	$> cat urls | removepro -subs`
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
	return result
}

// RemoveProtocol.
func RemoveProtocol(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[res+3:]
	} else {
		return input
	}
}

// GetOnlySubs.
func GetOnlySubs(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.Host
}

// ScanFlag.
func ScanFlag() bool {
	subsPtr := flag.Bool("subs", false, "Return only subdomains without protocols.")
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	return *subsPtr
}

// RemovePort.
func RemovePort(input string) string {
	res := strings.Index(input, ":")
	if res >= 0 {
		return input[:res-1]
	}
	return input
}
