# mauu

**m**ake **a**ll **u**rls **u**nique

Take as input on stdin a list of urls and print on stdout all the urls with unique parameters.

Inspired by uro (s0md3v)

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/mauu`
- `go build -o mauu`
- `sudo cp mauu /usr/bin`

### Usage

`cat urls | mauu`

### Sample output

```
https://sub1.example.com/path1?param1=value1
https://sub1.example.com/path2?param1=value1&param2=value2
https://sub2.example.com/path3/subpath1/?param1=value1
https://sub3.example.com/path4?param4=value4
```
