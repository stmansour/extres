DIRS = xrtest

# extres Makefile

extres:
	go vet
	# if [[ `uname` = "Darwin" ]]; then staticcheck; else golint; fi
	if [[ -f "${GOPATH}/bin/golint" ]]; then golint; else staticcheck; fi
	go build
	go install

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	go clean

test:
	for dir in $(DIRS); do make -C $$dir test;done

all: clean extres stats

try: clean extres

stats:
	@echo "GO SOURCE CODE STATISTICS"
	@echo "----------------------------------------"
	@find . -name "*.go" | srcstats
	@echo "----------------------------------------"
