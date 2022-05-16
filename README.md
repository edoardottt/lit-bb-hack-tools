<p align="center">
  <img src="https://github.com/edoardottt/images/blob/main/lit-bb-hack-tools/banner.png">
</p>

Command Line tools useful during Bug Bounty / Penetration testing. Focused on Web targets.

Installation 📥
-------

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools`
- `make build`

Then use the tools as described in the README in each tools folder.

- `make clean` (uninstall)
- `make update` (update)

Tools list 📃
-------

- [eae](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/eae) Take as input on stdin a list of urls and print on stdout all the extensions sorted.
- [heacoll](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/heacoll) Take as input on stdin a list of urls and print on stdout all the unique headers found.
- [removepro](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/removepro) Take as input on stdin a list of urls and print on stdout all the unique urls without protocols.
- [subtake](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/subtake) Take as input on stdin a list of urls and print on stdout CNAME records found with `dig`.
- [eap](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/eap) Take as input on stdin a list of urls and print on stdout all the protocols sorted.
- [gitdump](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/gitdump) It prints all the matches in a git repository with a specified pattern.
- [removehost](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/removehost) Take as input on stdin a list of urls and print on stdout all the unique queries without protocol and host.
- [chainredir](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/chainredir) Take as input a URL and print on stdout all the redirects.
- [tahm](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/tahm) Take as input on stdin a list of urls and print on stdout all the status codes and body sizes for HTTP methods.
- [robotstxt](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/robotstxt) Take as input on stdin a list of urls and print on stdout all the unique paths found in the robots.txt file.
- [cleanpath](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/cleanpath) Take as input on stdin a list of urls/paths and print on stdout all the unique paths (at any level).
- [eefjsf](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/eefjsf) Take as input on stdin a list of js file urls and print on stdout all the unique endpoints found. 
- [bbtargets](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/bbtargets) Produce as output on stdout all the bug bounty targets found on Chaos list by Project Discovery.
- [nrp](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/nrp) Take as input on stdin a list of domains and print on stdout all the unique domains without redirects.
- [eah](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/eah) Take as input on stdin a list of urls and print on stdout all the hosts sorted.
- [doomxss](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/doomxss) Take as input on stdin a list of html/js file urls and print on stdout all the possible DOM XSS sinks found.
- [eaparam](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/eaparam) Take as input on stdin a list of urls and print on stdout all the unique parameters.
- [bbscope](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/bbscope) Take as input on stdin a list of urls or subdomains and a BurpSuite Configuration file and print on stdout all in scope items.
- [eapath](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/eapath) Take as input on stdin a list of urls and print on stdout all the unique urls without queries.
- [rpfu](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/rpfu) Take as input on stdin a list of urls and print on stdout all the unique urls without ports (if 80 or 443).
- [rapwp](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/rapwp) Take as input on stdin a list of urls and a payload and print on stdout all the unique urls with ready to use payloads.
- [checkbypass](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/checkbypass) Take as input on stdin a payload and print on stdout all the successful WAF bypasses.
- [knoxssme](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/knoxssme) Take as input on stdin a list of urls and print on stdout the results from Knoxss.me API.
- [genscope](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/genscope) Take as input a file containing a list of (sub)domains (wildcards allowed) and produce a BurpSuite Configuration file.
- [kubemetrics](https://github.com/edoardottt/lit-bb-hack-tools/tree/main/kubemetrics) Take as input on stdin a list of urls and print on stdout all the unique paths and urls found in the /metrics endpoint.

Changelog 📌
-------
Detailed changes for each release are documented in the [release notes](https://github.com/edoardottt/lit-bb-hack-tools/releases).

Contributing 🤝
------
If you want to contribute to this project, you can start opening an [issue](https://github.com/edoardottt/lit-bb-hack-tools/issues).

[![](/CONTRIBUTORS.svg)](https://github.com/edoardottt/lit-bb-hack-tools/graphs/contributors)

License 📝
-------

This repository is under [GNU General Public License v3.0](https://github.com/edoardottt/lit-bb-hack-tools/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.
