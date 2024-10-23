all:
	go build -o goat -gcflags=all=-unusedresult=false main.go
	./goat
