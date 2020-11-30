.PHONY: install_deps
install_deps:
	go get -u github.com/google/addlicense
	docker pull quay.io/goswagger/swagger

.PHONY: download_swagger
download_swagger:
	curl -o ./swagger.json https://tripletex.no/v2/swagger.json

.PHONY: generate
generate:
	docker run --rm -it -e GOPATH=$$HOME/go:/go -v $$HOME:$$HOME -w $$(pwd) quay.io/goswagger/swagger \
	generate client -f \
	./swagger.json -t ./ --skip-validation
	$$HOME/go/bin/addlicense -c "Bjerk AS" **/*.go