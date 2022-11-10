

compile: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o ./dist/chat ./cmd/chat

run: compile
	docker-compose up --build
