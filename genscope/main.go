package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/edoardottt/golazy"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if len(os.Args) < 2 || *helpPtr {
		help()
	}
	domains := golazy.RemoveDuplicateValues(golazy.ReadFileLineByLine(os.Args[1]))
	GenerateDomains(domains)
}

// help shows the usage.
func help() {
	var usage = `Take as input a file containing a list of (sub)domains (wildcards allowed) and produce a BurpSuite Configuration file.
	$> genscope domains.txt`
	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
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

// GenerateDomains.
func GenerateDomains(input []string) {
	var domains []Domain
	for _, elem := range input {
		domain := "^" + strings.ReplaceAll(strings.ReplaceAll(elem, ".", "\\."), "*", ".*") + "$"
		// Here add logic for hosts.
		dom80 := Domain{Enabled: true, File: "^/.*", Host: domain, Port: "^80$", Protocol: "http"}
		dom443 := Domain{Enabled: true, File: "^/.*", Host: domain, Port: "^443$", Protocol: "https"}
		domains = append(domains, dom80)
		domains = append(domains, dom443)
	}
	var result = BurpSuiteConfiguration{Target: Target{Scope: Scope{Advanced_mode: true, Exclude: []Domain{}, Include: domains}}}

	file, _ := json.MarshalIndent(result, "", "	")

	_ = ioutil.WriteFile("genscope.json", file, 0644)
}
