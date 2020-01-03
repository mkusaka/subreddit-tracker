build:
	env GOOS=linux GOARCH=amd64 go build -o update main.go

run:
	go run main.go
