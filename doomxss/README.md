# doomxss

**dom** **xss** sinks finder

Take as input on stdin a list of html/js file urls and print on stdout all the possible DOM XSS sinks found. 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/doomxss`
- `go build -o doomxss`
- `sudo cp doomxss /usr/bin`

### Usage

- `cat urls | doomxss`

### Output

```
[ location.href= ] https://example.com/script.js
[ document.referrer= ] https://target.dom/alert.js
[ eval( ] https://bersaglio.it/checks.js
```
