# eefjsf

**e**xtract **e**ndpoints **f**rom **js** **f**iles

Take as input on stdin a list of js file urls and print on stdout all the unique endpoints found. 

Thanks to [@renniepak](https://twitter.com/renniepak/status/1288371394401783809).

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/eefjsf`
- `go build -o eefjsf`
- `sudo cp eefjsf /usr/bin`

### Usage

- `cat js-urls | eefjsf`

### Output

```
/api
/chat_panel/v2
/embed/video/
/chat/v2/conversation/download
/chat/v2
/mychats
/tunnel/UI/Login
/hidden
/hidden1
```
