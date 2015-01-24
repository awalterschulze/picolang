build:
	go install .
	picolang
	rm -rf ./out/packages/*
	cp -R ./funcs ./out/packages/funcs
	cp -R ./fun ./out/packages/fun
	(cd out && make stop && make build && make run && go run main.go)