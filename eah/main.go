package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

func main() {
	input := ScanTargets()
	set := make(map[string]int)
	for _, elem := range input {
		host := GetHost(elem)
		if host != "" {
			_, exists := set[host]
			if exists {
				set[host] += 1
			} else {
				set[host] = 1
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
	return result
}

//removeDuplicateValues
func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//GetHost >
func GetHost(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.Host
}
