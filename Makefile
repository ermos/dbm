build/bin:
	@go build -o ${GOPATH}/bin/dbm ./

prod/build:
	goreleaser --snapshot --skip-publish --rm-dist

prod/release:
	test $(VERSION)
	git tag -a v$(VERSION) -m "$(VERSION)"
	goreleaser --snapshot --skip-publish --rm-dist

delete-tag:
	test $(VERSION)
	git tag --delete v$(VERSION)