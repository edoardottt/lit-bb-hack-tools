package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	input := ScanTargets()
	conf := ScanBurpConfFile()
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls or subdomains and a BurpSuite Configuration file and print on stdout all in scope items.
	$> cat urls | bbscope target-scope.json
	$> cat subs | bbscope target-scope.json`
	fmt.Println(usage)
	os.Exit(0)
}

type BurpSuiteConfiguration struct {
	Target Target `json:"target"`
}

type Target struct {
	Scope Scope `json:"scope"`
}

type Scope struct {
	Advanced_mode bool     `json:"advanced_mode"`
	Exclude       []Domain `json:"exclude"`
	Include       []Domain `json:"include"`
}

type Domain struct {
	Enabled  bool   `json:"enabled"`
	File     string `json:"file"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}

//ScanTargets returns the array of elements
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

//ScanBurpConfFile returns the struct of the configuration file
func ScanBurpConfFile() BurpSuiteConfiguration {
	jsonFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf BurpSuiteConfiguration
	json.Unmarshal(byteValue, &conf)
	return conf
}

//RemoveDuplicateValues >
func RemoveDuplicateValues(strSlice []string) []string {
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

//GetProtocol >
func GetProtocol(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[:res]
	} else {
		return ""
	}
}
