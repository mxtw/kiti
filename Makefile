LDFLAGS="-w -s"

build:
	GOOS=linux go build -o ./dist/kiti ldflags=${LDFLAGS} ./cmd/kiti.go

build_windows:
	GOOS=windows go build -o ./dist/kiti.exe -ldflags=${LDFLAGS} ./cmd/kiti.go

all: build build_windows

clean:
	rm -rf ./dist

install:
	go install ./cmd/kiti.go
