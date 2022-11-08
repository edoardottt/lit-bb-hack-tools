# nrp

**n**o **r**edirect **p**lease

Take as input on stdin a list of domains and print on stdout all the unique domains without redirects.  
For example, if two domains (A and B) redirects to the same domain C, the output will be C. 

### Input

```
https://example.com/
http://example.com/
http://noredirect.com/nor
https://redirect-no.fr/blocked
```

### Usage

- `cat urls | nrp`

### Output

```
example.com 200
noredirect.com 200
redirect-no.fr 401
```
