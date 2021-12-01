.PHTONY: build
build:
	go build -o mojimoji .

.PHTONY: test
test:
	go build -o mojimoji .
	./mojimoji tmp


