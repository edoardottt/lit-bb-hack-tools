build:
	@echo ""
	@echo "=>  https://github.com/edoardottt/lit-bb-hack-tools v1.3.4 <="
	@echo ""
	@cd eae && go build -o eae && sudo mv eae /usr/bin && echo "[ + ] eae installed!"
	@cd heacoll && go build -o heacoll && sudo mv heacoll /usr/bin && echo "[ + ] heacoll installed!"
	@cd removepro && go build -o removepro && sudo mv removepro /usr/bin && echo "[ + ] removepro installed!"
	@cd subtake && chmod +x subtake && sudo cp subtake /usr/bin && echo "[ + ] subtake installed!"
	@cd eap && go build -o eap && sudo mv eap /usr/bin && echo "[ + ] eap installed!"
	@cd gitdump && chmod +x gitdump && sudo cp gitdump /usr/bin && echo "[ + ] gitdump installed!"
	@cd removehost && go build -o removehost && sudo mv removehost /usr/bin && echo "[ + ] removehost installed!"
	@cd chainredir && go build -o chainredir && sudo mv chainredir /usr/bin && echo "[ + ] chainredir installed!"
	@cd tahm && go build -o tahm && sudo mv tahm /usr/bin && echo "[ + ] tahm installed!"
	@cd robotstxt && go build -o robotstxt && sudo mv robotstxt /usr/bin && echo "[ + ] robotstxt installed!"
	@cd cleanpath && go build -o cleanpath && sudo mv cleanpath /usr/bin && echo "[ + ] cleanpath installed!"
	@cd eefjsf && go build -o eefjsf && sudo mv eefjsf /usr/bin && echo "[ + ] eefjsf installed!"
	@cd bbtargets && go build -o bbtargets && sudo mv bbtargets /usr/bin && echo "[ + ] bbtargets installed!"
	@cd nrp && go build -o nrp && sudo mv nrp /usr/bin && echo "[ + ] nrp installed!"
	@cd eah && go build -o eah && sudo mv eah /usr/bin && echo "[ + ] eah installed!"
	@cd doomxss && go build -o doomxss && sudo mv doomxss /usr/bin && echo "[ + ] doomxss installed!"
	@cd eaparam && go build -o eaparam && sudo mv eaparam /usr/bin && echo "[ + ] eaparam installed!"
	@cd bbscope && go build -o bbscope && sudo mv bbscope /usr/bin && echo "[ + ] bbscope installed!"
	@cd eapath && go build -o eapath && sudo mv eapath /usr/bin && echo "[ + ] eapath installed!"
	@cd rpfu && go build -o rpfu && sudo mv rpfu /usr/bin && echo "[ + ] rpfu installed!"
	@cd rapwp && go build -o rapwp && sudo mv rapwp /usr/bin && echo "[ + ] rapwp installed!"
	@cd checkbypass && go build -o checkbypass && sudo mv checkbypass /usr/bin && echo "[ + ] checkbypass installed!"
	@cd knoxssme && go build -o knoxssme && sudo mv knoxssme /usr/bin && echo "[ + ] knoxssme installed!"
	@cd genscope && go build -o genscope && sudo mv genscope /usr/bin && echo "[ + ] genscope installed!"
	@cd kubemetrics && go build -o kubemetrics && sudo mv kubemetrics /usr/bin && echo "[ + ] kubemetrics installed!"
	@echo Done!

clean:
	@sudo rm -rf /usr/bin/eae
	@sudo rm -rf /usr/bin/heacoll
	@sudo rm -rf /usr/bin/removepro
	@sudo rm -rf /usr/bin/subtake
	@sudo rm -rf /usr/bin/eap
	@sudo rm -rf /usr/bin/gitdump
	@sudo rm -rf /usr/bin/removehost
	@sudo rm -rf /usr/bin/chainredir
	@sudo rm -rf /usr/bin/tahm
	@sudo rm -rf /usr/bin/robotstxt
	@sudo rm -rf /usr/bin/cleanpath
	@sudo rm -rf /usr/bin/eefjsf
	@sudo rm -rf /usr/bin/bbtargets
	@sudo rm -rf /usr/bin/nrp
	@sudo rm -rf /usr/bin/eah
	@sudo rm -rf /usr/bin/doomxss
	@sudo rm -rf /usr/bin/eaparam
	@sudo rm -rf /usr/bin/bbscope
	@sudo rm -rf /usr/bin/eapath
	@sudo rm -rf /usr/bin/rpfu
	@sudo rm -rf /usr/bin/rapwp
	@sudo rm -rf /usr/bin/checkbypass
	@sudo rm -rf /usr/bin/knoxssme
	@sudo rm -rf /usr/bin/genscope
	@sudo rm -rf /usr/bin/kubemetrics
	@echo Cleaned everything!

update:
	@git pull
	@make clean
	@make build
