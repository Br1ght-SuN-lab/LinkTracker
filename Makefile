run:
	go run ./cmd/bot


test:
	go test ./internal/bot/application/dispatcher


all_tests:
	go test -v ./internal/bot/application/dispatcher