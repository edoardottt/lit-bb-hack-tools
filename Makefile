installfolder = /usr/local/bin
version = 1.4.0

updatemod:
	@go get -u ./...
	@go mod tidy -v
	@echo "Done."

lint:
	@golangci-lint run

build:
	@echo ""
	@echo "=>  https://github.com/edoardottt/lit-bb-hack-tools v$(version) <="
	@echo ""
	@cd eae && go build -o eae && sudo mv eae $(installfolder) && echo "[ + ] eae installed!"
	@cd heacoll && go build -o heacoll && sudo mv heacoll $(installfolder) && echo "[ + ] heacoll installed!"
	@cd removepro && go build -o removepro && sudo mv removepro $(installfolder) && echo "[ + ] removepro installed!"
	@cd subtake && chmod +x subtake && sudo cp subtake $(installfolder) && echo "[ + ] subtake installed!"
	@cd eap && go build -o eap && sudo mv eap $(installfolder) && echo "[ + ] eap installed!"
	@cd gitdump && chmod +x gitdump && sudo cp gitdump $(installfolder) && echo "[ + ] gitdump installed!"
	@cd removehost && go build -o removehost && sudo mv removehost $(installfolder) && echo "[ + ] removehost installed!"
	@cd chainredir && go build -o chainredir && sudo mv chainredir $(installfolder) && echo "[ + ] chainredir installed!"
	@cd tahm && go build -o tahm && sudo mv tahm $(installfolder) && echo "[ + ] tahm installed!"
	@cd robotstxt && go build -o robotstxt && sudo mv robotstxt $(installfolder) && echo "[ + ] robotstxt installed!"
	@cd cleanpath && go build -o cleanpath && sudo mv cleanpath $(installfolder) && echo "[ + ] cleanpath installed!"
	@cd eefjsf && go build -o eefjsf && sudo mv eefjsf $(installfolder) && echo "[ + ] eefjsf installed!"
	@cd bbtargets && go build -o bbtargets && sudo mv bbtargets $(installfolder) && echo "[ + ] bbtargets installed!"
	@cd nrp && go build -o nrp && sudo mv nrp $(installfolder) && echo "[ + ] nrp installed!"
	@cd eah && go build -o eah && sudo mv eah $(installfolder) && echo "[ + ] eah installed!"
	@cd doomxss && go build -o doomxss && sudo mv doomxss $(installfolder) && echo "[ + ] doomxss installed!"
	@cd eaparam && go build -o eaparam && sudo mv eaparam $(installfolder) && echo "[ + ] eaparam installed!"
	@cd bbscope && go build -o bbscope && sudo mv bbscope $(installfolder) && echo "[ + ] bbscope installed!"
	@cd eapath && go build -o eapath && sudo mv eapath $(installfolder) && echo "[ + ] eapath installed!"
	@cd rpfu && go build -o rpfu && sudo mv rpfu $(installfolder) && echo "[ + ] rpfu installed!"
	@cd rapwp && go build -o rapwp && sudo mv rapwp $(installfolder) && echo "[ + ] rapwp installed!"
	@cd checkbypass && go build -o checkbypass && sudo mv checkbypass $(installfolder) && echo "[ + ] checkbypass installed!"
	@cd knoxssme && go build -o knoxssme && sudo mv knoxssme $(installfolder) && echo "[ + ] knoxssme installed!"
	@cd genscope && go build -o genscope && sudo mv genscope $(installfolder) && echo "[ + ] genscope installed!"
	@cd kubemetrics && go build -o kubemetrics && sudo mv kubemetrics $(installfolder) && echo "[ + ] kubemetrics installed!"
	@cd earh && go build -o earh && sudo mv earh $(installfolder) && echo "[ + ] earh installed!"
	@echo Done!

clean:
	@sudo rm -rf $(installfolder)/eae
	@sudo rm -rf $(installfolder)/heacoll
	@sudo rm -rf $(installfolder)/removepro
	@sudo rm -rf $(installfolder)/subtake
	@sudo rm -rf $(installfolder)/eap
	@sudo rm -rf $(installfolder)/gitdump
	@sudo rm -rf $(installfolder)/removehost
	@sudo rm -rf $(installfolder)/chainredir
	@sudo rm -rf $(installfolder)/tahm
	@sudo rm -rf $(installfolder)/robotstxt
	@sudo rm -rf $(installfolder)/cleanpath
	@sudo rm -rf $(installfolder)/eefjsf
	@sudo rm -rf $(installfolder)/bbtargets
	@sudo rm -rf $(installfolder)/nrp
	@sudo rm -rf $(installfolder)/eah
	@sudo rm -rf $(installfolder)/doomxss
	@sudo rm -rf $(installfolder)/eaparam
	@sudo rm -rf $(installfolder)/bbscope
	@sudo rm -rf $(installfolder)/eapath
	@sudo rm -rf $(installfolder)/rpfu
	@sudo rm -rf $(installfolder)/rapwp
	@sudo rm -rf $(installfolder)/checkbypass
	@sudo rm -rf $(installfolder)/knoxssme
	@sudo rm -rf $(installfolder)/genscope
	@sudo rm -rf $(installfolder)/kubemetrics
	@sudo rm -rf $(installfolder)/earh
	@echo Cleaned everything!

update:
	@git pull
	@make clean
	@make build
