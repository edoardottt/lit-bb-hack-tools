# tahm

**t**est **a**ll **h**ttp **m**ethods

Take as input on stdin a list of urls and print on stdout all the status codes and body sizes for HTTP methods. 

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/tahm`
- `go build -o tahm`
- `sudo cp tahm /usr/bin`

### Usage

`cat urls | tahm`

### Output

```
= https://www.example.com/ =
METHOD   STATUS                  SIZE  
GET      200 OK                  1256  
POST     404 Not Found           445   
PUT      404 Not Found           1256  
DELETE   405 Method Not Allowed  0     
HEAD     200 OK                  0     
CONNECT  400 Bad Request         349   
OPTIONS  200 OK                  0     
TRACE    405 Method Not Allowed  0     
PATCH    405 Method Not Allowed  0     
---------------------------

```
