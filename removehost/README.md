# removehost

**remove** **host**

Take as input on stdin a list of urls and print on stdout all the unique queries without protocol and host.

### Input

```
google.com/identification/ok
google.com/user/password-reminder
google.com/st/articles/19237343?language=de
google.com/specials
google.com/xxx/paper/category/paper
google.com/xxx/paper/category/paper_tags
google.com/status-page/ie6.css
google.com/status-page/logo.png
google.com/status-page/reset.css
google.com/status-page/style.css
google.com/pub/enml.dtd
http://google.com/pub/enml2.dtd
```

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
