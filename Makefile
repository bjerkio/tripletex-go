.PHONY: generate
generate:
	swagger generate client -f \
	https://triplepub.s3-eu-west-1.amazonaws.com/swagger.json -t ./ \
	--skip-validation