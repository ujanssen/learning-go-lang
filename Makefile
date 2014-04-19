build:
	find . -type f -name "*.go" -exec go build {} \;

test:
	go test src/hue/hue_test.go
	go test src/task/task_test.go

clean:
	go clean
	rm -f *.png