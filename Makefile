build:
	go build -o ./dist/kiti ./cmd/kiti.go 

clean:
	rm -rf ./dist

install:
	go install ./cmd/kiti.go
