# removepro

**remove** **pro**tocols

Take as input on stdin a list of urls and print on stdout all the unique urls without protocols.  
With `-subs` you can output only domains without the queries.

### Input

```
https://google.com/id/ok
http://ciao.it
https://okokok.it/query?exec=ok
http://cleandomain.out#fragment1
```

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
