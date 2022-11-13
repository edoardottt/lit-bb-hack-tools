package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/edoardottt/golazy"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")

	flag.Parse()

	if *helpPtr {
		help()
	}

	set := make(map[string]int)

	input := ScanTargets()
	for _, elem := range input {
		protocol := GetProtocol(elem)
		if protocol != "" {
			_, exists := set[protocol]
			if exists {
				set[protocol] += 1
			} else {
				set[protocol] = 1
			}
		}
	}

	// sort reverse.
	var (
		n = map[int][]string{}
		a = []int{}
	)

	for k, v := range set {
		n[v] = append(n[v], k)
	}

	for k := range n {
		a = append(a, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("[ %d ] %s\n", k, s)
		}
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the protocols sorted.
	$> cat urls | eap`

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

// GetProtocol.
func GetProtocol(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[:res]
	} else {
		return ""
	}
}
