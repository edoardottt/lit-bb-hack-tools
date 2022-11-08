# doomxss

**dom** **xss** sinks finder

Take as input on stdin a list of html/js file urls and print on stdout all the possible DOM XSS sinks found. 

### Input

```
https://example.com/main.js
https://test2.test.com/user-display.js
https://sub2.sub1.test.com/user-display.html
https://example.com/script.js
https://bersaglio.it/checks.html
...
```

### Usage

- `cat urls | doomxss`

### Output

```
[ location.href= ] https://example.com/script.js
[ document.referrer= ] https://target.dom/alert.js
[ eval( ] https://bersaglio.it/checks.html
```
