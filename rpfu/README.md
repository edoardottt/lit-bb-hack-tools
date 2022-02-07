# rpfu

**r**emove **p**ort **f**rom **u**rls

Take as input on stdin a list of urls and print on stdout all the unique urls without ports (if 80 or 443). 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/rpfu`
- `go build -o rpfu`
- `sudo cp rpfu /usr/bin`

### Usage

`cat urls | rpfu`

### Output

```
https://google.com/id/ok
https://ciao.it
https://okokok.it/query?exec=ok
https://cleandomain.out#fragment1
```
