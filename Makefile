GOCMD=go
GOTEST=$(GOCMD) test
test:
	echo "Checking code formats"
	$(GOPATH)/bin/goimports -w pkg/ cmd/
	echo "Running unittests..."
	$(GOPATH)/bin/gocov test ./pkg/... | $(GOPATH)/bin/gocov-xml > coverage.xml
stat:
	$(GOPATH)/bin/gocyclo ./pkg/
api:
	go build -o bin/mimiron-api cmd/api/*