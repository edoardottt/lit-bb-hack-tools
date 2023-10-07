# knoxssme

**knoxss.me** client

Take as input on stdin a list of urls and print on stdout the results from Knoxss.me API.

  
> **Note**
> (Only for Linux) Create the configuration file `~/.config/knoxss/knoxss.key` and copy your API key there.  

### Input

```
https://brutelogic.com.br/xsx.php
https://brutelogic.com.br/xss.php
```

### Usage

- `cat urls | knoxssme`
- `cat urls | knoxssme -k exampleapikey-wbfkwfiuwlahlflvug`
- `cat urls | knoxssme -o output.txt`

If you are on Linux and you configured correctly the configuration file you don't need the option `-k`.


### Output

```
[ SAFE ] https://brutelogic.com.br/xsx.php
[ XSS! ] https://brutelogic.com.br/xss.php?a=1<!--><Svg OnLoad=_=confirm,_(1)-->
```
