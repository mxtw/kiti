LDFLAGS="-w -s"

build:
	GOOS=linux go build -o ./dist/kiti -ldflags=${LDFLAGS}

all: build

clean:
	rm -rf ./dist
