package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	input := ScanTargets()
	result := GetRobots(input)
	result = removeDuplicateValues(result)
	for _, elem := range result {
		fmt.Println(elem)
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
	return removeDuplicateValues(result)
}

//removeDuplicateValues
func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//GetRobots
func GetRobots(input []string) []string {
	var result []string
	for _, elem := range input {
		robots := GetRequest("https://" + elem + "/robots.txt")
		if robots != "" {
			s := strings.Split(robots, "\n")
			for _, line := range s {
				if strings.Contains(line, "Allow") || strings.Contains(line, "Disallow") {
					word := strings.Split(line, " ")
					result = append(result, word[1])
				}
			}
		}
	}
	return removeDuplicateValues(result)
}

//GetRequest performs a GET request
func GetRequest(target string) string {
	resp, err := http.Get(target)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	//Convert the body to type string
	sb := string(body)
	return sb
}
