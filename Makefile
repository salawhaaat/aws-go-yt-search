build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main src/main.go
deploy_prod: build
	serverless deploy --stage prod
