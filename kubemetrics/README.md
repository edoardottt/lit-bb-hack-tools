# kubemetrics

**kubernetes** **metrics** 

Take as input on stdin a list of urls and print on stdout all the unique paths and urls found in the /metrics endpoint.  

Inspired by [@ITSecurityguard](https://twitter.com/ITSecurityguard/status/1510951340763136005), thanks to [@remonsec](https://twitter.com/remonsec).

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/kubemetrics`
- `go build -o kubemetrics`
- `sudo cp kubemetrics /usr/bin`

### Usage

- `cat urls | kubemetrics`

### Output

```
/search
/search/about
/search/static
/sprint_xhtml
db.example.com:81/conf.txt
/sprint_wml
/pqa
/trends/beta
/trends/topics
```
