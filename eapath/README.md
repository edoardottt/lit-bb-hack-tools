# eapath

**e**xtract **a**ll **path**s

Take as input on stdin a list of urls and print on stdout all the unique urls without queries.

### Input

```
https://example.com/file?p=1
https://example.com/2?a="
https://www.example.com/file
https://example.com/file?1=2
https://example.com/2?d=2
https://www.example.com/file
```

### Usage

`cat urls | eapath`

### Output

```
https://example.com/file
https://example.com/2
https://www.example.com/file
```
