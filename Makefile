run:
	air

tidy:
	go mod tidy -v

publish:
	go build -o apiserver .
