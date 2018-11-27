

all: src/main.go src/index.go
	go run $^

clean:
	find . -name "*~" -exec rm {} \;
