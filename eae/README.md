# eae

**e**xtract **a**ll **e**xtensions

Take as input on stdin a list of urls and print on stdout all the extensions sorted. 

### Input

```
https://example.com/main.zip
https://example.com/main.php
https://test.it/main.bz2
https://test.it/
...
```

### Usage

`cat urls | eae`

### Output

```
[ 1434 ] .jpg
[ 583 ] .png
[ 571 ] .gif
[ 27 ] .exe
[ 24 ] .zip
[ 18 ] .bz2
[ 12 ] .html
[ 10 ] .wav
[ 7 ] .css
[ 7 ] .cab
[ 7 ] .pdf
[ 6 ] .txt
[ 5 ] .rar
[ 4 ] .c
[ 4 ] .mp3
[ 3 ] .php
[ 3 ] .avi
[ 3 ] .reg
[ 3 ] .wmv
```
