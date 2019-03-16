build:
	@go build -o ./out/generator ./generator
	@go build -o ./out/checker ./checker

build-generator:
	@go build -o ./out/generator ./generator

build-checker:
	@go build -o ./out/checker ./checker

run:
	@ ./out/generator $(alphabet) $(length) | ./out/checker
