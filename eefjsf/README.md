# eefjsf

**e**xtract **e**ndpoints **f**rom **js** **f**iles

Take as input on stdin a list of js file urls and print on stdout all the unique endpoints found. 

Thanks to [@renniepak](https://twitter.com/renniepak/status/1288371394401783809).

### Input

```
https://example.com/main.js
https://test2.test.com/user-display.js
```

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
