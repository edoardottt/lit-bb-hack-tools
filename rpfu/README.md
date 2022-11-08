# rpfu

**r**emove **p**ort **f**rom **u**rls

Take as input on stdin a list of urls and print on stdout all the unique urls without ports (if 80 or 443). 

### Input

```
https://google.com:443/id/ok
https://ciao.it:443
https://okokok.it:443/query?exec=ok
http://cleandomain.out:80#fragment1

...
```

### Usage

`cat urls | rpfu`

### Output

```
https://google.com/id/ok
https://ciao.it
https://okokok.it/query?exec=ok
http://cleandomain.out#fragment1
```
