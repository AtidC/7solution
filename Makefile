PORT=8000

test1:
	go run test1/main.go

test2:
	go run test2/main.go

test3:
	open http://localhost:$(PORT)/beef/summary 
	go run test3/main.go 

.PHONY: test1 test2 test3	