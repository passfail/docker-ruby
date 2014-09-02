build-container:
	docker build --tag passfail/docker-ruby .

build-lib:
	go build -o bin/docker-ruby
