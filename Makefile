BINARY_NAME=share-img

build:
 GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
 GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go

run: ./${BINARY_NAME}

build_and_run: build run

dev: compiledaemon -command='./${BINARY_NAME}' -directory='.'

clean: |
	# go clean
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux
	rm -f ${BINARY_NAME}-windows
	rm -f ${BINARY_NAME}
	rm -f gorm.db
	rm -rf storage/*