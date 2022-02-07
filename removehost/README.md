# removehost

**remove** **host**

Take as input on stdin a list of urls and print on stdout all the unique queries without protocol and host.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/removehost`
- `go build -o removehost`
- `sudo cp removehost /usr/bin`

### Usage

- `cat urls | removehost`

### Output

```
/identification/ok
/user/password-reminder
/st/articles/19237343?language=de
/specials
/xxx/paper/category/paper
/xxx/paper/category/paper_tags
/status-page/ie6.css
/status-page/logo.png
/status-page/reset.css
/status-page/style.css
/pub/enml.dtd
/pub/enml2.dtd
```
