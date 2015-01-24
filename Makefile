build:
	(cd lang && gocc lang.bnf)
	rm ./out/main.go || true
	go install ./...
	picolang -ip="192.168.59.103" -o="./out/" mapreduce.pico
	(cd out && make build)

run:
	(cd out && make run)

test:
	go test -v ./...

stop:
	(cd out && make stop)