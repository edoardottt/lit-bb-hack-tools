# eapath

**e**xtract **a**ll **path**s

Take as input on stdin a list of urls and print on stdout all the unique urls without queries.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/eapath`
- `go build -o eapath`
- `sudo cp eapath /usr/bin`

### Usage

`cat urls | eapath`

### Output

```
https://example.com/file
https://example.com/2
https://www.example.com/file
```
