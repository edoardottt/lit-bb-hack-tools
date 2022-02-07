# eaparam

**e**xtract **a**ll **param**eters

Take as input on stdin a list of urls and print on stdout all the unique parameters.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/eaparam`
- `go build -o eaparam`
- `sudo cp eaparam /usr/bin`

### Usage

`cat urls | eaparam`

### Output

```
id
page
url
redir_url
utm
hl
```
