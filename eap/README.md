# eap

**e**xtract **a**ll **p**rotocols

Take as input on stdin a list of urls and print on stdout all the protocols sorted. 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/eap`
- `go build -o eap`
- `sudo cp eap /usr/bin`

### Usage

`cat urls | eap`

### Sample output

```
[ 2458 ] http
[ 314 ] https
[ 21 ] ftp
[ 2 ] s3
```
