all: test build serve run close
.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	cd cli && go build -o goploader
	mv cli/goploader .

.PHONY: run
run: 
	./goploader --dir=example/files/ --url=http://localhost:8080/upload --method=post
	./goploader --file=example/files/testfile.json --url=http://localhost:8080/upload --method=post
	./goploader --file=example/files/testfile.json --url=http://localhost:8080/upload --method=put

.PHONY: serve
serve: 
	screen -d -m -S devses
	screen -dmS devses go run server/main.go
	sleep 2

.PHONE: close
close: 
	sh scripts/close-screens.sh
