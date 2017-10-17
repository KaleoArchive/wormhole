all: binary

.PHONY: binary
binary:
	@./scripts/build/binary.sh

.PHONY: cross
cross:
	@./scripts/build/cross.sh

.PHONY: vendor
vendor:
	@dep ensure

.PHONY: test
test:
	@./scripts/test.sh

.PHONY: image
image: 
	@./scripts/build/image.sh

.PHONY: push
push: 
	@./scripts/build/push.sh
