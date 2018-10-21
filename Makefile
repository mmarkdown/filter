mmark-filter: main.go plugins.go version.go
	go build -o mmark-filter

.PHONY: install
install:
	cp mmark-filter $$GOBIN/


.PHONY: clean
clean:
	rm mmark-filter
