SOURCE_NAME=$(shell grep -lr "func main" main.go)
BINARY_NAME=$(shell grep -lr "func main" main.go | awk -F '.' '{print $$1}')

linux:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} ${SOURCE_NAME}
m1:
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME} ${SOURCE_NAME}
mac:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} ${SOURCE_NAME}

run:
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}

docker-run:
	docker-compose -f docker-compose.yaml -p restapi-go up -d