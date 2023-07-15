build:
	@echo ""
	@echo "=>  https://github.com/edoardottt/lit-bb-hack-tools v1.3.5 <="
	@echo ""
	@cd eae && go build -o eae && sudo mv eae /usr/local/bin && echo "[ + ] eae installed!"
	@cd heacoll && go build -o heacoll && sudo mv heacoll /usr/local/bin && echo "[ + ] heacoll installed!"
	@cd removepro && go build -o removepro && sudo mv removepro /usr/local/bin && echo "[ + ] removepro installed!"
	@cd subtake && chmod +x subtake && sudo cp subtake /usr/local/bin && echo "[ + ] subtake installed!"
	@cd eap && go build -o eap && sudo mv eap /usr/local/bin && echo "[ + ] eap installed!"
	@cd gitdump && chmod +x gitdump && sudo cp gitdump /usr/local/bin && echo "[ + ] gitdump installed!"
	@cd removehost && go build -o removehost && sudo mv removehost /usr/local/bin && echo "[ + ] removehost installed!"
	@cd chainredir && go build -o chainredir && sudo mv chainredir /usr/local/bin && echo "[ + ] chainredir installed!"
	@cd tahm && go build -o tahm && sudo mv tahm /usr/local/bin && echo "[ + ] tahm installed!"
	@cd robotstxt && go build -o robotstxt && sudo mv robotstxt /usr/local/bin && echo "[ + ] robotstxt installed!"
	@cd cleanpath && go build -o cleanpath && sudo mv cleanpath /usr/local/bin && echo "[ + ] cleanpath installed!"
	@cd eefjsf && go build -o eefjsf && sudo mv eefjsf /usr/local/bin && echo "[ + ] eefjsf installed!"
	@cd bbtargets && go build -o bbtargets && sudo mv bbtargets /usr/local/bin && echo "[ + ] bbtargets installed!"
	@cd nrp && go build -o nrp && sudo mv nrp /usr/local/bin && echo "[ + ] nrp installed!"
	@cd eah && go build -o eah && sudo mv eah /usr/local/bin && echo "[ + ] eah installed!"
	@cd doomxss && go build -o doomxss && sudo mv doomxss /usr/local/bin && echo "[ + ] doomxss installed!"
	@cd eaparam && go build -o eaparam && sudo mv eaparam /usr/local/bin && echo "[ + ] eaparam installed!"
	@cd bbscope && go build -o bbscope && sudo mv bbscope /usr/local/bin && echo "[ + ] bbscope installed!"
	@cd eapath && go build -o eapath && sudo mv eapath /usr/local/bin && echo "[ + ] eapath installed!"
	@cd rpfu && go build -o rpfu && sudo mv rpfu /usr/local/bin && echo "[ + ] rpfu installed!"
	@cd rapwp && go build -o rapwp && sudo mv rapwp /usr/local/bin && echo "[ + ] rapwp installed!"
	@cd checkbypass && go build -o checkbypass && sudo mv checkbypass /usr/local/bin && echo "[ + ] checkbypass installed!"
	@cd knoxssme && go build -o knoxssme && sudo mv knoxssme /usr/local/bin && echo "[ + ] knoxssme installed!"
	@cd genscope && go build -o genscope && sudo mv genscope /usr/local/bin && echo "[ + ] genscope installed!"
	@cd kubemetrics && go build -o kubemetrics && sudo mv kubemetrics /usr/local/bin && echo "[ + ] kubemetrics installed!"
	@cd earh && go build -o earh && sudo mv earh /usr/local/bin && echo "[ + ] earh installed!"
	@echo Done!

clean:
	@sudo rm -rf /usr/local/bin/eae
	@sudo rm -rf /usr/local/bin/heacoll
	@sudo rm -rf /usr/local/bin/removepro
	@sudo rm -rf /usr/local/bin/subtake
	@sudo rm -rf /usr/local/bin/eap
	@sudo rm -rf /usr/local/bin/gitdump
	@sudo rm -rf /usr/local/bin/removehost
	@sudo rm -rf /usr/local/bin/chainredir
	@sudo rm -rf /usr/local/bin/tahm
	@sudo rm -rf /usr/local/bin/robotstxt
	@sudo rm -rf /usr/local/bin/cleanpath
	@sudo rm -rf /usr/local/bin/eefjsf
	@sudo rm -rf /usr/local/bin/bbtargets
	@sudo rm -rf /usr/local/bin/nrp
	@sudo rm -rf /usr/local/bin/eah
	@sudo rm -rf /usr/local/bin/doomxss
	@sudo rm -rf /usr/local/bin/eaparam
	@sudo rm -rf /usr/local/bin/bbscope
	@sudo rm -rf /usr/local/bin/eapath
	@sudo rm -rf /usr/local/bin/rpfu
	@sudo rm -rf /usr/local/bin/rapwp
	@sudo rm -rf /usr/local/bin/checkbypass
	@sudo rm -rf /usr/local/bin/knoxssme
	@sudo rm -rf /usr/local/bin/genscope
	@sudo rm -rf /usr/local/bin/kubemetrics
	@sudo rm -rf /usr/local/bin/earh
	@echo Cleaned everything!

update:
	@git pull
	@make clean
	@make build
