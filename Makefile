GOCMD=go
GOTEST=$(GOCMD) test
pre-test:
	go get -v github.com/axw/gocov/gocov
	go get -v github.com/marinbek/gocov-xml
	go get -v golang.org/x/tools/cmd/goimports
test: pre-test
	echo "Checking code formats"
	$(GOPATH)/bin/goimports -w pkg/ cmd/
	echo "Running unittests..."
	$(GOPATH)/bin/gocov test ./pkg/... | $(GOPATH)/bin/gocov-xml > coverage.xml
stat:
	$(GOPATH)/bin/gocyclo ./pkg/
api:
	go build -o bin/mimiron-api cmd/api/*