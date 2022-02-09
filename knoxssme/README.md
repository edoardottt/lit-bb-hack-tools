# knoxssme

**knoxss.me** client

Take as input on stdin a list of urls and print on stdout the results from Knoxss.me API.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/knoxssme`
- `go build -o knoxssme`
- `sudo cp knoxssme /usr/bin`
- (Only for Linux) Create the configuration file `~/.config/knoxss/knoxss.key` and copy your API key there.

### Usage

- `cat urls | knoxssme`
- `cat urls | knoxssme -k exampleapikeywbfkwfiuwlahlflvug`
- `cat urls | knoxssme -o output.txt`

### Output

```
[ SAFE ] https://brutelogic.com.br/xsx.php
[ XSS! ] https://brutelogic.com.br/xss.php?a=1<!--><Svg OnLoad=_=confirm,_(1)-->
```
