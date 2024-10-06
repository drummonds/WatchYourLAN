mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/drummonds/WatchYourLAN && \
	go mod tidy

run:
	cd cmd/WatchYourLAN/ && \
	sudo \
	go run . #-c /data/config #-n http://192.168.2.3:8850

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint