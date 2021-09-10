# nrp

**n**o **r**edirect **p**lease

Take as input on stdin a list of domains and print on stdout all the unique domains without redirects.  

For example, if two domains (A and B) redirects to the same domain C, the output will be C. 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/nrp`
- `go build -o nrp`
- `sudo cp nrp /usr/bin`

### Usage

- `cat urls | nrp`

### Sample output

```
example.com
noredirect.com
redirect-no.fr
```
