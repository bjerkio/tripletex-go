.PHONY: install_deps
install_deps:
	go get -u github.com/google/addlicense
	docker pull quay.io/goswagger/swagger

.PHONY: download_swagger
download_swagger:
	curl -o ./swagger.json https://triplepub.s3-eu-west-1.amazonaws.com/swagger.json

.PHONY: generate
generate:
	alias swagger="docker run --rm -it -e GOPATH=$$HOME/go:/go -v $$HOME:$$HOME -w $$(pwd) quay.io/goswagger/swagger"
	swagger generate client -f \
	./swagger.json -t ./ --skip-validation
	$$HOME/go/bin/addlicense -c "Bjerk AS" **/*.go