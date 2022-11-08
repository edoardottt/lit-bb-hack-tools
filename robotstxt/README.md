# robotstxt

**robots** **txt** 

Take as input on stdin a list of urls and print on stdout all the unique paths found in the robots.txt file.  

Inspired by [@remonsec](https://twitter.com/remonsec/status/1410661328977600517) 

### Input

```
https://google.com/search
https://google.com/search/about
https://google.com/search/static
https://google.com/pda/search?
https://google.com/sprint_xhtml
https://google.com/sprint_wml
https://google.com/pqa
https://google.com/books?*q=related:*
https://google.com/books?*q=editions:*
https://google.com/books?*q=subject:*
https://google.com/trends/beta
https://google.com/trends/topics
```

### Usage

- `cat urls | robotstxt`

### Output

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
