# eah

**e**xtract **a**ll **h**osts

Take as input on stdin a list of urls and print on stdout all the hosts sorted. 

### Input

```
https://example.com/file?p=1&id=1
https://example.com/2?a="&page=!&url=1
https://www.google.com/file?redir_url=2
https://example.com/file?1=2?utm=cc
https://example.com/2?d=2?hl=1
https://www.google.com/file?hl=2
...
```

### Usage

`cat urls | eah`

### Output

```
[ 1148 ] www.google.com
[ 934 ] linux.com
[ 73 ] store.microsoft.com
[ 21 ] hackerone.com
```
