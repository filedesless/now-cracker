

all: main.go index.go
	FLAG=flag{fake} go run $^

clean:
	find . -name "*~" -exec rm {} \;
