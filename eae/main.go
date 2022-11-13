package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
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
	extractExtensions(ScanTargets())
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the extensions sorted.
	$> cat urls | eae`
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

// extractExtensions.
func extractExtensions(input []string) {
	set := make(map[string]int)
	for _, elem := range input {
		u, err := url.Parse(elem)
		if err == nil {
			elem = u.Path
			firstIndex := strings.Index(elem, "?")
			if firstIndex > -1 {
				elem = elem[:firstIndex]
			}
			i := strings.LastIndex(elem, ".")
			if i >= 0 {
				extension := elem[i:]
				_, exists := set[extension]
				if exists {
					set[extension] += 1
				} else {
					set[extension] = 1
				}
			}
		}
	}
	// sort reverse.
	n := map[int][]string{}
	var a []int
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
