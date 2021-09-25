package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

//main
func main() {
	extractExtensions(ScanTargets())
}

//ScanTargets return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

	var result []string
	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return RemoveDuplicateValues(result)
}

//extractExtensions
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
	//sort reverse
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

//RemoveDuplicateValues
func RemoveDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
