build:
	find . -type f -name "*.go" -exec go build {} \;

test:
	go test src/hue/hue_test.go
	go test src/pngimage/pngimage_test.go

clean:
	go clean
	rm -f *.png