# bbscope

**b**ug **b**ounty **scope**

Take as input on stdin a list of urls or subdomains and a BurpSuite Configuration file and print on stdout all in scope items.

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/bbscope`
- `go build -o bbscope`
- `sudo cp bbscope /usr/bin`

### Usage

- `cat urls | bbscope target-scope.json`
- `cat subs | bbscope target-scope.json`

### Sample output

```
partners.shopify.com
accounts.shopify.com
1.shopifykloud.com
www.shopify.com
www.shopifycloud.com
...
```
