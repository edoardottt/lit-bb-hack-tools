# cleanpath

**clean** **path**s

Take as input on stdin a list of urls/paths and print on stdout all the unique paths (at any level).  

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/cleanpath`
- `go build -o cleanpath`
- `sudo cp cleanpath /usr/bin`

### Usage

- `cat input | cleanpath`

### Sample output

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
