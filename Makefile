build:
	go build -o bin/shell main.go

clean:
	rm -rf bin/ history.txt