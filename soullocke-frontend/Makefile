build:
	npm run build

run:
	npm run dev

clean-modules:
	rm -rf node_modules

clean-build:
	rm -rf build

generate:
	npm run generate

install:
	nvm use
	npm install

clean-install:
	nvm use
	$(MAKE) clean-modules
	npm install

clean-all:
	$(MAKE) clean-build
	$(MAKE) clean-modules

build-clean:
	$(MAKE) clean-build
	$(MAKE) build
