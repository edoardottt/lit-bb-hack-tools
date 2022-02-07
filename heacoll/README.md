# heacoll

**hea**ders **coll**ector

Take as input on stdin a list of urls and print on stdout all the unique headers found. 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/heacoll`
- `go build -o heacoll`
- `sudo cp heacoll /usr/bin`

### Usage

`cat urls | heacoll`

### Output

```
Date : [Sat, 08 May 2021 14:20:28 GMT Sat, 08 May 2021 14:20:29 GMT Sat, 08 May 2021 14:20:30 GMT]

X-Content-Type-Options : [nosniff]

X-Ua-Compatible : [IE=edge]

Permissions-Policy : [ch-ua-full-version=*, ch-ua-platform=*, ch-ua-platform-version=*, ch-ua-arch=*, ch-ua-model=*]

X-Frame-Options : [SAMEORIGIN]

Alt-Svc : [h3-29=":443"; ma=2592000,h3-T051=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"]

Accept-Ch : [ect,rtt,downlink]

P3p : [CP="This is not a P3P policy! See g.co/p3phelp for more info." CP="This is not a P3P policy! See http://support.google.com/accounts/answer/151657?hl=it for more info."]

Set-Cookie : [CONSENT=PENDING+069; expires=Fri, 01-Jan-2038 00:00:00 GMT; path=/; domain=.google.com; Secure session-id=13xx1-xxxx-xx; Domain=.amazon.com; Expires=Sun, 08-May-2022 14:20:29 GMT; Path=/; Secure session-id-time=xxxxxx; Domain=.amazon.com; Expires=Sun, 08-May-2022 14:20:29 GMT; Path=/; Secure i18n-prefs=USD; Domain=.amazon.com; Expires=Sun, 08-May-2022 14:20:29 GMT; Path=/ sp-cdn="L5Z9:IT"; Version=1; Domain=.amazon.com; Max-Age=31536000; Expires=Sun, 08-May-2022 14:20:29 GMT; Path=/; Secure; HttpOnly skin=noskin; path=/; domain=.amazon.com Domain=.youtube.com; Path=/; Secure; HttpOnly; SameSite=none CONSENT=PENDING+120; expires=Fri, 01-Jan-2038 00:00:00 GMT; path=/; domain=.youtube.com; Secure NID=215=xxcx; expires=Sun, 07-Nov-2021 14:20:30 GMT; path=/; domain=.google.it; HttpOnly CONSENT=PENDING+986; expires=Fri, 01-Jan-2038 00:00:00 GMT; path=/; domain=.google.it; Secure]

Cache-Control : [private, max-age=0 no-cache no-cache, no-store, max-age=0, must-revalidate]

Server : [gws Server ESF]

Vary : [Content-Type,Accept-Encoding,X-Amzn-CDN-Cache,X-Amzn-AX-Treatment,User-Agent]

Pragma : [no-cache]

Expires : [-1 Mon, 01 Jan 1990 00:00:00 GMT]

X-Xss-Protection : [0 1;]

Strict-Transport-Security : [max-age=47474747; includeSubDomains; preload max-age=31536000]

X-Amz-Rid : [xxxxxxxxxxx]

Content-Language : [en-US]

Accept-Ch-Lifetime : [86400]

Content-Type : [text/html; charset=ISO-8859-1 text/html;charset=UTF-8 text/html; charset=utf-8]

```
