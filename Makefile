GOFILES=`go list ./... | grep -v /vendor/`

all: check Temporal

# List all commands
.PHONY: ls
ls:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

# Run simple checks
.PHONY: check
check:
	@echo "===================      running checks     ==================="
	go vet ./...
	@echo "Executing dry run of tests..."
	@go test -run xxxx ./...
	@echo "===================          done           ==================="

# Build Temporal
Temporal:
	go build ./...

# Static analysis and style checks
.PHONY: lint
lint:
	go fmt ./...
	golint $(GOFILES)
	# Shellcheck disabled for now - too much to fix
	# shellcheck **/*.sh(e[' [[ ! `echo "$REPLY" | grep "vendor/" ` ]]'])

# Set up test environment
.PHONY: testenv
testenv:
	docker-compose -f test/docker-compose.yml up 

# Execute short tests
.PHONY: test
test: check
	@echo "===================  executing short tests  ==================="
	go test -race -cover -short ./...
	@echo "===================          done           ==================="

# Execute all tests
.PHONY: test
test-all: check
	@echo "===================   executing all tests   ==================="
	go test -race -cover ./...
	@echo "===================          done           ==================="

# Remove assets
.PHONY: clean
clean:
	rm -f Temporal
	docker-compose -f test/docker-compose.yml rm -f -s -v

# Rebuild vendored dependencies
.PHONY: vendor
vendor:
	@echo "=================== generating dependencies ==================="
	# Update standard dependencies
	dep ensure -v

	# Generate IPFS dependencies
	go get -u github.com/ipfs/go-ipfs
	( cd $(GOPATH)/src/github.com/ipfs/go-ipfs ; make install )

	# Vendor IPFS dependencies
	rm -rf vendor/github.com/ipfs/go-ipfs vendor/gx
	cp -r $(GOPATH)/src/github.com/ipfs/go-ipfs vendor/github.com/ipfs
	cp -r $(GOPATH)/src/gx vendor/github.com/ipfs vendor

	# Vendor ethereum - this step is required for some of the cgo components, as
	# dep doesn't seem to resolve them
	go get -u github.com/ethereum/go-ethereum
	cp -r $(GOPATH)/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1 \
  	./vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/
	@echo "===================          done           ==================="