# eap

**e**xtract **a**ll **p**rotocols

Take as input on stdin a list of urls and print on stdout all the protocols sorted by count. 

### Input

```
https://example.com/file?p=1&id=1
https://example.com/2?a="&page=!&url=1
https://www.example.com/file?redir_url=2
https://example.com/file?1=2?utm=cc
https://example.com/2?d=2?hl=1
https://www.example.com/file?hl=2
...
```

### Usage

`cat urls | eap`

### Output

```
[ 2458 ] http
[ 314 ] https
[ 21 ] ftp
[ 2 ] s3
```
