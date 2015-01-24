build:
	(cd lang && gocc lang.bnf)
	rm ./out/main.go || true
	go install ./...
	picolang
	(cd out && make stop && make build && make run)

run:
	(cd out && go run main.go)

test:
	go test -v ./...