build:
	GOOS=linux GOARCH=amd64 go build -o dist/signal-watcher-linux-amd64 ./cmd/signal-watcher
	GOOS=linux GOARCH=arm64 go build -o dist/signal-watcher-linux-arm64 ./cmd/signal-watcher

clean:
	rm -rf dist
