all: build run

build:
	cd src && go build -o ../tracker main.go

run:
	./tracker

clean:
	cd src && go clean
	@rm -f tracker