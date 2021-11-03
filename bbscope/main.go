package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if len(os.Args) < 2 || *helpPtr {
		help()
	}
	var result []string
	switch os.Args[1] {
	case "sub":
		input := ScanTargets()
		conf := ScanBurpConfFile()
		result = checkSubs(input, conf)
	case "url":
		input := ScanTargets()
		conf := ScanBurpConfFile()
		result = checkUrls(input, conf)
	default:
		help()
	}
	for _, elem := range result {
		fmt.Println(elem)
	}
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls or subdomains and a BurpSuite Configuration file and print on stdout all in scope items.
	$> cat urls | bbscope url target-scope.json
	$> cat subs | bbscope sub target-scope.json`
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
	jsonFile, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

//checkSubs returns a slice of string containing only the in scope subdomains
func checkSubs(input []string, conf BurpSuiteConfiguration) []string {
	var result []string
	for _, item := range input {
		if GetProtocol(item) == "" {
			item = "http://" + item
		}
		u, err := url.Parse(item)
		if err != nil {
			continue
		}
		subdomain := u.Host
		for _, excluded := range conf.Target.Scope.Exclude {
			r, _ := regexp.Compile(excluded.Host)
			if r.MatchString(subdomain) {
				break
			}
		}
		for _, included := range conf.Target.Scope.Include {
			r, _ := regexp.Compile(included.Host)
			if r.MatchString(subdomain) {
				result = append(result, subdomain)
				break
			}
		}
	}
	return result
}

//checkUrls returns a slice of string containing only the in scope urls
func checkUrls(input []string, conf BurpSuiteConfiguration) []string {
	var result []string
	for _, item := range input {
		if GetProtocol(item) == "" {
			continue
		}
		u, err := url.Parse(item)
		if err != nil {
			continue
		}
		var excludedItem = false
		for _, excluded := range conf.Target.Scope.Exclude {
			rHost, _ := regexp.Compile(excluded.Host)
			rFile, _ := regexp.Compile(excluded.File)
			if rHost.MatchString(u.Host) && rFile.MatchString(u.Path) {
				excludedItem = true
				break
			}
		}
		if excludedItem {
			continue
		}
		for _, included := range conf.Target.Scope.Include {
			rHost, _ := regexp.Compile(included.Host)
			rFile, _ := regexp.Compile(included.File)
			if rHost.MatchString(u.Host) && rFile.MatchString(u.Path) {
				result = append(result, item)
				break
			}
		}
	}
	return result
}
