package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/edoardottt/golazy"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	output := GetTargets()
	if len(output) == 0 {
		fmt.Println()
		fmt.Println("[ ! ] Error while retrieving targets.")
		fmt.Println()
		os.Exit(1)
	}
	for _, elem := range output {
		fmt.Println(elem)
	}
}

// help shows the usage.
func help() {
	var usage = `Produce as output on stdout all the bug bounty targets found on Chaos list by Project Discovery.
	$> bbtargets`
	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
	os.Exit(0)
}

// Target is a struct containing informations about
// a single bug bounty program.
type Target struct {
	Name    string   `json:"name"`
	Url     string   `json:"url"`
	Bounty  bool     `json:"bounty"`
	Domains []string `json:"domains"`
}

// Programs is a struct containing informations about
// all the programs.
type Programs struct {
	Targets []Target `json:"programs"`
}

// GetTargets is the function that actually retrieves
// the json file containing the informations.
func GetTargets() []string {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	var results Programs
	url := "https://raw.githubusercontent.com/projectdiscovery/public-bugbounty-programs/master/chaos-bugbounty-list.json"
	resp, err := client.Get(url)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		return []string{}
	}

	var output []string

	for _, res := range results.Targets {
		// only programs with bounty.
		if res.Bounty {
			if strings.Contains(res.Url, "hackerone") || strings.Contains(res.Url, "bugcrowd") ||
				strings.Contains(res.Url, "intigriti") || strings.Contains(res.Url, "yeswehack") {
				output = append(output, cleanIgnored(res.Domains)...)
			}
		}
	}
	return output
}

// cleanIgnored is the function that clean the results
// from ignored targets.
func cleanIgnored(domains []string) []string {
	var ignoredsubs []string
	if _, err := os.Stat("ignored.txt"); err == nil {
		var ignored = readFile("ignored.txt")
		for _, domain := range domains {
			for _, forb := range ignored {
				if strings.Contains(domain, forb) {
					ignoredsubs = append(ignoredsubs, domain)
				}
			}
		}
	}
	return Difference(domains, ignoredsubs)
}

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// readFile.
func readFile(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("failed to open %s ", inputFile)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	var dir = ""
	for scanner.Scan() {
		dir = scanner.Text()
		if len(dir) > 0 {
			text = append(text, dir)
		}
	}
	file.Close()
	text = golazy.RemoveDuplicateValues(text)
	return text
}
