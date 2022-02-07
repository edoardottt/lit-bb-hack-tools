# checkbypass

**check** **bypass**

Take as input on stdin a payload and print on stdout all the successful WAF bypasses.  

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/checkbypass`
- `go build -o checkbypass`
- `sudo cp checkbypass /usr/bin`

### Usage

- `checkbypass -p "<script>alert()</script>"`

### Sample output

```
[ BYPASSED ] Cloudflare   : https://www.cloudflare.com/?=test%3Dciaoooo
[ BYPASSED ] Akamai       : https://www.akamai.com/?=test%3Dciaoooo
[ BYPASSED ] F5           : https://www.f5.com/?=test%3Dciaoooo
[ BYPASSED ] CloudFront   : https://docs.aws.amazon.com/?=test%3Dciaoooo
[ BYPASSED ] Fortiweb     : https://www.fortinet.com/?=test%3Dciaoooo
[ BYPASSED ] Imperva      : https://www.imperva.com/?=test%3Dciaoooo
[ BYPASSED ] Wordfence    : https://www.wordfence.com/products/?=test%3Dciaoooo
```
