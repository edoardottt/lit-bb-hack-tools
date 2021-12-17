# dedup

**dedup**licate urls

Take as input on stdin a list of urls and print on stdout all the urls with unique parameters.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/dedup`
- `go build -o dedup`
- `sudo cp dedup /usr/bin`

### Usage

`cat urls | dedup`

### Sample output

```
https://sub1.example.com/path1?param1=value1
https://sub1.example.com/path2?param1=value1&param2=value2
https://sub2.example.com/path3/subpath1/?param1=value1
https://sub3.example.com/path4?param4=value4
```
