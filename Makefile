BINARY_DIR=./bin
BINARY_NAME=blogsrv

all: build test

air:
	air -build.cmd="make build" -build.bin="${BINARY_DIR}/${BINARY_NAME}"

build:
	templ generate
	mkdir -p ${BINARY_DIR}
	go build -o ${BINARY_DIR}/${BINARY_NAME} ./cmd/blogsrv

buildall: build
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_DIR}/${BINARY_NAME}-darwin ./cmd/blogsrv
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_DIR}/${BINARY_NAME}-linux ./cmd/blogsrv 
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_DIR}/${BINARY_NAME}-windows ./cmd/blogsrv

run: build
	${BINARY_DIR}/${BINARY_NAME}

test:
	go test -v ./...

clean:
	rm -rf ${BINARY_DIR}
