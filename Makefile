.PHONY: build build_bpf build_go build_docker clean run_docker push_docker test
.DEFAULT_GOAL = build

TAG := v0.0.1-pre

build_bpf:
	$(MAKE) -C bpf whatchadoin_kern.o

build_go:
	CGO_ENABLED=0 go build
	
build_docker:
	docker build -t=pemcconnell/whatchadoin:$(TAG) .

build: build_bpf build_go build_docker

clean:
	$(MAKE) -C bpf clean

test:
	go test -v ./...

run_docker:
	docker run --cap-add SYS_ADMIN -u0 --rm pemcconnell/whatchadoin:$(TAG)

push_docker: build_docker
	docker push pemcconnell/whatchadoin:$(TAG)
