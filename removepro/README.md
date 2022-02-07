# removepro

**remove** **pro**tocols

Take as input on stdin a list of urls and print on stdout all the unique urls without protocols.  
With `-subs` you can output only domains without the queries.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/removepro`
- `go build -o removepro`
- `sudo cp removepro /usr/bin`

### Usage

- `cat urls | removepro`
- `cat urls | removepro -subs`

### Output

```
google.com/id/ok
ciao.it
okokok.it/query?exec=ok
cleandomain.out#fragment1
```
