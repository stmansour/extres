xrtest: *.go config.json
	go vet
	golint
	go build

config.json:
	/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	cp confdev.json config.json

clean:
	go clean
	rm -f xrtest conf*.json

test:
	./xrtest

all: clean xrtest
