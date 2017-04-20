xrtest: *.go
	go vet
	golint
	go build

clean:
	go clean
	rm -f xrtest conf*.json

all: clean xrtest
