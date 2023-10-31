.PHONY: build
build:
	go build -o keyr main.go

.PHONY: protoc
protoc:
	find ./ -name "*.proto" -exec protoc --go_out=. --go_opt=paths=source_relative {} \;

