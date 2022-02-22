# genscope

**gen**erate **scope**

Take as input a file containing a list of (sub)domains (wildcards allowed) and produce a BurpSuite Configuration file.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/genscope`
- `go build -o genscope`
- `sudo cp genscope /usr/bin`

### Usage

- `genscope domains.txt`

### Input

```
example.com
*.example2.com
www.example3.com
*.example4.com
```