# tahm

**t**est **a**ll **h**ttp **m**ethods

Take as input on stdin a list of urls and print on stdout all the status codes and body sizes for HTTP methods. 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/tahm`
- `go build -o tahm`
- `sudo cp tahm /usr/bin`

### Usage

`cat urls | tahm`

### Sample output

```
= https://www.amazon.com =
 GET Status: 200 OK | Size: 390841
 POST Status: 405 Method Not Allowed | Size: 222
 PUT Status: 200 OK | Size: 395014
 DELETE Status: 500 Internal Server Error | Size: 171676
 HEAD Status: 200 OK | Size: 0
---------------------------

```
