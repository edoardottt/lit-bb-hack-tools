# subtake

**sub**domain **take**over

Take as input on stdin a list of urls and print on stdout CNAME records found with `dig`.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/subtake`
- `chmod +x subtake.sh`
- `mv subtake.sh subtake`
- `sudo cp subtake /usr/bin`

### Usage

- `cat urls | ./subtake.sh`

### Sample output

```
admin.stage2aa.paypal.com
admin.stage2aa.paypal.com. 2943	IN	CNAME	www.stage2aa.paypal.com.

---------------------------
admin.stage2b.paypal.com
admin.stage2b.paypal.com. 3599	IN	CNAME	www.stage2b.paypal.com.

---------------------------
admin.stage2.paypal.com
admin.stage2.paypal.com. 3599	IN	CNAME	www.stage2.paypal.com.

---------------------------
api-3t.paypal.com
api-3t.paypal.com.	3114	IN	CNAME	api-3t.glb.paypal.com.
```
