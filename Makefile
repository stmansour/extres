DIRS = xrtest

extres:
	go vet
	golint
	go build
	go install

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	go clean

all: clean extres stats

try: clean extres

stats:
	@echo "GO SOURCE CODE STATISTICS"
	@echo "----------------------------------------"
	@find . -name "*.go" | srcstats
	@echo "----------------------------------------"

