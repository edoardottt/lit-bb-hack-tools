# rapwp

**r**eplace **a**ll **p**arameters **w**ith **p**ayloads

Take as input on stdin a list of urls and a payload and print on stdout all the unique urls with ready to use payloads.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/rapwp`
- `go build -o rapwp`
- `sudo cp rapwp /usr/bin`

### Usage


`cat urls | rapwp -p "<svg onload=alert(1)>"`
`cat urls | rapwp -pL payloads.txt`
`cat urls | rapwp -pL payloads.txt -obo`

### Output

```
https://sub1.example.com/path1?param1=%3Csvg%20onload=%22alert(1)%22%3E
https://sub1.example.com/path2?param1=value1&param2=%3Csvg%20onload=%22alert(1)%22%3E
https://sub2.example.com/path3/subpath1/?param1=%3Csvg%20onload=%22alert(1)%22%3E
https://sub3.example.com/path4?param4=%3Csvg%20onload=%22alert(1)%22%3E
```
