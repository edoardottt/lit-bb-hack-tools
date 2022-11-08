# bbscope

**b**ug **b**ounty **scope**

Take as input on stdin a list of urls or subdomains and a BurpSuite Configuration file and print on stdout all in scope items.

### Input (subs)

```
partners.shopify.com
accounts.shopify.com
1.shopifykloud.com
```

### Input (urls)

```
https://partners.shopify.com/
https://accounts.shopify.com/admin
https://1.shopifykloud.com/1
```

### Usage

- `cat urls | bbscope url target-scope.json`
- `cat subs | bbscope sub target-scope.json`

### Output

```
partners.shopify.com
accounts.shopify.com
1.shopifykloud.com
www.shopify.com
www.shopifycloud.com
...
```
