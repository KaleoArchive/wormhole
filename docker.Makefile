all: binary

.PHONY: binary
binary:
	@./scripts/dev-shell.sh make binary

.PHONY: cross
cross:
	@./scripts/dev-shell.sh make cross

.PHONY: vendor
vendor:
	@./scripts/dev-shell.sh make vendor

.PHONY: test
test:
	@./scripts/dev-shell.sh make test
