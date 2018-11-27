

all:
	go run main.go index.go

clean:
	find . -name "*~" -exec rm {} \;
