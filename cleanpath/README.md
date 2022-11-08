# cleanpath

**clean** **path**s

Take as input on stdin a list of urls/paths and print on stdout all the unique paths (at any level).  

### Input

```
https://google.com/
https://example.com/api
http://example.com/books/all
https://example.com/books/all/1.pdf
books/all/2.pdf
http://noredirect.com/nor
https://redirect-no.fr/blocked
```

### Usage

- `cat input | cleanpath`

### Output

```
api
api/level1
api/level1/level2
static
search
search/advanced
about
books
books/all
books/all/1.pdf
books/all/2.pdf
trends
trends/1
trends/1/2
trends/1/2/3
```
