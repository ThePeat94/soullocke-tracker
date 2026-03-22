SHELL := /bin/bash
NVM_INIT := . ~/.nvm/nvm.sh && nvm use

.PHONY: build run clean-modules clean-build generate install clean-install clean-all build-clean

build:
	$(NVM_INIT) && npm run build

run:
	$(NVM_INIT) && npm run dev

clean-modules:
	rm -rf node_modules

clean-build:
	rm -rf build

generate:
	$(NVM_INIT) && npm run generate

install:
	$(NVM_INIT) && npm install

clean-install:
	$(MAKE) clean-modules
	$(NVM_INIT) && npm install

clean-all:
	$(MAKE) clean-build
	$(MAKE) clean-modules

build-clean:
	$(MAKE) clean-build
	$(MAKE) build
