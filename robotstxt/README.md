# robotstxt

**robots** **txt** 

Take as input on stdin a list of urls and print on stdout all the unique paths found in the robots.txt file.  

Inspired by [@remonsec](https://twitter.com/remonsec/status/1410481151433576449) 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/robotstxt`
- `go build -o robotstxt`
- `sudo cp robotstxt /usr/bin`

### Usage

- `cat urls | robotstxt`

### Sample output

```
/search
/search/about
/search/static
/pda/search?
/sprint_xhtml
/sprint_wml
/pqa
/books?*q=related:*
/books?*q=editions:*
/books?*q=subject:*
/trends/beta
/trends/topics
```
