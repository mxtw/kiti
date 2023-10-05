LDFLAGS="-w -s"

build:
	GOOS=linux go build -o ./dist/kiti -ldflags=${LDFLAGS}

build_windows:
	GOOS=windows go build -o ./dist/kiti.exe -ldflags=${LDFLAGS}

# TODO build other archs, build darwin

all: build build_windows

clean:
	rm -rf ./dist
