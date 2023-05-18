TARGET = whatchadoin
BPF_TARGET = ${TARGET:=_kern}
BPF_C = ${BPF_TARGET:=.c}
BPF_OBJ = ${BPF_C:.c=.o}

BPF_PINNED_PATH := /sys/fs/bpf/$(TARGET)

$(BPF_OBJ): %.o: %.c
	clang -S \
		-g \
		-target bpf \
	  -I../libbpf/src\
		-Wall \
		-Werror \
		-O2 -emit-llvm -c -o ${@:.o=.ll} $<
	llc -march=bpf -filetype=obj -O2 -o $@ ${@:.o=.ll}

clean:
	sudo rm -f $(BPF_PINNED_PATH)
	rm -f $(BPF_OBJ)
	rm -f ${BPF_OBJ:.o=.ll}
