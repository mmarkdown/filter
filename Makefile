mmark-filter: main.go plugins.go version.go
	go build -o mmark-filter

mmark-filter.1: mmark-filter.1.md
	pandoc mmark-filter.1.md -s -t man > mmark-filter.1

.PHONY: clean
clean:
	rm mmark-filter

.PHONY: install
install:
	cp mmark-filter $$GOBIN/
