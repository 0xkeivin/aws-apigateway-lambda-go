buildgo:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o build/bin/app lambda-go/main.go
