serve:
    air

run:
    go run .

build:
    go build -o bin/site

test:
    go test ./...

clean:
    rm -rf bin/
