build:
	@cd eae && go build -o eae && sudo mv eae /usr/bin && echo [ + ] eae installed!
	@cd heacoll && go build -o heacoll && sudo mv heacoll /usr/bin && echo [ + ] heacoll installed!
	@cd removepro && go build -o removepro && sudo mv removepro /usr/bin && echo [ + ] removepro installed!
	@cd subtake && chmod +x subtake && sudo cp subtake /usr/bin && echo [ + ] subtake installed!
	@cd eap && go build -o eap && sudo mv eap /usr/bin && echo [ + ] eap installed!
	@cd gitdump && chmod +x gitdump && sudo cp gitdump /usr/bin && echo [ + ] gitdump installed!
	@echo Done!

clean:
	@sudo rm -rf /usr/bin/eae
	@sudo rm -rf /usr/bin/heacoll
	@sudo rm -rf /usr/bin/removepro
	@sudo rm -rf /usr/bin/subtake
	@sudo rm -rf /usr/bin/eap
	@sudo rm -rf /usr/bin/gitdump
	@echo Cleaned everything!