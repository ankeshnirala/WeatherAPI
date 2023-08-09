build:
	@go build -o weatherapi

run: build
	@./weatherapi