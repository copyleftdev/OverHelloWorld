build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

run:
	go run ./cmd/server/main.go

docker:
	docker build -t overhelloworld .

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

loadtest:
	k6 run tests/load/hello-load-test.js
