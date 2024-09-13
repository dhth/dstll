PACKAGE_NAME := github.com/dhth/dstll
DOCKER_IMAGE = ghcr.io/goreleaser/goreleaser-cross:v1.22.7

.PHONY: release-dry-run
release-dry-run:
	@docker run \
		--rm \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		$(DOCKER_IMAGE) \
	  build --clean --snapshot

.PHONY: release
release:
	@docker run \
		--rm \
		-e CGO_ENABLED=1 \
		-e GITHUB_TOKEN=${GITHUB_TOKEN} \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		$(DOCKER_IMAGE) \
	  release --clean
