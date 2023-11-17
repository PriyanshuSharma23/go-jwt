build:
	go build -o ./bin/main main.go

clean:
	rm ./bin/main

run: build
	./bin/main

