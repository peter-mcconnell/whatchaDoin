whatchadoin
===========

A simple example project that shows how to build eBPF programs in C and load them with Golang; accompanying article: https://www.petermcconnell.com/posts/writing-an-ebpf-program-to-analyze-processes/.

This program will analyze behavior from processes on a machine.

![whatchadoin logo](./media/logo.png "whatchadoin")

build dependencies
------------------

This project was built and tested in the following environment:

```
 - x86_64
 - linux kernel 5.15.0-71-generic
 - llvm 16.0.3 with -bpf target support (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/llvm.yaml)
 - clang 16.0.3 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/llvm.yaml)
 - golang go1.20.2 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/golang.yaml)
 - docker 23.0.1 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/docker.yaml)
 - make 4.3
```

build everything (bpf program, go program, docker image)
--------------------------------------------------------

If you haven't already, ensure the libbpf submodule is pulled:
```sh
$ git submodule init
$ git submodule update
```

Then proceed to build:
```sh
$ make
```

running with docker
-------------------

```sh
$ make run_docker
# press Ctrl+C when you want to resume normality
```

Or you can run the image from docker hub (no need to pull this repo):

```sh
docker run --cap-add SYS_ADMIN -u0 --rm pemcconnell/whatchadoin:v0.0.1-pre
```

testing
-------

@TODO

clean
-----

```sh
$ make clean
```
