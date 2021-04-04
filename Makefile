os=$(shell sh -c "uname -s | tr '[:upper:]' '[:lower:]'")

.PHONY: snapshot
snapshot:
	goreleaser build --snapshot --rm-dist

.PHONY: install
install: snapshot
	cp dist/quote_$(os)_amd64/quote ~/.local/bin

