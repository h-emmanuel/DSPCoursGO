.PHONY: help
.SILENT:
.DEFAULT_GOAL: help

TARGET_OS = "windows" "linux" "darwin"
TARGET_ARCH = "amd64" "arm64"

help: #Pour générer automatiquement l'aide ## Display all commands available
	$(eval PADDING=$(shell grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk '{ print length($$1)-1 }' | sort -n | tail -n 1))
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-$(PADDING)s\033[0m %s\n", $$1, $$2}'

run: clean ## Build and run app
	go build -o ./build/app main.go
	./build/app test/source test/destination

clean: ## On clean le dossier de build
	rm -rf ./build/*

build: clean ## On fait un build 

	for os in $(TARGET_OS); do 															\
		for arch in $(TARGET_ARCH); do 													\
			env GOOS=$$os GOARCH=$$arch go build -o ./build/app-$$os-$$arch main.go ; 	\
		done; 																			\
	done;

commit:
	git add -A .
	git commit -m "auto commit"
	git push