NAME=monitor
# Output File Location
DIR=out
$(shell mkdir -p ${DIR})

# Go build flags
LDFLAGS=-ldflags "-s -w"

init:
	go mod tidy

serve:
	go run main.go

default:
	go build ${LDFLAGS} -o ${DIR}/${NAME} main.go

# Compile Server - Windows x64
windows:
	export GOOS=windows;export GOARCH=amd64;go build ${LDFLAGS} -o ${DIR}/${NAME}-Windows-x64.exe main.go

# Compile Server - Linux x64
linux:
	export GOOS=linux;export GOARCH=amd64;go build ${LDFLAGS} -o ${DIR}/${NAME}-Linux-x64 main.go

# Compile Server - Darwin x64
darwin:
	export GOOS=darwin;export GOARCH=amd64;go build ${LDFLAGS} -o ${DIR}/${NAME}-Darwin-x64 main.go

# clean
clean:
	rm -rf ${DIR}