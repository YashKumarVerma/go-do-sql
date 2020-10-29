build: internal/*
	go build -o build/go-sql internal/main.go

run:
	go run internal/main.go

clean:
	rm -rf build/*